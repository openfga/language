package graph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

func TestReverseGraph(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		model          string
		expectedOutput string // can visualize in https://dreampuf.github.io/GraphvizOnline
	}{
		`direct_assignment`: {
			model: `
				model
					schema 1.1
				type user
				type company
					relations
						define member: [user]
						define owner: [user]
						define approved_member: member or owner
				type group
					relations
						define approved_member: [user]
				type license
					relations
						define active_member: approved_member from owner
						define owner: [company, group]`,
			expectedOutput: `digraph {
graph [
rankdir=TB
];

// Node definitions.
0 [label=company];
1 [label="company#approved_member"];
2 [label=union];
3 [label="company#member"];
4 [label="company#owner"];
5 [label=user];
6 [label=group];
7 [label="group#approved_member"];
8 [label=license];
9 [label="license#active_member"];
10 [label="license#owner"];

// Edge definitions.
1 -> 2;
2 -> 3;
2 -> 4;
3 -> 5 [label=direct];
4 -> 5 [label=direct];
7 -> 5 [label=direct];
9 -> 1 [headlabel="(license#owner)"];
9 -> 7 [headlabel="(license#owner)"];
10 -> 0 [label=direct];
10 -> 6 [label=direct];
}`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model := language.MustTransformDSLToProto(testCase.model)
			graph, err := NewAuthorizationModelGraph(model)
			require.NoError(t, err)
			require.Equal(t, DrawingDirectionListObjects, graph.drawingDirection)
			reversedGraph, err := graph.Reversed()
			require.NoError(t, err)
			require.Equal(t, DrawingDirectionCheck, reversedGraph.drawingDirection)
			actualDOT := reversedGraph.GetDOT()
			actualSorted := getSorted(actualDOT)
			expectedSorted := getSorted(testCase.expectedOutput)

			diff := cmp.Diff(expectedSorted, actualSorted)

			require.Empty(t, diff, "expected %s\ngot\n%s", testCase.expectedOutput, actualDOT)
		})
	}
}
