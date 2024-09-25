package graph

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/graph/encoding"
)

type WeightedAuthorizationModelEdge struct {
	*AuthorizationModelEdge
	weights WeightMap

	// isLoop signals that this edge is a self-loop. The associated node will also have this flag set to true.
	isLoop bool
}

func NewWeightedAuthorizationModelEdge(edge *AuthorizationModelEdge, isLoop bool) *WeightedAuthorizationModelEdge {
	return &WeightedAuthorizationModelEdge{
		AuthorizationModelEdge: edge,
		weights:                make(WeightMap),
		isLoop:                 isLoop,
	}
}

var _ encoding.Attributer = (*WeightedAuthorizationModelEdge)(nil)

func (weightedEdge *WeightedAuthorizationModelEdge) Attributes() []encoding.Attribute {
	weightsStr := weightedEdge.weights
	labelSet := false
	attrs := make([]encoding.Attribute, 0, len(weightedEdge.AuthorizationModelEdge.Attributes()))
	for _, attr := range weightedEdge.AuthorizationModelEdge.Attributes() {
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

// assignWeightsToEdge assigns the weights of the edge based on the weights of the "to" node.
func (weightedEdge *WeightedAuthorizationModelEdge) assignWeightsToEdge() error {
	neighborNode, ok := weightedEdge.To().(*WeightedAuthorizationModelNode)
	if !ok {
		return fmt.Errorf("%w: could not cast to WeightedAuthorizationModelNode", ErrBuildingGraph)
	}

	for k, v := range neighborNode.weights {
		weightedEdge.weights[k] = v
	}

	switch weightedEdge.AuthorizationModelEdge.edgeType {
	case DirectEdge, TTUEdge:
		for k, v := range neighborNode.weights {
			if v != math.MaxInt {
				weightedEdge.weights[k] = v + 1
			}
		}
	case RewriteEdge, ComputedEdge:
		// do nothing
	}

	return nil
}
