package graph

import (
	"slices"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

// unresolvableCycles finds every cycle the builder rejects, from the graph's
// structure alone, as one RelationError per strongly connected component: model
// cycles, which no tuple data can resolve, and tuple cycles through an
// intersection or exclusion, which the weighted graph does not allow. It works
// on a graph whose weights could not be assigned. The components are the same
// whatever order they are found in, and Parse sorts the errors, so the same
// model gets the same errors on every run.
//
// Each error names every relation in its component, so a relation on an
// accepted cycle that shares a relation with a rejected one is named with it.
func (wg *WeightedAuthorizationModelGraph) unresolvableCycles() RelationErrors {
	var relationErrs RelationErrors

	// A component over the edges that read no tuple: going round it consumes
	// nothing, so it never terminates.
	inModelCycle := map[string]bool{}

	for _, component := range wg.components(func(edge *WeightedAuthorizationModelEdge) bool { return !readsTuple(edge) }) {
		relations := relationsIn(component, nil)
		if len(relations) == 0 {
			continue
		}

		for _, node := range component {
			inModelCycle[node.uniqueLabel] = true
		}

		relationErrs = append(relationErrs, &RelationError{Relations: relations, Err: ErrModelCycle})
	}

	// A tuple-reading component with an intersection or exclusion in it. The
	// tuple data may resolve a tuple cycle, but the weighted graph does not allow
	// an intersection or exclusion on one (ErrContrainstTupleCycle). Relations
	// already reported for a model cycle are not named again.
	for _, component := range wg.components(nil) {
		if !slices.ContainsFunc(component, isConstraint) {
			continue
		}

		if relations := relationsIn(component, inModelCycle); len(relations) > 0 {
			relationErrs = append(relationErrs, &RelationError{Relations: relations, Err: ErrContrainstTupleCycle})
		}
	}

	return relationErrs
}

// components returns the strongly connected components of two or more nodes,
// over the edges keep accepts, or over every edge when keep is nil. A single
// node is never one: the builder accepts a relation that only rewrites itself,
// leaving it with no weights, so it is not a refusal.
func (wg *WeightedAuthorizationModelGraph) components(keep func(*WeightedAuthorizationModelEdge) bool) [][]*WeightedAuthorizationModelNode {
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
			// A node's edge to itself cannot join it to any other node.
			if to := ids[edge.to.uniqueLabel]; to != from && (keep == nil || keep(edge)) {
				g.SetEdge(simple.Edge{F: simple.Node(from), T: simple.Node(to)})
			}
		}
	}

	var components [][]*WeightedAuthorizationModelNode

	for _, scc := range topo.TarjanSCC(g) {
		if len(scc) < 2 {
			continue
		}

		component := make([]*WeightedAuthorizationModelNode, len(scc))
		for i, node := range scc {
			component[i] = nodes[node.ID()]
		}

		components = append(components, component)
	}

	return components
}

// relationsIn returns the labels of the relations among nodes, sorted, leaving
// out the ones in skip. Operator and grouping nodes share a component with the
// relations they belong to and are not named.
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

// isConstraint reports whether node is an intersection or exclusion operator.
func isConstraint(node *WeightedAuthorizationModelNode) bool {
	return node.nodeType == OperatorNode && (node.label == IntersectionOperator || node.label == ExclusionOperator)
}

// readsTuple reports whether a step consumes a tuple: a tuple-to-userset edge,
// or a direct edge to another relation (a userset restriction). These are the
// steps that let a cycle terminate. The weight walk and the cycle scan both use
// it, so they agree on which cycles can be resolved.
func readsTuple(edge *WeightedAuthorizationModelEdge) bool {
	return edge.edgeType == TTUEdge ||
		(edge.edgeType == DirectEdge && edge.to.nodeType == SpecificTypeAndRelation)
}
