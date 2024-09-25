package graph

import (
	"fmt"
	"math"

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

func (weightedNode *WeightedAuthorizationModelNode) Attributes() []encoding.Attribute {
	weightsStr := weightedNode.weights
	labelSet := false
	attrs := make([]encoding.Attribute, 0, len(weightedNode.AuthorizationModelNode.Attributes()))
	for _, attr := range weightedNode.AuthorizationModelNode.Attributes() {
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

func (weightedNode *WeightedAuthorizationModelNode) assignWeightsToNode(outgoingEdges []*WeightedAuthorizationModelEdge) error {
	for _, edge := range outgoingEdges {
		for k, v := range edge.weights {
			weightedNode.weights[k] = max(weightedNode.weights[k], v)
			if weightedNode.isNested {
				weightedNode.weights[k] = math.MaxInt
			}
		}
	}

	return weightedNode.verifyNodeIsValid(outgoingEdges)
}

// verifyNodeIsValid checks that intersections and exclusions are correct. For example, an intersection operator that has
// a weight map such as "user:1, employee:2" may be valid, but if one of the edges/operands connecting it doesn't have a weight defined for
// "user", then the operator is invalid, because one of the operands of the intersection will never lead to a "user" type.// This function assumes that the input edges have weights already.
func (weightedNode *WeightedAuthorizationModelNode) verifyNodeIsValid(outgoingEdges []*WeightedAuthorizationModelEdge) error {
	if weightedNode.nodeType == OperatorNode && !(weightedNode.label == UnionOperator) {
		edgeWeights := make([]WeightMap, len(outgoingEdges))
		for i := 0; i < len(outgoingEdges); i++ {
			edgeWeights[i] = outgoingEdges[i].weights
		}
		intersect, err := IntersectionOfKeys(edgeWeights...)
		if err != nil {
			return err
		}

		if len(intersect) == 0 {
			return fmt.Errorf("%w: this operator has at leat one operand that cannot reach all types", ErrInvalidModel)
		}
	}

	return nil
}
