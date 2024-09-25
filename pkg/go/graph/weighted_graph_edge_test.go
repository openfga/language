package graph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/graph/multi"
)

func TestAssignWeightsToEdge(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		makeEdge              func() *WeightedAuthorizationModelEdge
		expectedWeightsOfEdge WeightMap
	}{
		`direct_edge_adds_one`: {
			makeEdge: func() *WeightedAuthorizationModelEdge {
				from := &AuthorizationModelNode{}
				weightedFromNode := NewWeightedAuthorizationModelNode(from, false)

				to := &AuthorizationModelNode{}
				weightedToNode := NewWeightedAuthorizationModelNode(to, false)
				weightedToNode.weights["group"] = 1

				edge := &AuthorizationModelEdge{
					Line:     multi.Line{F: weightedFromNode, T: weightedToNode},
					edgeType: DirectEdge,
				}

				return NewWeightedAuthorizationModelEdge(edge, false)
			},
			expectedWeightsOfEdge: map[string]int{
				"group": 2,
			},
		},
		`ttu_edge_adds_one`: {
			makeEdge: func() *WeightedAuthorizationModelEdge {
				from := &AuthorizationModelNode{}
				weightedFromNode := NewWeightedAuthorizationModelNode(from, false)

				to := &AuthorizationModelNode{}
				weightedToNode := NewWeightedAuthorizationModelNode(to, false)
				weightedToNode.weights["group"] = 1

				edge := &AuthorizationModelEdge{
					Line:     multi.Line{F: weightedFromNode, T: weightedToNode},
					edgeType: TTUEdge,
				}

				return NewWeightedAuthorizationModelEdge(edge, false)
			},
			expectedWeightsOfEdge: map[string]int{
				"group": 2,
			},
		},
		`direct_edge_infinity`: {
			makeEdge: func() *WeightedAuthorizationModelEdge {
				from := &AuthorizationModelNode{}
				weightedFromNode := NewWeightedAuthorizationModelNode(from, false)

				to := &AuthorizationModelNode{}
				weightedToNode := NewWeightedAuthorizationModelNode(to, false)
				weightedToNode.weights["group"] = math.MaxInt

				edge := &AuthorizationModelEdge{
					Line:     multi.Line{F: weightedFromNode, T: weightedToNode},
					edgeType: DirectEdge,
				}

				return NewWeightedAuthorizationModelEdge(edge, false)
			},
			expectedWeightsOfEdge: map[string]int{
				"group": math.MaxInt,
			},
		},
		`ttu_edge_infinity`: {
			makeEdge: func() *WeightedAuthorizationModelEdge {
				from := &AuthorizationModelNode{}
				weightedFromNode := NewWeightedAuthorizationModelNode(from, false)

				to := &AuthorizationModelNode{}
				weightedToNode := NewWeightedAuthorizationModelNode(to, false)
				weightedToNode.weights["group"] = math.MaxInt

				edge := &AuthorizationModelEdge{
					Line:     multi.Line{F: weightedFromNode, T: weightedToNode},
					edgeType: TTUEdge,
				}

				return NewWeightedAuthorizationModelEdge(edge, false)
			},
			expectedWeightsOfEdge: map[string]int{
				"group": math.MaxInt,
			},
		},
		`ttu_edge_infinity_multiple`: {
			makeEdge: func() *WeightedAuthorizationModelEdge {
				from := &AuthorizationModelNode{}
				weightedFromNode := NewWeightedAuthorizationModelNode(from, false)

				to := &AuthorizationModelNode{}
				weightedToNode := NewWeightedAuthorizationModelNode(to, false)
				weightedToNode.weights["group"] = math.MaxInt
				weightedToNode.weights["user"] = math.MaxInt

				edge := &AuthorizationModelEdge{
					Line:     multi.Line{F: weightedFromNode, T: weightedToNode},
					edgeType: TTUEdge,
				}

				return NewWeightedAuthorizationModelEdge(edge, false)
			},
			expectedWeightsOfEdge: map[string]int{
				"group": math.MaxInt,
				"user":  math.MaxInt,
			},
		},
		`computed_edge_equals`: {
			makeEdge: func() *WeightedAuthorizationModelEdge {
				from := &AuthorizationModelNode{}
				weightedFromNode := NewWeightedAuthorizationModelNode(from, false)

				to := &AuthorizationModelNode{}
				weightedToNode := NewWeightedAuthorizationModelNode(to, false)
				weightedToNode.weights["group"] = 1
				weightedToNode.weights["user"] = math.MaxInt

				edge := &AuthorizationModelEdge{
					Line:     multi.Line{F: weightedFromNode, T: weightedToNode},
					edgeType: ComputedEdge,
				}

				return NewWeightedAuthorizationModelEdge(edge, false)
			},
			expectedWeightsOfEdge: map[string]int{
				"group": 1,
				"user":  math.MaxInt,
			},
		},
		`rewrite_edge_equals`: {
			makeEdge: func() *WeightedAuthorizationModelEdge {
				from := &AuthorizationModelNode{}
				weightedFromNode := NewWeightedAuthorizationModelNode(from, false)

				to := &AuthorizationModelNode{}
				weightedToNode := NewWeightedAuthorizationModelNode(to, false)
				weightedToNode.weights["group"] = 1
				weightedToNode.weights["user"] = math.MaxInt

				edge := &AuthorizationModelEdge{
					Line:     multi.Line{F: weightedFromNode, T: weightedToNode},
					edgeType: RewriteEdge,
				}

				return NewWeightedAuthorizationModelEdge(edge, false)
			},
			expectedWeightsOfEdge: map[string]int{
				"group": 1,
				"user":  math.MaxInt,
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			edge := tc.makeEdge()
			err := edge.assignWeightsToEdge()
			require.NoError(t, err)

			require.Equal(t, tc.expectedWeightsOfEdge, edge.weights)
		})
	}
}
