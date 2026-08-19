package validation

import (
	"fmt"
	"strings"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
	"github.com/openfga/language/pkg/go/graph"
)

// TestCheckRelationCycleShapes covers the shapes the rewrite-tree walk has nothing to say
// about, and the legal cycles that must not be caught alongside them.
//
// The messages are asserted in full. A finding that named the right relation and worded the
// reason for a different shape would be worse than no finding, because a reader would go
// looking for the wrong thing.
func TestCheckRelationCycleShapes(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl  string
		want []string
	}{
		// Both relations hold a plain [user], so both have an entry point and the walk is
		// silent. The cycle between them reads no tuple.
		"rewrite cycle that reads no tuple": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`,
			want: []string{
				"[cyclic-relation] a line 5-5 column 11-12 " +
					"\"`a` on `document` takes part in a cycle that cannot be resolved: " +
					"no relation in it reads a tuple, so resolving it never terminates.\"",
				"[cyclic-relation] b line 6-6 column 11-12 " +
					"\"`b` on `document` takes part in a cycle that cannot be resolved: " +
					"no relation in it reads a tuple, so resolving it never terminates.\"",
			},
		},
		// The cycle does read a tuple, which is what makes the nesting on its own legal, but
		// a step of it is an operand of the exclusion.
		"tuple cycle through an exclusion": {
			dsl: `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not blocked
    define blocked: [user, group#member]
`,
			want: []string{
				"[cyclic-relation] blocked line 6-6 column 11-18 " +
					"\"`blocked` on `group` takes part in a cycle that cannot be resolved: " +
					"a relation in it is an operand of an `and` or a `but not`.\"",
				"[cyclic-relation] member line 5-5 column 11-17 " +
					"\"`member` on `group` takes part in a cycle that cannot be resolved: " +
					"a relation in it is an operand of an `and` or a `but not`.\"",
			},
		},
		// admin is an operand of the same intersection but is in no cycle, so it is not
		// reported. The finding is for taking part in the cycle, not for the operator.
		"tuple cycle through an intersection names only the relations in it": {
			dsl: `model
  schema 1.1
type user
type group
  relations
    define admin: [user]
    define member: [user, group#member] and admin
`,
			want: []string{
				"[cyclic-relation] member line 6-6 column 11-17 " +
					"\"`member` on `group` takes part in a cycle that cannot be resolved: " +
					"a relation in it is an operand of an `and` or a `but not`.\"",
			},
		},
		// The cycle reads a tuple at every step and nothing constrains it, which is the
		// nested-group model every deployment has.
		"userset nesting is not reported": {
			dsl: `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member]
`,
			want: []string{},
		},
		// Recursion through a tupleset reads a tuple too.
		"tupleset recursion is not reported": {
			dsl: `model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define viewer: [user] or viewer from parent
`,
			want: []string{},
		},
		// An exclusion with no cycle running through it is ordinary.
		"exclusion outside a cycle is not reported": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define blocked: [user]
    define viewer: [user] but not blocked
`,
			want: []string{},
		},
		// A resolvable cycle alongside an unresolvable one does not make the model legal, so
		// member is reported for the constrained cycle it is also in.
		"a legal cycle does not excuse an illegal one": {
			dsl: `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not blocked
    define blocked: [user, group#member]
    define owner: [user, group#owner]
`,
			want: []string{
				"[cyclic-relation] blocked line 6-6 column 11-18 " +
					"\"`blocked` on `group` takes part in a cycle that cannot be resolved: " +
					"a relation in it is an operand of an `and` or a `but not`.\"",
				"[cyclic-relation] member line 5-5 column 11-17 " +
					"\"`member` on `group` takes part in a cycle that cannot be resolved: " +
					"a relation in it is an operand of an `and` or a `but not`.\"",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.ElementsMatch(t, test.want, describeFindings(graphFindings(t, test.dsl)))
		})
	}
}

// TestCycleShapeFindingsCarryTheCallerContract checks the parts of a finding a consumer
// branches on rather than reads, and that the sentinel is not the entrypoint one.
//
// A relation reported here is satisfiable on its own, so a caller that treated the two as
// interchangeable would tell a user to give the relation an entrypoint it already has.
func TestCycleShapeFindingsCarryTheCallerContract(t *testing.T) {
	t.Parallel()

	findings := graphFindings(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`)
	require.Len(t, findings, 2)

	for _, finding := range findings {
		require.NotNil(t, finding.Metadata)

		assert.Equal(t, CyclicRelation, finding.Metadata.ErrorType)
		require.ErrorIs(t, finding, fgaerrors.ErrRelationInUnresolvableCycle)
		require.NotErrorIs(t, finding, fgaerrors.ErrNoEntrypoints,
			"the relation has an entrypoint, and a caller filtering on one must not see the other")
		require.NotErrorIs(t, finding, fgaerrors.ErrModelNotBuildable,
			"a positioned finding must not also read as the positionless refusal")

		assert.Equal(t, fgaerrors.SeverityError, finding.Severity)
		assert.Equal(t, fgaerrors.ErrorKindRelation, finding.Category)
		assert.True(t, isCriticalErrorType(finding.Metadata.ErrorType))

		// The whole point of computing this outside the graph is somewhere to put it.
		require.NotNil(t, finding.Line)
		require.NotNil(t, finding.Column)
	}
}

// TestCycleShapesAgreeWithTheBuilder holds the check to the graph, which stays the
// authority on whether a model is resolvable.
//
// The check reads the model rather than the graph, so it is a second implementation of a
// question the builder already answers, and the failure that costs a user most is a
// relation reported in a model the builder accepts. Every case in the shared corpus is run
// through both.
func TestCycleShapesAgreeWithTheBuilder(t *testing.T) {
	t.Parallel()

	var accepted, acceptedAndReported, refused, refusedAndReported int

	for _, entry := range corpusModels(t) {
		collector := NewErrorCollector(nil)
		checkRelationCycleShapes(collector, NewSemanticValidator(entry.Model),
			strings.Split(entry.Case.DSL, "\n"))

		reported := collector.CountAll() > 0

		if _, buildErr := graph.NewWeightedAuthorizationModelGraphBuilder().Build(entry.Model); buildErr == nil {
			accepted++

			assert.Falsef(t, reported,
				"the builder accepts %q, so no relation in it takes part in an unresolvable cycle:\n%s",
				entry.Case.Name, strings.Join(describeFindings(collector.AllFindings()), "\n"))

			if reported {
				acceptedAndReported++
			}

			continue
		}

		refused++

		if reported {
			refusedAndReported++
		}
	}

	// Without these the assertion above passes on a check that reports nothing at all.
	assert.Positive(t, accepted, "no corpus model built, so agreement was never tested")
	assert.Positive(t, refusedAndReported,
		"the check reported nothing anywhere, so agreeing with the builder cost it nothing")

	t.Logf("builder accepted %d (reported on %d), refused %d (reported on %d)",
		accepted, acceptedAndReported, refused, refusedAndReported)
}

// TestCycleSearchStaysWithinItsBudget pins the two things the cap has to be true of: it is
// nowhere near binding on a real model, and a model that does exhaust it gives up rather
// than reports half an answer.
func TestCycleSearchStaysWithinItsBudget(t *testing.T) {
	t.Parallel()

	t.Run("no corpus model comes close to the cap", func(t *testing.T) {
		t.Parallel()

		worstWalked, worstCase := 0, ""

		for _, entry := range corpusModels(t) {
			steps := relationDependencySteps(NewSemanticValidator(entry.Model))

			for _, typeDef := range entry.Model.GetTypeDefinitions() {
				for relationName := range typeDef.GetRelations() {
					start := typeDef.GetType() + "#" + relationName

					shape, walked := worstCycleThrough(steps, start, cycleStepBudget)
					if walked > worstWalked {
						worstWalked, worstCase = walked, entry.Case.Name+" "+start
					}

					// The answer for a real model cannot depend on the cap, or
					// tightening it later would silently change what is reported.
					tighter, _ := worstCycleThrough(steps, start, 1_000)
					assert.Equalf(t, shape, tighter,
						"%s in %q answers differently at a tighter budget", start, entry.Case.Name)
				}
			}
		}

		assert.Positive(t, worstWalked, "no relation walked a step, so the search was never exercised")
		assert.Lessf(t, worstWalked, cycleStepBudget/100,
			"the worst corpus relation walked %d steps, which is within two orders of the cap (%s)",
			worstWalked, worstCase)

		t.Logf("worst corpus relation walked %d of %d steps (%s)", worstWalked, cycleStepBudget, worstCase)
	})

	t.Run("the top of the order stops the search at once", func(t *testing.T) {
		t.Parallel()

		// Twelve relations rewriting each other is factorially many paths, and no path
		// reads a tuple, so the first cycle closed is already the worst there is.
		steps := relationDependencySteps(NewSemanticValidator(modelFromDSL(t, denselyRewritingModel(12))))

		shape, walked := worstCycleThrough(steps, "document#r0", cycleStepBudget)

		assert.Equal(t, cycleReadsNoTuple, shape)
		assert.Lessf(t, walked, 10, "the search kept walking after the answer could no longer change")
	})

	t.Run("a search that settles below the top reports what it found", func(t *testing.T) {
		t.Parallel()

		// Constrained tuple cycles are not the top of the order, so every path is walked
		// on the chance that a worse cycle is further in. At eight relations that finishes
		// inside the cap.
		dsl := denselyConstrainedModel(8)
		model := modelFromDSL(t, dsl)

		require.ErrorIs(t, buildRefusal(t, model), graph.ErrTupleCycle)

		shape, walked := worstCycleThrough(relationDependencySteps(NewSemanticValidator(model)),
			"document#r0", cycleStepBudget)

		assert.Equal(t, cycleUnderConstraint, shape)
		require.Lessf(t, walked, cycleStepBudget,
			"the search hit the cap, so settling inside it is not what this case tests")

		findings := findingsFrom(ValidateDSL(model, dsl, &EngineOptions{UseGraphValidation: true})).GetErrors()
		assert.Len(t, findings, 8, "one per relation in the cycle, admin excepted")

		for _, finding := range findings {
			require.NotNil(t, finding.Metadata)
			assert.Equal(t, CyclicRelation, finding.Metadata.ErrorType)
			assert.NotNil(t, finding.Line)
		}
	})

	t.Run("a search that exhausts the cap gives up rather than half answers", func(t *testing.T) {
		t.Parallel()

		// One more relation is around eight times the paths, which is over the cap.
		dsl := denselyConstrainedModel(9)
		model := modelFromDSL(t, dsl)

		require.ErrorIs(t, buildRefusal(t, model), graph.ErrTupleCycle)

		shape, walked := worstCycleThrough(relationDependencySteps(NewSemanticValidator(model)),
			"document#r0", cycleStepBudget)

		assert.Equal(t, cycleResolvable, shape,
			"a search cut short must not report the shape it had reached, which is a lower bound")
		require.Equal(t, cycleStepBudget, walked, "the cap did not bind, so giving up is not being tested")

		// What a caller loses is the position, not the answer. The refusal still reaches
		// them through the backstop.
		findings := findingsFrom(ValidateDSL(model, dsl, &EngineOptions{UseGraphValidation: true})).GetErrors()
		require.Len(t, findings, 1)
		require.NotNil(t, findings[0].Metadata)

		assert.Equal(t, GraphModelUnbuildable, findings[0].Metadata.ErrorType)
		require.ErrorIs(t, findings[0], graph.ErrTupleCycle)
	})
}

// buildRefusal returns the error the builder refused a model with, failing if it accepted
// it. A test that means to exercise a refusal is testing nothing once the model builds.
func buildRefusal(t *testing.T, model *openfgav1.AuthorizationModel) error {
	t.Helper()

	_, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.Error(t, err, "the builder accepted the model, so there is no refusal to fall back from")

	return err
}

// denselyRewritingModel builds relations that all rewrite each other and nothing else, so
// the paths through any one of them grow factorially and no path reads a tuple.
func denselyRewritingModel(relationCount int) string {
	return denseModel(relationCount, "r%d", " or ", "%s")
}

// denselyConstrainedModel builds the same density out of userset restrictions under an
// intersection, so every path reads a tuple and every step is constrained. That is the
// combination no cycle in it can be the top of the order for.
func denselyConstrainedModel(relationCount int) string {
	return denseModel(relationCount, "document#r%d", ", ", "[user, %s] and admin")
}

// denseModel writes relationCount relations, each referring to every other. Each reference
// is target formatted with the relation's index, the references are joined with separator,
// and the joined list is substituted into rewrite.
func denseModel(relationCount int, target, separator, rewrite string) string {
	var dsl strings.Builder

	dsl.WriteString("model\n  schema 1.1\ntype user\ntype document\n  relations\n    define admin: [user]\n")

	for i := range relationCount {
		references := make([]string, 0, relationCount-1)

		for j := range relationCount {
			if j != i {
				references = append(references, fmt.Sprintf(target, j))
			}
		}

		dsl.WriteString(fmt.Sprintf("    define r%d: "+rewrite+"\n", i, strings.Join(references, separator)))
	}

	return dsl.String()
}

// TestClassifyCycle covers the classifier on its own, over closed cycles rather than
// models, so the precedence between the two shapes is pinned where it is decided.
func TestClassifyCycle(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		cycle []cycleStep
		want  cycleShape
	}{
		"reads a tuple and nothing constrains it": {
			cycle: []cycleStep{{to: "group#member", readsTuple: true}},
			want:  cycleResolvable,
		},
		"reads no tuple": {
			cycle: []cycleStep{{to: "document#b"}, {to: "document#a"}},
			want:  cycleReadsNoTuple,
		},
		"reads a tuple under a constraint": {
			cycle: []cycleStep{{to: "group#blocked", constrained: true}, {to: "group#member", readsTuple: true}},
			want:  cycleUnderConstraint,
		},
		// One step reading a tuple is enough for the cycle to terminate, so the constraint
		// is what is left to report.
		"one step of several reads a tuple": {
			cycle: []cycleStep{{to: "a"}, {to: "b", readsTuple: true, constrained: true}, {to: "c"}},
			want:  cycleUnderConstraint,
		},
		// Reading no tuple outranks the constraint, and it has to outrank it here as well
		// as in the search, or a relation in both shapes would be worded by whichever the
		// search reached first.
		"reads no tuple under a constraint": {
			cycle: []cycleStep{{to: "a", constrained: true}, {to: "b"}},
			want:  cycleReadsNoTuple,
		},
		// Not a cycle any model produces, and the answer still has to be one of the three
		// rather than a zero value that happens to read as resolvable.
		"no steps": {
			cycle: nil,
			want:  cycleReadsNoTuple,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.want, classifyCycle(test.cycle))
		})
	}
}

// TestWorstCycleShapeIsTheTopOfTheOrder pins the constant against the order it names. It is
// what the search stops on, so a shape added below it would quietly stop the search early.
func TestWorstCycleShapeIsTheTopOfTheOrder(t *testing.T) {
	t.Parallel()

	for _, shape := range []cycleShape{cycleResolvable, cycleUnderConstraint, cycleReadsNoTuple} {
		assert.LessOrEqual(t, shape, worstCycleShape)
	}
}

// TestRelationDependencySteps covers what the steps map records, which is what every answer
// above is computed from.
//
// A missing step is a cycle never found and a spurious one is a cycle that is not there, so
// each case is a claim about one kind of rewrite rather than about a whole model.
func TestRelationDependencySteps(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl   string
		from  string
		want  []cycleStep
		other map[string][]cycleStep
	}{
		// A concrete type and a wildcard end a path, so neither is a step. Only the userset
		// restriction continues to another relation, and reading its tuple is what lets the
		// cycle terminate.
		"direct assignment records only the userset restrictions": {
			dsl: `model
  schema 1.1
type user
type group
  relations
    define member: [user, user:*, group#member]
`,
			from: "group#member",
			want: []cycleStep{{to: "group#member", readsTuple: true}},
		},
		// A rewrite reads nothing, which is the whole distinction the shapes turn on.
		"computed userset reads no tuple": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define editor: [user]
    define viewer: editor
`,
			from: "document#viewer",
			want: []cycleStep{{to: "document#editor"}},
		},
		// A tupleset expands to one step per assignable type of the tupleset relation, and
		// the computed relation is looked for on each. Recording only the first would miss
		// a cycle that runs through the second.
		"tupleset expands across every assignable type": {
			dsl: `model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]
type team
  relations
    define viewer: [user]
type document
  relations
    define parent: [folder, team]
    define viewer: viewer from parent
`,
			from: "document#viewer",
			want: []cycleStep{
				{to: "folder#viewer", readsTuple: true},
				{to: "team#viewer", readsTuple: true},
			},
		},
		// A union constrains nothing, so its operands are as constrained as the union is.
		"union operands are unconstrained": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define a: [user]
    define b: [user]
    define viewer: a or b
`,
			from: "document#viewer",
			want: []cycleStep{{to: "document#a"}, {to: "document#b"}},
		},
		"intersection operands are constrained": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define a: [user]
    define b: [user]
    define viewer: a and b
`,
			from: "document#viewer",
			want: []cycleStep{{to: "document#a", constrained: true}, {to: "document#b", constrained: true}},
		},
		// Both sides of an exclusion are constrained. The subtrahend obviously is, and the
		// base is too, because the resolver cannot finish it without the other.
		"both operands of an exclusion are constrained": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define a: [user]
    define b: [user]
    define viewer: a but not b
`,
			from: "document#viewer",
			want: []cycleStep{{to: "document#a", constrained: true}, {to: "document#b", constrained: true}},
		},
		// Nesting a union inside an exclusion keeps the exclusion's answer, so an operand
		// two levels down is still constrained.
		"a union nested in an exclusion stays constrained": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define a: [user]
    define b: [user]
    define c: [user]
    define viewer: (a or b) but not c
`,
			from: "document#viewer",
			want: []cycleStep{
				{to: "document#a", constrained: true},
				{to: "document#b", constrained: true},
				{to: "document#c", constrained: true},
			},
		},
		// A relation with nowhere to go has no steps rather than an empty one, so it cannot
		// close a cycle.
		"a terminal relation records no steps": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]
`,
			from: "document#viewer",
			want: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			steps := relationDependencySteps(NewSemanticValidator(modelFromDSL(t, test.dsl)))

			assert.ElementsMatchf(t, test.want, steps[test.from], "steps for %s:%s",
				test.from, describeCycleSteps(steps, test.from))
		})
	}
}

// TestCycleShapesRunOnlyBehindARefusedBuild pins where the check sits. It is a second
// answer to a question the graph already answers, so on a model the graph accepts the graph
// rules are what run.
func TestCycleShapesRunOnlyBehindARefusedBuild(t *testing.T) {
	t.Parallel()

	// Legal, and every relation in it is in a cycle the check has an opinion about only
	// because the check is never asked.
	dsl := `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member]
`
	model := modelFromDSL(t, dsl)

	_, buildErr := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.NoError(t, buildErr, "the build has to succeed for this to test the branch it means to")

	assert.Empty(t, describeFindings(graphFindings(t, dsl)))
}

// TestCycleShapesDoNotDisplaceTheWalk keeps the two fallbacks in order. The walk names the
// same relations for a model that is both cyclic and unsatisfiable, and it names them with
// the wording the corpus pins, so running the cycle check as well would double every
// finding.
func TestCycleShapesDoNotDisplaceTheWalk(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type document
  relations
    define admin: [user]
    define viewer: admin and editor
    define editor: viewer
`

	_, buildErr := graph.NewWeightedAuthorizationModelGraphBuilder().Build(modelFromDSL(t, dsl))
	require.Error(t, buildErr, "the model has to be one the builder refuses")

	collector := NewErrorCollector(nil)
	checkRelationCycleShapes(collector, NewSemanticValidator(modelFromDSL(t, dsl)),
		strings.Split(dsl, "\n"))
	require.NotEmpty(t, collector.AllFindings(),
		"the cycle check has to have something to say here, or the ordering is untested")

	for _, finding := range graphFindings(t, dsl) {
		require.NotNil(t, finding.Metadata)
		assert.Equal(t, RelationNoEntrypoint, finding.Metadata.ErrorType,
			"the walk found this relation first, so the cycle check must not report it again")
	}
}

// TestCheckRelationCycleShapesOnNilModel checks the guard, which is reached the same way
// validateWithGraph's is.
func TestCheckRelationCycleShapesOnNilModel(t *testing.T) {
	t.Parallel()

	collector := NewErrorCollector(nil)

	checkRelationCycleShapes(collector, nil, nil)
	assert.Equal(t, 0, collector.CountAll(), "no validator, so nothing to read")

	checkRelationCycleShapes(collector, NewSemanticValidator(nil), nil)
	assert.Equal(t, 0, collector.CountAll(), "a validator over a nil model has no relations")
}

// TestCycleShapeFindingsAreOrderedDeterministically pins the order findings come out in.
//
// The relations of a type live in a map, so without sorting them the same model would report
// the same cycle in a different order per run. That is not a wrong answer but it is an
// unusable one: a corpus comparison goes flaky and an editor redraws its diagnostics in a
// different order every keystroke.
func TestCycleShapeFindingsAreOrderedDeterministically(t *testing.T) {
	t.Parallel()

	// Relation names deliberately out of declaration order, so sorting them is visible
	// rather than incidentally matching the source.
	dsl := `model
  schema 1.1
type user
type document
  relations
    define zebra: [user] or alpha
    define alpha: [user] or mango
    define mango: [user] or zebra
`

	first := describeFindings(graphFindings(t, dsl))
	require.Len(t, first, 3, "all three relations are in the cycle")

	for range 20 {
		assert.Equal(t, first, describeFindings(graphFindings(t, dsl)))
	}

	symbols := make([]string, 0, len(first))

	for _, finding := range graphFindings(t, dsl) {
		require.NotNil(t, finding.Metadata)

		symbols = append(symbols, finding.Metadata.Symbol)
	}

	assert.Equal(t, []string{"alpha", "mango", "zebra"}, symbols,
		"findings within a type come out in relation-name order")
}

// TestCycleShapeFindingCarriesTheDeclaringFileAndModule covers the provenance an editor
// needs to put the diagnostic in the right document.
//
// A modular model reaches validation already flattened, with the file and module each
// relation came from left on the model's metadata, so a finding that dropped them would be
// unplaceable in the only case where placing it is hard. The relation-level source info is
// read where it is set and the type's is the fallback, so both are exercised here.
func TestCycleShapeFindingCarriesTheDeclaringFileAndModule(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`
	model := modelFromDSL(t, dsl)

	for _, typeDef := range model.GetTypeDefinitions() {
		if typeDef.GetType() != "document" {
			continue
		}

		require.NotNil(t, typeDef.GetMetadata())

		typeDef.Metadata.Module = "core"
		typeDef.Metadata.SourceInfo = &openfgav1.SourceInfo{File: "core.fga"}

		// Only a carries its own source info. b falls back to the type's.
		relationMetadata, ok := typeDef.GetMetadata().GetRelations()["a"]
		require.True(t, ok, "relation a has no metadata, so the direct read is untested")

		relationMetadata.Module = "documents"
		relationMetadata.SourceInfo = &openfgav1.SourceInfo{File: "documents.fga"}
	}

	byRelation := map[string]*ValidationError{}

	for _, finding := range findingsFrom(ValidateDSL(model, dsl,
		&EngineOptions{UseGraphValidation: true})).GetErrors() {
		require.NotNil(t, finding.Metadata)
		require.Equal(t, CyclicRelation, finding.Metadata.ErrorType)

		byRelation[finding.Metadata.Symbol] = finding
	}

	require.Contains(t, byRelation, "a")
	require.Contains(t, byRelation, "b")

	assert.Equal(t, "documents.fga", byRelation["a"].File, "a declares its own file")
	assert.Equal(t, "documents", byRelation["a"].Metadata.Module)

	assert.Equal(t, "core.fga", byRelation["b"].File, "b has none, so the type's file stands in")
	assert.Equal(t, "core", byRelation["b"].Metadata.Module)
}

// TestCycleShapeFindingHasNoPositionWithoutSourceText covers the JSON entry point. The
// relation is still named; only the position is missing, and it is nil rather than line
// zero.
func TestCycleShapeFindingHasNoPositionWithoutSourceText(t *testing.T) {
	t.Parallel()

	model := modelFromDSL(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`)

	findings := findingsFrom(ValidateJSON(model, &EngineOptions{UseGraphValidation: true})).GetErrors()
	require.Len(t, findings, 2)

	for _, finding := range findings {
		require.NotNil(t, finding.Metadata)
		assert.Equal(t, CyclicRelation, finding.Metadata.ErrorType)
		assert.NotEmpty(t, finding.Metadata.Symbol)
		assert.Nil(t, finding.Line)
		assert.Nil(t, finding.Column)
	}
}
