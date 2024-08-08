package graph

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

type EdgeType int64

const (
	DirectEdge   EdgeType = 0
	ComputedEdge EdgeType = 1
)

type AuthorizationModelEdge struct {
	graph.Line

	// custom attributes
	edgeType EdgeType
}

var _ encoding.Attributer = (*AuthorizationModelEdge)(nil)

func (n *AuthorizationModelEdge) Attributes() []encoding.Attribute {
	var attrs []encoding.Attribute

	if n.edgeType == DirectEdge {
		attrs = append(attrs, encoding.Attribute{
			Key:   "label",
			Value: "direct",
		})
	}

	if n.edgeType == ComputedEdge {
		attrs = append(attrs, encoding.Attribute{
			Key:   "style",
			Value: "dashed",
		})
	}

	return attrs
}
