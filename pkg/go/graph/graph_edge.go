package graph

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

type EdgeType int64

const (
	DirectEdge   EdgeType = 0
	RewriteEdge  EdgeType = 1
	TTUEdge      EdgeType = 2
	ComputedEdge EdgeType = 3
)

var _ graph.Line = (*AuthorizationModelEdge)(nil)

type AuthorizationModelEdge struct {
	Line graph.Line

	// custom attributes
	edgeType EdgeType

	// only when edgeType == TTUEdge
	conditionedOn string
}

func (n *AuthorizationModelEdge) From() graph.Node {
	return n.Line.From()
}

func (n *AuthorizationModelEdge) To() graph.Node {
	return n.Line.To()
}

func (n *AuthorizationModelEdge) ReversedLine() graph.Line {
	panic("implement me")
}

func (n *AuthorizationModelEdge) ID() int64 {
	return n.Line.ID()
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
