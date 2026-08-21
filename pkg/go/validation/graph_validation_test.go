package validation

import (
	"errors"
	"fmt"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
	"github.com/openfga/language/pkg/go/graph"
)

// graphFindings runs validation with the graph path selected and returns the findings.
func graphFindings(t *testing.T, dsl string) []*ValidationError {
	t.Helper()

	return findingsFrom(ValidateDSL(modelFromDSL(t, dsl), dsl,
		&EngineOptions{UseGraphValidation: true})).GetErrors()
}

// describeFindings renders findings as code, symbol, position and message, which is
// everything a caller can see.
func describeFindings(findings []*ValidationError) []string {
	described := []string{}

	for _, finding := range findings {
		errorType := ValidationErrorType("")
		symbol := ""

		if finding.Metadata != nil {
			errorType = finding.Metadata.ErrorType
			symbol = finding.Metadata.Symbol
		}

		described = append(described, fmt.Sprintf("[%s] %s%s %q",
			errorType, symbol, describePosition(finding.Line, finding.Column), finding.Message))
	}

	return described
}

// TestCheckRelationEntrypoints covers the rule that reads unreachable relations off the
// graph, one case per thing the rule has to get right.
//
// The messages are asserted in full, not by code alone. They are what the shared corpus
// compares across implementations, so a rule that found the right relation and worded it
// differently is a divergence.
func TestCheckRelationEntrypoints(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl  string
		want []string
	}{
		// A relation that rewrites only itself closes through a computed edge and reaches
		// nothing, which is the loop wording.
		"relation that rewrites itself": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`,
			want: []string{
				"[relation-no-entry-point] viewer line 5-5 column 11-17 " +
					"\"`viewer` is an impossible relation for `document` (potential loop).\"",
			},
		},
		// Every unreachable relation is reported, not just the first the walk reached.
		"two relations that rewrite themselves": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
    define reader: reader
`,
			want: []string{
				"[relation-no-entry-point] reader line 6-6 column 11-17 " +
					"\"`reader` is an impossible relation for `document` (potential loop).\"",
				"[relation-no-entry-point] viewer line 5-5 column 11-17 " +
					"\"`viewer` is an impossible relation for `document` (potential loop).\"",
			},
		},
		// Both wordings out of one model, which is the discriminator doing its work rather
		// than one variant happening to be right for every case.
		"a loop and a missing entrypoint in one model": {
			dsl: `model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define viewer: viewer
    define reader: reader from parent
`,
			want: []string{
				"[relation-no-entry-point] viewer line 6-6 column 11-17 " +
					"\"`viewer` is an impossible relation for `folder` (potential loop).\"",
				"[relation-no-entry-point] reader line 7-7 column 11-17 " +
					"\"`reader` is an impossible relation for `folder` (no entrypoint).\"",
			},
		},
		// A tupleset that reaches itself across two types, which is the shape the corpus
		// fires on rather than the single-relation one above.
		"tupleset that reaches itself across two types": {
			dsl: `model
  schema 1.1
type user
type team
  relations
    define parent: [group]
    define viewer: viewer from parent
type group
  relations
    define parent: [team]
    define viewer: viewer from parent
`,
			want: []string{
				"[relation-no-entry-point] viewer line 6-6 column 11-17 " +
					"\"`viewer` is an impossible relation for `team` (no entrypoint).\"",
				"[relation-no-entry-point] viewer line 10-10 column 11-17 " +
					"\"`viewer` is an impossible relation for `group` (no entrypoint).\"",
			},
		},
		// One reachable branch of a union is enough, so the unreachable one is not a
		// finding on its own.
		"union with one satisfiable branch is not reported": {
			dsl: `model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define viewer: viewer from parent or reader from parent
    define reader: [user]
`,
			want: []string{},
		},
		// A tupleset that can never be satisfied reaches a TTU edge, so it is a missing
		// entrypoint rather than a loop.
		"tupleset with no reachable terminal type": {
			dsl: `model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define viewer: viewer from parent
`,
			want: []string{
				"[relation-no-entry-point] viewer line 6-6 column 11-17 " +
					"\"`viewer` is an impossible relation for `folder` (no entrypoint).\"",
			},
		},
		// Recursion with a base case keeps its terminal types and takes weight Infinite,
		// so an empty weights map is not merely "this relation recurses".
		"recursion with a base case is not reported": {
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
		"direct assignment is not reported": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]
`,
			want: []string{},
		},
		// A wildcard is a terminal type, so it is an entrypoint like any other.
		"wildcard is an entrypoint": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: [user:*]
`,
			want: []string{},
		},
		// A relation reachable only through a userset still has a terminal type behind it.
		"userset restriction is an entrypoint": {
			dsl: `model
  schema 1.1
type user
type group
  relations
    define member: [user]
type document
  relations
    define viewer: [group#member]
`,
			want: []string{},
		},
		// The rule looks at relation nodes only. An operator node has no weights of its
		// own to speak of and must not be reported as a relation.
		"operator over reachable children is not reported": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define admin: [user]
    define editor: [user]
    define viewer: admin or editor
`,
			want: []string{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.ElementsMatch(t, test.want, describeFindings(graphFindings(t, test.dsl)))
		})
	}
}

// TestGraphEntrypointFindingsCarryTheSameContractAsTheTraversal checks the parts of a
// finding a caller branches on, rather than reads: the code, the sentinel and the
// criticality.
//
// Without this the rule could satisfy the corpus on message text alone while handing
// callers a finding they cannot classify.
func TestGraphEntrypointFindingsCarryTheSameContractAsTheTraversal(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`

	findings := graphFindings(t, dsl)
	require.Len(t, findings, 1)

	finding := findings[0]
	require.NotNil(t, finding.Metadata)

	assert.Equal(t, RelationNoEntrypoint, finding.Metadata.ErrorType)
	require.ErrorIs(t, finding, fgaerrors.ErrNoEntrypoints)
	assert.Equal(t, fgaerrors.SeverityError, finding.Severity)
	assert.True(t, isCriticalErrorType(finding.Metadata.ErrorType))

	// The same model down the traversal path, field for field. Anything the graph path
	// leaves unset that the traversal sets is a contract the caller loses by switching.
	fromTraversal := findingsFrom(ValidateDSL(modelFromDSL(t, dsl), dsl, DefaultEngineOptions())).GetErrors()
	require.Len(t, fromTraversal, 1)
	assert.Equal(t, describeFindings(fromTraversal), describeFindings(findings))
	assert.Equal(t, fromTraversal[0].Severity, finding.Severity)
	assert.Equal(t, fromTraversal[0].Category, finding.Category)
	assert.Equal(t, fromTraversal[0].Metadata, finding.Metadata)
}

// modelRefusedWithNothingToSay is the model the positionless backstop exists for: the
// builder refuses it, the rewrite-tree walk finds nothing, and the cycle-shape check finds
// nothing either.
//
// The walk requires the computed relation on at least one assignable type of the tupleset;
// the builder requires it on every one. Here folder has viewer and team does not. No cycle
// is involved, so naming a relation would take a rule this branch does not have, and the
// refusal is all there is to report.
var modelRefusedWithNothingToSay = struct {
	dsl         string
	wantMessage string
	wantCause   error
}{
	dsl: `model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]
type team
type document
  relations
    define parent: [folder, team]
    define viewer: viewer from parent
`,
	// The sentinel's own text. The builder's message goes on to name the type and relation
	// it stopped at, which is a detail this deliberately drops: it is the first of possibly
	// several and which one it is changes between runs.
	wantMessage: "the model cannot be built into a weighted graph: invalid model",
	wantCause:   graph.ErrInvalidModel,
}

// TestRaiseModelUnbuildableReachesBothSentinels covers the code the backstop raises.
//
// A caller has two questions about a refused build, whether it was refused and why, and
// the finding has to answer both through errors.Is. Chaining is the only reason the raise
// site overrides the table's cause.
func TestRaiseModelUnbuildableReachesBothSentinels(t *testing.T) {
	t.Parallel()

	dsl := modelRefusedWithNothingToSay.dsl

	require.Empty(t, findingsFrom(ValidateDSL(modelFromDSL(t, dsl), dsl,
		DefaultEngineOptions())).GetErrors(),
		"the walk has to be silent here or the refusal is not what is being tested")

	findings := graphFindings(t, dsl)
	require.Len(t, findings, 1, "nothing else found anything, so the refusal is the only finding")

	finding := findings[0]
	require.NotNil(t, finding.Metadata)

	assert.Equal(t, GraphModelUnbuildable, finding.Metadata.ErrorType)
	assert.Equal(t, modelRefusedWithNothingToSay.wantMessage, finding.Message)

	require.ErrorIs(t, finding, fgaerrors.ErrModelNotBuildable, "the refusal itself")
	require.ErrorIs(t, finding, modelRefusedWithNothingToSay.wantCause, "the reason the builder gave")

	// The builder answers a refusal with an error and no graph, so there is nothing to
	// resolve a position against and no relation named.
	assert.Nil(t, finding.Line)
	assert.Nil(t, finding.Column)
}

// TestRaiseModelUnbuildableChainsEveryGraphSentinel covers the raise site over every
// sentinel the builder can refuse with, rather than only the one that still reaches the
// backstop end to end.
//
// The cycle-shape check answers the other two before the backstop does, so without this
// the chaining would only ever be exercised for ErrInvalidModel and a caller matching on
// either cycle sentinel would have no test behind it.
func TestRaiseModelUnbuildableChainsEveryGraphSentinel(t *testing.T) {
	t.Parallel()

	for _, sentinel := range []error{graph.ErrModelCycle, graph.ErrTupleCycle, graph.ErrInvalidModel} {
		t.Run(sentinel.Error(), func(t *testing.T) {
			t.Parallel()

			cause := fmt.Errorf("%w: from the builder", sentinel)

			collector := NewErrorCollector(nil)
			collector.RaiseModelUnbuildable(describeRefusal(cause), cause)

			findings := collector.AllFindings()
			require.Len(t, findings, 1)

			require.ErrorIs(t, findings[0], fgaerrors.ErrModelNotBuildable, "the refusal itself")
			require.ErrorIs(t, findings[0], sentinel, "the reason the builder gave")

			assert.NotContains(t, findings[0].Message, "from the builder",
				"the builder's own text reached the message, which is what varies between runs")
		})
	}
}

// TestEveryGraphSentinelHasAWording keeps graphRefusalReasons exhaustive over what the
// builder can refuse with.
//
// A sentinel missing from the table falls back to wording that says nothing about the model,
// which is a worse finding than the one it replaced. Adding a sentinel to pkg/go/graph should
// fail here rather than quietly degrade a message.
func TestEveryGraphSentinelHasAWording(t *testing.T) {
	t.Parallel()

	// Every error value pkg/go/graph declares, read off its source rather than inferred, so
	// this list going stale is a compile failure.
	declared := map[string]error{
		"ErrModelCycle":           graph.ErrModelCycle,
		"ErrInvalidModel":         graph.ErrInvalidModel,
		"ErrTupleCycle":           graph.ErrTupleCycle,
		"ErrContrainstTupleCycle": graph.ErrContrainstTupleCycle,
	}

	for name, sentinel := range declared {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			reason := describeRefusal(fmt.Errorf("%w: detail the builder added", sentinel))

			assert.NotEqual(t, unrecognisedRefusal, reason, "no wording for %s", name)
			assert.NotContains(t, reason, "detail the builder added")

			// The sentinel's own text, not that of something it wraps. Checking only
			// that the fallback was avoided would pass on ErrContrainstTupleCycle
			// resolving to the plain tuple cycle it wraps, which says less.
			assert.Equalf(t, sentinel.Error(), reason,
				"%s resolved to the wording of a less specific sentinel", name)
		})
	}

	assert.Equal(t, unrecognisedRefusal, describeRefusal(errors.New("something else entirely")),
		"an error carrying no graph sentinel has to fall through rather than match one")
}

// TestRefusalWordingIsTheSentinelNotTheBuilderMessage pins the substitution itself.
//
// The builder's message for this model names the type and relation it stopped at. That text
// is what varies between runs on a model with several problems, so it has to be absent from
// the finding and reachable through the chain.
func TestRefusalWordingIsTheSentinelNotTheBuilderMessage(t *testing.T) {
	t.Parallel()

	dsl := modelRefusedWithNothingToSay.dsl
	model := modelFromDSL(t, dsl)

	_, buildErr := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.Error(t, buildErr)
	require.Contains(t, buildErr.Error(), "team type does not have defined viewer relation",
		"the builder stopped naming a relation, so there is a detail to drop")

	findings := graphFindings(t, dsl)
	require.Len(t, findings, 1)

	assert.NotContains(t, findings[0].Message, "team type does not have defined viewer relation",
		"the builder's first-problem detail reached the finding")
	assert.Contains(t, findings[0].Message, graph.ErrInvalidModel.Error())

	// Dropped from the message, not lost. A caller that wants it unwraps.
	require.ErrorIs(t, findings[0], graph.ErrInvalidModel)
	assert.Contains(t, findings[0].Unwrap().Error(), "team type does not have defined viewer relation",
		"the builder's text has to stay reachable through the chain")
}

// TestGraphPathOutputIsStableAcrossRuns is the guarantee all of the sorting in this package
// exists for, asserted end to end rather than per pass.
//
// Whatever the builder's map iteration does, the same model has to produce the same findings
// in the same order every run, or a corpus comparison is flaky and an editor redraws
// diagnostics in a different order on every keystroke.
func TestGraphPathOutputIsStableAcrossRuns(t *testing.T) {
	t.Parallel()

	// A model with two independent problems, which is the shape the builder picks between.
	// Both tuplesets are missing the computed relation on one of their assignable types, so
	// the rewrite-tree walk is silent and the refusal is what gets reported.
	dsl := `model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]
    define editor: [user]
type team
type squad
type document
  relations
    define parent: [folder, team]
    define viewer: viewer from parent
type file
  relations
    define owner: [folder, squad]
    define editor: editor from owner
`
	model := modelFromDSL(t, dsl)

	_, buildErr := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.Error(t, buildErr, "the model has to be one the builder refuses")

	first := describeFindings(findingsFrom(
		ValidateDSL(model, dsl, &EngineOptions{UseGraphValidation: true})).GetErrors())
	require.NotEmpty(t, first)

	for range 50 {
		assert.Equal(t, first, describeFindings(findingsFrom(
			ValidateDSL(model, dsl, &EngineOptions{UseGraphValidation: true})).GetErrors()))
	}
}

// TestUnbuildableFindingDoesNotClaimTheWrongSentinel checks the chain is specific. A
// finding that matched every graph sentinel would satisfy the test above and tell a
// caller nothing.
func TestUnbuildableFindingDoesNotClaimTheWrongSentinel(t *testing.T) {
	t.Parallel()

	findings := graphFindings(t, modelRefusedWithNothingToSay.dsl)
	require.Len(t, findings, 1)

	require.ErrorIs(t, findings[0], graph.ErrInvalidModel)
	require.NotErrorIs(t, findings[0], graph.ErrModelCycle)
	require.NotErrorIs(t, findings[0], graph.ErrTupleCycle)
	require.NotErrorIs(t, findings[0], fgaerrors.ErrNoEntrypoints,
		"a refused build is not an entrypoint finding, and a caller filtering on one must not see the other")
	require.NotErrorIs(t, findings[0], fgaerrors.ErrRelationInUnresolvableCycle,
		"the refusal is not a cycle finding, and the two share this branch")
}

// TestRefusedModelKeepsThePerRelationFindingsTheWalkHas pins what a refusal costs a
// caller, which is a position and a relation name rather than a finding.
//
// The relations here are impossible as well as caught in a cycle, so the walk names them
// and the refusal adds nothing. Reporting the refusal instead would replace two findings a
// consumer can put on a line with one it cannot.
func TestRefusedModelKeepsThePerRelationFindingsTheWalkHas(t *testing.T) {
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

	assert.Equal(t, describeFindings(findingsFrom(ValidateDSL(modelFromDSL(t, dsl), dsl,
		DefaultEngineOptions())).GetErrors()), describeFindings(graphFindings(t, dsl)),
		"a refused build reports what the walk reports, down to the position")

	for _, finding := range graphFindings(t, dsl) {
		require.NotNil(t, finding.Metadata)
		assert.Equal(t, RelationNoEntrypoint, finding.Metadata.ErrorType)
		assert.NotNil(t, finding.Line, "a consumer needs somewhere to put this")
	}
}

// TestGraphValidationIsOffByDefault pins the switch. The graph path is not the source of
// truth yet, so a model that reports differently under it must report the traversal's
// answer when nothing asked for the graph.
func TestGraphValidationIsOffByDefault(t *testing.T) {
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
	model := modelFromDSL(t, dsl)

	byDefault := findingsFrom(ValidateDSL(model, dsl, DefaultEngineOptions())).GetErrors()
	explicitlyOff := findingsFrom(ValidateDSL(model, dsl, &EngineOptions{UseGraphValidation: false})).GetErrors()
	nilOptions := findingsFrom(ValidateDSL(model, dsl, nil)).GetErrors()

	assert.Equal(t, describeFindings(byDefault), describeFindings(explicitlyOff))
	assert.Equal(t, describeFindings(byDefault), describeFindings(nilOptions))

	for _, finding := range byDefault {
		require.NotNil(t, finding.Metadata)
		assert.NotEqual(t, GraphModelUnbuildable, finding.Metadata.ErrorType,
			"the default path reported a graph finding, so the switch is not doing the switching")
	}
}

// TestGraphAndTraversalNeverBothRun checks the exclusion structurally rather than by
// counting findings, which a rule with exact parity would satisfy either way.
//
// A model whose relations are unreachable and whose build succeeds reports the same
// findings down both paths, so if the engine ran both it would report each one twice.
func TestGraphAndTraversalNeverBothRun(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
    define reader: reader
`
	model := modelFromDSL(t, dsl)

	weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.NoError(t, err, "the build has to succeed, or the graph path reports one finding for another reason")
	require.NotNil(t, weighted)

	// Both paths find these two, so a union would be four.
	require.Len(t, findingsFrom(ValidateDSL(model, dsl, DefaultEngineOptions())).GetErrors(), 2)
	assert.Len(t, findingsFrom(ValidateDSL(model, dsl, &EngineOptions{UseGraphValidation: true})).GetErrors(), 2)
}

// TestValidateWithGraphHandlesNoSourceText covers the JSON entry point, where there are
// no lines to resolve a position against. The finding is still raised, with its position
// left nil rather than resolved to line zero.
func TestValidateWithGraphHandlesNoSourceText(t *testing.T) {
	t.Parallel()

	model := modelFromDSL(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`)

	findings := findingsFrom(ValidateJSON(model, &EngineOptions{UseGraphValidation: true})).GetErrors()
	require.Len(t, findings, 1)
	require.NotNil(t, findings[0].Metadata)

	assert.Equal(t, RelationNoEntrypoint, findings[0].Metadata.ErrorType)
	assert.Equal(t, "`viewer` is an impossible relation for `document` (potential loop).", findings[0].Message)
	assert.Nil(t, findings[0].Line)
	assert.Nil(t, findings[0].Column)
}

// TestValidateWithGraphOnNilModel checks the guard. RunAllValidations returns before the
// phases on a nil model, so this reaches the function directly, both with no validator at
// all and with one built over a nil model, which is the shape the engine would hold.
func TestValidateWithGraphOnNilModel(t *testing.T) {
	t.Parallel()

	collector := NewErrorCollector(nil)
	validateWithGraph(collector, nil, nil)

	assert.Equal(t, 0, collector.CountAll(), "a nil model has nothing to build and nothing to report")

	validateWithGraph(collector, NewSemanticValidator(nil), nil)

	assert.Equal(t, 0, collector.CountAll(), "a validator over a nil model has nothing to report either")
}

// TestGraphRuleRegistryIsWellFormed keeps the registry usable as the place rules are
// added. Two rules under one id, or an entry with no function, would make a failure
// unattributable.
func TestGraphRuleRegistryIsWellFormed(t *testing.T) {
	t.Parallel()

	require.NotEmpty(t, graphRules)

	seen := make(map[string]struct{}, len(graphRules))

	for _, rule := range graphRules {
		assert.NotEmpty(t, rule.id, "a rule with no id cannot be named in failure output")
		assert.NotNil(t, rule.check, "a rule with no check silently passes")

		_, duplicate := seen[rule.id]
		assert.Falsef(t, duplicate, "two rules share the id %q", rule.id)
		seen[rule.id] = struct{}{}
	}
}

// TestSplitRelationLabel covers the label parsing, including the forms it has to refuse.
// A malformed label reaching the raise site would report a finding against a relation
// named "".
func TestSplitRelationLabel(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		label          string
		wantObjectType string
		wantRelation   string
		wantOK         bool
	}{
		"type and relation":     {label: "document#viewer", wantObjectType: "document", wantRelation: "viewer", wantOK: true},
		"module qualified type": {label: "core.document#viewer", wantObjectType: "core.document", wantRelation: "viewer", wantOK: true},
		"type only":             {label: "document"},
		"no relation":           {label: "document#"},
		"no type":               {label: "#viewer"},
		"empty":                 {label: ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			objectType, relation, ok := splitRelationLabel(test.label)

			assert.Equal(t, test.wantOK, ok)
			assert.Equal(t, test.wantObjectType, objectType)
			assert.Equal(t, test.wantRelation, relation)
		})
	}
}

// TestRelationForNodeResolvesEveryNode covers node attribution, which is what lets a rule
// report against a relation without the graph exposing one.
//
// Every node except a terminal type has to resolve, and it has to resolve to a relation
// node that is in the graph. Checking only that the label splits on a "#" is not enough:
// an operator label contains one too, so returning the node's own label unchanged would
// split cleanly and name a relation that does not exist.
func TestRelationForNodeResolvesEveryNode(t *testing.T) {
	t.Parallel()

	model := modelFromDSL(t, `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member]
type folder
  relations
    define parent: [folder]
    define viewer: [user] or viewer from parent
type document
  relations
    define parent: [folder]
    define admin: [user]
    define editor: [user, group#member] or admin
    define viewer: (editor or viewer from parent) but not blocked
    define blocked: [user]
`)

	weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.NoError(t, err)

	var terminals, relations, grouped int

	for _, nodeID := range sortedNodeIDs(weighted) {
		node, ok := weighted.GetNodeByID(nodeID)
		require.True(t, ok)

		definition, ok := relationForNode(weighted, node)

		if node.GetNodeType() == graph.SpecificType || node.GetNodeType() == graph.SpecificTypeWildcard {
			assert.Falsef(t, ok, "terminal node %q resolved to relation %q", nodeID, definition)

			terminals++

			continue
		}

		require.Truef(t, ok, "node %q resolved to no relation", nodeID)

		// The resolved label has to name a relation the graph holds, and one whose
		// declared type and relation the model can be searched for.
		resolvedNode, inGraph := weighted.GetNodeByID(definition)
		require.Truef(t, inGraph, "node %q resolved to %q, which is not a node in the graph", nodeID, definition)
		assert.Equalf(t, graph.SpecificTypeAndRelation, resolvedNode.GetNodeType(),
			"node %q resolved to %q, which is not a relation node", nodeID, definition)

		objectType, relation, split := splitRelationLabel(definition)
		require.Truef(t, split, "node %q resolved to %q, which is not a relation label", nodeID, definition)
		assert.NotNilf(t, relationMetaFor(model, objectType, relation),
			"node %q resolved to %q, which the model does not declare", nodeID, definition)

		if node.GetNodeType() == graph.SpecificTypeAndRelation {
			assert.Equalf(t, nodeID, definition, "relation node %q resolved to another relation", nodeID)

			relations++
		} else {
			grouped++
		}
	}

	assert.Positive(t, terminals, "no terminal nodes, so the refusal above was never exercised")
	assert.Positive(t, relations, "no relation nodes, so resolving to self was never exercised")
	assert.Positive(t, grouped,
		"no operator or logical nodes, so reading the relation off an outgoing edge was never exercised")
}

// relationMetaFor reports whether the model declares a relation, by locating it the way
// the rule's own position lookup does.
func relationMetaFor(model *openfgav1.AuthorizationModel, objectType, relation string) *openfgav1.Userset {
	for _, typeDef := range model.GetTypeDefinitions() {
		if typeDef.GetType() != objectType {
			continue
		}

		if userset, ok := typeDef.GetRelations()[relation]; ok {
			return userset
		}
	}

	return nil
}

// TestSortedNodeIDsIsStable checks the ordering the rule iterates in. The graph stores
// nodes in a map, so without this a model with two unreachable relations would report
// them in a different order per run and the corpus comparison would be flaky rather than
// wrong.
func TestSortedNodeIDsIsStable(t *testing.T) {
	t.Parallel()

	model := modelFromDSL(t, `model
  schema 1.1
type user
type document
  relations
    define admin: [user]
    define editor: [user]
    define viewer: admin or editor
`)

	weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.NoError(t, err)

	first := sortedNodeIDs(weighted)
	require.NotEmpty(t, first)
	assert.IsIncreasing(t, first)

	for range 20 {
		assert.Equal(t, first, sortedNodeIDs(weighted))
	}
}

// TestCyclesThroughRewritesPicksTheMessageVariant covers the discriminator on its own, so a
// change to it is a failure here rather than a message that reads oddly in the corpus.
func TestCyclesThroughRewritesPicksTheMessageVariant(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl      string
		nodeID   string
		wantOnly bool
	}{
		"self rewrite cycles through rewrites": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`,
			nodeID:   "document#viewer",
			wantOnly: true,
		},
		"tupleset does not": {
			dsl: `model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define viewer: viewer from parent
`,
			nodeID:   "folder#viewer",
			wantOnly: false,
		},
		"direct assignment does not": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]
`,
			nodeID:   "document#viewer",
			wantOnly: false,
		},
		// A terminal node has no outgoing edges at all, which is not the same as having
		// only rewrites and must not read as a loop.
		"a node with no outgoing edges does not": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define viewer: [user]
`,
			nodeID:   "user",
			wantOnly: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model := modelFromDSL(t, test.dsl)

			weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
			require.NoError(t, err)

			_, ok := weighted.GetNodeByID(test.nodeID)
			require.Truef(t, ok, "node %q is not in the graph, so this case tests nothing", test.nodeID)

			scope := &graphScope{weighted: weighted, model: model}
			assert.Equal(t, test.wantOnly, scope.cyclesThroughRewrites(test.nodeID))
		})
	}
}

// TestCyclesThroughRewritesTerminatesOnACycle checks the visited set. A relation that
// rewrites itself has an edge back to the node the walk started at, so without the visited
// set the queue would never empty.
//
// A self-rewrite is the only cycle this can reach. The walk stops at the first edge that
// is not a rewrite, so a longer cycle would have to be a chain of rewrites, and the
// builder refuses every one of those with a model cycle before a rule sees it.
func TestCyclesThroughRewritesTerminatesOnACycle(t *testing.T) {
	t.Parallel()

	model := modelFromDSL(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`)

	weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.NoError(t, err)

	edges, ok := weighted.GetEdgesFromNodeID("document#viewer")
	require.True(t, ok)
	require.NotEmpty(t, edges, "the node has no outgoing edge, so there is no cycle to terminate on")

	var closesOnItself bool

	for _, edge := range edges {
		if edge.GetTo().GetUniqueLabel() == "document#viewer" {
			closesOnItself = true
		}
	}

	require.True(t, closesOnItself, "the node does not reach itself, so this case tests nothing")

	scope := &graphScope{weighted: weighted, model: model}
	assert.True(t, scope.cyclesThroughRewrites("document#viewer"))
}

// TestGraphValidationRunsBehindTheCascadeGate pins where the phase sits. A model with an
// undefined reference is reported as that, and the build refusal it would also produce is
// not piled on top.
func TestGraphValidationRunsBehindTheCascadeGate(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: editor
`
	model := modelFromDSL(t, dsl)

	_, buildErr := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.Error(t, buildErr, "this model has to be one the builder refuses for the test to mean anything")

	findings := graphFindings(t, dsl)
	require.NotEmpty(t, findings)

	for _, finding := range findings {
		require.NotNil(t, finding.Metadata)
		assert.NotEqual(t, GraphModelUnbuildable, finding.Metadata.ErrorType,
			"the refusal was reported on top of the reference error that explains it")
	}
}

// TestGraphValidationHonoursSkipSemanticValidation checks the switch does not smuggle the
// phase past the skip it belongs to.
func TestGraphValidationHonoursSkipSemanticValidation(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`

	err := ValidateDSL(modelFromDSL(t, dsl), dsl, &EngineOptions{
		UseGraphValidation:     true,
		SkipSemanticValidation: true,
	})

	assert.NoError(t, err, "the semantic phase was skipped, so the graph rule must not have run")
}

// TestUnbuildableFindingIsRecoverableThroughTheEntryPoints checks the finding survives the
// route a consumer actually uses, errors.As on what ValidateDSL returned.
func TestUnbuildableFindingIsRecoverableThroughTheEntryPoints(t *testing.T) {
	t.Parallel()

	dsl := modelRefusedWithNothingToSay.dsl

	err := ValidateDSL(modelFromDSL(t, dsl), dsl, &EngineOptions{UseGraphValidation: true})
	require.Error(t, err)

	var collection *ValidationErrors
	require.ErrorAs(t, err, &collection)
	require.Len(t, collection.GetErrors(), 1)

	require.ErrorIs(t, err, fgaerrors.ErrModelNotBuildable, "the collection carries the sentinel through")
	require.ErrorIs(t, err, graph.ErrInvalidModel)
}
