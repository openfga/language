package graph

import (
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
	}{
		`direct_assignment`: {
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
rankdir=BT
];

// Node definitions.
1 [label="folder#viewer - weights:[user1=1,user2=1]"];
2 [label=user1];
3 [label=user2];

// Edge definitions.
2 -> 1 [label=direct];
3 -> 1 [label=direct];
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
rankdir=BT
];

// Node definitions.
1 [label="folder#rewrite - weights:[user=1]"];
2 [label="folder#viewer - weights:[user=1]"];
3 [label=user];

// Edge definitions.
2 -> 1 [style=dashed];
3 -> 2 [label=direct];
}`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model := language.MustTransformDSLToProto(testCase.model)
			weightedGraph, err := NewWeightedAuthorizationModelGraph(model)
			require.NoError(t, err)

			actualDOT := weightedGraph.GetDOT()
			actualSorted := getSorted(actualDOT)
			expectedSorted := getSorted(testCase.expectedOutput)

			diff := cmp.Diff(expectedSorted, actualSorted)

			require.Empty(t, diff, "expected %s\ngot\n%s", testCase.expectedOutput, actualDOT)
		})
	}
}
