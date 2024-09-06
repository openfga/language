package graph

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

type graphTestCase struct {
	model          string
	expectedOutput string
}

// TestGetDOTRepresentation also tests that the graph is built correctly.
//
//nolint:maintidx
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
		`direct_assignment_with_usersets_recursive`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define viewer: [user,folder#viewer]
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
1 -> 1 [label=direct];
2 -> 1 [label=direct];
}`,
		},
		`direct_assignment_with_conditions`: { // conditions are not represented and edges are de-duped
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define viewer: [user with condX, user]
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
		`tuple_to_userset_recursive`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define parent: [folder]
						define viewer: [user] or viewer from parent`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#parent"];
2 [label="folder#viewer"];
3 [label=union];
4 [label=user];

// Edge definitions.
0 -> 1 [label=direct];
2 -> 3 [headlabel="(folder#parent)"];
3 -> 2 [style=dashed];
4 -> 3 [label=direct];
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
		`intersection_of_relations`: {
			model: `
				model
					schema 1.1
				type user
				type folder
				   relations
					 define a: [user]
					 define b: [user]
					 define c: a and b`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label="folder#c"];
5 [label=intersection];

// Edge definitions.
1 -> 5 [style=dashed];
2 -> 1 [label=direct];
2 -> 3 [label=direct];
3 -> 5 [style=dashed];
5 -> 4 [style=dashed];
}`,
		},
		`intersection_of_relation_and_type`: {
			model: `
				model
					schema 1.1
				type user
				type folder
				   relations
					 define a: [user]
					 define b: [user] and a`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label=intersection];

// Edge definitions.
1 -> 4 [style=dashed];
2 -> 1 [label=direct];
2 -> 4 [label=direct];
4 -> 3 [style=dashed];
}`,
		},
		`intersection_with_parens`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define a: [user]
						define b: [user]
						define c: [user]
						define d: [user]
						define e: (a and b and c) and d`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label="folder#c"];
5 [label="folder#d"];
6 [label="folder#e"];
7 [label=intersection];
8 [label=intersection];

// Edge definitions.
1 -> 8 [style=dashed];
2 -> 1 [label=direct];
2 -> 3 [label=direct];
2 -> 4 [label=direct];
2 -> 5 [label=direct];
3 -> 8 [style=dashed];
4 -> 8 [style=dashed];
5 -> 7 [style=dashed];
7 -> 6 [style=dashed];
8 -> 7 [style=dashed];
}`,
		},
		`union_of_relations`: {
			model: `
				model
					schema 1.1
				type user
				type folder
				   relations
					 define a: [user]
					 define b: [user]
					 define c: a or b`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label="folder#c"];
5 [label=union];

// Edge definitions.
1 -> 5 [style=dashed];
2 -> 1 [label=direct];
2 -> 3 [label=direct];
3 -> 5 [style=dashed];
5 -> 4 [style=dashed];
}`,
		},
		`union_of_relation_and_type`: {
			model: `
				model
					schema 1.1
				type user
				type folder
				   relations
					 define a: [user]
					 define b: [user] or a`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label=union];

// Edge definitions.
1 -> 4 [style=dashed];
2 -> 1 [label=direct];
2 -> 4 [label=direct];
4 -> 3 [style=dashed];
}`,
		},
		`union_with_parens`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define a: [user]
						define b: [user]
						define c: [user]
						define d: [user]
						define e: (a or b or c) or d`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label="folder#c"];
5 [label="folder#d"];
6 [label="folder#e"];
7 [label=union];
8 [label=union];

// Edge definitions.
1 -> 8 [style=dashed];
2 -> 1 [label=direct];
2 -> 3 [label=direct];
2 -> 4 [label=direct];
2 -> 5 [label=direct];
3 -> 8 [style=dashed];
4 -> 8 [style=dashed];
5 -> 7 [style=dashed];
7 -> 6 [style=dashed];
8 -> 7 [style=dashed];
}`,
		},
		`multigraph`: {
			model: `
				model
				  schema 1.1
				
				type user
				
				type state
				  relations
					define can_view: [user]
				
				type transition
				  relations
					define start: [state]
					define end: [state]
					define can_apply: [user] and can_view from start and can_view from end`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=state];
1 [label="state#can_view"];
2 [label=user];
3 [label=transition];
4 [label="transition#can_apply"];
5 [label=intersection];
6 [label="transition#end"];
7 [label="transition#start"];

// Edge definitions.
0 -> 6 [label=direct];
0 -> 7 [label=direct];
1 -> 5 [headlabel="(transition#start)"];
1 -> 5 [headlabel="(transition#end)"];
2 -> 1 [label=direct];
2 -> 5 [label=direct];
5 -> 4 [style=dashed];
}`,
		},
		`exclusion_of_relations`: {
			model: `
				model
					schema 1.1
				type user
				type folder
				   relations
					 define a: [user]
					 define b: [user]
					 define c: a but not b`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label="folder#c"];
5 [label=exclusion];

// Edge definitions.
1 -> 5 [style=dashed];
2 -> 1 [label=direct];
2 -> 3 [label=direct];
3 -> 5 [style=dashed];
5 -> 4 [style=dashed];
}`,
		},
		`exclusion_of_relation_and_type`: {
			model: `
				model
					schema 1.1
				type user
				type folder
				   relations
					 define a: [user]
					 define b: [user] but not a`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label=exclusion];

// Edge definitions.
1 -> 4 [style=dashed];
2 -> 1 [label=direct];
2 -> 4 [label=direct];
4 -> 3 [style=dashed];
}`,
		},
		`exclusion_with_parens`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define a: [user]
						define b: [user]
						define c: [user]
						define d: [user]
						define e: (a or b or c) but not d`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=user];
3 [label="folder#b"];
4 [label="folder#c"];
5 [label="folder#d"];
6 [label="folder#e"];
7 [label=exclusion];
8 [label=union];

// Edge definitions.
1 -> 8 [style=dashed];
2 -> 1 [label=direct];
2 -> 3 [label=direct];
2 -> 4 [label=direct];
2 -> 5 [label=direct];
3 -> 8 [style=dashed];
4 -> 8 [style=dashed];
5 -> 7 [style=dashed];
7 -> 6 [style=dashed];
8 -> 7 [style=dashed];
}`,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model := language.MustTransformDSLToProto(testCase.model)
			graph, err := NewAuthorizationModelGraph(model)
			require.NoError(t, err)

			actualDOT := graph.GetDOT()
			actualSorted := getSorted(actualDOT)
			expectedSorted := getSorted(testCase.expectedOutput)

			diff := cmp.Diff(expectedSorted, actualSorted)

			require.Empty(t, diff, "expected %s\ngot\n%s", testCase.expectedOutput, actualDOT)
		})
	}
}

func TestGetDOTRepresentation_2(t *testing.T) {
	t.Parallel()
	rootFolder := "../../../tests/data/transformer"
	testCases := make(map[string]*graphTestCase)

	err := filepath.Walk(rootFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		fileName := info.Name()
		extension := filepath.Ext(fileName)
		pathWithoutExtension := strings.TrimSuffix(path, extension)

		if extension == ".json" {
			return nil
		}

		if _, ok := testCases[pathWithoutExtension]; !ok {
			testCases[pathWithoutExtension] = &graphTestCase{}
		}

		content, err := os.ReadFile(path)
		if err != nil {
			//nolint:wrapcheck
			return err
		}

		if extension == ".fga" {
			testCases[pathWithoutExtension].model = string(content)
		} else if extension == ".dot" {
			testCases[pathWithoutExtension].expectedOutput = string(content)
		}

		return nil
	})

	require.NoError(t, err)
	require.NotEmpty(t, testCases)

	for testname, test := range testCases {
		t.Run(testname, func(t *testing.T) {
			t.Parallel()
			if test.model == "" {
				t.Skip("empty")
			}
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
