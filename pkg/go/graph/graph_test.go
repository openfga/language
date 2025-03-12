package graph

import (
	"fmt"
	"strconv"
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
			reversedGraph, err := graph.Reversed()
			require.NoError(t, err)
			actualDOT := reversedGraph.GetDOT()
			actualSorted := getSorted(actualDOT)
			expectedSorted := getSorted(testCase.expectedOutput)

			diff := cmp.Diff(expectedSorted, actualSorted)

			require.Empty(t, diff, "expected %s\ngot\n%s", testCase.expectedOutput, actualDOT)
		})
	}
}

func TestGetDrawingDirection(t *testing.T) {
	t.Parallel()
	model := language.MustTransformDSLToProto(`
				model
					schema 1.1
				type user
				type company
					relations
						define member: [user]`)
	graph, err := NewAuthorizationModelGraph(model)
	require.NoError(t, err)
	require.Equal(t, DrawingDirectionListObjects, graph.GetDrawingDirection())
	reversedGraph, err := graph.Reversed()
	require.NoError(t, err)
	require.Equal(t, DrawingDirectionCheck, reversedGraph.GetDrawingDirection())
}

func TestGetNodeByLabel(t *testing.T) {
	t.Parallel()
	model := language.MustTransformDSLToProto(`
				model
					schema 1.1
				type user
				type company
					relations
						define member: [user with cond, user:* with cond]
						define owner: [user]
						define approved_member: member or owner
				type group
					relations
						define approved_member: [user]
				type license
					relations
						define active_member: approved_member from owner
						define owner: [company, group]`)
	graph, err := NewAuthorizationModelGraph(model)
	require.NoError(t, err)

	testCases := []struct {
		label         string
		expectedFound bool
	}{
		// found
		{"user", true},
		{"user:*", true},
		{"company", true},
		{"company#member", true},
		{"company#owner", true},
		{"company#approved_member", true},
		{"group", true},
		{"group#approved_member", true},
		{"license", true},
		{"license#active_member", true},
		{"license#owner", true},
		// not found
		{"unknown", false},
		{"unknown#unknown", false},
		{"user with cond", false},
		{"user:* with cond", false},
	}
	for i, testCase := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			node, err := graph.GetNodeByLabel(testCase.label)
			if testCase.expectedFound {
				require.NoError(t, err)
				require.NotNil(t, node)
			} else {
				require.ErrorIs(t, err, ErrQueryingGraph)
				require.Nil(t, node)
			}
		})
	}
}

func TestGetNodeTypes(t *testing.T) {
	t.Parallel()
	model := language.MustTransformDSLToProto(`
				model
					schema 1.1
				type user
				type group
					relations
						define member: [user]
				type company
					relations
						define wildcard: [user:*]
						define direct: [user]
						define userset: [group#member]
						define intersectionRelation: wildcard and direct
						define unionRelation: wildcard or direct
						define differenceRelation: wildcard but not direct`)
	graph, err := NewAuthorizationModelGraph(model)
	require.NoError(t, err)

	testCases := []struct {
		label            string
		expectedNodeType NodeType
	}{
		{"user", SpecificType},
		{"user:*", SpecificTypeWildcard},
		{"group", SpecificType},
		{"group#member", SpecificTypeAndRelation},
		{"company", SpecificType},
		{"company#wildcard", SpecificTypeAndRelation},
		{"company#direct", SpecificTypeAndRelation},
		{"company#userset", SpecificTypeAndRelation},
		{"company#intersectionRelation", SpecificTypeAndRelation},
		{"company#unionRelation", SpecificTypeAndRelation},
		{"company#differenceRelation", SpecificTypeAndRelation},
	}
	for _, testCase := range testCases {
		t.Run(testCase.label, func(t *testing.T) {
			t.Parallel()
			node, err := graph.GetNodeByLabel(testCase.label)
			require.NoError(t, err)
			require.NotNil(t, node)
			require.Equal(t, testCase.expectedNodeType, node.NodeType(), "expected node type %d but got %d", testCase.expectedNodeType, node.NodeType())
		})
	}

	// testing the operator nodes is not so straightforward...
	var unionNodes, differenceNodes, intersectionNodes []*AuthorizationModelNode

	iterNodes := graph.Nodes()
	for iterNodes.Next() {
		node, ok := iterNodes.Node().(*AuthorizationModelNode)
		require.True(t, ok)
		if node.nodeType != OperatorNode {
			continue
		}

		switch node.label {
		case UnionOperator:
			unionNodes = append(unionNodes, node)
		case IntersectionOperator:
			intersectionNodes = append(intersectionNodes, node)
		case ExclusionOperator:
			differenceNodes = append(differenceNodes, node)
		}
	}

	require.Len(t, unionNodes, 1)
	require.Len(t, differenceNodes, 1)
	require.Len(t, intersectionNodes, 1)
}

func TestPathExists(t *testing.T) {
	type pathTest struct {
		fromLabel   string
		toLabel     string
		expectPath  bool
		expectedErr error
	}
	tests := []struct {
		name      string
		model     string
		pathTests []pathTest
	}{
		{
			name: "userset_computed_userset",
			model: `
model
	schema 1.1
type other
type user
type wild
type employee
type group
	relations
		define rootMember: [user, user:*, employee, wild:*]
		define member: rootMember
type folder
	relations
		define viewer: [group#member]
`,
			pathTests: []pathTest{
				{
					fromLabel:  "user",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					fromLabel:  "group",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					fromLabel:   "foo",
					toLabel:     "group#member",
					expectPath:  false,
					expectedErr: ErrQueryingGraph,
				},
				{
					fromLabel:   "user",
					toLabel:     "group#undefined",
					expectPath:  false,
					expectedErr: ErrQueryingGraph,
				},
				{
					// TODO: ideally this should be false.  However, for now, this
					// will return true as there is some path.
					fromLabel:  "group#member",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "group#member",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "group#rootMember",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "user",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "folder#viewer",
					expectPath: false,
				},
			},
		},
		{
			name: "nested_computed_userset",
			model: `
model
	schema 1.1
type other
type user
type employee
type wild
type group
	relations
		define member: [user, user:*, employee, wild:*, group#member]
type folder
	relations
		define viewer: [group#member]
`,
			pathTests: []pathTest{
				{
					fromLabel:  "user",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					fromLabel:  "group",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					fromLabel:  "group#member",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "group#member",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "user",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "folder#viewer",
					expectPath: false,
				},
			},
		},
		{
			name: "union_relation",
			model: `
model
	schema 1.1
type other
type user
type employee
type wild
type group
	relations
		define child1: [user, user:*]
		define child2: [employee]
		define child3: [wild:*]
		define member: child1 or child2 or child3
type folder
	relations
		define viewer: [group#member]
`,
			pathTests: []pathTest{
				{
					fromLabel:  "user",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					fromLabel:  "group",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					// TODO: ideally this should be false.  However, for now, this
					// will return true as there is some path.
					fromLabel:  "group#member",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "group#member",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "user",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "folder#viewer",
					expectPath: false,
				},
			},
		},
		{
			name: "intersection_relation",
			model: `
model
	schema 1.1
type other
type user
type employee
type wild
type group
	relations
		define child1: [user]
		define child2: [user, employee, wild:*]
		define member: child1 and child2
type folder
	relations
		define viewer: [group#member]
`,
			pathTests: []pathTest{
				{
					fromLabel:  "user",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					// Ideally, we will reject employee because type needs to appear
					// in both child for an intersection. For now, the graph
					// package is not smart enough to handle intersection as a special case.
					fromLabel:  "employee",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					// Ideally, we will reject wild because type needs to appear
					// in both child for an intersection. For now, the graph
					// package is not smart enough to handle intersection as a special case.
					fromLabel:  "wild:*",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					fromLabel:  "group",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					// TODO: ideally this should be false.  However, for now, this
					// will return true as there is some path.
					fromLabel:  "group#member",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "group#member",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "user",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "folder#viewer",
					expectPath: false,
				},
			},
		},
		{
			name: "exclusion_relation",
			model: `
model
	schema 1.1
type other
type user
type employee
type group
	relations
		define child1: [user]
		define child2: [user, employee]
		define member: child1 but not child2
type folder
	relations
		define viewer: [group#member]
`,
			pathTests: []pathTest{
				{
					fromLabel:  "user",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					// Ideally, we will reject employee because type needs to appear
					// in both child for exclusion. For now, the graph
					// package is not smart enough to handle exclusion as a special case.
					fromLabel:  "employee",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					fromLabel:  "group",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					// TODO: ideally this should be false.  However, for now, this
					// will return true as there is some path.
					fromLabel:  "group#member",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "group#member",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "user",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "folder#viewer",
					expectPath: false,
				},
			},
		},
		{
			name: "ttu",
			model: `
model
	schema 1.1
type other
type user
type employee
type wild
type group
	relations
		define rootMember: [user, user:*, employee, wild:*]
		define member: rootMember
type folder
	relations
		define parent: [group]
		define viewer: member from parent
`,
			pathTests: []pathTest{
				{
					fromLabel:  "user",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					fromLabel:  "group",
					toLabel:    "group#member",
					expectPath: false,
				},
				{
					// TODO: ideally this should be false.  However, for now, this
					// will return true as there is some path.
					fromLabel:  "group#member",
					toLabel:    "group#member",
					expectPath: true,
				},
				{
					fromLabel:  "group#member",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "group#rootMember",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "user",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "employee",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "wild:*",
					toLabel:    "folder#viewer",
					expectPath: true,
				},
				{
					fromLabel:  "other",
					toLabel:    "folder#viewer",
					expectPath: false,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			model := language.MustTransformDSLToProto(test.model)
			graph, err := NewAuthorizationModelGraph(model)
			require.NoError(t, err)

			for _, test := range test.pathTests {
				t.Run(fmt.Sprintf("%s -> %s", test.fromLabel, test.toLabel), func(t *testing.T) {
					actual, err := graph.PathExists(test.fromLabel, test.toLabel)
					if test.expectedErr == nil {
						require.NoError(t, err)
						require.Equal(t, test.expectPath, actual)
					} else {
						require.ErrorIs(t, err, test.expectedErr)
					}
				})
			}
		})
	}
}
