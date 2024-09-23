package graph

import (
	"fmt"

	"gonum.org/v1/gonum/graph/encoding"
)

type WeightedAuthorizationModelNode struct {
	*AuthorizationModelNode
	weights WeightMap

	// isNested signals that this node has a self-loop. That edge will also have this flag set to true.
	isNested bool
}

func NewWeightedAuthorizationModelNode(node *AuthorizationModelNode, isNested bool) *WeightedAuthorizationModelNode {
	return &WeightedAuthorizationModelNode{
		AuthorizationModelNode: node,
		weights:                make(WeightMap),
		isNested:               isNested,
	}
}

var _ encoding.Attributer = (*WeightedAuthorizationModelNode)(nil)

func (wn *WeightedAuthorizationModelNode) Attributes() []encoding.Attribute {
	weightsStr := wn.weights
	labelSet := false
	attrs := make([]encoding.Attribute, 0, len(wn.AuthorizationModelNode.Attributes()))
	for _, attr := range wn.AuthorizationModelNode.Attributes() {
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
