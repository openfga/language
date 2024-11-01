package graph

type WeightedAuthorizationModelEdge struct {
	weights       map[string]int
	edgeType      EdgeType
	conditionedOn string
	from          *WeightedAuthorizationModelNode
	to            *WeightedAuthorizationModelNode
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
	return edge.from
}

// GetTo returns the to node.
func (edge *WeightedAuthorizationModelEdge) GetTo() *WeightedAuthorizationModelNode {
	return edge.to
}
