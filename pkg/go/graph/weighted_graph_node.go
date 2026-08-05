package graph

import (
	"strings"
	"sync"
)

type NodeType int64

const (
	SpecificType            NodeType = 0 // e.g. `group`
	SpecificTypeAndRelation NodeType = 1 // e.g. `group#viewer`
	OperatorNode            NodeType = 2 // e.g. union
	SpecificTypeWildcard    NodeType = 3 // e.g. `group:*`
	LogicalDirectGrouping   NodeType = 4 // e.g. `[user, employee, type1#rel]`
	LogicalTTUGrouping      NodeType = 5 // e.g. `member from parent, wherer parent can have multiple terminal types`

	UnionOperator        = "union"
	IntersectionOperator = "intersection"
	ExclusionOperator    = "exclusion"
)

type WeightedAuthorizationModelNode struct {
	weights            map[string]int
	nodeType           NodeType
	label              string   // e.g. "group#member", UnionOperator, IntersectionOperator, ExclusionOperator
	uniqueLabel        string   // e.g. "group#member", or "union:01JH0MR4H1MBFGVN37E4PRMPM3"
	wildcards          []string // e.g. "user". This means that from this node there is a path to node user:*
	recursiveRelation  string
	tupleCycle         bool
	usersetWeights     sync.Map
	directAssigns      []string // refers to the direct assignments that a node relation can have, this maps will allow to in o(1) know if a write is correct or a contextual tuple is correct
	relationDefinition string   // the relation this node belongs to, e.g. "document#parent". Empty for type and wildcard nodes; equal to uniqueLabel for relation nodes; for operator and logical nodes it is the relation whose rewrite created them.
}

// GetWeights returns the entire weights map.
func (node *WeightedAuthorizationModelNode) GetWeights() map[string]int {
	return node.weights
}

// GetRelationDefinition returns the relationDefinition field, e.g. "document#parent".
// It is empty for SpecificType and SpecificTypeWildcard nodes, which do not belong to a
// single relation. For every other node type it names the relation the node was created
// for, so a finding on an operator or logical node can be attributed back to a relation.
func (node *WeightedAuthorizationModelNode) GetRelationDefinition() string {
	return node.relationDefinition
}

// SplitRelationDefinition splits a relation definition of the form "objectType#relation" into
// its two parts, returning them in the order they appear. It returns two empty strings when
// the input is not exactly one object type and one relation, which includes the empty relation
// definition carried by type and wildcard nodes — so callers can treat "no attribution" and
// "malformed" the same way.
func SplitRelationDefinition(relationDefinition string) (objectType string, relation string) {
	parts := strings.Split(relationDefinition, "#")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", ""
	}

	return parts[0], parts[1]
}

func (node *WeightedAuthorizationModelNode) GetDirectAssigns() []string {
	return node.directAssigns
}

// GetWeight returns the weight for a specific type. It can return Infinite to indicate recursion.
func (node *WeightedAuthorizationModelNode) GetWeight(key string) (int, bool) {
	weight, exists := node.weights[key]
	return weight, exists
}

// GetNodeType returns the node type.
func (node *WeightedAuthorizationModelNode) GetNodeType() NodeType {
	return node.nodeType
}

// GetLabel returns the label, e.g. "user", "group#member", UnionOperator, IntersectionOperator or ExclusionOperator.
func (node *WeightedAuthorizationModelNode) GetLabel() string {
	return node.label
}

// GetUniqueLabel returns the unique label. It is the same as GetLabel, except for operation nodes,
// where it takes the form "operation:ULID".
func (node *WeightedAuthorizationModelNode) GetUniqueLabel() string {
	return node.uniqueLabel
}

// GetWildcards returns an array of types, e.g. "user". This means that from this node there is a path to node user:*.
func (node *WeightedAuthorizationModelNode) GetWildcards() []string {
	return node.wildcards
}

// GetRecursiveRelation returns a string of the recursive relation in a tuple cycle. A recursive relation only
// exists when the node is self-referential without any intermediate nodes of SpecificTypeAndRelation.
func (node *WeightedAuthorizationModelNode) GetRecursiveRelation() string {
	return node.recursiveRelation
}

// IsPartOfTupleCycle returns a true if the node is part of a cycle that involves more than one node of type SpecificTypeAndRelation.
func (node *WeightedAuthorizationModelNode) IsPartOfTupleCycle() bool {
	return node.tupleCycle
}
