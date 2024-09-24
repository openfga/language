package graph

import (
	"testing"

	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

func TestGetDrawingDirectionOfWeightedGraph(t *testing.T) {
	t.Parallel()
	model := language.MustTransformDSLToProto(`
				model
					schema 1.1
				type user
				type company
					relations
						define member: [user]`)
	graph, err := NewWeightedAuthorizationModelGraph(model)
	require.NoError(t, err)
	require.Equal(t, DrawingDirectionCheck, graph.GetDrawingDirection())
}
