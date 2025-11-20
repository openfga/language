package graph

import (
	"sync"
	"testing"
)

// TestConcurrentUsersetWeights tests that concurrent access to usersetWeights
// maps in nodes and edges doesn't cause race conditions
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
					userset := "test#userset" + string(rune('A'+id))
					weight := j
					graph.setUsersetWeightToNode(node, userset, weight)
				}
			}(i)
		}

		wg.Wait()

		// Verify results - each goroutine should have successfully set its last value
		for i := 0; i < numGoroutines; i++ {
			userset := "test#userset" + string(rune('A'+i))
			mutex := getUsersetNodeMutex(node)
			mutex.RLock()
			weight, exists := node.usersetWeights[userset]
			mutex.RUnlock()

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
					userset := "test#edgeset" + string(rune('A'+id))
					weight := j
					graph.setUsersetWeightToEdge(edge, userset, weight)
				}
			}(i)
		}

		wg.Wait()

		// Verify results - each goroutine should have successfully set its last value
		for i := 0; i < numGoroutines; i++ {
			userset := "test#edgeset" + string(rune('A'+i))
			mutex := getUsersetEdgeMutex(edge)
			mutex.RLock()
			weight, exists := edge.usersetWeights[userset]
			mutex.RUnlock()

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
		mutex := getUsersetNodeMutex(node)
		mutex.RLock()
		_, exists := node.usersetWeights[userset]
		mutex.RUnlock()

		if !exists {
			t.Errorf("expected shared userset weight to exist")
		}
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
					userset := "parallel#node" + string(rune('A'+id))
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
					userset := "parallel#edge" + string(rune('A'+id))
					weight := j
					graph.setUsersetWeightToEdge(edge, userset, weight)
				}
			}(i)
		}

		wg.Wait()

		// Verify results
		for i := 0; i < numGoroutines; i++ {
			// Check node results
			userset := "parallel#node" + string(rune('A'+i))
			mutex := getUsersetNodeMutex(node)
			mutex.RLock()
			weight, exists := node.usersetWeights[userset]
			mutex.RUnlock()

			if !exists {
				t.Errorf("expected node userset weight for %s to exist", userset)
			}
			if weight != iterations-1 {
				t.Errorf("expected node userset weight for %s to be %d, got %d", userset, iterations-1, weight)
			}

			// Check edge results
			userset = "parallel#edge" + string(rune('A'+i))
			mutex2 := getUsersetEdgeMutex(edge)
			mutex2.RLock()
			weight, exists = edge.usersetWeights[userset]
			mutex2.RUnlock()

			if !exists {
				t.Errorf("expected edge userset weight for %s to exist", userset)
			}
			if weight != iterations-1 {
				t.Errorf("expected edge userset weight for %s to be %d, got %d", userset, iterations-1, weight)
			}
		}
	})

	// Test interleaved read/write operations
	t.Run("InterleavedReadWrite", func(t *testing.T) {
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
					mutex := getUsersetNodeMutex(node)
					mutex.RLock()
					_, exists := node.usersetWeights[userset]
					mutex.RUnlock()

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

		// We expect all reads to find the value
		if readCount != numGoroutines*iterations {
			t.Errorf("expected %d successful reads, got %d", numGoroutines*iterations, readCount)
		}
	})
}
