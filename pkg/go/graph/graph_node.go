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

type AuthorizationModelNode struct {
	graph.Node

	// custom attributes
	label       string // e.g. `union`, for DOT
	nodeType    NodeType
	uniqueLabel string // e.g. `union:01J54ND7WHGAAJTGDMFWP4FZTR`
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

func (n *AuthorizationModelNode) Label() string {
	return n.label
}

func (n *AuthorizationModelNode) NodeType() NodeType {
	return n.nodeType
}
