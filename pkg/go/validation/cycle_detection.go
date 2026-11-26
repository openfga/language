package validation

import (
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// entryPointResult mirrors pkg/js's { hasEntry, loop }.
type entryPointResult struct {
	hasEntry bool
	loop     bool
}

// validateEntryPoints reports relations that have no entry point. Such a
// relation is impossible: either it never reaches a concrete assignable type
// (no entrypoint) or it forms a rewrite loop (potential loop).
func validateEntryPoints(idx *index, src source) error {
	var fs []*Finding

	entryPoints := resolveEntryPoints(idx)

	for _, typeDef := range idx.model.GetTypeDefinitions() {
		relations := typeDef.GetRelations()
		if len(relations) == 0 {
			continue
		}

		typeName := typeDef.GetType()
		typeLine := src.typeLine(typeName)

		for _, relationName := range slices.Sorted(maps.Keys(relations)) {
			result := entryPoints.of(typeName, relationName)
			if result.hasEntry {
				continue
			}

			file, module := relationMeta(typeDef, relationName)
			line := src.relationLine(relationName, typeLine)

			finding := noEntryPoint(relationName, typeName)
			if result.loop {
				finding = noEntryPointLoop(relationName, typeName)
			}

			finding.at(src, line)
			finding.File, finding.Metadata.Module = file, module
			fs = append(fs, finding)
		}
	}

	return joinFindings(fs...)
}

// relationRef names one relation of one type.
type relationRef struct {
	typeName     string
	relationName string
}

// entryPoints holds, for every relation of a model, the answer pkg/js's
// hasEntryPointOrLoop in validate-dsl.ts gives for it: whether the relation
// has an entry point, and if not, whether its walk ran into a loop.
//
// That walk follows every path out of a relation, skipping relations already
// on the path, which takes exponential time when relations share
// dependencies, as `a2: a1 and b1` does over `a1: a0 and b0` and
// `b1: a0 and b0`. The same answers are found here without walking paths:
//
//   - The relations with an entry point are the least fixed point of the
//     rewrites, which markEntryPoints builds up from the direct assignments.
//     Skipping the relations on its path does not stop the walk from finding
//     any of them, as a relation that can be satisfied can be satisfied
//     without going through itself.
//   - A walk from a relation with no entry point passes a loop up only through
//     failing computed usersets, each of which leads to another relation with
//     no entry point. With only such relations on its path, every operand
//     fails or succeeds as the fixed point says, so the usersets it follows are
//     fixed, and it reports a loop exactly when they lead to a cycle, which
//     markLoops finds from their strongly connected components.
type entryPoints struct {
	idx      *index
	hasEntry map[relationRef]bool
	loop     map[relationRef]bool
}

func resolveEntryPoints(idx *index) *entryPoints {
	e := &entryPoints{
		idx:      idx,
		hasEntry: map[relationRef]bool{},
		loop:     map[relationRef]bool{},
	}

	relations := e.relations()
	e.markEntryPoints(relations)
	e.markLoops(relations)

	return e
}

// of returns the answer for typeName#relationName. A relation the model does
// not define has no entry point and no loop.
func (e *entryPoints) of(typeName, relationName string) entryPointResult {
	ref := relationRef{typeName, relationName}

	return entryPointResult{hasEntry: e.hasEntry[ref], loop: e.loop[ref]}
}

// relations lists the relations of the model. One without a name never has
// an entry point, as in pkg/js, so it is left out.
func (e *entryPoints) relations() []relationRef {
	var refs []relationRef

	for _, typeName := range slices.Sorted(maps.Keys(e.idx.relations)) {
		for _, relationName := range slices.Sorted(maps.Keys(e.idx.relations[typeName])) {
			if relationName != "" {
				refs = append(refs, relationRef{typeName, relationName})
			}
		}
	}

	return refs
}

func (e *entryPoints) rewrite(ref relationRef) *openfgav1.Userset {
	return e.idx.userset(ref.typeName, ref.relationName)
}

// markEntryPoints marks the relations with an entry point: those whose rewrite
// is satisfied outright, and then, each time a relation is marked, any
// relation depending on it that the mark satisfies.
func (e *entryPoints) markEntryPoints(relations []relationRef) {
	dependents := map[relationRef][]relationRef{}

	for _, ref := range relations {
		for _, dependency := range e.dependencies(ref, e.rewrite(ref), nil) {
			dependents[dependency] = append(dependents[dependency], ref)
		}
	}

	pending := slices.Clone(relations)

	for len(pending) > 0 {
		ref := pending[len(pending)-1]
		pending = pending[:len(pending)-1]

		if e.hasEntry[ref] || !e.satisfied(ref, e.rewrite(ref)) {
			continue
		}

		e.hasEntry[ref] = true
		pending = append(pending, dependents[ref]...)
	}
}

// dependencies appends the relations whose entry point satisfied reads for
// rewrite, a rewrite of ref.
func (e *entryPoints) dependencies(ref relationRef, rewrite *openfgav1.Userset,
	refs []relationRef) []relationRef {
	switch rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		for _, restriction := range e.idx.directTypeRestrictions(ref.typeName, ref.relationName) {
			if restriction.GetRelation() != "" {
				refs = append(refs, relationRef{restriction.GetType(), restriction.GetRelation()})
			}
		}

	case *openfgav1.Userset_ComputedUserset:
		refs = append(refs, relationRef{ref.typeName, rewrite.GetComputedUserset().GetRelation()})

	case *openfgav1.Userset_TupleToUserset:
		ttu := rewrite.GetTupleToUserset()
		for _, restriction := range e.idx.directTypeRestrictions(ref.typeName, ttu.GetTupleset().GetRelation()) {
			refs = append(refs, relationRef{restriction.GetType(), ttu.GetComputedUserset().GetRelation()})
		}

	case *openfgav1.Userset_Union:
		for _, child := range rewrite.GetUnion().GetChild() {
			refs = e.dependencies(ref, child, refs)
		}

	case *openfgav1.Userset_Intersection:
		for _, child := range rewrite.GetIntersection().GetChild() {
			refs = e.dependencies(ref, child, refs)
		}

	case *openfgav1.Userset_Difference:
		refs = e.dependencies(ref, rewrite.GetDifference().GetBase(), refs)
		refs = e.dependencies(ref, rewrite.GetDifference().GetSubtract(), refs)
	}

	return refs
}

// satisfied reports whether rewrite, a rewrite of ref, reaches an entry point
// through the relations marked so far. It reads each operand as
// hasEntryPointOrLoop does.
func (e *entryPoints) satisfied(ref relationRef, rewrite *openfgav1.Userset) bool {
	switch rewrite.GetUserset().(type) {
	case *openfgav1.Userset_This:
		// A direct assignment has an entry point if any assignable type is a
		// concrete type or wildcard. A type#relation restriction only provides
		// an entry point if that referenced relation itself has one.
		for _, restriction := range e.idx.directTypeRestrictions(ref.typeName, ref.relationName) {
			if restriction.GetRelation() == "" || restriction.GetWildcard() != nil {
				return true
			}

			restricted := relationRef{restriction.GetType(), restriction.GetRelation()}
			if e.rewrite(restricted) == nil {
				// Matches validate-dsl.ts: returns on the first missing
				// reference rather than trying later types. Unreachable through
				// validate, which reports the missing relation and skips this
				// phase.
				return false
			}

			if e.hasEntry[restricted] {
				return true
			}
		}

		return false

	case *openfgav1.Userset_ComputedUserset:
		return e.hasEntry[relationRef{ref.typeName, rewrite.GetComputedUserset().GetRelation()}]

	case *openfgav1.Userset_TupleToUserset:
		ttu := rewrite.GetTupleToUserset()
		tupleset := ttu.GetTupleset().GetRelation()
		computed := ttu.GetComputedUserset().GetRelation()

		if tupleset == "" || !e.idx.relationDefined(ref.typeName, tupleset) {
			return false
		}

		for _, restriction := range e.idx.directTypeRestrictions(ref.typeName, tupleset) {
			if e.hasEntry[relationRef{restriction.GetType(), computed}] {
				return true
			}
		}

		return false

	case *openfgav1.Userset_Union:
		for _, child := range rewrite.GetUnion().GetChild() {
			if e.satisfied(ref, child) {
				return true
			}
		}

		return false

	case *openfgav1.Userset_Intersection:
		for _, child := range rewrite.GetIntersection().GetChild() {
			if !e.satisfied(ref, child) {
				return false
			}
		}

		return true

	case *openfgav1.Userset_Difference:
		diff := rewrite.GetDifference()

		return e.satisfied(ref, diff.GetBase()) && e.satisfied(ref, diff.GetSubtract())
	}

	return false
}

// markLoops marks the relations with no entry point whose walk runs into a
// loop: those from which the computed usersets the walk follows lead to a
// cycle. The strongly connected components of those usersets come from
// forEachComponent after every component they lead to, so a component loops if
// it is a cycle itself or leads to a component already marked.
func (e *entryPoints) markLoops(relations []relationRef) {
	follows := map[relationRef][]relationRef{}

	var stuck []relationRef

	for _, ref := range relations {
		if !e.hasEntry[ref] {
			stuck = append(stuck, ref)
			follows[ref] = e.followed(ref, e.rewrite(ref), nil)
		}
	}

	forEachComponent(stuck, follows, func(component []relationRef) {
		loop := len(component) > 1

		for _, ref := range component {
			for _, next := range follows[ref] {
				loop = loop || next == ref || e.loop[next]
			}
		}

		for _, ref := range component {
			e.loop[ref] = loop
		}
	})
}

// followed appends the relations the walk enters through a computed userset
// while walking rewrite, an unsatisfied rewrite of ref: those in every operand
// of a union, and in the first unsatisfied operand of an intersection or
// difference, where the walk stops.
func (e *entryPoints) followed(ref relationRef, rewrite *openfgav1.Userset,
	refs []relationRef) []relationRef {
	switch rewrite.GetUserset().(type) {
	case *openfgav1.Userset_ComputedUserset:
		computed := relationRef{ref.typeName, rewrite.GetComputedUserset().GetRelation()}
		if computed.relationName != "" && e.idx.relationDefined(computed.typeName, computed.relationName) {
			refs = append(refs, computed)
		}

	case *openfgav1.Userset_Union:
		for _, child := range rewrite.GetUnion().GetChild() {
			refs = e.followed(ref, child, refs)
		}

	case *openfgav1.Userset_Intersection:
		for _, child := range rewrite.GetIntersection().GetChild() {
			if !e.satisfied(ref, child) {
				return e.followed(ref, child, refs)
			}
		}

	case *openfgav1.Userset_Difference:
		diff := rewrite.GetDifference()
		if !e.satisfied(ref, diff.GetBase()) {
			return e.followed(ref, diff.GetBase(), refs)
		}

		return e.followed(ref, diff.GetSubtract(), refs)
	}

	return refs
}

// forEachComponent calls visit with each strongly connected component of the
// graph whose edges are next, found with Tarjan's algorithm, which completes a
// component only after every component it has an edge to.
func forEachComponent(nodes []relationRef, next map[relationRef][]relationRef,
	visit func(component []relationRef)) {
	order := map[relationRef]int{}
	low := map[relationRef]int{}
	onStack := map[relationRef]bool{}

	var stack []relationRef

	var connect func(node relationRef)
	connect = func(node relationRef) {
		order[node] = len(order)
		low[node] = order[node]
		stack = append(stack, node)
		onStack[node] = true

		for _, successor := range next[node] {
			if _, seen := order[successor]; !seen {
				connect(successor)
				low[node] = min(low[node], low[successor])
			} else if onStack[successor] {
				low[node] = min(low[node], order[successor])
			}
		}

		if low[node] != order[node] {
			return
		}

		var component []relationRef

		for {
			member := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			onStack[member] = false
			component = append(component, member)

			if member == node {
				break
			}
		}

		visit(component)
	}

	for _, node := range nodes {
		if _, seen := order[node]; !seen {
			connect(node)
		}
	}
}
