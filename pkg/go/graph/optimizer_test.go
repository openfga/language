package graph

import (
	"fmt"
	"strings"
	"testing"
	"time"

	language "github.com/openfga/language/pkg/go/transformer"
	"github.com/stretchr/testify/require"
)

func TestModelOptimizer_BasicOptimization(t *testing.T) {
	t.Run("simple_union_optimization", func(t *testing.T) {
		graph := createTestGraphWithCommonPatterns()

		nodesBefore := len(graph.GetNodes())
		edgesBefore := countTotalEdges(graph)

		optimizedGraph, changes, _ := graph.OptimizeAuthorizationModelGraph()

		require.True(t, changes, "Expected optimizations to be applied")
		nodesAfter := len(optimizedGraph.GetNodes())
		edgesAfter := countTotalEdges(optimizedGraph)

		require.Equal(t, nodesBefore, 12)
		require.Equal(t, edgesBefore, 16)
		// Should reuse nodes
		require.Equal(t, nodesAfter, 14)
		require.Equal(t, edgesAfter, 16)
	})

	t.Run("canonical_reuse", func(t *testing.T) {
		graph := createGraphWithCanonicalPattern()

		nodesBefore := len(graph.GetNodes())
		edgesBefore := countTotalEdges(graph)

		optimizedGraph, changes, _ := graph.OptimizeAuthorizationModelGraph()

		require.True(t, changes, "Expected optimizations to be applied")
		nodesAfter := len(optimizedGraph.GetNodes())
		edgesAfter := countTotalEdges(optimizedGraph)

		require.Equal(t, nodesBefore, 11)
		require.Equal(t, edgesBefore, 13)
		// Should reuse nodes
		require.Equal(t, nodesAfter, 10)
		require.Equal(t, edgesAfter, 10)
	})
}

func TestModelOptimizer_PerformanceOptimizations(t *testing.T) {
	t.Run("string_builder_performance", func(t *testing.T) {
		// Test that canonical key generation is efficient
		node := &WeightedAuthorizationModelNode{
			uniqueLabel: "test-node",
			label:       UnionOperator,
		}

		logicEdges := make([]*LogicEdge, 100)
		for i := 0; i < 100; i++ {
			logicEdges[i] = &LogicEdge{
				ID: fmt.Sprintf("edge_%d", i),
			}
		}

		start := time.Now()
		for i := 0; i < 1000; i++ {
			_ = createCanonicalKey("document", node, logicEdges)
		}
		duration := time.Since(start)

		require.Less(t, duration, 100*time.Millisecond, "Canonical key generation should be fast")
	})

	t.Run("combination_generation_performance", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()
		optimizer := &ModelOptimizer{graph: graph}
		optimizer.cleanUp()

		// Create a node with many logic edges
		node := &WeightedAuthorizationModelNode{
			uniqueLabel: "test-union",
			label:       UnionOperator,
		}

		logicEdges := make([]*LogicEdge, 10)
		for i := 0; i < 10; i++ {
			logicEdges[i] = &LogicEdge{
				ID: fmt.Sprintf("edge_%d", i),
			}
		}

		start := time.Now()
		optimizer.mineFrequentPattern(node, logicEdges, "document", 2)
		optimizer.mineFrequentPattern(node, logicEdges, "document", 3)
		duration := time.Since(start)

		require.Less(t, duration, 10*time.Millisecond, "Pattern mining should be fast")

		// Verify patterns were created
		require.Greater(t, len(optimizer.frequentPatterns), 0)
	})
}

func TestModelOptimizer_EdgeCases(t *testing.T) {
	t.Run("empty_graph", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()
		optimizedGraph, changes, _ := graph.OptimizeAuthorizationModelGraph()
		require.False(t, changes, "Expected no optimizations to be applied")
		require.Equal(t, 0, len(optimizedGraph.GetNodes()))
		require.Equal(t, 0, len(optimizedGraph.GetEdges()))
	})

	t.Run("single_node_graph", func(t *testing.T) {
		model := `
		model
	  		schema 1.1
			type user
		`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, _ := wgb.Build(authorizationModel)
		optimizedGraph, changes, _ := graph.OptimizeAuthorizationModelGraph()
		require.False(t, changes, "Expected no optimizations to be applied")
		require.Equal(t, 1, len(optimizedGraph.GetNodes()))
	})

	t.Run("no_patterns_to_optimize", func(t *testing.T) {
		model := `
			model
	  			schema 1.1
				type user
				type document
					relations
						define owner: [user]
					    define admin: [user]
						define viewer: [user] or owner or admin
						define editor: owner and admin
						define contributor: editor and viewer
						define reader: viewer
						define group: admin
						define collaborator: contributor or reader
						define auditor: viewer from parent
						define parent: [document]
		`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, _ := wgb.Build(authorizationModel)

		_, changes, _ := graph.OptimizeAuthorizationModelGraph()
		require.False(t, changes, "Expected no optimizations to be applied")
	})

	t.Run("max_iterations_reached", func(t *testing.T) {
		// This test ensures the algorithm doesn't run indefinitely
		graph := createComplexGraphWithManyPatterns()

		edgesBefore := countTotalEdges(graph)
		nodesBefore := len(graph.GetNodes())

		optimizedGraph, changes, _ := graph.OptimizeAuthorizationModelGraph()
		require.Equal(t, edgesBefore, 52)
		require.Equal(t, nodesBefore, 25)
		require.True(t, changes, "Expected optimizations to be applied")
		edgesAfter := countTotalEdges(optimizedGraph)
		nodesAfter := len(optimizedGraph.GetNodes())
		require.Equal(t, nodesAfter, 27)
		require.Equal(t, edgesAfter, 34)

		require.NotNil(t, optimizedGraph)
	})
}

func TestModelOptimizer_SpecificOptimizations(t *testing.T) {

	t.Run("ttu_edge_aggregation", func(t *testing.T) {
		model := `
	model
  		schema 1.1
		type user
        type document
            relations
                define owner: [user]
			    define admin: [user]
				define parent: [document, group]
				define writer: owner or admin from parent
				define viewer: admin or admin from parent or owner
		 type group
		    relations
				define admin: [user]
	`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, _ := wgb.Build(authorizationModel)
		edgesBefore := countTotalEdges(graph)
		nodesBefore := len(graph.GetNodes())

		optimizedGraph, changes, _ := graph.OptimizeAuthorizationModelGraph()
		require.True(t, changes)
		edgesAfter := countTotalEdges(optimizedGraph)
		nodesAfter := len(optimizedGraph.GetNodes())
		require.Equal(t, edgesBefore, 14)
		require.Equal(t, nodesBefore, 11)
		// TTU edges should be aggregated
		require.Equal(t, edgesAfter, 12)
		require.Equal(t, nodesAfter, 11)

		// Should handle TTU edges correctly
		require.NotNil(t, optimizedGraph)
	})
}

func TestModelOptimizer_MultipleIterations(t *testing.T) {
	t.Run("multiple_iterations_with_diff_frequency", func(t *testing.T) {
		model := `
	model
  		schema 1.1
		type user
        type document
            relations
                define relp1: rel1 or rel2 or rel3
				define relp2: rel2 or rel3 or rel1
				define relp3: rel3 or rel1
				define relp4: rel3 or rel2 or rel5
				define relp5: rel2 or rel6 or rel3
			    define rel1: [user]
				define rel2: [user]
				define rel3: [user]
				define rel4: [user]
				define rel5: [user]
				define rel6: [user]
	`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, _ := wgb.Build(authorizationModel)
		edgesBefore := countTotalEdges(graph)
		nodesBefore := len(graph.GetNodes())
		optimizedGraph, changes, _ := graph.OptimizeAuthorizationModelGraph()

		require.True(t, changes)
		require.Equal(t, edgesBefore, 25)
		require.Equal(t, nodesBefore, 18)

		// Duplicate should be removed
		edgesAfter := countTotalEdges(optimizedGraph)
		nodesAfter := len(graph.GetNodes())
		require.Equal(t, edgesAfter, 22)
		require.Equal(t, nodesAfter, 19)
	})

}

func TestModelOptimizer_LogicEdgeOperations(t *testing.T) {
	t.Run("add_compute_logic_edge", func(t *testing.T) {
		optimizer := &ModelOptimizer{}
		logicEdges := make(map[string]*LogicEdge)

		edge := &WeightedAuthorizationModelEdge{
			to: &WeightedAuthorizationModelNode{uniqueLabel: "document#owner"},
		}

		// First addition should succeed
		success := optimizer.addComputeLogicEdge(logicEdges, edge)
		require.True(t, success)
		require.Len(t, logicEdges, 1)

		// Second addition of same edge should fail (duplicate)
		success = optimizer.addComputeLogicEdge(logicEdges, edge)
		require.False(t, success)
		require.Len(t, logicEdges, 1)
	})

	t.Run("add_ttu_logic_edge", func(t *testing.T) {
		optimizer := &ModelOptimizer{}
		logicEdges := make(map[string]*LogicEdge)

		edge1 := &WeightedAuthorizationModelEdge{
			to:               &WeightedAuthorizationModelNode{uniqueLabel: "folder#member"},
			tuplesetRelation: "parent",
		}

		edge2 := &WeightedAuthorizationModelEdge{
			to:               &WeightedAuthorizationModelNode{uniqueLabel: "group#member"},
			tuplesetRelation: "parent",
		}

		optimizer.addTTULogicEdge(logicEdges, edge1)
		optimizer.addTTULogicEdge(logicEdges, edge2)

		// Should aggregate TTU edges with same tupleset relation
		require.Len(t, logicEdges, 1)

		// The aggregated edge should contain both original edges
		for _, logicEdge := range logicEdges {
			require.Len(t, logicEdge.Edges, 2)
		}
	})
}

func TestModelOptimizer_CanonicalKeyGeneration(t *testing.T) {
	t.Run("short_key", func(t *testing.T) {
		node := &WeightedAuthorizationModelNode{
			uniqueLabel: "test-node",
			label:       UnionOperator,
		}

		logicEdges := []*LogicEdge{
			{ID: "document#owner"},
			{ID: "document#admin"},
		}

		key := createCanonicalKey("document", node, logicEdges)
		expected := "document@union:document#owner,document#admin"
		require.Equal(t, expected, key)
	})

	t.Run("long_key_with_hash", func(t *testing.T) {
		node := &WeightedAuthorizationModelNode{
			uniqueLabel: "test-node",
			label:       UnionOperator,
		}

		// Create very long edge IDs to trigger hashing
		logicEdges := make([]*LogicEdge, 10)
		for i := 0; i < 10; i++ {
			logicEdges[i] = &LogicEdge{
				ID: fmt.Sprintf("document#very_long_edge_name_that_will_trigger_hashing_%d", i),
			}
		}

		key := createCanonicalKey("document", node, logicEdges)
		require.True(t, strings.HasPrefix(key, "hash_"))
		// max size will be 21 characters
		require.Less(t, len(key), 22) // Should be much shorter than the original content

	})
}

// Helper functions for creating test graphs

func createGraphWithCanonicalPattern() *WeightedAuthorizationModelGraph {
	model := `
	model
  		schema 1.1
		type user
        type document
            relations
                define owner: [user]
			    define admin: [user]
				define writer: [user]
				define viewer2: owner or admin
				define viewer: owner or admin
				define viewer3: owner or admin or writer
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, _ := wgb.Build(authorizationModel)
	return graph
}

func createTestGraphWithCommonPatterns() *WeightedAuthorizationModelGraph {
	model := `
	model
  		schema 1.1
		type user
        type document
            relations
                define owner: [user]
			    define admin: [user]
				define writer: [user]
				define parent: [document]
				define viewer2: [user] or owner or admin 
				define viewer: writer from parent or owner or admin
				define viewer3: admin or owner or writer
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, _ := wgb.Build(authorizationModel)
	return graph
}

func createComplexGraphWithManyPatterns() *WeightedAuthorizationModelGraph {
	model := `
		model
  			schema 1.1
			type user
			type document
				relations
					define rel1: [user]
				    define rel2: [user]
					define rel3: [user]
					define rel5: [user]
				    define rel6: [user]
					define rel7: [user]
				    define rel8: [user]
					define rel9: [user]
				    define rel10: [user]
					define relA: rel1 or rel2 or rel3
					define relB: rel3 or rel2 or rel5
					define relC: rel3 or rel1 or rel6 or rel2
					define relD: rel7 or rel3 or rel1 or rel6 or rel2
					define relE: rel7 or rel8 or rel1 or rel6 or rel2 or rel3
					define relF: rel7 or rel8 or rel1 or rel6 or rel2 or rel3 or rel9
					define relX: rel10 or rel7 or rel8 or rel1 or rel6 or rel2 or rel3 or rel9				
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, _ := wgb.Build(authorizationModel)
	return graph
}

func countTotalEdges(graph *WeightedAuthorizationModelGraph) int {
	total := 0
	for _, edges := range graph.GetEdges() {
		total += len(edges)
	}
	return total
}

// Benchmark tests
func BenchmarkCanonicalKeyGeneration(b *testing.B) {
	node := &WeightedAuthorizationModelNode{
		uniqueLabel: "test-node",
		label:       UnionOperator,
	}

	logicEdges := make([]*LogicEdge, 10)
	for i := 0; i < 10; i++ {
		logicEdges[i] = &LogicEdge{
			ID: fmt.Sprintf("document#relation_%d", i),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = createCanonicalKey("document", node, logicEdges)
	}
}

func BenchmarkPatternMining(b *testing.B) {
	graph := NewWeightedAuthorizationModelGraph()
	optimizer := &ModelOptimizer{graph: graph}
	optimizer.cleanUp()

	node := &WeightedAuthorizationModelNode{
		uniqueLabel: "test-union",
		label:       UnionOperator,
	}

	logicEdges := make([]*LogicEdge, 8)
	for i := 0; i < 8; i++ {
		logicEdges[i] = &LogicEdge{
			ID: fmt.Sprintf("edge_%d", i),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		optimizer.mineFrequentPattern(node, logicEdges, "document", 2)
		optimizer.mineFrequentPattern(node, logicEdges, "document", 3)
	}
}
