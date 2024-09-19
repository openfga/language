package graph

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

type NodeType int64

const (
	SpecificType            NodeType = 0 // e.g. `group`
	SpecificTypeAndRelation NodeType = 1 // e.g. `group#viewer`
	OperatorNode            NodeType = 2 // e.g. union
)

var _ graph.Node = (*AuthorizationModelNode)(nil)

type AuthorizationModelNode struct {
	Node graph.Node

	// custom attributes
	label       string // e.g. `union`, for DOT
	nodeType    NodeType
	uniqueLabel string // e.g. `union:01J54ND7WHGAAJTGDMFWP4FZTR`
}

func (n *AuthorizationModelNode) ID() int64 {
	return n.Node.ID()
}

var _ encoding.Attributer = (*AuthorizationModelNode)(nil)

func (n *AuthorizationModelNode) Attributes() []encoding.Attribute {
	var attrs []encoding.Attribute

	attrs = append(attrs, encoding.Attribute{
		Key:   "label",
		Value: n.label,
	})

	return attrs
}
