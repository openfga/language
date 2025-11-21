package graph

import (
	"fmt"
	"sync"
	"testing"
)

// TestConcurrentUsersetWeights tests that concurrent access to usersetWeights
// maps in nodes and edges doesn't cause race conditions (panics) or data corruption.
// It verifies that sync.Map correctly handles high-frequency concurrent operations.
func TestConcurrentUsersetWeights(t *testing.T) {
	// Create a test graph
	graph := NewWeightedAuthorizationModelGraph()

	// Add a node and edge for testing
	graph.AddNode("type#rel", "type#rel", SpecificTypeAndRelation)
	node, _ := graph.GetNodeByID("type#rel")

	graph.AddNode("type2#rel2", "type2#rel2", SpecificTypeAndRelation)

	graph.AddEdge("type#rel", "type2#rel2", DirectEdge, "rel", "", []string{NoCond})
	edges, _ := graph.GetEdgesFromNodeId("type#rel")
	edge := edges[0]

	// Test concurrent writes to node
	t.Run("ConcurrentNodeWrites", func(t *testing.T) {
		const iterations = 1000
		const numGoroutines = 10

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				for j := 0; j < iterations; j++ {
					// Use different keys to avoid interference between goroutines
					userset := fmt.Sprintf("test#userset%c", 'A'+id)
					weight := j
					// Atomic write using sync.Map.Store
					graph.setUsersetWeightToNode(node, userset, weight)
				}
			}(i)
		}

		wg.Wait()

		// Verify results - each goroutine should have successfully set its last value
		for i := 0; i < numGoroutines; i++ {
			userset := fmt.Sprintf("test#userset%c", 'A'+i)

			// Atomic read using sync.Map.Load
			weightVal, exists := graph.getUsersetWeightFromNode(node, userset)

			// Cast the interface{} result back to int
			weight := weightVal

			if !exists {
				t.Errorf("expected userset weight for %s to exist", userset)
			}
			if weight != iterations-1 {
				t.Errorf("expected userset weight for %s to be %d, got %d", userset, iterations-1, weight)
			}
		}
	})

	// Test concurrent writes to edge
	t.Run("ConcurrentEdgeWrites", func(t *testing.T) {
		const iterations = 1000
		const numGoroutines = 10

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				for j := 0; j < iterations; j++ {
					// Use different keys to avoid interference between goroutines
					userset := fmt.Sprintf("test#edgeset%c", 'A'+id)
					weight := j
					// Atomic write using sync.Map.Store
					graph.setUsersetWeightToEdge(edge, userset, weight)
				}
			}(i)
		}

		wg.Wait()

		// Verify results - each goroutine should have successfully set its last value
		for i := 0; i < numGoroutines; i++ {
			userset := fmt.Sprintf("test#edgeset%c", 'A'+i)

			// Atomic read using sync.Map.Load
			weightVal, exists := graph.getUsersetWeightFromEdge(edge, userset)
			weight := weightVal

			if !exists {
				t.Errorf("expected userset weight for %s to exist", userset)
			}
			if weight != iterations-1 {
				t.Errorf("expected userset weight for %s to be %d, got %d", userset, iterations-1, weight)
			}
		}
	})

	// Test concurrent writes to same key
	t.Run("ConcurrentSameKeyWrites", func(t *testing.T) {
		const iterations = 1000
		const numGoroutines = 10
		const userset = "shared#userset"

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				for j := 0; j < iterations; j++ {
					// All goroutines write to the same key
					graph.setUsersetWeightToNode(node, userset, id*iterations+j)
				}
			}(i)
		}

		wg.Wait()

		// We can't predict the exact final value, but it should exist
		_, exists := graph.getUsersetWeightFromNode(node, userset)

		if !exists {
			t.Errorf("expected shared userset weight to exist")
		}
		// The key test here is the lack of panic and the successful atomic update.
	})

	// Test parallel node and edge operations
	t.Run("ParallelNodeAndEdgeWrites", func(t *testing.T) {
		const iterations = 1000
		const numGoroutines = 5 // 5 for nodes, 5 for edges

		var wg sync.WaitGroup
		wg.Add(numGoroutines * 2)

		// Node writes
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				for j := 0; j < iterations; j++ {
					userset := fmt.Sprintf("parallel#node%c", 'A'+id)
					weight := j
					graph.setUsersetWeightToNode(node, userset, weight)
				}
			}(i)
		}

		// Edge writes
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				for j := 0; j < iterations; j++ {
					userset := fmt.Sprintf("parallel#edge%c", 'A'+id)
					weight := j
					graph.setUsersetWeightToEdge(edge, userset, weight)
				}
			}(i)
		}

		wg.Wait()

		// Verify results
		for i := 0; i < numGoroutines; i++ {
			// Check node results
			userset := fmt.Sprintf("parallel#node%c", 'A'+i)
			weightVal, exists := graph.getUsersetWeightFromNode(node, userset)
			weight := weightVal

			if !exists {
				t.Errorf("expected node userset weight for %s to exist", userset)
			}
			if weight != iterations-1 {
				t.Errorf("expected node userset weight for %s to be %d, got %d", userset, iterations-1, weight)
			}

			// Check edge results
			userset = fmt.Sprintf("parallel#edge%c", 'A'+i)
			weightVal, exists = graph.getUsersetWeightFromEdge(edge, userset)
			weight = weightVal

			if !exists {
				t.Errorf("expected edge userset weight for %s to exist", userset)
			}
			if weight != iterations-1 {
				t.Errorf("expected edge userset weight for %s to be %d, got %d", userset, iterations-1, weight)
			}
		}
	})

	// Test interleaved read/write operations on a Node. This validates that the non-blocking read path
	// (sync.Map.Load) remains functional while writers (sync.Map.Store) are updating the same map.
	t.Run("InterleavedNodeReadWriteConcurrency", func(t *testing.T) {
		const iterations = 1000
		const numGoroutines = 5 // 5 writers, 5 readers
		const userset = "readwrite#userset"

		// Initialize the value
		graph.setUsersetWeightToNode(node, userset, 0)

		var wg sync.WaitGroup
		wg.Add(numGoroutines * 2)

		// Writers
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < iterations; j++ {
					// Atomic write
					graph.setUsersetWeightToNode(node, userset, j)
				}
			}()
		}

		// Readers
		readCount := 0
		var readCountMutex sync.Mutex
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				localReadCount := 0
				for j := 0; j < iterations; j++ {
					// Atomic, non-blocking read
					_, exists := graph.getUsersetWeightFromNode(node, userset)

					if exists {
						localReadCount++
					}
				}
				readCountMutex.Lock()
				readCount += localReadCount
				readCountMutex.Unlock()
			}()
		}

		wg.Wait()

		// We expect the value to always be present after the initial write, and no panic should occur.
		if readCount != numGoroutines*iterations {
			t.Errorf("expected %d successful reads, got %d", numGoroutines*iterations, readCount)
		}
	})

	// Test interleaved read/write operations on an Edge.
	t.Run("InterleavedEdgeReadWriteConcurrency", func(t *testing.T) {
		const iterations = 1000
		const numGoroutines = 5 // 5 writers, 5 readers
		const userset = "readwrite#edgeuserset"

		// Initialize the value
		graph.setUsersetWeightToEdge(edge, userset, 0)

		var wg sync.WaitGroup
		wg.Add(numGoroutines * 2)

		// Writers
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < iterations; j++ {
					// Atomic write
					graph.setUsersetWeightToEdge(edge, userset, j)
				}
			}()
		}

		// Readers
		readCount := 0
		var readCountMutex sync.Mutex
		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				localReadCount := 0
				for j := 0; j < iterations; j++ {
					// Atomic, non-blocking read
					_, exists := graph.getUsersetWeightFromEdge(edge, userset)

					if exists {
						localReadCount++
					}
				}
				readCountMutex.Lock()
				readCount += localReadCount
				readCountMutex.Unlock()
			}()
		}

		wg.Wait()

		// We expect the value to always be present after the initial write, and no panic should occur.
		if readCount != numGoroutines*iterations {
			t.Errorf("expected %d successful reads, got %d", numGoroutines*iterations, readCount)
		}
	})
}
