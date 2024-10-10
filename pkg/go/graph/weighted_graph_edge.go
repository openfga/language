package graph

type WeightedAuthorizationModelEdge struct {
	from          *WeightedAuthorizationModelNode
	to            *WeightedAuthorizationModelNode
	edgeType      EdgeType
	conditionedOn string
	weights       map[string]int
}
