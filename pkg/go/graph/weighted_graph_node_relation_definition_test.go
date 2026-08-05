package graph

import (
	"testing"

	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

// TestNodeRelationDefinition pins the relation attributed to every kind of node the builder
// creates. This is what lets a graph-derived finding name the relation it came from, and then
// be resolved to a line and column.
func TestNodeRelationDefinition(t *testing.T) {
	t.Parallel()

	model := `
	model
  schema 1.1
type user
type group
  relations
    define member: [user]
type folder
  relations
    define parent: [folder]
    define viewer: [user, group#member] or viewer from parent
type document
  relations
    define parent: [folder]
    define owner: [user]
    define editor: [user, user:*]
    define blocked: [user]
    define viewer: (owner or editor or viewer from parent) but not blocked
`

	authModel := language.MustTransformDSLToProto(model)
	graph, err := NewWeightedAuthorizationModelGraphBuilder().Build(authModel)
	require.NoError(t, err)

	t.Run("type nodes carry no relation", func(t *testing.T) {
		t.Parallel()

		for _, id := range []string{"user", "group", "folder", "document"} {
			node, ok := graph.GetNodeByID(id)
			require.True(t, ok, "expected node %s", id)
			require.Equal(t, SpecificType, node.GetNodeType())
			require.Empty(t, node.GetRelationDefinition(),
				"type node %s should not be attributed to a relation", id)
		}
	})

	t.Run("wildcard nodes carry no relation", func(t *testing.T) {
		t.Parallel()

		node, ok := graph.GetNodeByID("user:*")
		require.True(t, ok)
		require.Equal(t, SpecificTypeWildcard, node.GetNodeType())
		require.Empty(t, node.GetRelationDefinition(),
			"a wildcard is reachable from many relations, so it belongs to none")
	})

	t.Run("relation nodes are their own relation", func(t *testing.T) {
		t.Parallel()

		for _, id := range []string{
			"group#member", "folder#parent", "folder#viewer",
			"document#parent", "document#owner", "document#editor",
			"document#blocked", "document#viewer",
		} {
			node, ok := graph.GetNodeByID(id)
			require.True(t, ok, "expected node %s", id)
			require.Equal(t, SpecificTypeAndRelation, node.GetNodeType())
			require.Equal(t, id, node.GetRelationDefinition())
		}
	})

	t.Run("operator nodes name the relation whose rewrite created them", func(t *testing.T) {
		t.Parallel()

		// document#viewer is "(owner or editor or viewer from parent) but not blocked", so it
		// produces nested operators. Every one of them must attribute to document#viewer and
		// not to the enclosing operator, which is the case that makes this field load-bearing.
		var operators int

		for _, node := range graph.GetNodes() {
			if node.GetNodeType() != OperatorNode {
				continue
			}

			operators++

			require.NotEmpty(t, node.GetRelationDefinition(),
				"operator node %s must be attributed to a relation", node.GetUniqueLabel())

			objectType, relation := SplitRelationDefinition(node.GetRelationDefinition())
			require.NotEmpty(t, objectType, "node %s", node.GetUniqueLabel())
			require.NotEmpty(t, relation, "node %s", node.GetUniqueLabel())

			// The owning relation must itself be a relation node in the graph.
			owner, ok := graph.GetNodeByID(node.GetRelationDefinition())
			require.True(t, ok, "relation %s named by node %s does not exist",
				node.GetRelationDefinition(), node.GetUniqueLabel())
			require.Equal(t, SpecificTypeAndRelation, owner.GetNodeType())
		}

		require.NotZero(t, operators, "model should produce operator nodes")
	})

	t.Run("nested exclusion operators all attribute to document#viewer", func(t *testing.T) {
		t.Parallel()

		var found int

		for _, node := range graph.GetNodes() {
			if node.GetNodeType() == OperatorNode && node.GetRelationDefinition() == "document#viewer" {
				found++
			}
		}

		// One exclusion plus one union, both belonging to document#viewer.
		require.GreaterOrEqual(t, found, 2,
			"both the exclusion and the nested union should attribute to document#viewer")
	})

	t.Run("logical grouping nodes name their relation", func(t *testing.T) {
		t.Parallel()

		for _, node := range graph.GetNodes() {
			switch node.GetNodeType() {
			case LogicalDirectGrouping, LogicalTTUGrouping:
				require.NotEmpty(t, node.GetRelationDefinition(),
					"logical node %s must be attributed to a relation", node.GetUniqueLabel())

				owner, ok := graph.GetNodeByID(node.GetRelationDefinition())
				require.True(t, ok, "relation %s named by logical node %s does not exist",
					node.GetRelationDefinition(), node.GetUniqueLabel())
				require.Equal(t, SpecificTypeAndRelation, owner.GetNodeType())
			case SpecificType, SpecificTypeAndRelation, SpecificTypeWildcard, OperatorNode:
				// covered by the other subtests
			}
		}
	})
}

// TestNodeRelationDefinitionIsStable asserts the attribution does not depend on which
// relation happened to reach a shared node first in a single build.
func TestNodeRelationDefinitionIsStable(t *testing.T) {
	t.Parallel()

	model := `
	model
  schema 1.1
type user
type group
  relations
    define member: [user]
type document
  relations
    define a: [group#member]
    define b: [group#member]
`

	authModel := language.MustTransformDSLToProto(model)

	first, err := NewWeightedAuthorizationModelGraphBuilder().Build(authModel)
	require.NoError(t, err)

	second, err := NewWeightedAuthorizationModelGraphBuilder().Build(authModel)
	require.NoError(t, err)

	// group#member is referenced by both document#a and document#b, and is a relation node,
	// so it must be attributed to itself either way rather than to whichever referenced it.
	for _, graph := range []*WeightedAuthorizationModelGraph{first, second} {
		node, ok := graph.GetNodeByID("group#member")
		require.True(t, ok)
		require.Equal(t, "group#member", node.GetRelationDefinition())
	}
}

func TestSplitRelationDefinition(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		relationDefinition string
		wantObjectType     string
		wantRelation       string
	}{
		{
			name:               "type and relation",
			relationDefinition: "document#parent",
			wantObjectType:     "document",
			wantRelation:       "parent",
		},
		{
			name:               "empty is not a relation definition",
			relationDefinition: "",
			wantObjectType:     "",
			wantRelation:       "",
		},
		{
			name:               "missing separator",
			relationDefinition: "document",
			wantObjectType:     "",
			wantRelation:       "",
		},
		{
			name:               "more than one separator",
			relationDefinition: "document#parent#extra",
			wantObjectType:     "",
			wantRelation:       "",
		},
		{
			name:               "relation names may contain underscores",
			relationDefinition: "tier#assignee_sub",
			wantObjectType:     "tier",
			wantRelation:       "assignee_sub",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			objectType, relation := SplitRelationDefinition(tt.relationDefinition)
			require.Equal(t, tt.wantObjectType, objectType)
			require.Equal(t, tt.wantRelation, relation)
		})
	}
}
