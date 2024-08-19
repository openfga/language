package graph

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	language "github.com/openfga/language/pkg/go/transformer"
)

// TestGetDOTRepresentation also tests that the graph is built correctly.
//
//nolint:maintidx
func TestGetDOTRepresentation(t *testing.T) {
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
		`complex_scenario`: {
			model: `
model
  schema 1.1
type user
type employee

# yes, wildcard is a userset instead of a direct, makes it easier and tests the userset path

type directs-user
  relations
    define direct: [user]
    define direct_cond: [user with xcond]
    define direct_wild: [user:*]
    define direct_wild_cond: [user:* with xcond]

    define direct_and_direct_cond: [user, user with xcond, employee]
    define direct_and_direct_wild: [user, user:*, employee:*]
    define direct_and_direct_wild_cond: [user, user:* with xcond]
    define direct_cond_and_direct_wild: [user with xcond, user:*]
    define direct_cond_and_direct_wild_cond: [user with xcond, user:* with xcond]
    define direct_wildcard_and_direct_wildcard_cond: [user:*, user:* with xcond]

    define computed: direct
    define computed_cond: direct_cond
    define computed_wild: direct_wild
    define computed_wild_cond: direct_wild_cond
    define computed_computed: computed
    define computed_computed_computed: computed_computed
    
    define or_computed: computed or computed_cond or direct_wild
    define and_computed: computed_cond and computed_wild
    define butnot_computed: computed_wild_cond but not computed_computed

    define tuple_cycle2: [user, usersets-user#tuple_cycle2, employee]  
    define tuple_cycle3: [user, complexity3#cycle_nested]
    define compute_tuple_cycle3: tuple_cycle3

type directs-employee
  relations
    define direct: [employee]
    define computed: direct
    define direct_cond: [employee with xcond]
    define direct_wild: [employee:*]
    define direct_wild_cond: [employee:* with xcond]
    
type usersets-user
  relations
    define userset: [directs-user#direct, directs-employee#direct]
    define userset_to_computed: [directs-user#computed, directs-employee#computed]
    define userset_to_computed_cond: [directs-user#computed_cond, directs-employee#direct_cond]
    define userset_to_computed_wild: [directs-user#computed_wild, directs-employee#direct_wild]
    define userset_to_computed_wild_cond: [directs-user#direct_wild_cond, directs-employee#direct_wild_cond]

    define userset_cond: [directs-user#direct with xcond]
    define userset_cond_to_computed: [directs-user#computed with xcond]
    define userset_cond_to_computed_cond: [directs-user#computed_cond with xcond]
    define userset_cond_to_computed_wild: [directs-user#computed_wild with xcond]
    define userset_cond_to_computed_wild_cond: [directs-user#computed_wild_cond with xcond]

    define userset_to_or_computed: [directs-user#or_computed]
    define userset_to_butnot_computed: [directs-user#butnot_computed]
    define userset_to_and_computed:[directs-user#and_computed]
    
    define userset_recursive: [user, usersets-user#userset_recursive]
    
    define or_userset: userset or userset_to_computed_cond
    define and_userset: userset_to_computed_cond and userset_to_computed_wild
    define butnot_userset: userset_cond_to_computed_wild but not userset_cond
    
    define nested_or_userset: userset_to_or_computed or userset_to_butnot_computed
    define nested_and_userset: userset_to_and_computed and userset_to_or_computed

    define ttu_direct_userset: [ttus#direct_pa_direct_ch]
    define ttu_direct_cond_userset: [ttus#direct_cond_pa_direct_ch]
    define ttu_or_direct_userset: [ttus#or_comp_from_direct_parent]
    define ttu_and_direct_userset: [ttus#and_comp_from_direct_parent]
    
    define tuple_cycle2: [ttus#tuple_cycle2]
    define tuple_cycle3: [directs-user#compute_tuple_cycle3]

type ttus
  relations
    define direct_parent: [directs-user]
    define mult_parent_types: [directs-user, directs-employee]
    define mult_parent_types_cond: [directs-user with xcond, directs-employee with xcond]
    define direct_cond_parent: [directs-user with xcond]

    define userset_parent: [usersets-user]
    define userset_cond_parent: [usersets-user with xcond]

    define tuple_cycle2: tuple_cycle2 from direct_parent
    define tuple_cycle3: tuple_cycle3 from userset_parent

    define direct_pa_direct_ch: direct from mult_parent_types
    define direct_cond_pa_direct_ch: direct from mult_parent_types_cond
    define or_comp_from_direct_parent: or_computed from direct_parent
    define and_comp_from_direct_parent: and_computed from direct_cond_parent

    define userset_pa_userset_ch: userset from userset_parent
    define userset_pa_userset_comp_ch: userset_to_computed from userset_parent
    define userset_pa_userset_comp_cond_ch: userset_to_computed_cond from userset_parent
    define userset_pa_userset_comp_wild_ch: userset_to_computed_wild from userset_parent
    define userset_pa_userset_comp_wild_cond_ch: userset_to_computed_wild_cond from userset_parent

    define userset_cond_userset_ch: userset from userset_cond_parent
    define userset_cond_userset_comp_ch: userset_to_computed from userset_cond_parent
    define userset_cond_userset_comp_cond_ch: userset_to_computed_cond from userset_cond_parent
    define userset_cond_userset_comp_wild_ch: userset_to_computed_wild from userset_cond_parent
    define userset_cond_userset_comp_wild_cond_ch: userset_to_computed_wild_cond from userset_cond_parent

    define or_ttu: direct_pa_direct_ch or direct_cond_pa_direct_ch
    define and_ttu: or_comp_from_direct_parent and direct_pa_direct_ch
    define nested_butnot_ttu: or_comp_from_direct_parent but not userset_pa_userset_comp_wild_ch

type complexity3
  relations
    define ttu_parent: [ttus]
    define userset_parent: [usersets-user]
    
    define ttu_userset_ttu: ttu_direct_userset from userset_parent
    define ttu_ttu_userset: userset_pa_userset_ch from ttu_parent
    define userset_ttu_userset: [ttus#userset_pa_userset_ch]
    define userset_userset_ttu: [usersets-user#ttu_direct_userset] 
    
    define compute_ttu_userset_ttu: ttu_userset_ttu
    define compute_userset_ttu_userset: userset_ttu_userset
    
    define or_compute_complex3: compute_ttu_userset_ttu or compute_userset_ttu_userset
    define and_nested_complex3: [ttus#and_ttu] and compute_ttu_userset_ttu 

    define cycle_nested: [ttus#tuple_cycle3]   

type complexity4
  relations
    define userset_ttu_userset_ttu: [complexity3#ttu_userset_ttu]
    define ttu_ttu_ttu_userset: ttu_ttu_userset from parent
    
    define userset_or_compute_complex3: [complexity3#or_compute_complex3]
    define ttu_and_nested_complex3: and_nested_complex3 from parent
    
    define or_complex4: userset_or_compute_complex3 or ttu_and_nested_complex3
    define parent: [complexity3]
 
condition xcond(x: string) {
  x == '1'
}`, expectedOutput: ``,
		},
	}

	for name, test := range testCases {
		test := test

		t.Run(name, func(t *testing.T) {
			model := language.MustTransformDSLToProto(test.model)
			graph, err := NewAuthorizationModelGraph(model)
			require.NoError(t, err)
			fmt.Println(graph.GetAllTruthyPaths())

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
