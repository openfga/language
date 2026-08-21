package validation

import (
	"fmt"
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// Cycles among relations are not all the same, and the rewrite-tree walk in
// cycle_detection.go only reports the ones that leave a relation with no way in. Two other
// shapes make a model unresolvable while every relation in them stays satisfiable, so that
// walk answers has-an-entry-point for all of them and reports nothing:
//
//	define a: [user] or b
//	define b: [user] or a
//
// Both hold a plain [user], so both have an entry point. The cycle between them reads no
// tuple, so resolving either one never gets closer to an answer.
//
//	define member: [user, group#member] but not blocked
//	define blocked: [user, group#member]
//
// Both hold a plain [user] again. Here the cycle does read a tuple, which is what makes
// `member: [user, group#member]` on its own a legal nesting, but a step of it is an
// operand of the exclusion, and the resolver cannot subtract a set it is still computing.
//
// The weighted graph refuses both, with ErrModelCycle and ErrTupleCycle, and names no
// relation in either. What follows finds the same two shapes over the model and names the
// relations, so a caller gets a position rather than a sentence about the whole model.
//
// The graph stays the authority on whether a model is resolvable.
// TestCycleShapesAgreeWithTheBuilder holds this to it: it may only report relations in
// models the builder refuses.

// cycleStep is one relation depending on another.
type cycleStep struct {
	to string

	// readsTuple marks a step that consumes a tuple, which is what lets a cycle
	// terminate. A direct `[type#relation]` restriction and a tuple-to-userset both read
	// one; rewriting `define x: y` does not.
	readsTuple bool

	// constrained marks a step written as an operand of an intersection or an exclusion.
	constrained bool
}

// cycleShape names why a cycle cannot be resolved.
type cycleShape int

// The values are ordered by how badly the cycle fails, least first. A relation can sit in
// several cycles and the worst decides what it is reported for, so the order is what makes
// that answer independent of which cycle the search reaches first.
const (
	cycleResolvable cycleShape = iota
	// The graph's ErrTupleCycle for AND and BUT NOT: the cycle terminates, but a step
	// of it is an operand of an intersection or an exclusion.
	cycleUnderConstraint
	// The graph's ErrModelCycle: going round the cycle consumes nothing, so it never
	// terminates. It ranks above the constraint because it fails whether or not an
	// operator encloses it.
	cycleReadsNoTuple
)

// worstCycleShape is the top of that order. A search that reaches it can stop, because
// nothing it has left to walk can change the answer.
const worstCycleShape = cycleReadsNoTuple

// cycleStepBudget caps the steps one relation's search may walk. Enumerating simple cycles
// is exponential in the worst case and a model is not required to be small.
//
// Exhausting it gives up rather than guesses, so the relation goes unreported and the
// refusal falls through to the positionless finding the graph's own error carries. A caller
// loses the position, not the answer. TestCycleSearchStaysWithinItsBudget pins both that
// the shared corpus stays far below the cap and that a model over it still degrades that
// way.
const cycleStepBudget = 200_000

// checkRelationCycleShapes reports every relation that takes part in a cycle the resolver
// cannot work through.
//
// A relation is reported for the cycle it is in rather than for anything wrong with the
// relation, so a satisfiable relation is reported when the cycle it belongs to is not
// resolvable. That is the case the builder refuses the whole model over.
func checkRelationCycleShapes(collector *ErrorCollector, validator *SemanticValidator, lines []string) {
	if validator == nil || validator.model == nil {
		return
	}

	steps := relationDependencySteps(validator)

	for _, typeDef := range validator.model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		if len(relations) == 0 {
			continue
		}

		typeName := typeDef.GetType()
		typeLineIndex := GetTypeLineNumber(typeName, lines, nil)

		for _, relationName := range slices.Sorted(maps.Keys(relations)) {
			shape, _ := worstCycleThrough(steps, typeName+"#"+relationName, cycleStepBudget)
			if shape == cycleResolvable {
				continue
			}

			collector.RaiseCyclicRelation(relationName, typeName, describeCycleShape(shape),
				relationMeta(typeDef, relationName),
				GetRelationLineNumber(relationName, lines, typeLineIndex))
		}
	}
}

// describeCycleShape is the clause the finding ends with, and it is the only place the two
// shapes are worded.
func describeCycleShape(shape cycleShape) string {
	if shape == cycleReadsNoTuple {
		return "no relation in it reads a tuple, so resolving it never terminates"
	}

	return "a relation in it is an operand of an `and` or a `but not`"
}

// relationDependencySteps builds, for every relation in the model, the relations it depends
// on and how.
//
// Only the steps between relations are recorded. Terminal types and wildcards end a path
// rather than continuing one, so they cannot be part of a cycle and are left out.
func relationDependencySteps(validator *SemanticValidator) map[string][]cycleStep {
	steps := map[string][]cycleStep{}

	for _, typeDef := range validator.model.GetTypeDefinitions() {
		typeName := typeDef.GetType()

		// Not sorted, unlike the pass in checkRelationCycleShapes. Each relation writes
		// one key of its own and the steps within a key come out in rewrite-tree order,
		// so the map this builds is the same map whatever order it is filled in. What
		// findings come out in is decided where they are raised.
		for relationName, rewrite := range typeDef.GetRelations() {
			from := typeName + "#" + relationName
			collectCycleSteps(validator, typeName, relationName, rewrite, false, func(step cycleStep) {
				steps[from] = append(steps[from], step)
			})
		}
	}

	return steps
}

// collectCycleSteps walks one relation's rewrite and hands every relation it reaches to
// emit.
//
// Whether a step is constrained is carried down rather than recomputed: a union's operands
// are as constrained as the union itself, while an intersection's and an exclusion's
// operands are constrained whatever encloses them, so nesting a union inside an exclusion
// keeps the exclusion's answer.
func collectCycleSteps(validator *SemanticValidator, typeName, relationName string,
	rewrite *openfgav1.Userset, constrained bool, emit func(cycleStep)) {
	if rewrite == nil {
		return
	}

	switch rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		for _, tr := range directlyRelatedTypes(validator, typeName, relationName) {
			// A concrete type or a wildcard is where a path ends. Only a userset
			// restriction continues to another relation.
			if tr.GetRelation() == "" || tr.GetWildcard() != nil {
				continue
			}

			emit(cycleStep{to: tr.GetType() + "#" + tr.GetRelation(), readsTuple: true, constrained: constrained})
		}

	case *openfgav1.Userset_ComputedUserset:
		if computed := rewrite.GetComputedUserset().GetRelation(); computed != "" {
			emit(cycleStep{to: typeName + "#" + computed, readsTuple: false, constrained: constrained})
		}

	case *openfgav1.Userset_TupleToUserset:
		ttu := rewrite.GetTupleToUserset()

		tupleset := ttu.GetTupleset().GetRelation()
		computed := ttu.GetComputedUserset().GetRelation()

		if tupleset == "" || computed == "" {
			return
		}

		for _, tr := range directlyRelatedTypes(validator, typeName, tupleset) {
			if tr.GetType() == "" {
				continue
			}

			emit(cycleStep{to: tr.GetType() + "#" + computed, readsTuple: true, constrained: constrained})
		}

	case *openfgav1.Userset_Union:
		for _, child := range rewrite.GetUnion().GetChild() {
			collectCycleSteps(validator, typeName, relationName, child, constrained, emit)
		}

	case *openfgav1.Userset_Intersection:
		for _, child := range rewrite.GetIntersection().GetChild() {
			collectCycleSteps(validator, typeName, relationName, child, true, emit)
		}

	case *openfgav1.Userset_Difference:
		difference := rewrite.GetDifference()
		collectCycleSteps(validator, typeName, relationName, difference.GetBase(), true, emit)
		collectCycleSteps(validator, typeName, relationName, difference.GetSubtract(), true, emit)
	}
}

// directlyRelatedTypes returns the type restrictions declared for a relation.
func directlyRelatedTypes(validator *SemanticValidator, typeName,
	relationName string) []*openfgav1.RelationReference {
	typeDef := validator.GetTypeDefinition(typeName)
	if typeDef == nil {
		return nil
	}

	metadata, ok := typeDef.GetMetadata().GetRelations()[relationName]
	if !ok {
		return nil
	}

	return metadata.GetDirectlyRelatedUserTypes()
}

// worstCycleThrough returns the least resolvable cycle that passes through start, and the
// steps it walked to decide.
//
// A relation can sit in several cycles at once, a resolvable one and an unresolvable one,
// and a legal cycle alongside an illegal one does not make the model legal. So every cycle
// back to start is classified and the worst one wins, which is also why the search cannot
// stop at the first cycle it finds. It stops only once the answer can no longer change, at
// the top of the order or out of budget.
//
// The budget is a parameter rather than the constant so a test can pin that the answer for
// a real model does not depend on it.
func worstCycleThrough(steps map[string][]cycleStep, start string, budget int) (cycleShape, int) {
	worst := cycleResolvable
	walked := 0

	// onPath holds the steps taken from start to the relation being expanded, so a step
	// back to start closes a cycle whose properties are the disjunction over the path.
	var onPath []cycleStep

	inPath := map[string]bool{start: true}

	var walk func(from string)

	walk = func(from string) {
		for _, step := range steps[from] {
			// Checked per step rather than on entry so a decided search unwinds
			// here instead of walking every remaining sibling first.
			if worst == worstCycleShape || walked >= budget {
				return
			}

			walked++

			if step.to == start {
				if shape := classifyCycle(append(onPath, step)); shape > worst {
					worst = shape
				}

				continue
			}

			// A relation already on this path leads only to cycles that do not
			// pass through start, which belong to that relation's own search.
			if inPath[step.to] {
				continue
			}

			inPath[step.to] = true
			onPath = append(onPath, step)

			walk(step.to)

			onPath = onPath[:len(onPath)-1]
			delete(inPath, step.to)
		}
	}

	walk(start)

	// Reaching the top of the order settles it whatever the budget did, because no cycle
	// left unwalked could be worse.
	if worst == worstCycleShape {
		return worst, walked
	}

	// Otherwise a search cut short has not seen every cycle, so what it holds is a lower
	// bound. Reporting it would give one relation of a cycle a finding and the next one
	// nothing, depending on where each search ran out.
	if walked >= budget {
		return cycleResolvable, walked
	}

	return worst, walked
}

// classifyCycle decides why, if at all, a closed cycle cannot be resolved.
//
// Reading no tuple is checked first because it is the more basic failure: such a cycle
// does not terminate whether or not an operator encloses it.
func classifyCycle(cycle []cycleStep) cycleShape {
	readsTuple, constrained := false, false

	for _, step := range cycle {
		readsTuple = readsTuple || step.readsTuple
		constrained = constrained || step.constrained
	}

	switch {
	case !readsTuple:
		return cycleReadsNoTuple
	case constrained:
		return cycleUnderConstraint
	default:
		return cycleResolvable
	}
}

// describeCycleSteps renders a relation's outgoing steps. It exists for test failure
// output, where a wrong answer is unreadable without them.
func describeCycleSteps(steps map[string][]cycleStep, from string) string {
	rendered := ""

	for _, step := range steps[from] {
		rendered += fmt.Sprintf(" -> %s(readsTuple=%v constrained=%v)", step.to, step.readsTuple, step.constrained)
	}

	return rendered
}
