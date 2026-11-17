package graph

import (
	"fmt"
	"slices"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

// rejectedCycles returns an error for each model cycle and each tuple cycle through an
// intersection or exclusion, and the nodes of the cycles it reports.
func (wg *WeightedAuthorizationModelGraph) rejectedCycles(components [][]*WeightedAuthorizationModelNode) (RelationErrors, map[string]bool) {
	var relationErrs RelationErrors

	failed := map[string]bool{}

	// Without TTU and userset edges, every cycle is a model cycle. A single node is not
	// one: the builder accepts a relation that only rewrites itself.
	notTTUOrUserset := func(edge *WeightedAuthorizationModelEdge) bool { return !isTTUOrUsersetEdge(edge) }
	for _, component := range wg.stronglyConnectedComponents(notTTUOrUserset) {
		if len(component) < 2 {
			continue
		}

		relations := relationsIn(component, nil)
		if len(relations) == 0 {
			continue
		}

		markFailed(failed, component)

		relationErrs = append(relationErrs, &RelationError{Relations: relations, Err: ErrModelCycle})
	}

	// Cycles through an intersection or exclusion, less relations already in a model cycle.
	for _, component := range components {
		if len(component) < 2 || !slices.ContainsFunc(component, isIntersectionOrExclusion) {
			continue
		}

		if relations := relationsIn(component, failed); len(relations) > 0 {
			markFailed(failed, component)

			relationErrs = append(relationErrs, &RelationError{Relations: relations, Err: ErrContrainstTupleCycle})
		}
	}

	return relationErrs, failed
}

// assignWeightsByComponent walks components in order, skipping any that reach a failed
// node, and returns an error for each component whose walk fails.
func (wg *WeightedAuthorizationModelGraph) assignWeightsByComponent(components [][]*WeightedAuthorizationModelNode, failed map[string]bool) RelationErrors {
	visited := make(map[string]bool)
	ancestorPath := make([]*WeightedAuthorizationModelEdge, 0)
	tupleCycleDependencies := make(map[string][]*WeightedAuthorizationModelEdge)

	// The loop of AssignWeights, over one component's relations.
	assignWeights := func(relations []string) error {
		for _, nodeID := range relations {
			if visited[nodeID] {
				continue
			}

			tupleCycles, err := wg.calculateNodeWeight(nodeID, visited, ancestorPath, tupleCycleDependencies)
			if err != nil {
				return err
			}
			if len(tupleCycles) > 0 {
				return fmt.Errorf("%w: %d tuple cycles found without resolution", ErrTupleCycle, len(tupleCycles))
			}
		}
		return nil
	}

	var relationErrs RelationErrors

	for _, component := range components {
		if wg.reachesFailed(component, failed) {
			markFailed(failed, component)

			continue
		}

		relations := relationsIn(component, nil)
		if err := assignWeights(relations); err != nil {
			markFailed(failed, component)

			relationErrs = append(relationErrs, &RelationError{Relations: relations, Err: err})
		}
	}

	return relationErrs
}

// reachesFailed reports whether component holds a failed node or has an edge to one.
func (wg *WeightedAuthorizationModelGraph) reachesFailed(component []*WeightedAuthorizationModelNode, failed map[string]bool) bool {
	for _, node := range component {
		if failed[node.uniqueLabel] {
			return true
		}

		for _, edge := range wg.edges[node.uniqueLabel] {
			if failed[edge.to.uniqueLabel] {
				return true
			}
		}
	}

	return false
}

func markFailed(failed map[string]bool, component []*WeightedAuthorizationModelNode) {
	for _, node := range component {
		failed[node.uniqueLabel] = true
	}
}

// stronglyConnectedComponents returns the graph's strongly connected components over the
// edges keep accepts, or every edge if keep is nil, each after the components it reaches.
func (wg *WeightedAuthorizationModelGraph) stronglyConnectedComponents(keep func(*WeightedAuthorizationModelEdge) bool) [][]*WeightedAuthorizationModelNode {
	g := simple.NewDirectedGraph()
	nodes := make([]*WeightedAuthorizationModelNode, 0, len(wg.nodes))
	ids := make(map[string]int64, len(wg.nodes))

	for label, node := range wg.nodes {
		ids[label] = int64(len(nodes))
		nodes = append(nodes, node)
		g.AddNode(simple.Node(ids[label]))
	}

	for label, edges := range wg.edges {
		from := ids[label]

		for _, edge := range edges {
			// SetEdge panics on a self-edge, which can't join two nodes anyway.
			if to := ids[edge.to.uniqueLabel]; to != from && (keep == nil || keep(edge)) {
				g.SetEdge(simple.Edge{F: simple.Node(from), T: simple.Node(to)})
			}
		}
	}

	sccs := topo.TarjanSCC(g)
	components := make([][]*WeightedAuthorizationModelNode, len(sccs))
	// Each node is in exactly one component, so one array holds them all.
	members := make([]*WeightedAuthorizationModelNode, 0, len(nodes))

	for i, scc := range sccs {
		start := len(members)
		for _, node := range scc {
			members = append(members, nodes[node.ID()])
		}

		components[i] = members[start:len(members):len(members)]
	}

	return components
}

// relationsIn returns the sorted labels of the relations among nodes, except those in skip.
func relationsIn(nodes []*WeightedAuthorizationModelNode, skip map[string]bool) []string {
	var relations []string

	for _, node := range nodes {
		if node.nodeType == SpecificTypeAndRelation && !skip[node.uniqueLabel] {
			relations = append(relations, node.uniqueLabel)
		}
	}

	slices.Sort(relations)

	return relations
}

func isIntersectionOrExclusion(node *WeightedAuthorizationModelNode) bool {
	return node.nodeType == OperatorNode && (node.label == IntersectionOperator || node.label == ExclusionOperator)
}

// isTTUOrUsersetEdge reports whether edge is a TTU edge or a direct edge to a userset,
// the edges that make a cycle a tuple cycle.
func isTTUOrUsersetEdge(edge *WeightedAuthorizationModelEdge) bool {
	return edge.edgeType == TTUEdge ||
		(edge.edgeType == DirectEdge && edge.to.nodeType == SpecificTypeAndRelation)
}
