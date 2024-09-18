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
1 [label="folder#viewer - weights:[user=1]"];
2 [label=user];

// Edge definitions.
1 -> 2 [label="direct - weights:[user=1]"];
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
1 [label="folder#viewer - weights:[folder#viewer=+∞,user=1]"];
2 [label=user];

// Edge definitions.
1 -> 1 [label="direct - weights:[folder#viewer=+∞]"];
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
			expectedOutput: ``,
			expectedError:  ErrInvalidModel,
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
			expectedOutput: ``,
			expectedError:  ErrInvalidModel,
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
