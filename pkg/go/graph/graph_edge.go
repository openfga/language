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

type AuthorizationModelEdge struct {
	graph.Line

	// custom attributes
	edgeType EdgeType

	// only when edgeType == TTUEdge
	conditionedOn string
}

var _ encoding.Attributer = (*AuthorizationModelEdge)(nil)

func (n *AuthorizationModelEdge) EdgeType() EdgeType {
	return n.edgeType
}

// ConditionedOn returns the TTU relation. For example, relation
// define viewer: viewer from parent
// gives the graph "document#viewer" -> "document#viewer" and the edge
// is conditioned on "document#parent".
func (n *AuthorizationModelEdge) ConditionedOn() string {
	return n.conditionedOn
}

func (n *AuthorizationModelEdge) Attributes() []encoding.Attribute {
	switch n.edgeType {
	case DirectEdge:
		return []encoding.Attribute{
			{
				Key:   "label",
				Value: "direct",
			},
		}
	case ComputedEdge:
		return []encoding.Attribute{
			{
				Key:   "style",
				Value: "dashed",
			},
		}
	case TTUEdge:
		headLabelAttrValue := n.conditionedOn
		if headLabelAttrValue == "" {
			headLabelAttrValue = "missing"
		}

		return []encoding.Attribute{
			{
				Key:   "headlabel",
				Value: "(" + headLabelAttrValue + ")",
			},
		}
	case RewriteEdge:
		return []encoding.Attribute{}
	default:
		return []encoding.Attribute{}
	}
}
