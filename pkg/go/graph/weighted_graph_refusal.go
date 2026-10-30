package graph

import (
	"cmp"
	"fmt"
	"maps"
	"slices"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// CycleError is the refusal for a cycle the builder cannot resolve, naming
// the relations in it. It wraps the matching sentinel.
type CycleError struct {
	kind error
	text string

	// Members holds the unique labels ("type#relation") of the relations in
	// the cycle, sorted and deduplicated. The builder stops at the first
	// cycle it cannot resolve; Build re-walks a refusal in sorted node order,
	// so the same model names the same cycle every run.
	Members []string
}

func (e *CycleError) Error() string { return e.text }

func (e *CycleError) Unwrap() error { return e.kind }

// relationMembers filters node labels to the relations among them, sorted and
// deduplicated. Operator and logical nodes share a cycle with the relations
// they belong to and are not named.
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

// cycleMembers collects the node labels on the cycle that closing closes,
// walking ancestorPath backwards until it reaches the closing node.
func cycleMembers(ancestorPath []*WeightedAuthorizationModelEdge, closing string) []string {
	labels := []string{closing}

	for i := len(ancestorPath) - 1; i >= 0; i-- {
		from := ancestorPath[i].from.uniqueLabel
		if from == closing {
			break
		}

		labels = append(labels, from)
	}

	return labels
}

// refuseDeterministically rebuilds the graph for model and re-runs the weight
// walk over the nodes in sorted order, so which cycle the refusal names does
// not depend on map iteration. Build calls it only after a refusal; the
// accepting path never runs any of this.
//
// The construction below deliberately duplicates Build rather than sharing
// code with it, so the accepting path is not restructured for the sake of the
// error path. If the two drift, the cycle-error tests fail.
func (wgb *WeightedAuthorizationModelGraphBuilder) refuseDeterministically(model *openfgav1.AuthorizationModel) error {
	wb := NewWeightedAuthorizationModelGraph()

	sortedTypeDefs := make([]*openfgav1.TypeDefinition, len(model.GetTypeDefinitions()))
	copy(sortedTypeDefs, model.GetTypeDefinitions())

	slices.SortFunc(sortedTypeDefs, func(a, b *openfgav1.TypeDefinition) int {
		return cmp.Compare(a.GetType(), b.GetType())
	})

	for _, typeDef := range sortedTypeDefs {
		wb.GetOrAddNode(typeDef.GetType(), typeDef.GetType(), SpecificType)

		sortedRelations := make([]string, 0, len(typeDef.GetRelations()))
		for relationName := range typeDef.GetRelations() {
			sortedRelations = append(sortedRelations, relationName)
		}

		slices.Sort(sortedRelations)

		for _, relation := range sortedRelations {
			uniqueLabel := typeDef.GetType() + "#" + relation
			parentNode := wb.GetOrAddNode(uniqueLabel, uniqueLabel, SpecificTypeAndRelation)

			operatorIndex := 0
			if err := wgb.parseRewrite(wb, parentNode, model, typeDef.GetRelations()[relation], typeDef, relation, &operatorIndex); err != nil {
				return err
			}
		}
	}

	return wb.assignWeightsSorted()
}

// assignWeightsSorted is AssignWeights over the nodes in sorted order, naming
// the refused cycle in a CycleError.
func (wg *WeightedAuthorizationModelGraph) assignWeightsSorted() error {
	visited := make(map[string]bool)
	ancestorPath := make([]*WeightedAuthorizationModelEdge, 0)
	tupleCycleDependencies := make(map[string][]*WeightedAuthorizationModelEdge)

	for _, nodeID := range slices.Sorted(maps.Keys(wg.nodes)) {
		node := wg.nodes[nodeID]
		if wg.isLogicalOperator(node) {
			continue
		}

		if visited[nodeID] {
			continue
		}

		tupleCyles, err := wg.calculateNodeWeight(nodeID, visited, ancestorPath, tupleCycleDependencies)
		if err != nil {
			return err
		}
		if len(tupleCyles) > 0 {
			return &CycleError{
				kind:    ErrTupleCycle,
				text:    fmt.Sprintf("tuple cycle: %d tuple cycles found without resolution", len(tupleCyles)),
				Members: wg.relationMembers(tupleCyles),
			}
		}
	}
	return nil
}
