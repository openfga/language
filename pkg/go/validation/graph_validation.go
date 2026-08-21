package validation

import (
	"errors"
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
// Only a graph Build accepted is read. Weight assignment marks a node visited before it walks
// that node's edges and stops at the first node it cannot weight, so in a graph it refused the
// relations left without weights are the ones the walk had not reached yet as much as the ones
// nothing can satisfy. An accepted graph has no such ambiguity: every node was weighted, and a
// relation node with no weights reaches no terminal type, which is what the rewrite-tree walk
// calls an impossible relation.
//
// A refused model is answered by the rewrite-tree walk, which resolves the same relations as
// this rule does for every model in the shared corpus the graph accepts. Where the walk finds
// nothing the refusal is reported as itself: the model holds a pattern the walk does not look
// for, and the refusal names that pattern without naming a relation.
func validateWithGraph(collector *ErrorCollector, semantic *SemanticValidator, lines []string) {
	if semantic == nil || semantic.model == nil {
		return
	}

	model := semantic.model

	weighted, err := graph.NewWeightedAuthorizationModelGraphBuilder().Build(model)
	if err != nil {
		validateCyclesAndEntryPoints(collector, semantic, lines)

		// Only where the walk found nothing, because the relations it names are the same
		// relations an unresolvable cycle runs through and one finding per relation is
		// enough. Where it found nothing the refusal is about a cycle shape it does not
		// look for, which is what this names.
		if !collector.HasErrors() {
			checkRelationCycleShapes(collector, semantic, lines)
		}

		if !collector.HasErrors() {
			collector.RaiseModelUnbuildable(describeRefusal(err), err)
		}

		return
	}

	scope := &graphScope{weighted: weighted, model: model, collector: collector, lines: lines}
	for _, rule := range graphRules {
		rule.check(scope)
	}
}

// graphRefusalReasons words a build refusal from the sentinel it carries.
//
// The wording is each sentinel's own text, which is a fixed string, rather than the message
// the builder returned. The builder reports the first problem it meets and chooses that
// problem by ranging over a map, so the message names one of possibly several broken
// relations and names a different one on the next run. Reading the sentinel instead gives the
// same finding every run for the same model, which is what lets a consumer cache it, diff it,
// or compare it against a corpus.
//
// Ordered most specific first, because the constrained tuple cycle wraps the plain one and
// its text says more. TestEveryGraphSentinelHasAWording keeps this exhaustive.
var graphRefusalReasons = []struct {
	sentinel error
	reason   string
}{
	{sentinel: graph.ErrContrainstTupleCycle, reason: graph.ErrContrainstTupleCycle.Error()},
	{sentinel: graph.ErrTupleCycle, reason: graph.ErrTupleCycle.Error()},
	{sentinel: graph.ErrModelCycle, reason: graph.ErrModelCycle.Error()},
	{sentinel: graph.ErrInvalidModel, reason: graph.ErrInvalidModel.Error()},
}

// unrecognisedRefusal is the wording for a refusal carrying no sentinel this package knows.
//
// It says nothing about the model on purpose. The alternative is falling back to the
// builder's message, which is the text this exists to keep out of a finding, and a sentinel
// missing from the table is a gap to close rather than a case to paper over.
const unrecognisedRefusal = "a reason this version does not recognise"

// describeRefusal returns the wording for the sentinel a refusal carries.
func describeRefusal(err error) string {
	for _, refusal := range graphRefusalReasons {
		if errors.Is(err, refusal.sentinel) {
			return refusal.reason
		}
	}

	return unrecognisedRefusal
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

		if scope.cyclesThroughRewrites(nodeID) {
			scope.collector.RaiseNoEntryPointLoop(relation, objectType, meta, lineIndex)

			continue
		}

		scope.collector.RaiseNoEntryPoint(relation, objectType, meta, lineIndex)
	}
}

// cyclesThroughRewrites reports whether nodeID reaches itself following only the edges that
// rewrite one relation into another.
//
// That is what separates the two wordings the traversal uses. A relation that comes back to
// itself through rewrites is a potential loop; one that cannot be satisfied for any other
// reason has no entrypoint. Edges that are not rewrites are skipped rather than treated as
// disqualifying, because a relation can hold both a rewrite that loops and a direct
// assignment that does not, and the loop is still what makes it impossible. That mirrors how
// the rewrite-tree walk propagates a loop out of an operator if any one child loops, while
// reading only reachability off its direct and tupleset children.
func (s *graphScope) cyclesThroughRewrites(nodeID string) bool {
	visited := map[string]bool{}
	queue := []string{nodeID}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		edges, _ := s.weighted.GetEdgesFromNodeID(current)
		for _, edge := range edges {
			if !rewritesAnotherRelation(edge.GetEdgeType()) {
				continue
			}

			next := edge.GetTo().GetUniqueLabel()
			if next == nodeID {
				return true
			}

			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}

// rewritesAnotherRelation reports whether an edge kind stands for one relation being
// rewritten as another, as against a tuple being read or a set of edges being grouped.
func rewritesAnotherRelation(kind graph.EdgeType) bool {
	return kind == graph.ComputedEdge || kind == graph.RewriteEdge
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
