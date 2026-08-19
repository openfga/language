package validation

import (
	"slices"
	"strings"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"

	"github.com/openfga/language/pkg/go/graph"
)

// validateWithGraph resolves entrypoints and cycles from the weighted graph rather than
// by walking the rewrite tree.
//
// It builds through the exported WeightedAuthorizationModelGraphBuilder rather than
// reimplementing the walk, so what validation reports and what anything else reading that
// graph concludes cannot drift apart the way two implementations of one question would.
//
// The graph reports in two ways and only one of them can carry a position. Build either
// returns a graph or refuses the model, and on refusal it returns no graph, so there is
// nothing left to enumerate: a model with three broken relations yields one error naming
// none of them. Every rule therefore runs only on a model that built, and a refused model
// produces the single finding RaiseModelUnbuildable writes.
func validateWithGraph(collector *ErrorCollector, model *openfgav1.AuthorizationModel, lines []string) {
	if model == nil {
		return
	}

	weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	if err != nil {
		collector.RaiseModelUnbuildable(err)

		return
	}

	scope := &graphScope{weighted: weighted, model: model, collector: collector, lines: lines}
	for _, rule := range graphRules {
		rule.check(scope)
	}
}

// graphRule is one check over a built graph. Every rule collects all of its hits rather
// than returning on the first, which is the whole reason these run outside pkg/go/graph.
type graphRule struct {
	// id names the rule in tests and in failure output. It is not a wire value; what
	// reaches a caller is the ValidationErrorType the rule raises.
	id    string
	check func(*graphScope)
}

// graphRules is the registry every rule joins. A new rule is one entry here plus one
// function, so rules written in parallel meet only on this line.
var graphRules = []graphRule{
	{id: "entrypoints", check: checkRelationEntrypoints},
}

// graphScope is what a rule gets: the built graph, the model behind it, and the means to
// report a finding against a relation.
type graphScope struct {
	weighted  *graph.WeightedAuthorizationModelGraph
	model     *openfgav1.AuthorizationModel
	collector *ErrorCollector
	lines     []string
}

// checkRelationEntrypoints reports relations that can never be satisfied.
//
// A relation node carries one weight per terminal type it can reach. Reaching none means
// no tuple can ever satisfy the relation, which is what the rewrite-tree traversal calls
// an impossible relation. Recursion alone does not empty the weights: a relation with a
// base case keeps its terminal types and takes weight Infinite, so
// `[user] or viewer from parent` is untouched here while `viewer from parent` alone is
// reported.
//
// Only SpecificTypeAndRelation nodes are considered. A type node has no weights either,
// because a type is what weights are counted to rather than from.
func checkRelationEntrypoints(scope *graphScope) {
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

		meta := scope.metaFor(objectType, relation)
		lineIndex := scope.relationLine(objectType, relation)

		if scope.reachesOnlyRewrites(nodeID) {
			scope.collector.RaiseNoEntryPointLoop(relation, objectType, meta, lineIndex)

			continue
		}

		scope.collector.RaiseNoEntryPoint(relation, objectType, meta, lineIndex)
	}
}

// reachesOnlyRewrites reports whether every edge reachable from nodeID rewrites another
// relation, with no direct assignment or tupleset anywhere.
//
// That separates the two things the traversal words differently. `define viewer: viewer`
// closes on itself through a computed rewrite and nothing else, which it calls a potential
// loop. `define viewer: viewer from parent` reaches a tupleset it can never satisfy, which
// it calls a missing entrypoint.
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

// relationLine resolves the source line a relation is declared on, through the same
// helpers the other phases use, so a graph finding and a traversal finding for the same
// relation land in the same place.
func (s *graphScope) relationLine(objectType, relation string) *int {
	if len(s.lines) == 0 {
		return nil
	}

	typeLine := GetTypeLineNumber(objectType, s.lines, nil)

	return GetRelationLineNumber(relation, s.lines, typeLine)
}

// metaFor resolves the file and module a relation was declared in. The graph knows
// neither: a modular model reaches Build already flattened, with its provenance left on
// the model's metadata.
func (s *graphScope) metaFor(objectType, relation string) *Meta {
	for _, typeDef := range s.model.GetTypeDefinitions() {
		if typeDef.GetType() != objectType {
			continue
		}

		return relationMeta(typeDef, relation)
	}

	return nil
}

// splitRelationLabel splits a relation definition into its object type and relation,
// e.g. "document#viewer".
func splitRelationLabel(definition string) (objectType, relation string, ok bool) {
	objectType, relation, found := strings.Cut(definition, "#")
	if !found || objectType == "" || relation == "" {
		return "", "", false
	}

	return objectType, relation, true
}

// relationForNode resolves the relation a node belongs to.
//
// A relation node is its own answer. An operator or logical node is not: its label holds
// the relation but also an operator and an index, and the grouping labels are built
// differently again. Rather than parse those forms, this reads the relation off an
// outgoing edge, which records the relation the edge was written for. Operator and logical
// nodes exist to group edges, so they always have one.
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

// sortedNodeIDs returns the graph's node IDs in a stable order, so the same model raises
// findings in the same order. The graph stores nodes in a map, which Go iterates randomly.
func sortedNodeIDs(weighted *graph.WeightedAuthorizationModelGraph) []string {
	nodes := weighted.GetNodes()
	ids := make([]string, 0, len(nodes))

	for id := range nodes {
		ids = append(ids, id)
	}

	slices.Sort(ids)

	return ids
}
