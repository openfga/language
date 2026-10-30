package graph

import (
	"maps"
	"slices"
)

// CycleError is the refusal Build returns for a model whose relations sit in
// cycles the builder cannot resolve. It wraps the error the weight walk
// refused with, so errors.Is and the message are unchanged, and adds every
// unresolvable cycle with the relations in it.
type CycleError struct {
	err error

	// Cycles holds every unresolvable cycle in the model, sorted by their
	// first member. Each cycle names the relations in its strongly connected
	// component, so a relation on a resolvable cycle that shares a relation
	// with a broken one is named with it.
	Cycles []Cycle
}

// Cycle is one unresolvable cycle: the sentinel naming why it cannot be
// resolved, and the relations in it.
type Cycle struct {
	// Kind is ErrModelCycle for a cycle no step of which reads a tuple, and
	// ErrContrainstTupleCycle for a tuple-reading cycle constrained by an
	// intersection or an exclusion.
	Kind error

	// Members holds the unique labels ("type#relation") of the relations in
	// the cycle, sorted and deduplicated.
	Members []string
}

func (e *CycleError) Error() string { return e.err.Error() }

func (e *CycleError) Unwrap() error { return e.err }

// unresolvableCycles finds every cycle the builder cannot resolve, reading
// the graph's structure. The graph is fully constructed before weight
// assignment starts, so this works on a refused build; it runs in linear
// time and is deterministic regardless of the order the weight walk failed
// in.
func (wg *WeightedAuthorizationModelGraph) unresolvableCycles() []Cycle {
	var cycles []Cycle

	// A cycle no step of which reads a tuple: a strongly connected component
	// over the rewrite-only edges. Going around it consumes nothing, so it
	// never terminates.
	inModelCycle := map[string]bool{}

	for _, component := range stronglyConnected(wg, false) {
		members := wg.relationMembers(component)
		if len(members) == 0 {
			continue
		}

		for _, member := range members {
			inModelCycle[member] = true
		}

		cycles = append(cycles, Cycle{Kind: ErrModelCycle, Members: members})
	}

	// A tuple-reading cycle with an intersection or exclusion on it: the
	// resolver cannot subtract or intersect a set it is still computing.
	// Relations already reported for a model cycle are not named again.
	for _, component := range stronglyConnected(wg, true) {
		if !containsConstraint(wg, component) {
			continue
		}

		members := make([]string, 0, len(component))

		for _, member := range wg.relationMembers(component) {
			if !inModelCycle[member] {
				members = append(members, member)
			}
		}

		if len(members) > 0 {
			cycles = append(cycles, Cycle{Kind: ErrContrainstTupleCycle, Members: members})
		}
	}

	slices.SortFunc(cycles, func(a, b Cycle) int {
		return slices.Compare(a.Members, b.Members)
	})

	return cycles
}

// readsTuple mirrors the weight walk's test for a step that consumes a tuple:
// a tuple-to-userset edge, or a direct edge to another relation (a userset
// restriction). These are the steps that let a cycle terminate.
func readsTuple(edge *WeightedAuthorizationModelEdge) bool {
	return edge.edgeType == TTUEdge ||
		(edge.edgeType == DirectEdge && edge.to.nodeType == SpecificTypeAndRelation)
}

// containsConstraint reports whether the component holds an intersection or
// exclusion operator node.
func containsConstraint(wg *WeightedAuthorizationModelGraph, component []string) bool {
	for _, label := range component {
		node := wg.nodes[label]
		if node != nil && node.nodeType == OperatorNode &&
			(node.label == IntersectionOperator || node.label == ExclusionOperator) {
			return true
		}
	}

	return false
}

// relationMembers filters component labels to the relations among them,
// sorted and deduplicated. Operator and grouping nodes share a cycle with the
// relations they belong to and are not named.
func (wg *WeightedAuthorizationModelGraph) relationMembers(labels []string) []string {
	members := make([]string, 0, len(labels))
	seen := make(map[string]bool, len(labels))

	for _, label := range labels {
		if seen[label] {
			continue
		}
		seen[label] = true

		if node, ok := wg.nodes[label]; ok && node.nodeType == SpecificTypeAndRelation {
			members = append(members, label)
		}
	}

	slices.Sort(members)

	return members
}

// stronglyConnected returns the cyclic strongly connected components of the
// graph: components of two or more nodes. A single node is never one — a
// relation rewriting itself is the builder's recursive-edge special case and
// resolves. With tupleEdges false, edges that read a tuple are left out,
// which is the subgraph where a cycle can never terminate.
//
// Tarjan's algorithm, iterative so model depth is not bounded by the stack.
func stronglyConnected(wg *WeightedAuthorizationModelGraph, tupleEdges bool) [][]string {
	index := 0
	indices := map[string]int{}
	low := map[string]int{}
	onStack := map[string]bool{}
	stack := []string{}

	var components [][]string

	neighbors := func(from string) []string {
		var out []string

		for _, edge := range wg.edges[from] {
			if !tupleEdges && readsTuple(edge) {
				continue
			}

			out = append(out, edge.to.uniqueLabel)
		}

		return out
	}

	type frame struct {
		node string
		next int
	}

	for _, start := range slices.Sorted(maps.Keys(wg.nodes)) {
		if _, seen := indices[start]; seen {
			continue
		}

		frames := []frame{{node: start}}
		indices[start], low[start] = index, index
		index++
		stack = append(stack, start)
		onStack[start] = true

		for len(frames) > 0 {
			f := &frames[len(frames)-1]
			ns := neighbors(f.node)

			if f.next < len(ns) {
				next := ns[f.next]
				f.next++

				if _, seen := indices[next]; !seen {
					indices[next], low[next] = index, index
					index++
					stack = append(stack, next)
					onStack[next] = true
					frames = append(frames, frame{node: next})
				} else if onStack[next] {
					low[f.node] = min(low[f.node], indices[next])
				}

				continue
			}

			if low[f.node] == indices[f.node] {
				var component []string

				for {
					top := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					onStack[top] = false
					component = append(component, top)

					if top == f.node {
						break
					}
				}

				if len(component) > 1 {
					components = append(components, component)
				}
			}

			frames = frames[:len(frames)-1]

			if len(frames) > 0 {
				parent := &frames[len(frames)-1]
				low[parent.node] = min(low[parent.node], low[f.node])
			}
		}
	}

	return components
}
