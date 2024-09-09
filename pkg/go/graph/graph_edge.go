package graph

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

type EdgeType int64

const (
	DirectEdge  EdgeType = 0
	RewriteEdge EdgeType = 1
	TTUEdge     EdgeType = 2
)

type AuthorizationModelEdge struct {
	graph.WeightedLine

	// custom attributes
	edgeType EdgeType

	// only when edgeType == TTUEdge
	conditionedOn string
}

var _ encoding.Attributer = (*AuthorizationModelEdge)(nil)

func (n *AuthorizationModelEdge) Attributes() []encoding.Attribute {
	var attrs []encoding.Attribute

	if n.edgeType == DirectEdge {
		attrs = append(attrs, encoding.Attribute{
			Key:   "label",
			Value: fmt.Sprintf("w: %.1f", n.Weight()),
		})
	}

	if n.edgeType == RewriteEdge {
		attrs = append(attrs, encoding.Attribute{
			Key:   "style",
			Value: "dashed",
		})
	}

	if n.edgeType == TTUEdge {
		headLabelAttrValue := n.conditionedOn
		if headLabelAttrValue == "" {
			headLabelAttrValue = "missing"
		}

		attrs = append(attrs, encoding.Attribute{
			Key:   "headlabel",
			Value: headLabelAttrValue,
		})
	}

	return attrs
}
