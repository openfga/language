package graph

type WeightedAuthorizationModelNode struct {
	weights     map[string]int
	nodeType    NodeType
	label       string // e.g. "group#member", UnionOperator, IntersectionOperator, ExclusionOperator
	uniqueLabel string
}

// GetWeights returns the entire weights map.
func (node *WeightedAuthorizationModelNode) GetWeights() map[string]int {
	return node.weights
}

// GetWeight returns the weight for a specific key.
func (node *WeightedAuthorizationModelNode) GetWeight(key string) (int, bool) {
	weight, exists := node.weights[key]
	return weight, exists
}

// GetNodeType returns the node type.
func (node *WeightedAuthorizationModelNode) GetNodeType() NodeType {
	return node.nodeType
}

// GetLabel returns the label.
func (node *WeightedAuthorizationModelNode) GetLabel() string {
	return node.label
}

// GetUniqueLabel returns the unique label.
func (node *WeightedAuthorizationModelNode) GetUniqueLabel() string {
	return node.uniqueLabel
}
