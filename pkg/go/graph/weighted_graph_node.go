package graph

import (
	"fmt"

	"gonum.org/v1/gonum/graph/encoding"
)

type WeightedAuthorizationModelNode struct {
	*AuthorizationModelNode
	weights   WeightMap
	wildcards []string
}

func NewWeightedAuthorizationModelNode(node *AuthorizationModelNode) *WeightedAuthorizationModelNode {
	n := &WeightedAuthorizationModelNode{
		AuthorizationModelNode: node,
		weights:                make(WeightMap),
		wildcards:              make([]string, 0),
	}

	if node.nodeType == SpecificTypeWildcard {
		n.wildcards = append(n.wildcards, node.uniqueLabel[:len(node.uniqueLabel)-2])
	}

	return n
}

var _ encoding.Attributer = (*WeightedAuthorizationModelNode)(nil)

func (node *WeightedAuthorizationModelNode) Attributes() []encoding.Attribute {
	weightsStr := node.weights
	labelSet := false
	attrs := make([]encoding.Attribute, 0, len(node.AuthorizationModelNode.Attributes()))
	for _, attr := range node.AuthorizationModelNode.Attributes() {
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

// GetWildcards returns the wildcards.
func (node *WeightedAuthorizationModelNode) GetWildcards() []string {
	return node.wildcards
}
