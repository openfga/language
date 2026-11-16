package graph

import (
	"fmt"
	"slices"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

// unresolvableCycles finds every cycle the builder rejects, from the graph's
// structure alone, as one RelationError per strongly connected component: model
// cycles, which no tuple data can resolve, and tuple cycles through an
// intersection or exclusion, which the weighted graph does not allow. It reads
// no weights, so it works on a graph whose weights are not assigned. It takes
// the graph's components over every edge, as components(nil) returns them.
// The components are the same whatever order they are found in, and Parse sorts
// the errors, so the same model gets the same errors on every run.
//
// Each error names every relation in its component, so a relation on an
// accepted cycle that shares a relation with a rejected one is named with it.
//
// It also returns every node of the components it reports, for the weight walk
// to treat as failed.
func (wg *WeightedAuthorizationModelGraph) unresolvableCycles(components [][]*WeightedAuthorizationModelNode) (RelationErrors, map[string]bool) {
	var relationErrs RelationErrors

	failed := map[string]bool{}

	// A component over the edges that read no tuple: going round it consumes
	// nothing, so it never terminates. A single node is never one: the builder
	// accepts a relation that only rewrites itself, leaving it with no weights,
	// so it is not a refusal.
	for _, component := range wg.components(func(edge *WeightedAuthorizationModelEdge) bool { return !readsTuple(edge) }) {
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

	// A tuple-reading component with an intersection or exclusion in it. The
	// tuple data may resolve a tuple cycle, but the weighted graph does not allow
	// an intersection or exclusion on one (ErrContrainstTupleCycle). Relations
	// already reported for a model cycle are not named again.
	for _, component := range components {
		if len(component) < 2 || !slices.ContainsFunc(component, isConstraint) {
			continue
		}

		if relations := relationsIn(component, failed); len(relations) > 0 {
			markFailed(failed, component)

			relationErrs = append(relationErrs, &RelationError{Relations: relations, Err: ErrContrainstTupleCycle})
		}
	}

	return relationErrs, failed
}

// assignWeightsByComponent assigns weights to a graph Build rejected one
// strongly connected component at a time, and returns an error for each
// component that fails instead of stopping at the first. It takes the graph's
// components over every edge, as components(nil) returns them, so a component
// is walked only once everything it reaches has been, and the nodes known to
// fail, those of the cycles unresolvableCycles reports, to which it adds the
// nodes of every component that fails.
//
// A component that holds a failed node, or has an edge to one, fails without an
// error of its own: its problem is the one it reaches. Any other component has
// its relations walked in name order, and the first error is returned with the
// component's relations. So every error is a root cause.
//
// Nothing that reaches a failed component is walked, so the weights a failed
// walk leaves half assigned there are never read. The only nodes outside it that
// a failed walk can leave without weights are those that group a tuple to
// userset, such as viewer from parent, for the relations of one type that use
// it. A relation that reaches one later fails on it with ErrModelCycle, as it
// does when it reaches it first. So what happens to a component depends only on
// the components it reaches, and the errors do not depend on the order of
// components that do not reach each other.
//
// A component with no relation, such as a type or an operator, is not walked on
// its own. An operator is walked from its relation, as in AssignWeights, so its
// error is reported with that relation. The call for each relation and the
// tuple cycle error are those of AssignWeights. Keep the two loops in step.
func (wg *WeightedAuthorizationModelGraph) assignWeightsByComponent(components [][]*WeightedAuthorizationModelNode, failed map[string]bool) RelationErrors {
	visited := make(map[string]bool)
	ancestorPath := make([]*WeightedAuthorizationModelEdge, 0)
	tupleCycleDependencies := make(map[string][]*WeightedAuthorizationModelEdge)

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
		if wg.blocked(component, failed) {
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

// blocked reports whether component holds a failed node or has an edge to one.
func (wg *WeightedAuthorizationModelGraph) blocked(component []*WeightedAuthorizationModelNode, failed map[string]bool) bool {
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

// markFailed adds the nodes of component to failed.
func markFailed(failed map[string]bool, component []*WeightedAuthorizationModelNode) {
	for _, node := range component {
		failed[node.uniqueLabel] = true
	}
}

// components returns the strongly connected components of the graph, single
// nodes included, over the edges keep accepts, or over every edge when keep is
// nil. They come in the order topo.TarjanSCC finds them, which lists a
// component only after every component it reaches, the reverse of the order
// topo.Sort gives: what a relation is defined by comes before it. Components
// that do not reach each other come in an order that changes from run to run.
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

	sccs := topo.TarjanSCC(g)
	components := make([][]*WeightedAuthorizationModelNode, len(sccs))
	// Every node is in exactly one component, so one array holds them all.
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
