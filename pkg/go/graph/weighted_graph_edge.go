package graph

type EdgeType int64

const (
	DirectEdge        EdgeType = 0
	RewriteEdge       EdgeType = 1
	TTUEdge           EdgeType = 2
	ComputedEdge      EdgeType = 3
	DirectLogicalEdge EdgeType = 4
	TTULogicalEdge    EdgeType = 5
	// When an edge does not have cond in the model, it will have a condition with value none.
	// This is required to differentiate when an edge need to support condition and no condition
	// like define rel1: [user, user with condX], in this case the edge will have [none, condX]
	// or an edge needs to support only condition like define rel1: [user with condX], the edge will have [condX]
	// in the case the edge does not have any condition like define rel1: [user], the edge will have [none].
	NoCond string = "none"
)

type WeightedAuthorizationModelEdge struct {
	weights            map[string]int
	edgeType           EdgeType
	tuplesetRelation   string // only present when the edgeType is a TTUEdge
	from               *WeightedAuthorizationModelNode
	to                 *WeightedAuthorizationModelNode
	wildcards          []string // e.g. "user". This means that in the direction of this edge there is a path to node user:*
	recursiveRelation  string
	tupleCycle         bool
	relationDefinition string // the relation definition that generated this edge
	// conditions on the edge. This is a flattened graph with deduplicated edges,
	// if you have a node with multiple edges to another node will be deduplicate and instead
	// only one edge but with multiple conditions,
	// define rel1: [user, user with condX]
	// then the node rel1 will have an edge pointing to the node user and with two conditions
	// one that will be none and another one that will be condX
	conditions []string
}

// GetWeights returns the entire weights map.
func (edge *WeightedAuthorizationModelEdge) GetWeights() map[string]int {
	return edge.weights
}

// GetWeight returns the weight for a specific type. It can return Infinite to indicate recursion.
func (edge *WeightedAuthorizationModelEdge) GetWeight(key string) (int, bool) {
	weight, exists := edge.weights[key]
	return weight, exists
}

// GetEdgeType returns the edge type.
func (edge *WeightedAuthorizationModelEdge) GetEdgeType() EdgeType {
	return edge.edgeType
}

// GetTuplesetRelation returns the tuplesetRelation field, e.g. "document#parent".
func (edge *WeightedAuthorizationModelEdge) GetTuplesetRelation() string {
	return edge.tuplesetRelation
}

// GetRelationDefinition returns the relationDefinition field, e.g. "document#parent".
func (edge *WeightedAuthorizationModelEdge) GetRelationDefinition() string {
	return edge.relationDefinition
}

// GetConditions returns the conditions field, e.g. "none, condX".
func (edge *WeightedAuthorizationModelEdge) GetConditions() []string {
	return edge.conditions
}

// GetFrom returns the from node.
func (edge *WeightedAuthorizationModelEdge) GetFrom() *WeightedAuthorizationModelNode {
	return edge.from
}

// GetTo returns the to node.
func (edge *WeightedAuthorizationModelEdge) GetTo() *WeightedAuthorizationModelNode {
	return edge.to
}

// GetWildcards returns an array of types, e.g. "user". This means that in the direction of this edge there is a path to node user:*.
func (edge *WeightedAuthorizationModelEdge) GetWildcards() []string {
	return edge.wildcards
}

// GetRecursiveRelation returns a string of the recursive relation in a tuple cycle. A recursive relation only
// exists when the node is self-referential without any intermediate nodes of SpecificTypeAndRelation.
func (edge *WeightedAuthorizationModelEdge) GetRecursiveRelation() string {
	return edge.recursiveRelation
}

// IsPartOfTupleCycle returns a true if the edge is part of a cycle path that involves more than one node of type SpecificTypeAndRelation.
func (edge *WeightedAuthorizationModelEdge) IsPartOfTupleCycle() bool {
	return edge.tupleCycle
}
