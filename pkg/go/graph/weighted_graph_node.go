package graph

type WeightedAuthorizationModelNode struct {
	weights     map[string]int
	nodeType    NodeType
	label       string // e.g. "group#member", UnionOperator, IntersectionOperator, ExclusionOperator
	uniqueLabel string
}
