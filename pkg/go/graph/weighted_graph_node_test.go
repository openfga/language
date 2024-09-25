package graph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssignWeightsToNode(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		makeNode              func() *WeightedAuthorizationModelNode
		makeOutgoingEdges     func() []*WeightedAuthorizationModelEdge
		expectedWeightsOfNode WeightMap
	}{
		`one_edge_and_not_nested`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				weights := make(WeightMap)
				weights["user"] = 1

				return []*WeightedAuthorizationModelEdge{
					{weights: weights},
				}
			},
			expectedWeightsOfNode: map[string]int{
				"user": 1,
			},
		},
		`two_edges_and_not_nested_and_different_types`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				weights1 := make(WeightMap)
				weights1["user"] = 1
				weights2 := make(WeightMap)
				weights2["group"] = 1

				return []*WeightedAuthorizationModelEdge{
					{weights: weights1},
					{weights: weights2},
				}
			},
			expectedWeightsOfNode: map[string]int{
				"user":  1,
				"group": 1,
			},
		},
		`two_edges_and_not_nested_picks_maximum`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				weights1 := make(WeightMap)
				weights1["user"] = 1
				weights2 := make(WeightMap)
				weights2["user"] = 2

				return []*WeightedAuthorizationModelEdge{
					{weights: weights1},
					{weights: weights2},
				}
			},
			expectedWeightsOfNode: map[string]int{
				"user": 2,
			},
		},
		`two_edges_and_nested`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{}, true)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				weights := make(WeightMap)
				weights["user"] = 1
				weights["group"] = 1

				return []*WeightedAuthorizationModelEdge{
					{weights: weights},
				}
			},
			expectedWeightsOfNode: map[string]int{
				"user":  math.MaxInt,
				"group": math.MaxInt,
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			node := tc.makeNode()
			node.assignWeightsToNode(tc.makeOutgoingEdges())

			require.Equal(t, tc.expectedWeightsOfNode, node.weights)
		})
	}
}

func TestVerifyIntersectionAndExclusionNodes(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		makeNode          func() *WeightedAuthorizationModelNode
		makeOutgoingEdges func() []*WeightedAuthorizationModelEdge
		expectError       bool
	}{
		`union`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{nodeType: OperatorNode, label: UnionOperator}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				return nil
			},
			expectError: false,
		},
		`intersection_good`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{nodeType: OperatorNode, label: IntersectionOperator}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				weights1 := make(WeightMap)
				weights1["user"] = 1
				weights2 := make(WeightMap)
				weights2["user"] = 2

				return []*WeightedAuthorizationModelEdge{
					{weights: weights1},
					{weights: weights2},
				}
			},
			expectError: false,
		},
		`intersection_bad`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{nodeType: OperatorNode, label: IntersectionOperator}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				weights1 := make(WeightMap)
				weights1["user"] = 1
				weights2 := make(WeightMap)
				weights2["group"] = 1

				return []*WeightedAuthorizationModelEdge{
					{weights: weights1},
					{weights: weights2},
				}
			},
			expectError: true,
		},
		`difference_good`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{nodeType: OperatorNode, label: "difference"}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				weights1 := make(WeightMap)
				weights1["user"] = 1
				weights2 := make(WeightMap)
				weights2["user"] = 2

				return []*WeightedAuthorizationModelEdge{
					{weights: weights1},
					{weights: weights2},
				}
			},
			expectError: false,
		},
		`difference_bad`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{nodeType: OperatorNode, label: "difference"}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				weights1 := make(WeightMap)
				weights1["user"] = 1
				weights2 := make(WeightMap)
				weights2["group"] = 1

				return []*WeightedAuthorizationModelEdge{
					{weights: weights1},
					{weights: weights2},
				}
			},
			expectError: true,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			node := tc.makeNode()
			node.assignWeightsToNode(tc.makeOutgoingEdges())
			err := node.verifyNodeIsValid(tc.makeOutgoingEdges())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
