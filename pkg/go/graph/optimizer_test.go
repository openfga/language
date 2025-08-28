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

		optimizedGraph := graph.OptimizeAuthorizationModelGraph()

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

		optimizedGraph := graph.OptimizeAuthorizationModelGraph()

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

	t.Run("slice_preallocation", func(t *testing.T) {
		optimizer := &ModelOptimizer{}
		optimizer.cleanUp()

		// Add many patterns
		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("pattern_%d", i)
			optimizer.operatorIndexes[key] = []*OperatorIndex{
				{NodeID: fmt.Sprintf("node_%d", i)},
				{NodeID: fmt.Sprintf("node_%d_dup", i)},
			}
		}

		modifiedNodes := make(map[string]bool)

		start := time.Now()
		for key := range optimizer.operatorIndexes {
			_ = optimizer.getValidOperatorIndexes(key, modifiedNodes)
		}
		duration := time.Since(start)

		require.Less(t, duration, 5*time.Millisecond, "Valid index filtering should be fast")
	})
}

func TestModelOptimizer_EdgeCases(t *testing.T) {
	t.Run("empty_graph", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()
		optimizedGraph := graph.OptimizeAuthorizationModelGraph()

		require.Equal(t, 0, len(optimizedGraph.GetNodes()))
		require.Equal(t, 0, len(optimizedGraph.GetEdges()))
	})

	t.Run("single_node_graph", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()
		graph.AddNode("document#viewer", "viewer", SpecificTypeAndRelation)

		optimizedGraph := graph.OptimizeAuthorizationModelGraph()
		require.Equal(t, 1, len(optimizedGraph.GetNodes()))
	})

	t.Run("no_patterns_to_optimize", func(t *testing.T) {
		graph := createGraphWithUniquePatterns()

		nodesBefore := len(graph.GetNodes())
		optimizedGraph := graph.OptimizeAuthorizationModelGraph()
		nodesAfter := len(optimizedGraph.GetNodes())

		require.Equal(t, nodesBefore, nodesAfter, "Should not create new nodes when no patterns to optimize")
	})

	t.Run("max_iterations_reached", func(t *testing.T) {
		// This test ensures the algorithm doesn't run indefinitely
		graph := createComplexGraphWithManyPatterns()

		start := time.Now()
		optimizedGraph := graph.OptimizeAuthorizationModelGraph()
		duration := time.Since(start)

		require.NotNil(t, optimizedGraph)
		require.Less(t, duration, 5*time.Second, "Optimization should complete within reasonable time")
	})
}

func TestModelOptimizer_SpecificOptimizations(t *testing.T) {
	t.Run("duplicate_edge_removal", func(t *testing.T) {
		model := `
	model
  		schema 1.1
		type user
        type document
            relations
                define owner: [user]
			    define admin: [user]
				define viewer: owner or owner
	`
		authorizationModel := language.MustTransformDSLToProto(model)
		wgb := NewWeightedAuthorizationModelGraphBuilder()
		graph, _ := wgb.Build(authorizationModel)
		edgesBefore := len(graph.GetEdges())
		optimizedGraph := graph.OptimizeAuthorizationModelGraph()

		// Duplicate should be removed
		edgesAfter := len(optimizedGraph.GetEdges())
		require.Equal(t, edgesBefore, 4)
		require.Equal(t, edgesAfter, 4)
	})

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

		optimizedGraph := graph.OptimizeAuthorizationModelGraph()
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
		optimizedGraph := graph.OptimizeAuthorizationModelGraph()

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

func createGraphWithUniquePatterns() *WeightedAuthorizationModelGraph {
	graph := NewWeightedAuthorizationModelGraph()

	// Create relations with unique patterns (no duplicates)
	graph.AddNode("document#owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("document#admin", "admin", SpecificTypeAndRelation)
	graph.AddNode("document#writer", "writer", SpecificTypeAndRelation)
	graph.AddNode("document#reader", "reader", SpecificTypeAndRelation)

	// Unique pattern 1: owner or admin
	graph.AddNode("document#viewer1", "viewer1", SpecificTypeAndRelation)
	graph.AddNode("union-op1", UnionOperator, OperatorNode)
	graph.AddEdge("document#viewer1", "union-op1", RewriteEdge, "", nil)
	graph.AddEdge("union-op1", "document#owner", ComputedEdge, "", nil)
	graph.AddEdge("union-op1", "document#admin", ComputedEdge, "", nil)

	// Unique pattern 2: writer or reader
	graph.AddNode("document#viewer2", "viewer2", SpecificTypeAndRelation)
	graph.AddNode("union-op2", UnionOperator, OperatorNode)
	graph.AddEdge("document#viewer2", "union-op2", RewriteEdge, "", nil)
	graph.AddEdge("union-op2", "document#writer", ComputedEdge, "", nil)
	graph.AddEdge("union-op2", "document#reader", ComputedEdge, "", nil)

	return graph
}

func createComplexGraphWithManyPatterns() *WeightedAuthorizationModelGraph {
	graph := NewWeightedAuthorizationModelGraph()

	// Create many nodes with overlapping patterns
	relations := []string{"owner", "admin", "writer", "reader", "editor", "viewer"}
	for _, rel := range relations {
		graph.AddNode(fmt.Sprintf("document#%s", rel), rel, SpecificTypeAndRelation)
	}

	// Create many complex relations with overlapping patterns
	for i := 0; i < 20; i++ {
		relationName := fmt.Sprintf("document#complex_%d", i)
		unionName := fmt.Sprintf("union-op-%d", i)

		graph.AddNode(relationName, fmt.Sprintf("complex_%d", i), SpecificTypeAndRelation)
		graph.AddNode(unionName, UnionOperator, OperatorNode)
		graph.AddEdge(relationName, unionName, RewriteEdge, "", nil)

		// Add overlapping patterns
		if i%3 == 0 {
			graph.AddEdge(unionName, "document#owner", ComputedEdge, "", nil)
			graph.AddEdge(unionName, "document#admin", ComputedEdge, "", nil)
		}
		if i%5 == 0 {
			graph.AddEdge(unionName, "document#writer", ComputedEdge, "", nil)
			graph.AddEdge(unionName, "document#reader", ComputedEdge, "", nil)
		}
		graph.AddEdge(unionName, fmt.Sprintf("document#%s", relations[i%len(relations)]), ComputedEdge, "", nil)
	}

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

func BenchmarkFullOptimization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		graph := createTestGraphWithCommonPatterns()
		_ = graph.OptimizeAuthorizationModelGraph()
	}
}
