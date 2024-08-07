package graph

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

type EdgeType int64

const (
	DirectEdge EdgeType = 0 // e.g. `group`
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

	return attrs
}
