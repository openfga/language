package graph_test

import (
	"slices"
	"testing"

	"github.com/openfga/language/pkg/go/graph"
	"github.com/openfga/language/pkg/go/transformer"

	"github.com/stretchr/testify/require"
)

const large string = `
	model
	  schema 1.1

	type t1

	type t2
	  relations
	    define rel_a: [t1] or rel_m or rel_a from rel_p
	    define rel_b: [t1] or rel_m or rel_b from rel_p
	    define rel_c: rel_t
	    define rel_d: rel_n
	    define rel_e: rel_t
	    define rel_f: rel_b
	    define rel_g: rel_b
	    define rel_h: rel_b
	    define rel_i: rel_n
	    define rel_j: rel_n or rel_b or rel_s or rel_a
	    define rel_k: rel_n or rel_b or rel_s or rel_a
	    define rel_l: rel_t
	    define rel_m: [t1] or rel_m from rel_p
	    define rel_n: [t1] or rel_m or rel_n from rel_p
	    define rel_o: [t1]
	    define rel_p: [t2]
	    define rel_q: [t1] or rel_m or rel_q from rel_p
	    define rel_r: [t1] or rel_m or rel_r from rel_p
	    define rel_s: [t1] or rel_m or rel_s from rel_p
	    define rel_t: [t1] or rel_m or rel_t from rel_p

	type t3
	  relations
	    define rel_a: ([t1] or rel_m or rel_a from rel_p) but not rel_1
	    define rel_b: ([t1] or rel_m or rel_b from rel_p) but not rel_1
	    define rel_1: [t1:*]
	    define rel_2: rel_a or rel_b or rel_s
	    define rel_3: rel_a or rel_s
	    define rel_4: rel_s
	    define rel_5: rel_a or rel_b or rel_s or rel_n
	    define rel_6: rel_a or rel_b or rel_s or rel_n or rel_q or rel_r or rel_t
	    define rel_7: rel_a
	    define rel_8: rel_a
	    define rel_9: rel_a
	    define rel_10: rel_a
	    define rel_11: rel_s
	    define rel_12: rel_s
	    define rel_c: rel_t
	    define rel_13: rel_b
	    define rel_14: rel_s
	    define rel_15: rel_s
	    define rel_16: rel_m
	    define rel_17: rel_m
	    define rel_18: rel_s
	    define rel_19: rel_m
	    define rel_20: rel_s
	    define rel_21: rel_s
	    define rel_d: rel_n
	    define rel_e: rel_t
	    define rel_22: rel_m
	    define rel_23: rel_s
	    define rel_24: rel_s
	    define rel_f: rel_b
	    define rel_25: rel_m
	    define rel_26: rel_m
	    define rel_27: rel_a or rel_s
	    define rel_28: rel_a
	    define rel_29: rel_a
	    define rel_30: rel_a
	    define rel_31: rel_s
	    define rel_32: rel_m
	    define rel_33: rel_s
	    define rel_34: rel_s
	    define rel_35: rel_m
	    define rel_36: rel_s
	    define rel_37: rel_s
	    define rel_38: rel_s
	    define rel_39: rel_s
	    define rel_40: rel_s
	    define rel_41: rel_m
	    define rel_42: rel_s
	    define rel_43: rel_m
	    define rel_44: rel_s or rel_a
	    define rel_45: rel_a
	    define rel_46: rel_a
	    define rel_47: rel_a or rel_b or rel_s or rel_n or rel_q or rel_r or rel_t
	    define rel_48: rel_a
	    define rel_49: rel_a or rel_b
	    define rel_50: rel_a
	    define rel_51: rel_m
	    define rel_52: rel_a
	    define rel_53: rel_s
	    define rel_54: rel_a
	    define rel_55: rel_a or rel_n or rel_s or rel_b
	    define rel_56: rel_b or rel_n or rel_a
	    define rel_57: rel_a or rel_b
	    define rel_58: rel_s or rel_a
	    define rel_59: rel_b
	    define rel_60: rel_s
	    define rel_61: rel_s
	    define rel_62: rel_b
	    define rel_63: rel_a or rel_b or rel_s
	    define rel_64: rel_b
	    define rel_65: rel_s
	    define rel_66: rel_a
	    define rel_67: rel_a
	    define rel_68: rel_m
	    define rel_69: rel_s
	    define rel_70: rel_a or rel_s
	    define rel_71: rel_a
	    define rel_72: rel_s
	    define rel_y: rel_a or rel_b or rel_n or rel_s or rel_q or rel_r or rel_t
	    define rel_73: rel_a
	    define rel_74: rel_m
	    define rel_75: rel_a or rel_b or rel_s
	    define rel_76: rel_a
	    define rel_77: rel_a
	    define rel_78: rel_a
	    define rel_79: rel_r
	    define rel_80: rel_q
	    define rel_h: rel_b
	    define rel_81: rel_b or rel_s
	    define rel_82: rel_s
	    define rel_83: rel_s
	    define rel_84: rel_s
	    define rel_85: rel_s or rel_b
	    define rel_86: rel_m
	    define rel_87: rel_a or rel_b or rel_n or rel_s or rel_q or rel_r or rel_t
	    define rel_z: rel_s or rel_a
	    define rel_88: rel_a or rel_b or rel_n or rel_s or rel_q or rel_r or rel_t
	    define rel_89: rel_s
	    define rel_90: rel_s
	    define rel_91: rel_a or rel_b or rel_n or rel_s or rel_q or rel_r or rel_t
	    define rel_92: rel_a or rel_b or rel_n or rel_s or rel_q or rel_r or rel_t
	    define rel_93: rel_m
	    define rel_94: rel_s
	    define rel_95: rel_m
	    define rel_96: rel_n
	    define rel_j: rel_n or rel_b or rel_s or rel_a
	    define rel_k: rel_n or rel_b or rel_s or rel_a
	    define rel_97: rel_s or rel_a
	    define rel_98: rel_s or rel_a or rel_b
	    define rel_99: rel_s
	    define rel_100: rel_s
	    define rel_101: rel_m or rel_a or rel_b or rel_n or rel_s or rel_q or rel_r or rel_t
	    define rel_l: rel_t
	    define rel_102: rel_m
	    define rel_103: rel_s
	    define rel_104: rel_a or rel_b or rel_s
	    define rel_105: rel_s or rel_a
	    define rel_106: rel_a
	    define rel_107: rel_s
	    define rel_m: ([t1] or rel_m from rel_p) but not rel_1
	    define rel_n: ([t1] or rel_m or rel_n from rel_p) but not rel_1
	    define rel_p: [t2]
	    define rel_q: ([t1] or rel_m or rel_q from rel_p) but not rel_1
	    define rel_r: ([t1] or rel_m or rel_r from rel_p) but not rel_1
	    define rel_s: ([t1] or rel_m or rel_s from rel_p) but not rel_1
	    define rel_t: ([t1] or rel_m or rel_t from rel_p) but not rel_1

	type t4
	  relations
	    define rel_a: rel_a from rel_p but not rel_1
	    define rel_b: rel_b from rel_p but not rel_1
	    define rel_1: rel_1 from rel_p
	    define rel_20: rel_s
	    define rel_z: rel_s or rel_a
	    define rel_m: rel_m from rel_p but not rel_1
	    define rel_n: rel_n from rel_p but not rel_1
	    define rel_p: [t3]
	    define rel_q: rel_q from rel_p but not rel_1
	    define rel_r: rel_r from rel_p but not rel_1
	    define rel_s: rel_s from rel_p but not rel_1
	    define rel_t: rel_t from rel_p but not rel_1
		
`

const small string = `
	model
		schema 1.1

	type user

	type org
		relations
			define rel_o: [user with is_ok, user with is_also_ok, user:*]

	type team
		relations
			define rel_p: [org]
			define rel_o: [user] and rel_o from rel_p

	type group
		relations
			define rel_p: [team]
			define rel_o: [user] and rel_o from rel_p

	type folder
		relations
			define owner: [group]
			define rel_p: [folder]
			define viewer: rel_o from owner or viewer from rel_p

	type document
		relations
			define rel_p: [folder]
			define editor: [user, group#rel_o]
			define viewer: [user] and viewer from rel_p

	condition is_ok(ok: bool) {
		ok
	}

	condition is_also_ok(ok: bool) {
		ok
	}
`

func dfsWeightedGraph(g *graph.WeightedAuthorizationModelGraph, source, leaf string) int {
	sourceNode, ok := g.GetNodeByID(source)
	if !ok {
		panic("source node does not exist")
	}

	leafNode, ok := g.GetNodeByID(leaf)
	if !ok {
		panic("leaf node does not exist")
	}

	visited := make(map[*graph.WeightedAuthorizationModelNode]struct{})

	stack := make([]*graph.WeightedAuthorizationModelNode, 0, 10)
	stack = append(stack, sourceNode)

	var leafCount int

	for len(stack) > 0 {
		ndx := len(stack) - 1
		node := stack[ndx]
		stack = stack[:ndx]

		if node == leafNode {
			leafCount++
		}

		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = struct{}{}

		if _, ok := node.GetWeight(leaf); !ok {
			continue
		}

		edges, ok := g.GetEdgesFromNode(node)
		if !ok {
			continue
		}

		for _, edge := range slices.Backward(edges) {
			stack = append(stack, edge.GetTo())
		}
	}
	return leafCount
}

func dfsCSR(g *graph.CSR, source, leaf string) int {
	sourceNode := g.MustNodeFromID(source)
	leafNode := g.MustNodeFromID(leaf)

	visited := make(map[graph.Node]struct{})

	stack := make([]graph.Node, 0, 10)
	stack = append(stack, sourceNode)

	neighbors := make([]graph.Node, 0, 10)

	var leafCount int

	for len(stack) > 0 {
		ndx := len(stack) - 1
		node := stack[ndx]
		stack = stack[:ndx]

		if node == leafNode {
			leafCount++
		}

		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = struct{}{}

		if _, ok := g.NodeWeight(node, leafNode); !ok {
			continue
		}

		for _, edge := range g.EdgesFromNode(node) {
			targetNode := g.EdgeTarget(edge)
			neighbors = append(neighbors, targetNode)
		}

		slices.Reverse(neighbors)
		stack = append(stack, neighbors...)
		neighbors = neighbors[:0]
	}
	return leafCount
}

func BenchmarkTraversal(b *testing.B) {
	model := transformer.MustTransformDSLToProto(large)
	builder := graph.NewWeightedAuthorizationModelGraphBuilder()

	b.Run("compressed sparse row graph", func(b *testing.B) {
		g, err := builder.Build(model)
		require.NoError(b, err)

		b.Run("build", func(b *testing.B) {
			for b.Loop() {
				_ = graph.NewCSR(g)
			}
		})

		csr := graph.NewCSR(g)

		b.Run("encode", func(b *testing.B) {
			for b.Loop() {
				_, _ = csr.MarshalBinary()
			}
		})

		b.Run("decode", func(b *testing.B) {
			data, err := csr.MarshalBinary()
			require.NoError(b, err)

			var ncsr graph.CSR
			err = ncsr.UnmarshalBinary(data)
			require.NoError(b, err)

			for b.Loop() {
				var ncsr graph.CSR
				_ = ncsr.UnmarshalBinary(data)
			}
		})

		b.Run("traverse", func(b *testing.B) {
			for b.Loop() {
				_ = dfsCSR(csr, "t3#rel_101", "t1")
			}
		})
	})

	b.Run("weighted graph", func(b *testing.B) {
		b.Run("build", func(b *testing.B) {
			for b.Loop() {
				_, _ = builder.Build(model)
			}
		})

		b.Run("traverse", func(b *testing.B) {
			wg, err := builder.Build(model)
			require.NoError(b, err)

			for b.Loop() {
				_ = dfsWeightedGraph(wg, "t3#rel_101", "t1")
			}
		})
	})
}

func TestCSR(t *testing.T) {
	model := transformer.MustTransformDSLToProto(small)
	builder := graph.NewWeightedAuthorizationModelGraphBuilder()
	g, err := builder.Build(model)
	require.NoError(t, err)

	original := graph.NewCSR(g)

	data, err := original.MarshalBinary()
	require.NoError(t, err)

	var decoded graph.CSR
	err = decoded.UnmarshalBinary(data)
	require.NoError(t, err)

	cases := []struct {
		name string
		csr  *graph.CSR
	}{
		{
			name: "build",
			csr:  original,
		},
		{
			name: "decoded",
			csr:  &decoded,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("all nodes report the correct id", func(t *testing.T) {
				var expectedIDs []string
				for id := range g.GetNodes() {
					expectedIDs = append(expectedIDs, id)
				}

				var resultIDs []string
				for _, node := range tc.csr.Nodes() {
					resultIDs = append(resultIDs, tc.csr.NodeID(node))
				}
				require.ElementsMatch(t, expectedIDs, resultIDs)
			})

			t.Run("node reports the correct weights", func(t *testing.T) {
				sourceNode, ok := tc.csr.NodeFromID("document#viewer")
				require.True(t, ok)

				leafNode, ok := tc.csr.NodeFromID("user")
				require.True(t, ok)

				weight, ok := tc.csr.NodeWeight(sourceNode, leafNode)
				require.True(t, ok)
				require.Equal(t, graph.Infinite, weight)

				_, ok = tc.csr.NodeRecursiveRelation(sourceNode)
				require.False(t, ok)

				sourceNode, ok = tc.csr.NodeFromID("document#editor")
				require.True(t, ok)

				weight, ok = tc.csr.NodeWeight(sourceNode, leafNode)
				require.True(t, ok)
				require.Equal(t, 4, weight)

				folderViewerNode, ok := tc.csr.NodeFromID("folder#viewer")
				require.True(t, ok)

				recursiveNode, ok := tc.csr.NodeRecursiveRelation(folderViewerNode)
				require.True(t, ok)

				require.Equal(t, folderViewerNode, recursiveNode)
			})

			t.Run("node reports the correct wildcard types", func(t *testing.T) {
				sourceNode, ok := tc.csr.NodeFromID("document#editor")
				require.True(t, ok)

				require.True(t, tc.csr.NodeReachesWildcard(sourceNode, "user"))

				sourceNode, ok = tc.csr.NodeFromID("document#viewer")
				require.True(t, ok)

				require.True(t, tc.csr.NodeReachesWildcard(sourceNode, "user"))
			})

			t.Run("node reports the correct outgoing edges", func(t *testing.T) {
				sourceNode, ok := tc.csr.NodeFromID("document#viewer")
				require.True(t, ok)

				var edges []graph.Edge

				for _, edge := range tc.csr.EdgesFromNode(sourceNode) {
					edges = append(edges, edge)
				}

				require.Len(t, edges, 1)
				edge := edges[0]
				require.Equal(t, graph.RewriteEdge, tc.csr.EdgeType(edge))
				toNode := tc.csr.EdgeTarget(edge)
				require.Equal(t, graph.OperatorNode, tc.csr.NodeType(toNode))
				require.Equal(t, graph.IntersectionOperator, tc.csr.NodeLabel(toNode))
			})

			t.Run("edge reports the correct conditions", func(t *testing.T) {
				sourceNode, ok := tc.csr.NodeFromID("org#rel_o")
				require.True(t, ok)

				targetNode, ok := tc.csr.NodeFromID("user")
				require.True(t, ok)

				var edge graph.Edge
				var found bool

				for _, e := range tc.csr.EdgesFromNode(sourceNode) {
					if tc.csr.EdgeTarget(e) == targetNode {
						edge = e
						found = true
						break
					}
				}
				require.True(t, found)

				var conditions []string

				for _, condition := range tc.csr.EdgeConditions(edge) {
					conditions = append(conditions, condition)
				}
				require.Len(t, conditions, 2)
				require.ElementsMatch(t, []string{"is_ok", "is_also_ok"}, conditions)
			})
		})
	}
}
