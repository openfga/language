package graph

type WeightedAuthorizationModelNode struct {
	label       string // e.g. "group#member", UnionOperator, IntersectionOperator, ExclusionOperator
	nodeType    NodeType
	uniqueLabel string
	weights     map[string]int
}
