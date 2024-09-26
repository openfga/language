package graph

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

func TestWeightedGraphBuilder(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		model          string
		expectedOutput string // can visualize in https://dreampuf.github.io/GraphvizOnline
		expectedError  error
	}{
		`direct_assignment_to_one_type`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define viewer: [user]`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer - weights:[user=1]"];
2 [label=user];

// Edge definitions.
1 -> 2 [label="direct - weights:[user=1]"];
}`,
		},
		`direct_assignment_to_one_userset`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define editor: [user]
						define viewer: [folder#editor]`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=folder];
1 [label="folder#editor - weights:[user=1]"];
2 [label=user];
3 [label="folder#viewer - weights:[user=2]"];

// Edge definitions.
1 -> 2 [label="direct - weights:[user=1]"];
3 -> 1 [label="direct - weights:[user=2]"];
}`,
		},
		`direct_assignment_to_multiple_types`: {
			model: `
				model
					schema 1.1
				type user1
				type user2
				type folder
					relations
						define viewer: [user1, user2]`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer - weights:[user1=1,user2=1]"];
2 [label=user1];
3 [label=user2];

// Edge definitions.
1 -> 2 [label="direct - weights:[user1=1]"];
1 -> 3 [label="direct - weights:[user2=1]"];
}`,
		},
		`direct_assignment_to_type_and_wildcard`: {
			model: `
				model
				  schema 1.1
				
				type user
				
				type folder
				  relations
					define viewer: [user,user:*]
	`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer - weights:[user=1]"];
2 [label=user];
3 [label="user:*"];

// Edge definitions.
1 -> 2 [label="direct - weights:[user=1]"];
1 -> 3 [label=direct];
}`,
		},
		`computed_rewrite`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define viewer: [user]
						define rewrite: viewer`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=folder];
1 [label="folder#rewrite - weights:[user=1]"];
2 [label="folder#viewer - weights:[user=1]"];
3 [label=user];

// Edge definitions.
1 -> 2 [
style=dashed
label="weights:[user=1]"
];
2 -> 3 [label="direct - weights:[user=1]"];
}`,
		},
		`nested_userset`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define viewer: [user, folder#viewer]`,

			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer - weights:[user=+∞]"];
2 [label=user];

// Edge definitions.
1 -> 1 [label="direct - weights:[user=+∞]"];
1 -> 2 [label="direct - weights:[user=1]"];
}`,
		},
		`simple_ttu`: {
			model: `
				model
					schema 1.1
				type user
				type company
					relations
						define approved_member: [user]
				type license
					relations
						define member: approved_member from owner
						define owner: [company]`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=company];
1 [label="company#approved_member - weights:[user=1]"];
2 [label=user];
3 [label=license];
4 [label="license#member - weights:[user=2]"];
5 [label="license#owner - weights:[company=1]"];

// Edge definitions.
1 -> 2 [label="direct - weights:[user=1]"];
4 -> 1 [
headlabel="(license#owner)"
label="weights:[user=2]"
];
5 -> 0 [label="direct - weights:[company=1]"];
}`,
		},
		`complex_ttu`: {
			model: `
				model
				  schema 1.1
				
				type user
				
				type folder
				  relations
					define viewer: approved from org
					define org: [org]
				
				type org
				  relations
					define approved: direct or viewer1
					define direct: [user]
					define viewer1: viewer2 from parent
					define parent: [org]
					define viewer2: viewer3 from parent
					define viewer3: [user]`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=folder];
1 [label="folder#org - weights:[org=1]"];
2 [label=org];
3 [label="folder#viewer - weights:[user=4]"];
4 [label="org#approved - weights:[user=3]"];
5 [label="union - weights:[user=3]"];
6 [label="org#direct - weights:[user=1]"];
7 [label="org#viewer1 - weights:[user=3]"];
8 [label=user];
9 [label="org#parent - weights:[org=1]"];
10 [label="org#viewer2 - weights:[user=2]"];
11 [label="org#viewer3 - weights:[user=1]"];

// Edge definitions.
1 -> 2 [label="direct - weights:[org=1]"];
3 -> 4 [
headlabel="(folder#org)"
label="weights:[user=4]"
];
4 -> 5 [label="weights:[user=3]"];
5 -> 6 [label="weights:[user=1]"];
5 -> 7 [label="weights:[user=3]"];
6 -> 8 [label="direct - weights:[user=1]"];
7 -> 10 [
headlabel="(org#parent)"
label="weights:[user=3]"
];
9 -> 2 [label="direct - weights:[org=1]"];
10 -> 11 [
headlabel="(org#parent)"
label="weights:[user=2]"
];
11 -> 8 [label="direct - weights:[user=1]"];
}`,
		},
		`multigraph`: {
			model: `
				model
					schema 1.1
				type user
				type state
					relations
						define can_view: [user] or member
						define member: [user]
				type transition
					relations
						define start: [state]
						define end: [state]
						define can_apply: [user] and can_view from start and can_view from end
				type group
					relations
						define owner: [user, transition#can_apply]
						define max_owner: [group#owner, group#max_owner]`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=group];
1 [label="group#max_owner - weights:[user=+∞]"];
2 [label="group#owner - weights:[user=3]"];
3 [label=user];
4 [label="transition#can_apply - weights:[user=2]"];
5 [label=state];
6 [label="state#can_view - weights:[user=1]"];
7 [label="union - weights:[user=1]"];
8 [label="state#member - weights:[user=1]"];
9 [label=transition];
10 [label="intersection - weights:[user=2]"];
11 [label="transition#end - weights:[state=1]"];
12 [label="transition#start - weights:[state=1]"];

// Edge definitions.
1 -> 1 [label="direct - weights:[user=+∞]"];
1 -> 2 [label="direct - weights:[user=4]"];
2 -> 3 [label="direct - weights:[user=1]"];
2 -> 4 [label="direct - weights:[user=3]"];
4 -> 10 [label="weights:[user=2]"];
6 -> 7 [label="weights:[user=1]"];
7 -> 3 [label="direct - weights:[user=1]"];
7 -> 8 [label="weights:[user=1]"];
8 -> 3 [label="direct - weights:[user=1]"];
10 -> 3 [label="direct - weights:[user=1]"];
10 -> 6 [
headlabel="(transition#start)"
label="weights:[user=2]"
];
10 -> 6 [
headlabel="(transition#end)"
label="weights:[user=2]"
];
11 -> 5 [label="direct - weights:[state=1]"];
12 -> 5 [label="direct - weights:[state=1]"];
}`,
		},
		`error_if_intersection_leads_to_two_types`: {
			model: `
				model
					schema 1.1
				type user1
				type user2
				type folder
					relations
						define viewer: a and b
						define a: [user1]
						define b: [user2]`,
			expectedError: ErrInvalidModel,
		},
		`error_if_difference_leads_to_two_types`: {
			model: `
				model
					schema 1.1
				type user1
				type user2
				type folder
					relations
						define viewer: a but not b
						define a: [user1]
						define b: [user2]`,
			expectedError: ErrInvalidModel,
		},
		`no_error_if_difference_leads_to_at_least_one_type`: {
			model: `
				model
					schema 1.1
				type user1
				type user2
				type folder
					relations
						define viewer: a but not b
						define a: [user1, user2]
						define b: [user2]`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=folder];
1 [label="folder#a - weights:[user1=1,user2=1]"];
2 [label=user1];
3 [label=user2];
4 [label="folder#b - weights:[user2=1]"];
5 [label="folder#viewer - weights:[user1=1,user2=1]"];
6 [label="exclusion - weights:[user1=1,user2=1]"];

// Edge definitions.
1 -> 2 [label="direct - weights:[user1=1]"];
1 -> 3 [label="direct - weights:[user2=1]"];
4 -> 3 [label="direct - weights:[user2=1]"];
5 -> 6 [label="weights:[user1=1,user2=1]"];
6 -> 1 [label="weights:[user1=1,user2=1]"];
6 -> 4 [label="weights:[user2=1]"];
}`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model := language.MustTransformDSLToProto(testCase.model)
			weightedGraph, err := NewWeightedAuthorizationModelGraph(model)
			if testCase.expectedError != nil {
				require.ErrorIs(t, err, testCase.expectedError)
			} else {
				require.NoError(t, err)

				actualDOT := weightedGraph.GetDOT()
				actualSorted := getSorted(actualDOT)
				expectedSorted := getSorted(testCase.expectedOutput)

				diff := cmp.Diff(expectedSorted, actualSorted)

				require.Empty(t, diff, "expected %s\ngot\n%s", testCase.expectedOutput, actualDOT)
			}
		})
	}
}

func TestGetOutgoingEdges(t *testing.T) {
	t.Parallel()
	// use a model that is a multigraph.
	model := language.MustTransformDSLToProto(`model
  schema 1.1
type user
type state
  relations
      define can_view: [user]
type transition
  relations
      define start: [state]
      define end: [state]
      define can_apply: [user] and can_view from start and can_view from end`)

	weightedGraphBuilder := NewWeightedAuthorizationModelGraphBuilder()
	_, err := weightedGraphBuilder.Build(model)
	require.NoError(t, err)

	iterNodes := weightedGraphBuilder.Nodes()
	for iterNodes.Next() {
		node, ok := iterNodes.Node().(*WeightedAuthorizationModelNode)
		require.True(t, ok)
		if node.nodeType != OperatorNode {
			continue
		}

		edges, err := weightedGraphBuilder.getOutgoingEdges(node)
		require.NoError(t, err)
		require.Len(t, edges, 3)

		var targets []string
		for _, edge := range edges {
			toNode, ok := edge.To().(*WeightedAuthorizationModelNode)
			require.True(t, ok)
			targets = append(targets, toNode.AuthorizationModelNode.label)
		}

		sort.Strings(targets)
		require.Equal(t, "state#can_view", targets[0])
		require.Equal(t, "state#can_view", targets[1])
		require.Equal(t, "user", targets[2])
	}
}

func TestAssignWeightsToLoopEdges(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		makeNode             func() *WeightedAuthorizationModelNode
		makeOutgoingEdges    func() []*WeightedAuthorizationModelEdge
		assertWeightsOfEdges func(edges []*WeightedAuthorizationModelEdge)
	}{
		`node_is_not_nested_does_nothing`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				return NewWeightedAuthorizationModelNode(&AuthorizationModelNode{}, false)
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				return []*WeightedAuthorizationModelEdge{
					{weights: make(WeightMap)},
				}
			},
			assertWeightsOfEdges: func(edges []*WeightedAuthorizationModelEdge) {
				require.Empty(t, edges[0].weights)
			},
		},

		`assigns_weight_on_nested_edges`: {
			makeNode: func() *WeightedAuthorizationModelNode {
				node := NewWeightedAuthorizationModelNode(&AuthorizationModelNode{}, true)
				node.weights["user"] = 1

				return node
			},
			makeOutgoingEdges: func() []*WeightedAuthorizationModelEdge {
				return []*WeightedAuthorizationModelEdge{
					{weights: make(WeightMap), isLoop: true},
					{weights: make(WeightMap), isLoop: false},
				}
			},
			assertWeightsOfEdges: func(edges []*WeightedAuthorizationModelEdge) {
				expectedMap := make(WeightMap)
				expectedMap["user"] = 1
				require.Equal(t, expectedMap, edges[0].weights)
				require.Empty(t, edges[1].weights)
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			edges := tc.makeOutgoingEdges()
			assignWeightsToLoopEdges(tc.makeNode(), edges)
			tc.assertWeightsOfEdges(edges)
		})
	}
}
