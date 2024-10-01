package graph

import (
	"testing"

	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

func TestEdgeConditionedOn(t *testing.T) {
	t.Parallel()
	model := language.MustTransformDSLToProto(`
		model
		  schema 1.1
		type user
		type folder
		type document
			relations
				define viewer: viewer from parent
				define parent: [document, folder]`)
	graph, err := NewAuthorizationModelGraph(model)
	require.NoError(t, err)

	relation, err := graph.GetNodeByLabel("document#viewer")
	require.NoError(t, err)

	nodesFromRelation := graph.From(relation.ID())
	require.Equal(t, 1, nodesFromRelation.Len())
	nodesFromRelation.Next()
	neighbor, ok := nodesFromRelation.Node().(*AuthorizationModelNode)
	require.True(t, ok)

	require.Equal(t, relation, neighbor)

	edges := graph.Lines(relation.ID(), neighbor.ID())
	require.Equal(t, 1, edges.Len())

	edges.Next()

	edge, ok := edges.Line().(*AuthorizationModelEdge)
	require.True(t, ok)

	require.Equal(t, "document#parent", edge.ConditionedOn())
}
