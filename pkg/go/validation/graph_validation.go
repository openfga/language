package validation

import (
	"errors"
	"slices"
	"strings"

	"github.com/openfga/language/pkg/go/graph"
)

// validateWithGraph resolves entry-point reachability from the weighted graph
// rather than by walking the rewrite tree. It builds through the exported
// WeightedAuthorizationModelGraphBuilder, so what validation reports and what
// anything else reading that graph concludes cannot drift apart.
func validateWithGraph(idx *index, src source) error {
	if idx == nil || idx.model == nil {
		return nil
	}

	weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(idx.model)
	if err != nil {
		return graphRefusal(idx, src, err)
	}

	scope := &graphScope{weighted: weighted, idx: idx, src: src}

	var fs []*Finding
	for _, rule := range graphRules {
		fs = append(fs, rule.check(scope)...)
	}

	return joinFindings(fs...)
}

// graphRefusal reports a model the builder refused, by the first of three
// answers that applies.
//
// The rewrite-tree walk answers first: where it reports anything, those are
// the findings. Where it is silent, the refusal is over a cycle shape the walk
// does not look for, and the relations in that cycle are named as
// cyclic-relation findings. A refusal neither can attribute becomes one
// positionless graph-model-unbuildable finding, with a fixed reason per
// builder sentinel.
func graphRefusal(idx *index, src source, cause error) error {
	if err := validateEntryPoints(idx, src); err != nil {
		return err
	}

	if fs := relationCycleShapes(idx, src); len(fs) > 0 {
		return joinFindings(fs...)
	}

	return joinFindings(modelUnbuildable(describeRefusal(cause)))
}

// graphRefusalReasons words a refusal from the sentinel it carries rather than
// from the builder's message: the builder picks which broken relation to name
// by ranging over a map, so its text changes run to run for the same model.
// Ordered most specific first, because the constrained tuple cycle wraps the
// plain one and its text says more.
var graphRefusalReasons = []error{
	graph.ErrContrainstTupleCycle,
	graph.ErrTupleCycle,
	graph.ErrModelCycle,
	graph.ErrInvalidModel,
}

// unrecognizedRefusal is the wording for a refusal carrying no sentinel this
// package knows: a gap to close rather than builder text to pass through.
const unrecognizedRefusal = "a reason this version does not recognize"

// describeRefusal returns the wording for the sentinel a refusal carries.
func describeRefusal(err error) string {
	for _, sentinel := range graphRefusalReasons {
		if errors.Is(err, sentinel) {
			return sentinel.Error()
		}
	}

	return unrecognizedRefusal
}

// graphRule is one check over a built graph. Every rule collects all of its hits
// rather than returning on the first, which is the whole reason these run
// outside pkg/go/graph.
type graphRule struct {
	// id names the rule in tests and in failure output. It is not a wire value;
	// what reaches a caller is the Kind the rule's findings carry.
	id    string
	check func(*graphScope) []*Finding
}

// graphRules is the registry every rule joins. A new rule is one entry here plus
// one function, so rules written in parallel meet only on this line.
var graphRules = []graphRule{
	{id: "entrypoints", check: checkRelationEntrypoints},
}

// graphScope is what a rule gets: the built graph, the indexed model behind it,
// and the source text a finding's position is stamped from.
type graphScope struct {
	weighted *graph.WeightedAuthorizationModelGraph
	idx      *index
	src      source
}

// checkRelationEntrypoints reports relations that can never be satisfied.
//
// A relation node carries one weight per terminal type it can reach. Reaching
// none means no tuple can ever satisfy the relation, which is what the
// rewrite-tree traversal calls an impossible relation. Recursion alone does not
// empty the weights: a relation with a base case keeps its terminal types and
// takes weight Infinite, so `[user] or viewer from parent` is untouched here
// while `viewer from parent` alone is reported.
//
// Only SpecificTypeAndRelation nodes are considered. A type node has no weights
// either, because a type is what weights are counted to rather than from.
func checkRelationEntrypoints(scope *graphScope) []*Finding {
	var fs []*Finding

	for _, nodeID := range sortedNodeIDs(scope.weighted) {
		node, ok := scope.weighted.GetNodeByID(nodeID)
		if !ok || node.GetNodeType() != graph.SpecificTypeAndRelation {
			continue
		}

		if len(node.GetWeights()) > 0 {
			continue
		}

		definition, ok := relationForNode(scope.weighted, node)
		if !ok {
			continue
		}

		objectType, relation, ok := splitRelationLabel(definition)
		if !ok {
			continue
		}

		finding := noEntryPoint(relation, objectType)
		if scope.reachesOnlyRewrites(nodeID) {
			finding = noEntryPointLoop(relation, objectType)
		}

		// Stamp position and provenance through the same helpers the rewrite-tree
		// phases use, so a graph finding and a traversal finding for the same
		// relation land in the same place.
		file, module := relationMeta(scope.idx.typeDef(objectType), relation)
		line := scope.src.relationLine(relation, scope.src.typeLine(objectType))

		finding.at(scope.src, line)
		finding.File, finding.Metadata.Module = file, module
		fs = append(fs, finding)
	}

	return fs
}

// reachesOnlyRewrites reports whether every edge reachable from nodeID rewrites
// another relation, with no direct assignment or tupleset anywhere.
//
// That separates the two things the traversal words differently. `define viewer:
// viewer` closes on itself through a computed rewrite and nothing else, which it
// calls a potential loop. `define viewer: viewer from parent` reaches a tupleset
// it can never satisfy, which it calls a missing entrypoint.
func (s *graphScope) reachesOnlyRewrites(nodeID string) bool {
	visited := map[string]bool{nodeID: true}
	queue := []string{nodeID}
	sawEdge := false

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		edges, _ := s.weighted.GetEdgesFromNodeID(current)
		for _, edge := range edges {
			sawEdge = true

			if edge.GetEdgeType() != graph.ComputedEdge && edge.GetEdgeType() != graph.RewriteEdge {
				// A direct, tupleset or grouping edge means something outside this
				// relation feeds it, so the problem is a missing entrypoint.
				return false
			}

			next := edge.GetTo().GetUniqueLabel()
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return sawEdge
}

// splitRelationLabel splits a relation definition into its object type and
// relation, e.g. "document#viewer".
func splitRelationLabel(definition string) (objectType, relation string, ok bool) {
	objectType, relation, found := strings.Cut(definition, "#")
	if !found || objectType == "" || relation == "" {
		return "", "", false
	}

	return objectType, relation, true
}

// relationForNode resolves the relation a node belongs to.
//
// A relation node is its own answer. An operator or logical node is not: its
// label holds the relation but also an operator and an index, and the grouping
// labels are built differently again. Rather than parse those forms, this reads
// the relation off an outgoing edge, which records the relation the edge was
// written for. Operator and logical nodes exist to group edges, so they always
// have one.
func relationForNode(weighted *graph.WeightedAuthorizationModelGraph,
	node *graph.WeightedAuthorizationModelNode) (string, bool) {
	switch node.GetNodeType() {
	case graph.SpecificType, graph.SpecificTypeWildcard:
		// A terminal node belongs to no relation.
		return "", false
	case graph.SpecificTypeAndRelation:
		return node.GetUniqueLabel(), true
	}

	edges, _ := weighted.GetEdgesFromNodeID(node.GetUniqueLabel())
	for _, edge := range edges {
		if definition := edge.GetRelationDefinition(); definition != "" {
			return definition, true
		}
	}

	return "", false
}

// sortedNodeIDs returns the graph's node IDs in a stable order, so the same
// model reports findings in the same order. The graph stores nodes in a map,
// which Go iterates randomly.
func sortedNodeIDs(weighted *graph.WeightedAuthorizationModelGraph) []string {
	nodes := weighted.GetNodes()
	ids := make([]string, 0, len(nodes))

	for id := range nodes {
		ids = append(ids, id)
	}

	slices.Sort(ids)

	return ids
}
