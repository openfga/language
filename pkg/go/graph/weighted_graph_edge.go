package graph

type WeightedAuthorizationModelEdge struct {
	weights       map[string]int
	edgeType      EdgeType
	conditionedOn string
	from          *WeightedAuthorizationModelNode
	to            *WeightedAuthorizationModelNode
}
