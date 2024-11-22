package graph

import (
	"fmt"

	"gonum.org/v1/gonum/graph/encoding"
)

type WeightedAuthorizationModelEdge struct {
	*AuthorizationModelEdge
	weights   WeightMap
	wildcards []string
}

func NewWeightedAuthorizationModelEdge(edge *AuthorizationModelEdge) *WeightedAuthorizationModelEdge {
	return &WeightedAuthorizationModelEdge{
		AuthorizationModelEdge: edge,
		weights:                make(WeightMap),
		wildcards:              make([]string, 0),
	}
}

var _ encoding.Attributer = (*WeightedAuthorizationModelEdge)(nil)

func (edge *WeightedAuthorizationModelEdge) Attributes() []encoding.Attribute {
	weightsStr := edge.weights
	labelSet := false
	attrs := make([]encoding.Attribute, 0, len(edge.AuthorizationModelEdge.Attributes()))
	for _, attr := range edge.AuthorizationModelEdge.Attributes() {
		if attr.Key == "label" {
			labelSet = true
			if len(weightsStr) > 0 {
				attr.Value += fmt.Sprintf(" - %v", weightsStr)
			}
		}
		attrs = append(attrs, attr)
	}

	if !labelSet {
		attrs = append(attrs, encoding.Attribute{
			Key:   "label",
			Value: fmt.Sprintf("%v", weightsStr),
		})
	}

	return attrs
}

// GetWeights returns the entire weights map.
func (edge *WeightedAuthorizationModelEdge) GetWeights() map[string]int {
	return edge.weights
}

// GetWeight returns the weight for a specific key.
func (edge *WeightedAuthorizationModelEdge) GetWeight(key string) (int, bool) {
	weight, exists := edge.weights[key]
	return weight, exists
}

// GetEdgeType returns the edge type.
func (edge *WeightedAuthorizationModelEdge) GetEdgeType() EdgeType {
	return edge.edgeType
}

// GetConditionedOn returns the conditionedOn field.
func (edge *WeightedAuthorizationModelEdge) GetConditionedOn() string {
	return edge.conditionedOn
}

// GetFrom returns the from node.
func (edge *WeightedAuthorizationModelEdge) GetFrom() *WeightedAuthorizationModelNode {
	from, _ := edge.AuthorizationModelEdge.From().(*WeightedAuthorizationModelNode)
	return from
}

// GetTo returns the to node.
func (edge *WeightedAuthorizationModelEdge) GetTo() *WeightedAuthorizationModelNode {
	to, _ := edge.AuthorizationModelEdge.To().(*WeightedAuthorizationModelNode)
	return to
}

// GetWildcards returns the wildcards.
func (edge *WeightedAuthorizationModelEdge) GetWildcards() []string {
	return edge.wildcards
}
