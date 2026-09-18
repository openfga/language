package validation

import (
	"fmt"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/graph"
)

// graphFindings runs the public graph entry point over the model parsed from
// dsl and recovers the findings.
func graphFindings(t *testing.T, dsl string) []*Finding {
	t.Helper()

	return findingsOf(ValidateDSLWithGraph(mustParse(t, dsl), dsl))
}

// describeFindings renders findings as code, symbol, position and message, which
// is everything a caller can see.
func describeFindings(findings []*Finding) []string {
	described := []string{}

	for _, finding := range findings {
		described = append(described, fmt.Sprintf("[%s] %s%s %q",
			finding.Metadata.Kind, finding.Metadata.Symbol,
			describePosition(finding.Line, finding.Column), finding.Message))
	}

	return described
}

// TestCheckRelationEntrypoints covers the rule that reads unreachable relations
// off the graph, one case per thing the rule has to get right.
//
// The messages are asserted in full, not by code alone. They are what the shared
// corpus compares across implementations, so a rule that found the right relation
// and worded it differently is a divergence.
func TestCheckRelationEntrypoints(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl  string
		want []string
	}{
		// A relation that rewrites only itself closes through a computed edge and
		// reaches nothing, which is the loop wording.
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
		// Both wordings out of one model, which is the discriminator doing its work
		// rather than one variant happening to be right for every case.
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
		// A tupleset that reaches itself across two types, which is the shape the
		// corpus fires on rather than the single-relation one above.
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
		// A tupleset that can never be satisfied reaches a TTU edge, so it is a
		// missing entrypoint rather than a loop.
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
		// Recursion with a base case keeps its terminal types and takes weight
		// Infinite, so an empty weights map is not merely "this relation recurses".
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
		// A relation reachable only through a userset still has a terminal type
		// behind it.
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
		// The rule looks at relation nodes only. An operator node has no weights of
		// its own to speak of and must not be reported as a relation.
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

// TestGraphEntrypointFindingsMatchTheTraversal checks that the graph path reports
// the same finding as the tree walk for the same model, field for field.
//
// Without this the rule could satisfy the corpus on message text alone while
// handing callers a finding whose code, symbol or position the traversal set and
// it left different.
func TestGraphEntrypointFindingsMatchTheTraversal(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`

	fromGraph := graphFindings(t, dsl)
	require.Len(t, fromGraph, 1)

	finding := fromGraph[0]
	assert.Equal(t, RelationNoEntrypoint, finding.Metadata.Kind)

	// The same model down the traversal path, field for field. Anything the graph
	// path leaves different is a contract a caller loses by switching.
	fromTraversal := findingsOf(ValidateDSL(mustParse(t, dsl), dsl))
	require.Len(t, fromTraversal, 1)

	assert.Equal(t, describeFindings(fromTraversal), describeFindings(fromGraph))
	assert.Equal(t, fromTraversal[0].Metadata, finding.Metadata)
	assert.Equal(t, fromTraversal[0].Line, finding.Line)
	assert.Equal(t, fromTraversal[0].Column, finding.Column)
}

// TestModelUnbuildableFinding covers the one code this path adds. A model the
// builder refuses yields a single positionless finding, since Build returns no
// graph to enumerate relations from.
func TestModelUnbuildableFinding(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl         string
		wantMessage string
	}{
		"cycle through an intersection": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define admin: [user]
    define viewer: admin and editor
    define editor: viewer
`,
			wantMessage: "the model cannot be built into a weighted graph: model cycle",
		},
		"cycle through an exclusion": {
			dsl: `model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define viewer: [user] but not banned
    define banned: viewer from parent
`,
			wantMessage: "the model cannot be built into a weighted graph: tuple cycle: " +
				"operands AND or BUT NOT cannot be involved in a cycle",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			findings := graphFindings(t, test.dsl)
			require.Len(t, findings, 1, "a refused build yields one finding, there being no graph to enumerate")

			finding := findings[0]
			assert.Equal(t, GraphModelUnbuildable, finding.Metadata.Kind)
			assert.Equal(t, test.wantMessage, finding.Message)

			// The builder stops at the first problem and returns no graph with its
			// error, so there is nothing to resolve a position against.
			assert.Nil(t, finding.Line)
			assert.Nil(t, finding.Column)
		})
	}
}

// TestGraphValidationIsNotWiredIntoTheDefaultPipeline pins the switch. The graph
// path is not the source of truth yet, so the default entry points must report
// the traversal's answer and never the graph-only code.
func TestGraphValidationIsNotWiredIntoTheDefaultPipeline(t *testing.T) {
	t.Parallel()

	// A model the graph refuses but the traversal reports on relation by relation:
	// the two paths give different findings, so a default that leaked the graph
	// answer would show here.
	dsl := `model
  schema 1.1
type user
type document
  relations
    define admin: [user]
    define viewer: admin and editor
    define editor: viewer
`

	byDefault := findingsOf(ValidateDSL(mustParse(t, dsl), dsl))
	require.NotEmpty(t, byDefault)

	for _, finding := range byDefault {
		assert.NotEqual(t, GraphModelUnbuildable, finding.Metadata.Kind,
			"the default path reported a graph finding, so it is wired to the graph phase")
	}
}

// TestValidateJSONWithGraphHandlesNoSourceText covers the JSON entry point, where
// there are no lines to resolve a position against. The finding is still raised,
// with its position left nil rather than resolved to line zero.
func TestValidateJSONWithGraphHandlesNoSourceText(t *testing.T) {
	t.Parallel()

	model := mustParse(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`)

	findings := findingsOf(ValidateJSONWithGraph(model))
	require.Len(t, findings, 1)

	assert.Equal(t, RelationNoEntrypoint, findings[0].Metadata.Kind)
	assert.Equal(t, "`viewer` is an impossible relation for `document` (potential loop).", findings[0].Message)
	assert.Nil(t, findings[0].Line)
	assert.Nil(t, findings[0].Column)
}

// TestValidateWithGraphOnNilModel checks the guard, reached directly since the
// pipeline returns before its phases on a nil model.
func TestValidateWithGraphOnNilModel(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateWithGraph(nil, source{}))
	require.NoError(t, validateWithGraph(newIndex(nil), source{}))
}

// TestGraphRuleRegistryIsWellFormed keeps the registry usable as the place rules
// are added. Two rules under one id, or an entry with no function, would make a
// failure unattributable.
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

// TestSplitRelationLabel covers the label parsing, including the forms it has to
// refuse. A malformed label reaching the raise site would report a finding
// against a relation named "".
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

// TestRelationForNodeResolvesEveryNode covers node attribution, which is what
// lets a rule report against a relation without the graph exposing one.
//
// Every node except a terminal type has to resolve, and it has to resolve to a
// relation node that is in the graph. Checking only that the label splits on a
// "#" is not enough: an operator label contains one too, so returning the node's
// own label unchanged would split cleanly and name a relation that does not exist.
func TestRelationForNodeResolvesEveryNode(t *testing.T) {
	t.Parallel()

	model := mustParse(t, `model
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

// relationMetaFor reports whether the model declares a relation, by locating it
// the way the rule's own position lookup does.
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

// TestSortedNodeIDsIsStable checks the ordering the rule iterates in. The graph
// stores nodes in a map, so without this a model with two unreachable relations
// would report them in a different order per run and the corpus comparison would
// be flaky rather than wrong.
func TestSortedNodeIDsIsStable(t *testing.T) {
	t.Parallel()

	model := mustParse(t, `model
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

// TestReachesOnlyRewritesPicksTheMessageVariant covers the discriminator on its
// own, so a change to it is a failure here rather than a message that reads oddly
// in the corpus.
func TestReachesOnlyRewritesPicksTheMessageVariant(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl      string
		nodeID   string
		wantOnly bool
	}{
		"self rewrite reaches only rewrites": {
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
		// A terminal node has no outgoing edges at all, which is not the same as
		// having only rewrites and must not read as a loop.
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

			weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(mustParse(t, test.dsl))
			require.NoError(t, err)

			_, ok := weighted.GetNodeByID(test.nodeID)
			require.Truef(t, ok, "node %q is not in the graph, so this case tests nothing", test.nodeID)

			scope := &graphScope{weighted: weighted}
			assert.Equal(t, test.wantOnly, scope.reachesOnlyRewrites(test.nodeID))
		})
	}
}

// TestReachesOnlyRewritesTerminatesOnACycle checks the visited set. A relation
// that rewrites itself has an edge back to the node the walk started at, so
// without the visited set the queue would never empty.
func TestReachesOnlyRewritesTerminatesOnACycle(t *testing.T) {
	t.Parallel()

	weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(mustParse(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
`))
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

	scope := &graphScope{weighted: weighted}
	assert.True(t, scope.reachesOnlyRewrites("document#viewer"))
}

// TestGraphValidationRunsBehindTheCascadeGate pins where the phase sits. A model
// with an undefined reference is reported as that, and the build refusal it would
// also produce is not piled on top.
func TestGraphValidationRunsBehindTheCascadeGate(t *testing.T) {
	t.Parallel()

	dsl := `model
  schema 1.1
type user
type document
  relations
    define viewer: editor
`

	findings := graphFindings(t, dsl)
	require.NotEmpty(t, findings)

	for _, finding := range findings {
		assert.NotEqual(t, GraphModelUnbuildable, finding.Metadata.Kind,
			"the refusal was reported on top of the reference error that explains it")
	}
}
