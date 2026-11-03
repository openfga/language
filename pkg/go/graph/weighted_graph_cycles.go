package graph

import (
	"maps"
	"slices"
)

// unresolvableCycles finds every cycle the builder cannot resolve, from the
// graph's structure alone, as one RelationError per strongly connected
// component. It works on a graph whose weights could not be assigned, runs in
// linear time and returns the same errors on every run, whatever order the
// weight walk failed in.
//
// Each error names every relation in its component, so a relation on a
// resolvable cycle that shares a relation with a broken one is named with it.
func (wg *WeightedAuthorizationModelGraph) unresolvableCycles() []RelationError {
	index := newCycleIndex(wg)

	var relationErrs []RelationError

	// A component over the edges that read no tuple: going round it consumes
	// nothing, so it never terminates.
	inModelCycle := make([]bool, len(index.labels))

	for _, component := range index.components(false) {
		relations := index.relations(component, nil)
		if len(relations) == 0 {
			continue
		}

		for _, node := range component {
			inModelCycle[node] = true
		}

		relationErrs = append(relationErrs, RelationError{Relations: relations, Err: ErrModelCycle})
	}

	// A tuple-reading component with an intersection or exclusion in it: the
	// resolver cannot intersect or subtract a set it is still computing.
	// Relations already reported for a model cycle are not named again.
	for _, component := range index.components(true) {
		if !index.constrained(component) {
			continue
		}

		if relations := index.relations(component, inModelCycle); len(relations) > 0 {
			relationErrs = append(relationErrs, RelationError{Relations: relations, Err: ErrContrainstTupleCycle})
		}
	}

	slices.SortFunc(relationErrs, func(a, b RelationError) int {
		return slices.Compare(a.Relations, b.Relations)
	})

	return relationErrs
}

// cycleIndex is the graph as integer arrays for the component search: nodes in
// label order, and the edges out of node n as edgeTo[edgeStart[n]:edgeStart[n+1]].
type cycleIndex struct {
	labels     []string
	relation   []bool // per node: a type#relation node
	constraint []bool // per node: an intersection or exclusion operator
	edgeStart  []int
	edgeTo     []int
	readsTuple []bool // per edge
}

func newCycleIndex(wg *WeightedAuthorizationModelGraph) *cycleIndex {
	labels := slices.Sorted(maps.Keys(wg.nodes))
	ids := make(map[string]int, len(labels))
	edgeCount := 0

	for i, label := range labels {
		ids[label] = i
		edgeCount += len(wg.edges[label])
	}

	index := &cycleIndex{
		labels:     labels,
		relation:   make([]bool, len(labels)),
		constraint: make([]bool, len(labels)),
		edgeStart:  make([]int, len(labels)+1),
		edgeTo:     make([]int, 0, edgeCount),
		readsTuple: make([]bool, 0, edgeCount),
	}

	for i, label := range labels {
		node := wg.nodes[label]
		index.relation[i] = node.nodeType == SpecificTypeAndRelation
		index.constraint[i] = node.nodeType == OperatorNode &&
			(node.label == IntersectionOperator || node.label == ExclusionOperator)

		for _, edge := range wg.edges[label] {
			index.edgeTo = append(index.edgeTo, ids[edge.to.uniqueLabel])
			index.readsTuple = append(index.readsTuple, readsTuple(edge))
		}

		index.edgeStart[i+1] = len(index.edgeTo)
	}

	return index
}

// components returns the strongly connected components of two or more nodes,
// over every edge or, with tupleEdges false, over the edges that read no tuple.
// A single node is never one: the builder accepts a relation that only rewrites
// itself, leaving it with no weights, so it is not a refusal.
//
// Tarjan's algorithm, iterative so model depth is not bounded by the stack.
func (index *cycleIndex) components(tupleEdges bool) [][]int {
	n := len(index.labels)
	order := make([]int, n) // visit order, from 1; 0 is unvisited
	low := make([]int, n)
	onStack := make([]bool, n)

	type frame struct {
		node, edge int
	}

	var (
		stack      []int
		frames     []frame
		components [][]int
		visited    int
	)

	visit := func(node int) {
		visited++
		order[node], low[node] = visited, visited
		stack = append(stack, node)
		onStack[node] = true
		frames = append(frames, frame{node: node, edge: index.edgeStart[node]})
	}

	for root := range n {
		if order[root] != 0 {
			continue
		}

		visit(root)

		for len(frames) > 0 {
			top := &frames[len(frames)-1]
			node := top.node

			if top.edge < index.edgeStart[node+1] {
				edge := top.edge
				top.edge++

				if !tupleEdges && index.readsTuple[edge] {
					continue
				}

				if next := index.edgeTo[edge]; order[next] == 0 {
					visit(next)
				} else if onStack[next] {
					low[node] = min(low[node], order[next])
				}

				continue
			}

			if low[node] == order[node] {
				start := len(stack) - 1
				for stack[start] != node {
					start--
				}

				for _, member := range stack[start:] {
					onStack[member] = false
				}

				if len(stack)-start > 1 {
					components = append(components, slices.Clone(stack[start:]))
				}

				stack = stack[:start]
			}

			frames = frames[:len(frames)-1]
			if len(frames) > 0 {
				parent := frames[len(frames)-1].node
				low[parent] = min(low[parent], low[node])
			}
		}
	}

	return components
}

// relations returns the labels of the relations among nodes, sorted, leaving
// out the nodes skip marks. Operator and grouping nodes share a component with
// the relations they belong to and are not named.
func (index *cycleIndex) relations(nodes []int, skip []bool) []string {
	var relations []string

	for _, node := range nodes {
		if index.relation[node] && (skip == nil || !skip[node]) {
			relations = append(relations, index.labels[node])
		}
	}

	slices.Sort(relations)

	return relations
}

// constrained reports whether nodes hold an intersection or exclusion operator.
func (index *cycleIndex) constrained(nodes []int) bool {
	return slices.ContainsFunc(nodes, func(node int) bool { return index.constraint[node] })
}

// readsTuple mirrors the weight walk's test for a step that consumes a tuple:
// a tuple-to-userset edge, or a direct edge to another relation (a userset
// restriction). These are the steps that let a cycle terminate.
func readsTuple(edge *WeightedAuthorizationModelEdge) bool {
	return edge.edgeType == TTUEdge ||
		(edge.edgeType == DirectEdge && edge.to.nodeType == SpecificTypeAndRelation)
}
