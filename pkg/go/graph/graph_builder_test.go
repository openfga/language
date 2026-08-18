package graph

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"

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
		model            string
		cycleInformation CycleInformation
		expectedOutput   string // can visualize in https://dreampuf.github.io/GraphvizOnline
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
			// NOTE: For now, we will not consider this case as cycle. We may want to reevaluate in the future.
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
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: true,
				canHaveCyclesAtRuntime: false,
			},
		},
		`computed_relation_with_size_two`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define x: y
						define y: x`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#x"];
2 [label="folder#y"];

// Edge definitions.
1 -> 2 [style=dashed];
2 -> 1 [style=dashed];
}`,
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: true,
				canHaveCyclesAtRuntime: false,
			},
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
		`tuple_to_userset_conditional`: {
			model: `
				model
					schema 1.1
				type user
				type document
					relations
						define parentt: [folder, folder with condX]
						define viewer: admin from parentt
				type folder
					relations
						define admin: [user]
				condition condX (x:int) {
					x > 0
				}`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=document];
1 [label="document#parentt"];
2 [label=folder];
3 [label="document#viewer"];
4 [label="folder#admin"];
5 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
4 -> 3 [headlabel="(document#parentt)"];
5 -> 4 [label=direct];
}`,
		},
		`userset_conditional`: {
			model: `
				model
					schema 1.1
				type folder
					relations
						define viewer: [group#member, group#member with condX]
				type group
					relations
						define member: [user]
				type user
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
2 [label="group#member"];
3 [label=group];
4 [label=user];

// Edge definitions.
2 -> 1 [label=direct];
4 -> 2 [label=direct];
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
3 -> 2;
4 -> 3 [label=direct];
}`,
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: false,
				canHaveCyclesAtRuntime: true,
			},
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
1 -> 5;
2 -> 1 [label=direct];
2 -> 3 [label=direct];
3 -> 5;
5 -> 4;
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
1 -> 4;
2 -> 1 [label=direct];
2 -> 4 [label=direct];
4 -> 3;
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
1 -> 8;
2 -> 1 [label=direct];
2 -> 3 [label=direct];
2 -> 4 [label=direct];
2 -> 5 [label=direct];
3 -> 8;
4 -> 8;
5 -> 7;
7 -> 6;
8 -> 7;
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
1 -> 5;
2 -> 1 [label=direct];
2 -> 3 [label=direct];
3 -> 5;
5 -> 4;
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
1 -> 4;
2 -> 1 [label=direct];
2 -> 4 [label=direct];
4 -> 3;
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
1 -> 8;
2 -> 1 [label=direct];
2 -> 3 [label=direct];
2 -> 4 [label=direct];
2 -> 5 [label=direct];
3 -> 8;
4 -> 8;
5 -> 7;
7 -> 6;
8 -> 7;
}`,
		},
		`unionRuntimeRecursive`: {
			model: `
				model
					schema 1.1
				type user
				type folder
				   relations
					 define a: [user] or b
					 define b: [user] or c
					 define c: [user] or a`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=union];
3 [label=user];
4 [label="folder#b"];
5 [label=union];
6 [label="folder#c"];
7 [label=union];

// Edge definitions.
1 -> 7;
2 -> 1;
3 -> 2 [label=direct];
3 -> 5 [label=direct];
3 -> 7 [label=direct];
4 -> 2;
5 -> 4;
6 -> 5;
7 -> 6;
}`,
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: false,
				canHaveCyclesAtRuntime: true,
			},
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
5 -> 4;
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
1 -> 5;
2 -> 1 [label=direct];
2 -> 3 [label=direct];
3 -> 5;
5 -> 4;
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
1 -> 4;
2 -> 1 [label=direct];
2 -> 4 [label=direct];
4 -> 3;
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
1 -> 8;
2 -> 1 [label=direct];
2 -> 3 [label=direct];
2 -> 4 [label=direct];
2 -> 5 [label=direct];
3 -> 8;
4 -> 8;
5 -> 7;
7 -> 6;
8 -> 7;
}`,
		},
		`multiple_cycles`: {
			model: `
				model
					schema 1.1
				type user
				type folder
					relations
						define x: y
						define y: x
						define a: [user] or x or b
						define b: [user] or c
						define c: [user] or a`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=folder];
1 [label="folder#a"];
2 [label=union];
3 [label=user];
4 [label="folder#x"];
5 [label="folder#b"];
6 [label=union];
7 [label="folder#c"];
8 [label=union];
9 [label="folder#y"];

// Edge definitions.
1 -> 8;
2 -> 1;
3 -> 2 [label=direct];
3 -> 6 [label=direct];
3 -> 8 [label=direct];
4 -> 2;
4 -> 9 [style=dashed];
5 -> 2;
6 -> 5;
7 -> 6;
8 -> 7;
9 -> 4 [style=dashed];
}`,
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: true,
				canHaveCyclesAtRuntime: true,
			},
		},
		`potential_cycle_or_but_not`: {
			model: `
				model
      schema 1.1
    type user
    type resource
      relations
        define x: [user] but not y
        define y: [user] but not z
        define z: [user] or x`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=resource];
1 [label="resource#x"];
2 [label=exclusion];
3 [label=user];
4 [label="resource#y"];
5 [label=exclusion];
6 [label="resource#z"];
7 [label=union];

// Edge definitions.
1 -> 7;
2 -> 1;
3 -> 2 [label=direct];
3 -> 5 [label=direct];
3 -> 7 [label=direct];
4 -> 2;
5 -> 4;
6 -> 5;
7 -> 6;
}`,
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: false,
				canHaveCyclesAtRuntime: true,
			},
		},
		`potential_cycle_four_union`: {
			model: `
				model
      schema 1.1
    type user
    type group
      relations
        define member: [user] or memberA or memberB or memberC
        define memberA: [user] or member or memberB or memberC
        define memberB: [user] or member or memberA or memberC
        define memberC: [user] or member or memberA or memberB`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=group];
1 [label="group#member"];
2 [label=union];
3 [label=user];
4 [label="group#memberA"];
5 [label="group#memberB"];
6 [label="group#memberC"];
7 [label=union];
8 [label=union];
9 [label=union];

// Edge definitions.
1 -> 7;
1 -> 8;
1 -> 9;
2 -> 1;
3 -> 2 [label=direct];
3 -> 7 [label=direct];
3 -> 8 [label=direct];
3 -> 9 [label=direct];
4 -> 2;
4 -> 8;
4 -> 9;
5 -> 2;
5 -> 7;
5 -> 9;
6 -> 2;
6 -> 7;
6 -> 8;
7 -> 4;
8 -> 5;
9 -> 6;
}`,
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: false,
				canHaveCyclesAtRuntime: true,
			},
		},
		`potential_cycle_four_union_with_one_member_no_union`: {
			model: `
				model
      schema 1.1
    type user
    type account
      relations
        define admin: [user] or member or super_admin or owner
        define member: [user] or owner or admin or super_admin
        define super_admin: [user] or admin or member or owner
        define owner: [user]`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=account];
1 [label="account#admin"];
2 [label=union];
3 [label=user];
4 [label="account#member"];
5 [label="account#super_admin"];
6 [label="account#owner"];
7 [label=union];
8 [label=union];

// Edge definitions.
1 -> 7;
1 -> 8;
2 -> 1;
3 -> 2 [label=direct];
3 -> 6 [label=direct];
3 -> 7 [label=direct];
3 -> 8 [label=direct];
4 -> 2;
4 -> 8;
5 -> 2;
5 -> 7;
6 -> 2;
6 -> 7;
6 -> 8;
7 -> 4;
8 -> 5;
}`,
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: false,
				canHaveCyclesAtRuntime: true,
			},
		},
		`intersection`: {
			model: `
				model
      schema 1.1
    type user

    type document
      relations
        define admin: [user]
        define action1: admin and action2 and action3
        define action2: admin and action1 and action3
        define action3: admin and action1 and action2`,
			expectedOutput: `digraph {
graph [
rankdir=BT
];

// Node definitions.
0 [label=document];
1 [label="document#action1"];
2 [label=intersection];
3 [label="document#admin"];
4 [label="document#action2"];
5 [label="document#action3"];
6 [label=intersection];
7 [label=intersection];
8 [label=user];

// Edge definitions.
1 -> 6;
1 -> 7;
2 -> 1;
3 -> 2;
3 -> 6;
3 -> 7;
4 -> 2;
4 -> 7;
5 -> 2;
5 -> 6;
6 -> 4;
7 -> 5;
8 -> 3 [label=direct];
}`,
			cycleInformation: CycleInformation{
				hasCyclesAtCompileTime: false,
				canHaveCyclesAtRuntime: true,
			},
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
			require.Equal(t, testCase.cycleInformation, graph.GetCycles())

			// assert that dot(reverse(reverse(graph)) == dot(graph)
			reverseGraph1, err := graph.Reversed()
			require.NoError(t, err)
			reverseGraph2, err := reverseGraph1.Reversed()
			require.NoError(t, err)
			reversedDot := reverseGraph2.GetDOT()
			reversedSorted := getSorted(reversedDot)

			diff = cmp.Diff(expectedSorted, reversedSorted)

			require.Empty(t, diff, "expected %s\ngot\n%s", testCase.expectedOutput, reversedDot)
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

			// assert that dot(reverse(reverse(graph)) == dot(graph)
			reverseGraph1, err := graph.Reversed()
			require.NoError(t, err)
			reverseGraph2, err := reverseGraph1.Reversed()
			require.NoError(t, err)
			reversedDot := reverseGraph2.GetDOT()
			reversedSorted := getSorted(reversedDot)

			diff = cmp.Diff(expectedSorted, reversedSorted)

			require.Empty(t, diff, "expected %s\ngot\n%s", test.expectedOutput, reversedDot)
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

// TestParseThisMalformedRelationReference tests that parseThis handles degenerate
// RelationReference shapes that survive proto validation but cannot be expressed in the DSL.
// These tests use hand-constructed protos because language.MustTransformDSLToProto cannot
// produce an empty relation name (the grammar requires at least one character after #).
func TestParseThisMalformedRelationReference(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		model                 *openfgav1.AuthorizationModel
		expectPanic           bool
		expectedUserNode      bool     // whether a "user" node should exist
		expectedGroupNode     bool     // whether a "group" node should exist
		expectedUserEdgeCond  []string // conditions on user -> document#viewer edge
		expectedGroupEdgeCond []string // conditions on group -> document#viewer edge
	}{
		`empty_relation_in_oneof_does_not_panic`: {
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: "1.1",
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{Type: "user"},
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"viewer": {Userset: &openfgav1.Userset_This{}},
						},
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{
											Type: "user",
											RelationOrWildcard: &openfgav1.RelationReference_Relation{
												Relation: "",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectPanic:           true, // Should not panic; parseThis must handle all oneof states
			expectedUserNode:      true,
			expectedGroupNode:     false,
			expectedUserEdgeCond:  []string{""},
			expectedGroupEdgeCond: nil,
		},
		`nil_wildcard_in_oneof_does_not_panic`: {
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: "1.1",
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{Type: "user"},
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"viewer": {Userset: &openfgav1.Userset_This{}},
						},
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{
											Type: "user",
											RelationOrWildcard: &openfgav1.RelationReference_Wildcard{
												Wildcard: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectPanic:           true, // Should not panic; parseThis must handle all oneof states
			expectedUserNode:      true,
			expectedGroupNode:     false,
			expectedUserEdgeCond:  []string{""},
			expectedGroupEdgeCond: nil,
		},
		`degenerate_entry_does_not_inherit_previous_node`: {
			// This is the silent graph-corruption variant: when a valid entry precedes
			// a degenerate one, curNode is declared outside the loop so the degenerate
			// entry reuses the previous iteration's node. The edge carrying condition "cond1"
			// should originate from "group", not "user".
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: "1.1",
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{Type: "user"},
					{
						Type: "group",
						Relations: map[string]*openfgav1.Userset{
							"member": {Userset: &openfgav1.Userset_This{}},
						},
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"member": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{Type: "user"},
									},
								},
							},
						},
					},
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"viewer": {Userset: &openfgav1.Userset_This{}},
						},
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{Type: "user"}, // valid entry first
										{
											Type:      "group",
											Condition: "cond1",
											RelationOrWildcard: &openfgav1.RelationReference_Relation{
												Relation: "", // degenerate entry second
											},
										},
									},
								},
							},
						},
					},
				},
				Conditions: map[string]*openfgav1.Condition{
					"cond1": {
						Name:       "cond1",
						Expression: "x == 1",
						Parameters: map[string]*openfgav1.ConditionParamTypeRef{
							"x": {TypeName: openfgav1.ConditionParamTypeRef_TYPE_NAME_INT},
						},
					},
				},
			},
			expectPanic:      false, // no panic, but curNode must not leak across loop iterations
			expectedUserNode: true,
			expectedGroupNode: true, // group node must exist; cond1 must be on group edge not user
			expectedUserEdgeCond: []string{""}, // user edge should have no condition
			expectedGroupEdgeCond: []string{"cond1"}, // group edge must have cond1; not merged onto user
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if testCase.expectPanic {
				// Test verifies parseThis handles degenerate RelationReference without panic
				require.NotPanics(t, func() {
					graph, err := NewAuthorizationModelGraph(testCase.model)
					require.NoError(t, err)
					require.NotNil(t, graph)
				})
			} else {
				graph, err := NewAuthorizationModelGraph(testCase.model)
				require.NoError(t, err)
				require.NotNil(t, graph)

				// Verify the expected nodes exist
				userNode, _ := graph.GetNodeByLabel("user")
				groupNode, _ := graph.GetNodeByLabel("group")
				viewerNode, err := graph.GetNodeByLabel("document#viewer")
				require.NoError(t, err, "document#viewer node should exist")

				if testCase.expectedUserNode {
					require.NotNil(t, userNode, "user node should exist")
				} else {
					require.Nil(t, userNode, "user node should not exist")
				}

				if testCase.expectedGroupNode {
					require.NotNil(t, groupNode, "group node should exist")
				} else {
					require.Nil(t, groupNode, "group node should not exist")
				}

				require.NotNil(t, viewerNode, "document#viewer node should exist")

				// Verify edge conditions
				if testCase.expectedUserNode && len(testCase.expectedUserEdgeCond) > 0 {
					userToViewer := findEdge(t, graph, userNode, viewerNode)
					require.NotNil(t, userToViewer, "edge from user to document#viewer should exist")
					require.ElementsMatch(t, testCase.expectedUserEdgeCond, userToViewer.conditions,
						"user -> document#viewer edge should have expected conditions")
				}

				if testCase.expectedGroupNode && len(testCase.expectedGroupEdgeCond) > 0 {
					groupToViewer := findEdge(t, graph, groupNode, viewerNode)
					require.NotNil(t, groupToViewer, "edge from group to document#viewer should exist")
					require.ElementsMatch(t, testCase.expectedGroupEdgeCond, groupToViewer.conditions,
						"group -> document#viewer edge should have expected conditions")
				}
			}
		})
	}
}

// findEdge returns the edge from fromNode to toNode, or nil if it doesn't exist.
func findEdge(t *testing.T, graph *AuthorizationModelGraph, fromNode, toNode *AuthorizationModelNode) *AuthorizationModelEdge {
	t.Helper()
	if fromNode == nil || toNode == nil {
		return nil
	}
	iter := graph.Lines(fromNode.ID(), toNode.ID())
	for iter.Next() {
		if edge, ok := iter.Line().(*AuthorizationModelEdge); ok {
			return edge
		}
	}
	return nil
}

// TestIsNilNode verifies that isNilNode correctly detects both nil and typed-nil *AuthorizationModelNode.
func TestIsNilNode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		node     graph.Node
		expected bool
	}{
		{
			name:     "nil interface",
			node:     nil,
			expected: true,
		},
		{
			name: "typed-nil *AuthorizationModelNode",
			node: func() graph.Node {
				var n *AuthorizationModelNode
				return n
			}(),
			expected: true,
		},
		{
			name: "real *AuthorizationModelNode",
			node: &AuthorizationModelNode{
				Node:        &simpleNode{id: 1},
				label:       "test",
				nodeType:    SpecificType,
				uniqueLabel: "test",
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := isNilNode(testCase.node)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

// TestUpsertEdgeWithTypedNil verifies that upsertEdge with a typed-nil from is a no-op instead of a panic.
func TestUpsertEdgeWithTypedNil(t *testing.T) {
	t.Parallel()

	graphBuilder := &AuthorizationModelGraphBuilder{
		multi.NewDirectedGraph(), map[string]int64{},
	}

	validNode := graphBuilder.getOrAddNode("user", "user", SpecificType)

	var nilNode *AuthorizationModelNode
	var nilNodeAsInterface graph.Node = nilNode

	// This should not panic, just return early
	require.NotPanics(t, func() {
		graphBuilder.upsertEdge(nilNodeAsInterface, validNode, DirectEdge, "", "")
	})

	// Verify no edge was created
	require.False(t, graphBuilder.hasEdge(nilNodeAsInterface, validNode, DirectEdge, ""))
}

// simpleNode is a minimal graph.Node implementation for testing.
type simpleNode struct {
	id int64
}

func (n *simpleNode) ID() int64 {
	return n.id
}
