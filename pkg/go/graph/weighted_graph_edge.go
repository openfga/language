package graph

import (
	"fmt"

	"gonum.org/v1/gonum/graph/encoding"
)

type WeightedAuthorizationModelEdge struct {
	*AuthorizationModelEdge
	weights WeightMap

	// isNested signals that this edge is a self-loop. The associated node will also have this flag set to true.
	isNested bool
}

func NewWeightedAuthorizationModelEdge(edge *AuthorizationModelEdge, isNested bool) *WeightedAuthorizationModelEdge {
	return &WeightedAuthorizationModelEdge{
		AuthorizationModelEdge: edge,
		weights:                make(WeightMap),
		isNested:               isNested,
	}
}

var _ encoding.Attributer = (*WeightedAuthorizationModelEdge)(nil)

func (wn *WeightedAuthorizationModelEdge) Attributes() []encoding.Attribute {
	weightsStr := wn.weights
	labelSet := false
	attrs := make([]encoding.Attribute, 0, len(wn.AuthorizationModelEdge.Attributes()))
	for _, attr := range wn.AuthorizationModelEdge.Attributes() {
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
