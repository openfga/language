package graph

import (
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

// TestGetDOTRepresentation also tests that the graph is built correctly.
func TestGetDOTRepresentation(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		model          string
		expectedOutput string // can visualize in https://dreampuf.github.io/GraphvizOnline
	}{
		`direct_assignment`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define viewer: [user]
				type user`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer"];
2 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
}`,
		},
		`direct_assignment_with_wildcard`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define viewer: [user:*]
				type user`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer"];
2 [label="user:*"];
3 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
}`,
		},
		`direct_assignment_with_wildcard_and_type`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define viewer: [user:*, user]
				type user`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer"];
2 [label="user:*"];
3 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
3 -> 1 [label=direct];
}`,
		},
		`direct_assignment_with_usersets`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define viewer: [group#member]
				type group
					relations
						define member: [user]
				type user`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer"];
2 [label="group#member"];
3 [label=group];
4 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
4 -> 2 [label=direct];
}`,
		},
		`direct_assignment_with_conditions`: { // conditions are not represented
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define viewer: [user with condX]
				condition condX (x:int) {
					x > 0
				}`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#viewer"];
2 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
}`,
		},
		`computed_relation`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define x: y
						define y: [user]
				type user`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#x"];
2 [label="folder#y"];
3 [label=user];

// Edge definitions.
2 -> 1 [style=dashed];
3 -> 2 [label=direct];
}`,
		},
		`computed_relation_with_cycle`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define x: y
						define y: z
						define z: x`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#x"];
2 [label="folder#y"];
3 [label="folder#z"];

// Edge definitions.
1 -> 3 [style=dashed];
2 -> 1 [style=dashed];
3 -> 2 [style=dashed];
}`,
		},
		`tuple_to_userset_one_related_type`: {
			model: `
				model
					schema 1.1
				type user
				type document
					relations
						define parent: [folder]
						define viewer: admin from parent
				type folder
					relations
						define admin: [user]`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=document];
1 [label="document#parent"];
2 [label=folder];
3 [label="document#viewer"];
4 [label="folder#admin"];
5 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
4 -> 3 [headlabel="(document#parent)"];
5 -> 4 [label=direct];
}`,
		},
		`tuple_to_userset_two_related_types`: {
			model: `
				model
					schema 1.1
				type user
				type document
					relations
						define parent: [folder, folder2]
						define viewer: admin from parent
				type folder
					relations
						define admin: [user]
				type folder2
					relations
						define admin: [user]`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=document];
1 [label="document#parent"];
2 [label=folder];
3 [label=folder2];
4 [label="document#viewer"];
5 [label="folder#admin"];
6 [label="folder2#admin"];
7 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
3 -> 1 [label=direct];
5 -> 4 [headlabel="(document#parent)"];
6 -> 4 [headlabel="(document#parent)"];
7 -> 5 [label=direct];
7 -> 6 [label=direct];
}`,
		},
		`tuple_to_userset_one_related_type_the_other_not`: {
			model: `
				model
					schema 1.1
				type user
				type document
					relations
						define parent: [folder, folder2]
						define viewer: admin from parent
				type folder
					relations
						define admin: [user]
				type folder2`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=document];
1 [label="document#parent"];
2 [label=folder];
3 [label=folder2];
4 [label="document#viewer"];
5 [label="folder#admin"];
6 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
3 -> 1 [label=direct];
5 -> 4 [headlabel="(document#parent)"];
6 -> 5 [label=direct];
}`,
		},
	}

	for name, test := range testCases {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model := language.MustTransformDSLToProto(test.model)
			graph, err := NewAuthorizationModelGraph(model)
			require.NoError(t, err)

			actualDOT := graph.GetDOT()
			actualSorted := getSorted(actualDOT)
			expectedSorted := getSorted(test.expectedOutput)

			diff := cmp.Diff(expectedSorted, actualSorted)

			require.Empty(t, diff, "expected %s\ngot\n%s", test.expectedOutput, actualDOT)
		})
	}
}

// getSorted assumes the input has multiple lines and returns the sorted version of it.
func getSorted(input string) string {
	lines := strings.FieldsFunc(input, func(r rune) bool {
		return r == '\n'
	})

	sort.Strings(lines)

	return strings.Join(lines, "\n")
}
