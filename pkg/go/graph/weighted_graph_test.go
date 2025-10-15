package graph

import (
	"math"
	"strings"
	"testing"

	"github.com/openfga/language/pkg/go/errors"
	"github.com/stretchr/testify/require"
)

// addNode is a helper function that adds a node to the graph and handles the error.
// It's used to maintain backward compatibility with the old AddNode method in tests.
func addNode(t *testing.T, graph *WeightedAuthorizationModelGraph, uniqueLabel, label string, nodeType NodeType, relationDefinition string) {
	_, err := graph.AddNode(uniqueLabel, label, nodeType, relationDefinition)
	require.NoError(t, err, "Failed to add node with unique label: %s", uniqueLabel)
}

/*
type user
type state

	  relations
	      define can_view: [user] or member or owner
	      define member: [user]
	      define owner: [user] and approved_member
	      define approved_member: [user:*]
		  define can_apply: (approved_member but not can_view) or owner
*/
func TestValidWeight1OneTerminalType(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "state-can_view", "can_view", SpecificTypeAndRelation, "state-can_view")
	addNode(t, graph, "state-can_view-or", UnionOperator, OperatorNode, "state-can_view")
	addNode(t, graph, "state-member", "member", SpecificTypeAndRelation, "state-member")
	addNode(t, graph, "state-owner", "owner", SpecificTypeAndRelation, "state-owner")
	addNode(t, graph, "state-owner-and", IntersectionOperator, OperatorNode, "state-owner")
	addNode(t, graph, "state-approved_member", "approved_member", SpecificTypeAndRelation, "state-approved_member")
	addNode(t, graph, "state-can_apply", "can_apply", SpecificTypeAndRelation, "state-can_apply")
	addNode(t, graph, "state-can_apply-but", ExclusionOperator, OperatorNode, "state-can_apply")
	addNode(t, graph, "state-can_apply-or", UnionOperator, OperatorNode, "state-can_apply")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "user:*", "user:*", SpecificTypeWildcard, "")

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "state-can_view", "", nil)
	graph.AddEdge("state-member", "user", DirectEdge, "state-member", "", nil)
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "user", DirectEdge, "state-owner", "", nil)
	graph.AddEdge("state-approved_member", "user:*", DirectEdge, "state-approved_member", "", nil)
	graph.AddEdge("state-can_apply", "state-can_apply-or", RewriteEdge, "state-can_apply", "", nil)
	graph.AddEdge("state-can_apply-or", "state-can_apply-but", RewriteEdge, "state-can_apply", "", nil)
	graph.AddEdge("state-can_apply-or", "state-owner", ComputedEdge, "state-can_apply", "", nil)
	graph.AddEdge("state-can_apply-but", "state-approved_member", ComputedEdge, "state-can_apply", "", nil)
	graph.AddEdge("state-can_apply-but", "state-can_view", ComputedEdge, "state-can_apply", "", nil)

	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, 1, graph.nodes["state-can_view"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_view-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner-and"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-approved_member"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_apply"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_apply-but"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_apply-or"].weights["user"])

	require.Len(t, graph.nodes["state-can_view"].weights, 1)
	require.Len(t, graph.nodes["state-member"].weights, 1)
	require.Len(t, graph.nodes["state-owner"].weights, 1)
	require.Len(t, graph.nodes["state-can_view-or"].weights, 1)
	require.Len(t, graph.nodes["state-owner-and"].weights, 1)
	require.Len(t, graph.nodes["state-approved_member"].weights, 1)
	require.Len(t, graph.nodes["state-can_apply"].weights, 1)
	require.Len(t, graph.nodes["state-can_apply-but"].weights, 1)
	require.Len(t, graph.nodes["state-can_apply-or"].weights, 1)

	require.Len(t, graph.nodes["state-approved_member"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-approved_member"].wildcards[0])
	require.Len(t, graph.nodes["state-owner"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-owner"].wildcards[0])
	require.Len(t, graph.nodes["state-owner-and"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-owner-and"].wildcards[0])
	require.Len(t, graph.nodes["state-can_apply"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-can_apply"].wildcards[0])
	require.Len(t, graph.nodes["state-can_apply-but"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-can_apply-but"].wildcards[0])
	require.Len(t, graph.nodes["state-can_apply-or"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-can_apply-or"].wildcards[0])
	require.Empty(t, graph.nodes["state-member"].wildcards)
	require.Len(t, graph.nodes["state-can_view-or"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-can_view-or"].wildcards[0])
}

/*
type user
type state

	relations
	    define can_view: [user] or member or owner
	    define member: [user]
	    define owner: [user] and can_view
*/
func TestInvalidWeight1WithAndModelCycle(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "state-can_view", "can_view", SpecificTypeAndRelation, "state-can_view")
	addNode(t, graph, "state-can_view-or", UnionOperator, OperatorNode, "state-can_view")
	addNode(t, graph, "state-member", "member", SpecificTypeAndRelation, "state-member")
	addNode(t, graph, "state-owner", "owner", SpecificTypeAndRelation, "state-owner")
	addNode(t, graph, "state-owner-and", IntersectionOperator, OperatorNode, "state-owner")
	addNode(t, graph, "user", "user", SpecificType, "")

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "state-can_view", "", nil)
	graph.AddEdge("state-member", "user", DirectEdge, "state-member", "", nil)
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "state-can_view", ComputedEdge, "state-owner", "", nil)

	require.Empty(t, graph.nodes["state-can_view"].wildcards)
	require.Empty(t, graph.nodes["state-can_view-or"].wildcards)
	require.Empty(t, graph.nodes["state-member"].wildcards)
	require.Empty(t, graph.nodes["state-owner"].wildcards)
	require.Empty(t, graph.nodes["state-owner-and"].wildcards)
	require.Empty(t, graph.nodes["user"].wildcards)

	err := graph.AssignWeights()
	require.ErrorIs(t, err, errors.ErrCycle)
}

/*
type user
type employee
type state

	  relations
	      define can_view: [user] or member or owner
	      define member: [user]
	      define owner: [employee] and approved_member
	      define approved_member: [employee]
		  define can_apply: (approved_member but not can_view) or member
*/
func TestValidWeight1MultipleTerminalTypes(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "state-can_view", "can_view", SpecificTypeAndRelation, "state-can_view")
	addNode(t, graph, "state-can_view-or", UnionOperator, OperatorNode, "state-can_view")
	addNode(t, graph, "state-member", "member", SpecificTypeAndRelation, "state-member")
	addNode(t, graph, "state-owner", "owner", SpecificTypeAndRelation, "state-owner")
	addNode(t, graph, "state-owner-and", IntersectionOperator, OperatorNode, "state-owner")
	addNode(t, graph, "state-approved_member", "approved_member", SpecificTypeAndRelation, "state-approved_member")
	addNode(t, graph, "state-can_apply", "can_apply", SpecificTypeAndRelation, "state-can_apply")
	addNode(t, graph, "state-can_apply-but", ExclusionOperator, OperatorNode, "state-can_apply")
	addNode(t, graph, "state-can_apply-or", UnionOperator, OperatorNode, "state-can_apply")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "employee", "employee", SpecificType, "")

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "state-can_view", "", nil)
	graph.AddEdge("state-member", "user", DirectEdge, "state-member", "", nil)
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "employee", DirectEdge, "state-owner", "", nil)
	graph.AddEdge("state-approved_member", "employee", DirectEdge, "state-approved_member", "", nil)
	graph.AddEdge("state-can_apply", "state-can_apply-or", RewriteEdge, "state-can_apply", "", nil)
	graph.AddEdge("state-can_apply-or", "state-can_apply-but", RewriteEdge, "state-can_apply", "", nil)
	graph.AddEdge("state-can_apply-or", "state-member", ComputedEdge, "state-can_apply", "", nil)
	graph.AddEdge("state-can_apply-but", "state-approved_member", ComputedEdge, "state-can_apply", "", nil)
	graph.AddEdge("state-can_apply-but", "state-can_view", ComputedEdge, "state-can_apply", "", nil)

	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, 1, graph.nodes["state-can_view"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_view"].weights["employee"])
	require.Equal(t, 1, graph.nodes["state-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner"].weights["employee"])
	require.Equal(t, 1, graph.nodes["state-can_view-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_view-or"].weights["employee"])
	require.Equal(t, 1, graph.nodes["state-owner-and"].weights["employee"])
	require.Equal(t, 1, graph.nodes["state-approved_member"].weights["employee"])
	require.Equal(t, 1, graph.nodes["state-can_apply"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_apply"].weights["employee"])
	require.Equal(t, 1, graph.nodes["state-can_apply-but"].weights["employee"])
	require.Equal(t, 1, graph.nodes["state-can_apply-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_apply-or"].weights["employee"])

	require.Len(t, graph.nodes["state-can_view"].weights, 2)
	require.Len(t, graph.nodes["state-member"].weights, 1)
	require.Len(t, graph.nodes["state-owner"].weights, 1)
	require.Len(t, graph.nodes["state-can_view-or"].weights, 2)
	require.Len(t, graph.nodes["state-owner-and"].weights, 1)
	require.Len(t, graph.nodes["state-approved_member"].weights, 1)
	require.Len(t, graph.nodes["state-can_apply"].weights, 2)
	require.Len(t, graph.nodes["state-can_apply-but"].weights, 1)
	require.Len(t, graph.nodes["state-can_apply-or"].weights, 2)

	require.Empty(t, graph.nodes["state-can_view"].wildcards)
	require.Empty(t, graph.nodes["state-owner"].wildcards)
	require.Empty(t, graph.nodes["state-member"].wildcards)
	require.Empty(t, graph.nodes["state-can_view-or"].wildcards)
	require.Empty(t, graph.nodes["state-owner-and"].wildcards)
	require.Empty(t, graph.nodes["state-approved_member"].wildcards)
	require.Empty(t, graph.nodes["state-can_apply"].wildcards)
	require.Empty(t, graph.nodes["state-can_apply-but"].wildcards)
	require.Empty(t, graph.nodes["state-can_apply-or"].wildcards)
}

/*
type user
type employee
type state

	relations
	    define can_view: [user] or member or owner
	    define member: [user]
	    define owner: [employee] and approved_member
	    define approved_member: [user]
*/
func TestInvalidWeight1NotMatchingTerminalTypes(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "state-can_view", "can_view", SpecificTypeAndRelation, "state-can_view")
	addNode(t, graph, "state-can_view-or", UnionOperator, OperatorNode, "state-can_view")
	addNode(t, graph, "state-member", "member", SpecificTypeAndRelation, "state-member")
	addNode(t, graph, "state-owner", "owner", SpecificTypeAndRelation, "state-owner")
	addNode(t, graph, "state-owner-and", IntersectionOperator, OperatorNode, "state-owner")
	addNode(t, graph, "state-approved_member", "approved_member", SpecificTypeAndRelation, "state-approved_member")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "employee", "employee", SpecificType, "")

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "state-can_view", "", nil)
	graph.AddEdge("state-member", "user", DirectEdge, "state-member", "", nil)
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "employee", DirectEdge, "state-owner", "", nil)
	graph.AddEdge("state-approved_member", "user", DirectEdge, "state-approved_member", "", nil)

	err := graph.AssignWeights()
	require.ErrorIs(t, err, errors.ErrNoEntrypoints)
	require.True(t, strings.Contains(err.Error(), "no entrypoints defined"))
}

/*
type user
type state

	relations
	    define can_view: [user] or member or owner
	    define member: [user:*]
	    define owner: [user] and approved_member
	    define approved_member: [user]

type transition

	relations
	    define start: [state]
	    define end: [state]
	    define can_apply: [user] and can_view from start and can_view from end
*/
func TestValidWeight2TTUOneTerminalType(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "state-can_view", "can_view", SpecificTypeAndRelation, "state-can_view")
	addNode(t, graph, "state-can_view-or", UnionOperator, OperatorNode, "state-can_view")
	addNode(t, graph, "state-member", "member", SpecificTypeAndRelation, "state-member")
	addNode(t, graph, "state-owner", "owner", SpecificTypeAndRelation, "state-owner")
	addNode(t, graph, "state-owner-and", IntersectionOperator, OperatorNode, "state-owner")
	addNode(t, graph, "state-approved_member", "approved_member", SpecificTypeAndRelation, "state-approved_member")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "state", "state", SpecificType, "")
	addNode(t, graph, "user:*", "user:*", SpecificTypeWildcard, "")
	addNode(t, graph, "transition-start", "start", SpecificTypeAndRelation, "transition-start")
	addNode(t, graph, "transition-end", "end", SpecificTypeAndRelation, "transition-end")
	addNode(t, graph, "transition-can_apply", "can_apply", SpecificTypeAndRelation, "transition-can_apply")
	addNode(t, graph, "transition-can_apply-and", IntersectionOperator, OperatorNode, "transition-can_apply")

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "state-can_view", "", nil)
	graph.AddEdge("state-member", "user:*", DirectEdge, "state-member", "", nil)
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "state-owner-and", "", nil)
	graph.AddEdge("state-owner-and", "user", DirectEdge, "state-owner-and", "", nil)
	graph.AddEdge("state-approved_member", "user", DirectEdge, "state-approved_member", "", nil)
	graph.AddEdge("transition-start", "state", DirectEdge, "transition-start", "", nil)
	graph.AddEdge("transition-end", "state", DirectEdge, "transition-end", "", nil)
	graph.AddEdge("transition-can_apply", "transition-can_apply-and", RewriteEdge, "transition-can_apply", "", nil)
	graph.AddEdge("transition-can_apply-and", "user", DirectEdge, "transition-can_apply", "", nil)
	graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-can_apply", "transition-start", nil)
	graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-can_apply", "transition-end", nil)

	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, 1, graph.nodes["state-can_view"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_view-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner-and"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-approved_member"].weights["user"])
	require.Equal(t, 1, graph.nodes["transition-start"].weights["state"])
	require.Equal(t, 1, graph.nodes["transition-end"].weights["state"])
	require.Equal(t, 2, graph.nodes["transition-can_apply"].weights["user"])
	require.Equal(t, 2, graph.nodes["transition-can_apply-and"].weights["user"])

	require.Len(t, graph.nodes["state-can_view"].weights, 1)
	require.Len(t, graph.nodes["state-member"].weights, 1)
	require.Len(t, graph.nodes["state-owner"].weights, 1)
	require.Len(t, graph.nodes["state-can_view-or"].weights, 1)
	require.Len(t, graph.nodes["state-owner-and"].weights, 1)
	require.Len(t, graph.nodes["state-approved_member"].weights, 1)
	require.Len(t, graph.nodes["transition-end"].weights, 1)
	require.Len(t, graph.nodes["transition-start"].weights, 1)
	require.Len(t, graph.nodes["transition-can_apply"].weights, 1)
	require.Len(t, graph.nodes["transition-can_apply-and"].weights, 1)

	require.Len(t, graph.nodes["state-can_view"].wildcards, 1)
	require.Len(t, graph.nodes["state-member"].wildcards, 1)
	require.Len(t, graph.nodes["state-can_view-or"].wildcards, 1)
	require.Len(t, graph.nodes["transition-can_apply"].wildcards, 1)
	require.Len(t, graph.nodes["transition-can_apply-and"].wildcards, 1)
	require.Empty(t, graph.nodes["state-owner"].wildcards)
	require.Empty(t, graph.nodes["state-owner-and"].wildcards)
	require.Empty(t, graph.nodes["state-approved_member"].wildcards)
	require.Empty(t, graph.nodes["transition-end"].wildcards)
	require.Empty(t, graph.nodes["transition-start"].wildcards)
}

/*
type user
type state

	relations
	    define can_view: [user] or member or owner
	    define member: [user:*]
	    define owner: [user] and approved_member
	    define approved_member: [user]

type transition

	relations
	    define can_apply: [user, state#can_view]
*/
func TestValidWeight2UserSetOneTerminalType(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "state-can_view", "can_view", SpecificTypeAndRelation, "state-can_view")
	addNode(t, graph, "state-can_view-or", UnionOperator, OperatorNode, "state-can_view")
	addNode(t, graph, "state-member", "member", SpecificTypeAndRelation, "state-member")
	addNode(t, graph, "state-owner", "owner", SpecificTypeAndRelation, "state-owner")
	addNode(t, graph, "state-owner-and", IntersectionOperator, OperatorNode, "state-owner")
	addNode(t, graph, "state-approved_member", "approved_member", SpecificTypeAndRelation, "state-approved_member")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "state", "state", SpecificType, "")
	addNode(t, graph, "user:*", "user:*", SpecificTypeWildcard, "")
	addNode(t, graph, "transition-can_apply", "can_apply", SpecificTypeAndRelation, "transition-can_apply")

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "state-can_view", "", nil)
	graph.AddEdge("state-member", "user:*", DirectEdge, "state-member", "", nil)
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "state-owner", "", nil)
	graph.AddEdge("state-owner-and", "user", DirectEdge, "state-owner", "", nil)
	graph.AddEdge("state-approved_member", "user", DirectEdge, "state-approved_member", "", nil)
	graph.AddEdge("transition-can_apply", "user", DirectEdge, "transition-can_apply", "", nil)
	graph.AddEdge("transition-can_apply", "state-can_view", DirectEdge, "transition-can_apply", "", nil)

	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, 1, graph.nodes["state-can_view"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_view-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner-and"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-approved_member"].weights["user"])
	require.Equal(t, 2, graph.nodes["transition-can_apply"].weights["user"])

	require.Len(t, graph.nodes["state-can_view"].weights, 1)
	require.Len(t, graph.nodes["state-member"].weights, 1)
	require.Len(t, graph.nodes["state-owner"].weights, 1)
	require.Len(t, graph.nodes["state-can_view-or"].weights, 1)
	require.Len(t, graph.nodes["state-owner-and"].weights, 1)
	require.Len(t, graph.nodes["state-approved_member"].weights, 1)
	require.Len(t, graph.nodes["transition-can_apply"].weights, 1)

	require.Len(t, graph.nodes["state-can_view"].wildcards, 1)
	require.Len(t, graph.nodes["state-member"].wildcards, 1)
	require.Len(t, graph.nodes["state-can_view-or"].wildcards, 1)
	require.Len(t, graph.nodes["transition-can_apply"].wildcards, 1)
	require.Empty(t, graph.nodes["state-owner"].wildcards)
	require.Empty(t, graph.nodes["state-owner-and"].wildcards)
	require.Empty(t, graph.nodes["state-approved_member"].wildcards)
}

/*
type user
type employee
type company

	relations
	  define member: [user]
	  define owner: [user]
	  define approved_member: company#member or owner

type group

	relations
	  define approved_member: [employee]

type license

	relations
	  define active_member: approved_member from owner
	  define owner: [company, group]
*/
func TestValidWeight2MultipleTerminalType(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "company-member", "member", SpecificTypeAndRelation, "company-member")
	addNode(t, graph, "company-owner", "owner", SpecificTypeAndRelation, "company-owner")
	addNode(t, graph, "company-approved_member", "approved_member", SpecificTypeAndRelation, "company-approved_member")
	addNode(t, graph, "company-approved_member-or", UnionOperator, OperatorNode, "company-approved_member")
	addNode(t, graph, "group-approved_member", "approved_member", SpecificTypeAndRelation, "group-approved_member")
	addNode(t, graph, "employee", "employee", SpecificType, "")
	addNode(t, graph, "company", "company", SpecificType, "")
	addNode(t, graph, "group", "group", SpecificType, "")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "license-active_member", "active_member", SpecificTypeAndRelation, "license-active_member")
	addNode(t, graph, "license-owner", "owner", SpecificTypeAndRelation, "license-owner")

	graph.AddEdge("company-member", "user", DirectEdge, "company-member", "", nil)
	graph.AddEdge("company-owner", "user", DirectEdge, "company-owner", "", nil)
	graph.AddEdge("company-approved_member", "company-approved_member-or", RewriteEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member-or", "company-member", DirectEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member-or", "company-owner", ComputedEdge, "company-approved_member", "", nil)
	graph.AddEdge("group-approved_member", "employee", DirectEdge, "group-approved_member", "", nil)
	graph.AddEdge("license-active_member", "company-approved_member", TTUEdge, "license-active_member", "company-owner", nil)
	graph.AddEdge("license-active_member", "group-approved_member", TTUEdge, "license-active_member", "company-owner", nil)
	graph.AddEdge("license-owner", "company", DirectEdge, "license-owner", "", nil)
	graph.AddEdge("license-owner", "group", DirectEdge, "license-owner", "", nil)

	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, 1, graph.nodes["company-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["company-owner"].weights["user"])
	require.Equal(t, 2, graph.nodes["company-approved_member"].weights["user"])
	require.Equal(t, 2, graph.nodes["company-approved_member-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["group-approved_member"].weights["employee"])
	require.Equal(t, 3, graph.nodes["license-active_member"].weights["user"])
	require.Equal(t, 2, graph.nodes["license-active_member"].weights["employee"])
	require.Equal(t, 1, graph.nodes["license-owner"].weights["company"])
	require.Equal(t, 1, graph.nodes["license-owner"].weights["group"])

	require.Len(t, graph.nodes["company-member"].weights, 1)
	require.Len(t, graph.nodes["company-owner"].weights, 1)
	require.Len(t, graph.nodes["company-approved_member"].weights, 1)
	require.Len(t, graph.nodes["company-approved_member-or"].weights, 1)
	require.Len(t, graph.nodes["group-approved_member"].weights, 1)
	require.Len(t, graph.nodes["license-active_member"].weights, 2)
	require.Len(t, graph.nodes["license-owner"].weights, 2)

	require.Empty(t, graph.nodes["company-member"].wildcards)
	require.Empty(t, graph.nodes["company-owner"].wildcards)
	require.Empty(t, graph.nodes["company-approved_member"].wildcards)
	require.Empty(t, graph.nodes["company-approved_member-or"].wildcards)
	require.Empty(t, graph.nodes["group-approved_member"].wildcards)
	require.Empty(t, graph.nodes["license-active_member"].wildcards)
	require.Empty(t, graph.nodes["license-owner"].wildcards)
}

/*
type user
type company

	relations
	  define member: [user]
	  define owner: approved_member
	  define approved_member: company#member or owner
*/
func TestInvalidWeight2ModelCycle(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "company-member", "member", SpecificTypeAndRelation, "company-member")
	addNode(t, graph, "company-owner", "owner", SpecificTypeAndRelation, "company-owner")
	addNode(t, graph, "company-approved_member", "approved_member", SpecificTypeAndRelation, "company-approved_member")
	addNode(t, graph, "company-approved_member-or", UnionOperator, OperatorNode, "company-approved_member")
	addNode(t, graph, "user", "user", SpecificType, "")

	graph.AddEdge("company-member", "user", DirectEdge, "company-member", "", nil)
	graph.AddEdge("company-owner", "company-approved_member", ComputedEdge, "company-owner", "", nil)
	graph.AddEdge("company-approved_member", "company-approved_member-or", RewriteEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member-or", "company-member", DirectEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member-or", "company-owner", ComputedEdge, "company-approved_member", "", nil)

	err := graph.AssignWeights()
	require.ErrorIs(t, err, errors.ErrCycle)
}

/*
type user
type company

	relations
	  define member: [user]
	  define executive: [employee]
	  define approved_member: [user, company#member]
	  define can_execute: executive but not approved_member
*/
func TestWeight2ButNotMistmatchType(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "company-member", "member", SpecificTypeAndRelation, "company-member")
	addNode(t, graph, "company-executive", "executive", SpecificTypeAndRelation, "company-executive")
	addNode(t, graph, "company-approved_member", "approved_member", SpecificTypeAndRelation, "company-approved_member")
	addNode(t, graph, "company-can_execute", "can_execute", SpecificTypeAndRelation, "company-can_execute")
	addNode(t, graph, "employee", "employee", SpecificType, "")
	addNode(t, graph, "company-can_execute-but", ExclusionOperator, OperatorNode, "company-can_execute")
	addNode(t, graph, "user", "user", SpecificType, "")

	graph.AddEdge("company-member", "user", DirectEdge, "company-member", "", nil)
	graph.AddEdge("company-executive", "employee", DirectEdge, "company-executive", "", nil)
	graph.AddEdge("company-approved_member", "user", DirectEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member", "company-member", DirectEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-can_execute", "company-can_execute-but", RewriteEdge, "company-can_execute", "", nil)
	graph.AddEdge("company-can_execute-but", "company-executive", ComputedEdge, "company-can_execute", "", nil)
	graph.AddEdge("company-can_execute-but", "company-approved_member", ComputedEdge, "company-can_execute", "", nil)

	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, 1, graph.nodes["company-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["company-executive"].weights["employee"])
	require.Equal(t, 2, graph.nodes["company-approved_member"].weights["user"])
	require.Equal(t, 1, graph.nodes["company-can_execute"].weights["employee"])
	_, found := graph.nodes["company-can_execute"].weights["user"]
	require.False(t, found)
}

/*
type document
 relations
  define rel1: rel2 or rel3 or rel4
  define rel2: [user]
  define rel3: rel5
  define rel5: [employee]
  define rel4: rel6 or rel7
  define rel6: [document#rel1]
  define rel7: [document#rel4]
*/

func TestValidTupleCycle(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "document-rel4", "rel4", SpecificTypeAndRelation, "document-rel4")
	addNode(t, graph, "document-rel4-or", UnionOperator, OperatorNode, "document-rel4")
	addNode(t, graph, "document-rel6", "rel6", SpecificTypeAndRelation, "document-rel6")
	addNode(t, graph, "document-rel7", "rel7", SpecificTypeAndRelation, "document-rel7")
	addNode(t, graph, "document-rel1", "rel1", SpecificTypeAndRelation, "document-rel1")
	addNode(t, graph, "document-rel1-or", UnionOperator, OperatorNode, "document-rel1")
	addNode(t, graph, "document-rel2", "rel2", SpecificTypeAndRelation, "document-rel2")
	addNode(t, graph, "document-rel3", "rel3", SpecificTypeAndRelation, "document-rel3")
	addNode(t, graph, "document-rel5", "rel5", SpecificTypeAndRelation, "document-rel5")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "employee", "employee", SpecificType, "")

	graph.AddEdge("document-rel4", "document-rel4-or", RewriteEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel6", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel7", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel6", "document-rel1", DirectEdge, "document-rel6", "", nil)
	graph.AddEdge("document-rel7", "document-rel4", DirectEdge, "document-rel7", "", nil)
	graph.AddEdge("document-rel1", "document-rel1-or", RewriteEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-or", "document-rel2", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-or", "document-rel3", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-or", "document-rel4", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel2", "user", DirectEdge, "document-rel2", "", nil)
	graph.AddEdge("document-rel3", "document-rel5", ComputedEdge, "document-rel3", "", nil)
	graph.AddEdge("document-rel5", "employee", DirectEdge, "document-rel5", "", nil)

	Infinite := math.MaxInt32
	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, Infinite, graph.nodes["document-rel1"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel1"].weights["employee"])
	require.Equal(t, 1, graph.nodes["document-rel2"].weights["user"])
	require.Equal(t, 1, graph.nodes["document-rel3"].weights["employee"])
	require.Equal(t, 1, graph.nodes["document-rel5"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["document-rel4"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["document-rel4"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel6"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["document-rel6"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel7"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["document-rel7"].weights["user"])

	require.Len(t, graph.nodes["document-rel1"].weights, 2)
	require.Len(t, graph.nodes["document-rel2"].weights, 1)
	require.Len(t, graph.nodes["document-rel3"].weights, 1)
	require.Len(t, graph.nodes["document-rel4"].weights, 2)
	require.Len(t, graph.nodes["document-rel5"].weights, 1)
	require.Len(t, graph.nodes["document-rel6"].weights, 2)
	require.Len(t, graph.nodes["document-rel7"].weights, 2)

	require.Empty(t, graph.nodes["document-rel1"].wildcards)
	require.Empty(t, graph.nodes["document-rel2"].wildcards)
	require.Empty(t, graph.nodes["document-rel3"].wildcards)
	require.Empty(t, graph.nodes["document-rel4"].wildcards)
	require.Empty(t, graph.nodes["document-rel5"].wildcards)
	require.Empty(t, graph.nodes["document-rel6"].wildcards)
	require.Empty(t, graph.nodes["document-rel7"].wildcards)
}

/*
type document
 relations
  define rel1: rel2 or rel3 or rel4
     define rel2: [user]
   define rel3: rel5
  define rel5: [employee]
  define rel4: (rel6 or rel7) and rel5
  define rel6: [document#rel1]
  define rel 7: [document#rel4]
*/

func TestInvalidTupleCycleWithInterceptionOfTerminalTypes(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "document-rel4", "rel4", SpecificTypeAndRelation, "document-rel4")
	addNode(t, graph, "document-rel4-or", UnionOperator, OperatorNode, "document-rel4")
	addNode(t, graph, "document-rel4-and", IntersectionOperator, OperatorNode, "document-rel4")
	addNode(t, graph, "document-rel6", "rel6", SpecificTypeAndRelation, "document-rel6")
	addNode(t, graph, "document-rel7", "rel7", SpecificTypeAndRelation, "document-rel7")
	addNode(t, graph, "document-rel1", "rel1", SpecificTypeAndRelation, "document-rel1")
	addNode(t, graph, "document-rel1-or", UnionOperator, OperatorNode, "document-rel1")
	addNode(t, graph, "document-rel2", "rel2", SpecificTypeAndRelation, "document-rel2")
	addNode(t, graph, "document-rel3", "rel3", SpecificTypeAndRelation, "document-rel3")
	addNode(t, graph, "document-rel5", "rel5", SpecificTypeAndRelation, "document-rel5")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "employee", "employee", SpecificType, "")

	graph.AddEdge("document-rel4", "document-rel4-and", RewriteEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-and", "document-rel4-or", RewriteEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-and", "document-rel5", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel6", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel7", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel6", "document-rel1", DirectEdge, "document-rel6", "", nil)
	graph.AddEdge("document-rel7", "document-rel4", DirectEdge, "document-rel7", "", nil)
	graph.AddEdge("document-rel1", "document-rel1-or", RewriteEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-or", "document-rel2", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-or", "document-rel3", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-or", "document-rel4", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel2", "user", DirectEdge, "document-rel2", "", nil)
	graph.AddEdge("document-rel3", "document-rel5", ComputedEdge, "document-rel3", "", nil)
	graph.AddEdge("document-rel5", "employee", DirectEdge, "document-rel5", "", nil)

	err := graph.AssignWeights()
	require.ErrorIs(t, err, errors.ErrConstraintTupleCycle)
}

/*
type document

	relations
	 define rel1: rel2 and rel3 and rel4
	 define rel2: [user]
	 define rel3: [employee, user]
	 define rel4: rel5 or rel6 or rel7
	 define rel5: [user]
	 define rel6: [employee]
	 define rel7: [document#rel4]
*/
func TestValidTupleCycleWithInterceptionNotInCycle(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "document-rel4", "rel4", SpecificTypeAndRelation, "document-rel4")
	addNode(t, graph, "document-rel4-or", UnionOperator, OperatorNode, "document-rel4")
	addNode(t, graph, "document-rel6", "rel6", SpecificTypeAndRelation, "document-rel6")
	addNode(t, graph, "document-rel7", "rel7", SpecificTypeAndRelation, "document-rel7")
	addNode(t, graph, "document-rel1", "rel1", SpecificTypeAndRelation, "document-rel1")
	addNode(t, graph, "document-rel1-and", IntersectionOperator, OperatorNode, "document-rel1")
	addNode(t, graph, "document-rel2", "rel2", SpecificTypeAndRelation, "document-rel2")
	addNode(t, graph, "document-rel3", "rel3", SpecificTypeAndRelation, "document-rel3")
	addNode(t, graph, "document-rel5", "rel5", SpecificTypeAndRelation, "document-rel5")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "employee", "employee", SpecificType, "")

	graph.AddEdge("document-rel4", "document-rel4-or", RewriteEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel6", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel7", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel5", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel6", "employee", DirectEdge, "document-rel6", "", nil)
	graph.AddEdge("document-rel7", "document-rel4", DirectEdge, "document-rel7", "", nil)
	graph.AddEdge("document-rel1", "document-rel1-and", RewriteEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-and", "document-rel2", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-and", "document-rel3", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel1-and", "document-rel4", ComputedEdge, "document-rel1", "", nil)
	graph.AddEdge("document-rel2", "user", DirectEdge, "document-rel2", "", nil)
	graph.AddEdge("document-rel3", "user", DirectEdge, "document-rel3", "", nil)
	graph.AddEdge("document-rel3", "employee", DirectEdge, "document-rel3", "", nil)
	graph.AddEdge("document-rel5", "user", DirectEdge, "document-rel5", "", nil)
	graph.AddEdge("document-rel6", "employee", DirectEdge, "document-rel6", "", nil)

	Infinite := math.MaxInt32
	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, Infinite, graph.nodes["document-rel1"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel1-and"].weights["user"])
	require.Equal(t, 1, graph.nodes["document-rel2"].weights["user"])
	require.Equal(t, 1, graph.nodes["document-rel3"].weights["employee"])
	require.Equal(t, 1, graph.nodes["document-rel3"].weights["user"])
	require.Equal(t, 1, graph.nodes["document-rel5"].weights["user"])
	require.Equal(t, 1, graph.nodes["document-rel6"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["document-rel4"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["document-rel4"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel4-or"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["document-rel4-or"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel7"].weights["employee"])
	require.Equal(t, Infinite, graph.nodes["document-rel7"].weights["user"])

	require.Len(t, graph.nodes["document-rel1"].weights, 1)
	require.Len(t, graph.nodes["document-rel1-and"].weights, 1)
	require.Len(t, graph.nodes["document-rel2"].weights, 1)
	require.Len(t, graph.nodes["document-rel3"].weights, 2)
	require.Len(t, graph.nodes["document-rel4"].weights, 2)
	require.Len(t, graph.nodes["document-rel4-or"].weights, 2)
	require.Len(t, graph.nodes["document-rel5"].weights, 1)
	require.Len(t, graph.nodes["document-rel6"].weights, 1)
	require.Len(t, graph.nodes["document-rel7"].weights, 2)

	require.Empty(t, graph.nodes["document-rel1"].wildcards)
	require.Empty(t, graph.nodes["document-rel2"].wildcards)
	require.Empty(t, graph.nodes["document-rel3"].wildcards)
	require.Empty(t, graph.nodes["document-rel4"].wildcards)
	require.Empty(t, graph.nodes["document-rel5"].wildcards)
	require.Empty(t, graph.nodes["document-rel6"].wildcards)
	require.Empty(t, graph.nodes["document-rel7"].wildcards)
	require.Empty(t, graph.nodes["document-rel1-and"].wildcards)
	require.Empty(t, graph.nodes["document-rel4-or"].wildcards)
}

/*
type document
	relations
	 define rel3: [employee, user, document#rel3]
	 define rel4: rel5 or rel4 from rel6
	 define rel5: [user]
	 define rel6: [document]
*/

func TestValidRecursionUsersetAndTTU(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "document-rel4", "rel4", SpecificTypeAndRelation, "document-rel4")
	addNode(t, graph, "document-rel4-or", UnionOperator, OperatorNode, "document-rel4")
	addNode(t, graph, "document-rel6", "rel6", SpecificTypeAndRelation, "document-rel6")
	addNode(t, graph, "document-rel3", "rel3", SpecificTypeAndRelation, "document-rel3")
	addNode(t, graph, "document-rel5", "rel5", SpecificTypeAndRelation, "document-rel5")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "employee", "employee", SpecificType, "")
	addNode(t, graph, "document", "document", SpecificType, "")

	graph.AddEdge("document-rel4", "document-rel4-or", RewriteEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel5", ComputedEdge, "document-rel4", "", nil)
	graph.AddEdge("document-rel4-or", "document-rel4", TTUEdge, "document-rel4", "document-rel6", nil)
	graph.AddEdge("document-rel6", "document", DirectEdge, "document-rel6", "", nil)
	graph.AddEdge("document-rel5", "user", DirectEdge, "document-rel5", "", nil)
	graph.AddEdge("document-rel3", "user", DirectEdge, "document-rel3", "", nil)
	graph.AddEdge("document-rel3", "employee", DirectEdge, "document-rel3", "", nil)
	graph.AddEdge("document-rel3", "document-rel3", DirectEdge, "document-rel3", "", nil)

	Infinite := math.MaxInt32
	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, Infinite, graph.nodes["document-rel3"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel3"].weights["employee"])
	require.Equal(t, 1, graph.nodes["document-rel5"].weights["user"])
	require.Equal(t, 1, graph.nodes["document-rel6"].weights["document"])
	require.Equal(t, Infinite, graph.nodes["document-rel4"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel4-or"].weights["user"])

	require.Len(t, graph.nodes["document-rel3"].weights, 2)
	require.Len(t, graph.nodes["document-rel4"].weights, 1)
	require.Len(t, graph.nodes["document-rel4-or"].weights, 1)
	require.Len(t, graph.nodes["document-rel5"].weights, 1)
	require.Len(t, graph.nodes["document-rel6"].weights, 1)

	require.Empty(t, graph.nodes["document-rel3"].wildcards)
	require.Empty(t, graph.nodes["document-rel4"].wildcards)
	require.Empty(t, graph.nodes["document-rel5"].wildcards)
	require.Empty(t, graph.nodes["document-rel6"].wildcards)
	require.Empty(t, graph.nodes["document-rel4-or"].wildcards)
}

/*
type user
type state
  relations
      define can_view: [user] or member
      define member: [user]
type transition
  relations
      define start: [state]
      define end: [state]
      define can_apply: [user] and can_view from start and can_view from end
type group
  relations
    define owner: [user, transition#can_apply]
    define max_owner: [group#owner, group#max_owner]
*/

func TestValidRecursionWithWeight3(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "state-can_view", "can_view", SpecificTypeAndRelation, "state-can_view")
	addNode(t, graph, "state-can_view-or", UnionOperator, OperatorNode, "state-can_view")
	addNode(t, graph, "state-member", "member", SpecificTypeAndRelation, "state-member")
	addNode(t, graph, "user", "user", SpecificType, "user")
	addNode(t, graph, "state", "state", SpecificType, "state")
	addNode(t, graph, "transition-start", "start", SpecificTypeAndRelation, "transition-start")
	addNode(t, graph, "transition-end", "end", SpecificTypeAndRelation, "transition-end")
	addNode(t, graph, "transition-can_apply", "can_apply", SpecificTypeAndRelation, "transition-can_apply")
	addNode(t, graph, "transition-can_apply-and", IntersectionOperator, OperatorNode, "transition-can_apply")
	addNode(t, graph, "group-owner", "owner", SpecificTypeAndRelation, "group-owner")
	addNode(t, graph, "group-max_owner", "max_owner", SpecificTypeAndRelation, "group-max_owner")

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "state-can_view", "", nil)
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "state-can_view", "", nil)
	graph.AddEdge("state-member", "user", DirectEdge, "state-member", "", nil)
	graph.AddEdge("transition-start", "state", DirectEdge, "transition-start", "", nil)
	graph.AddEdge("transition-end", "state", DirectEdge, "transition-end", "", nil)
	graph.AddEdge("transition-can_apply", "transition-can_apply-and", RewriteEdge, "transition-can_apply", "", nil)
	graph.AddEdge("transition-can_apply-and", "user", DirectEdge, "transition-can_apply", "", nil)
	graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-can_apply", "transition-start", nil)
	graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-can_apply", "transition-end", nil)
	graph.AddEdge("group-owner", "user", DirectEdge, "group-owner", "", nil)
	graph.AddEdge("group-owner", "transition-can_apply", DirectEdge, "group-owner", "", nil)
	graph.AddEdge("group-max_owner", "group-owner", DirectEdge, "group-max_owner", "", nil)
	graph.AddEdge("group-max_owner", "group-max_owner", DirectEdge, "group-max_owner", "", nil)

	Infinite := math.MaxInt32
	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, 1, graph.nodes["state-can_view"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_view-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["transition-start"].weights["state"])
	require.Equal(t, 1, graph.nodes["transition-end"].weights["state"])
	require.Equal(t, 2, graph.nodes["transition-can_apply-and"].weights["user"])
	require.Equal(t, 2, graph.nodes["transition-can_apply"].weights["user"])
	require.Equal(t, 3, graph.nodes["group-owner"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["group-max_owner"].weights["user"])

	require.Len(t, graph.nodes["group-owner"].weights, 1)
	require.Len(t, graph.nodes["state-member"].weights, 1)
	require.Len(t, graph.nodes["state-can_view"].weights, 1)
	require.Len(t, graph.nodes["state-can_view-or"].weights, 1)
	require.Len(t, graph.nodes["transition-start"].weights, 1)
	require.Len(t, graph.nodes["transition-end"].weights, 1)
	require.Len(t, graph.nodes["group-max_owner"].weights, 1)
	require.Len(t, graph.nodes["transition-can_apply-and"].weights, 1)
	require.Len(t, graph.nodes["transition-can_apply"].weights, 1)

	require.Empty(t, graph.nodes["group-owner"].wildcards)
	require.Empty(t, graph.nodes["state-can_view"].wildcards)
	require.Empty(t, graph.nodes["state-can_view-or"].wildcards)
	require.Empty(t, graph.nodes["transition-start"].wildcards)
	require.Empty(t, graph.nodes["transition-end"].wildcards)
	require.Empty(t, graph.nodes["state-member"].wildcards)
	require.Empty(t, graph.nodes["group-max_owner"].wildcards)
	require.Empty(t, graph.nodes["transition-can_apply-and"].wildcards)
	require.Empty(t, graph.nodes["transition-can_apply"].wildcards)
}

/*
model
  schema 1.1
type user
type employee
type company
  relations
    define approved_member: member and user_in_context
    define member: [user, employee, user:*]
    define facilitator: member
    define user_in_context: [user with x_less_than]
type group
  relations
    define approved_member: member but not user_in_context
    define member: [user, user with x_greater_than]
    define user_in_context: [user]
    define reader: member
    define assignee: reader
type license
  relations
    define active_holder: holder or holder_member
    define holder_member: member from owner
    define holder_approved_member: approved_member from owner
    define holder: [user]
    define parent: [license]
    define trust_holder: [user, user:*] or trust_holder from parent
    define owner: [company, group, group with x_condition]
type tier
  relations
   define subscriber: [company#facilitator]
   define assignee: [group#assignee, group#user_in_context, group#user_in_context with x_bigger_than]
   define subtier_owner: [user, tier#subtier_owner]
   define assignee_sub: subscriber and assignee and subtier_owner
type module
  relations
    define associated_license: [license]
    define module_holder: holder_member from associated_license
    define module_user: active_holder from associated_license
type feature
  relations
    define associated_module: [module]
    define associated_tier: [tier]
    define tier_can_access: subscriber from associated_tier
    define can_access: module_user from associated_module and subscriber from associated_tier

*/

func TestValidRecursionWithMultipleWeightsAndTypes(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()

	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "user:*", "user:*", SpecificTypeWildcard, "")
	addNode(t, graph, "employee", "employee", SpecificType, "")
	addNode(t, graph, "company", "company", SpecificType, "")
	addNode(t, graph, "group", "group", SpecificType, "")
	addNode(t, graph, "license", "license", SpecificType, "")
	addNode(t, graph, "tier", "tier", SpecificType, "")
	addNode(t, graph, "module", "module", SpecificType, "")
	addNode(t, graph, "feature", "feature", SpecificType, "")

	addNode(t, graph, "company-approved_member", "approved_member", SpecificTypeAndRelation, "company-approved_member")
	addNode(t, graph, "company-approved_member-and", IntersectionOperator, OperatorNode, "company-approved_member")
	addNode(t, graph, "company-member", "member", SpecificTypeAndRelation, "company-member")
	addNode(t, graph, "company-facilitator", "facilitator", SpecificTypeAndRelation, "company-facilitator")
	addNode(t, graph, "company-user_in_context", "user_in_context", SpecificTypeAndRelation, "company-user_in_context")

	addNode(t, graph, "group-approved_member", "approved_member", SpecificTypeAndRelation, "group-approved_member")
	addNode(t, graph, "group-approved_member-butnot", ExclusionOperator, OperatorNode, "group-approved_member")
	addNode(t, graph, "group-member", "member", SpecificTypeAndRelation, "group-member")
	addNode(t, graph, "group-user_in_context", "user_in_context", SpecificTypeAndRelation, "group-user_in_context")
	addNode(t, graph, "group-reader", "reader", SpecificTypeAndRelation, "group-reader")
	addNode(t, graph, "group-assignee", "assignee", SpecificTypeAndRelation, "group-assignee")

	addNode(t, graph, "license-active_holder", "active_holder", SpecificTypeAndRelation, "license-active_holder")
	addNode(t, graph, "license-active_holder-or", UnionOperator, OperatorNode, "license-active_holder")
	addNode(t, graph, "license-holder_member", "holder_member", SpecificTypeAndRelation, "license-holder_member")
	addNode(t, graph, "license-holder_approved_member", "holder_approved_member", SpecificTypeAndRelation, "license-holder_approved_member")
	addNode(t, graph, "license-holder", "holder", SpecificTypeAndRelation, "license-holder")
	addNode(t, graph, "license-parent", "parent", SpecificTypeAndRelation, "license-parent")
	addNode(t, graph, "license-trust_holder", "trust_holder", SpecificTypeAndRelation, "license-trust_holder")
	addNode(t, graph, "license-trust_holder-or", UnionOperator, OperatorNode, "license-trust_holder")
	addNode(t, graph, "license-owner", "owner", SpecificTypeAndRelation, "license-owner")

	addNode(t, graph, "tier-subscriber", "subscriber", SpecificTypeAndRelation, "tier-subscriber")
	addNode(t, graph, "tier-assignee", "assignee", SpecificTypeAndRelation, "tier-assignee")
	addNode(t, graph, "tier-subtier_owner", "subtier_owner", SpecificTypeAndRelation, "tier-subtier_owner")
	addNode(t, graph, "tier-assignee_sub", "assignee_sub", SpecificTypeAndRelation, "tier-assignee_sub")
	addNode(t, graph, "tier-assignee_sub-and", IntersectionOperator, OperatorNode, "tier-assignee_sub")

	addNode(t, graph, "module-associated_license", "associated_license", SpecificTypeAndRelation, "module-associated_license")
	addNode(t, graph, "module-module_holder", "module_holder", SpecificTypeAndRelation, "module-module_holder")
	addNode(t, graph, "module-module_user", "module_user", SpecificTypeAndRelation, "module-module_user")

	addNode(t, graph, "feature-associated_module", "associated_module", SpecificTypeAndRelation, "feature-associated_module")
	addNode(t, graph, "feature-associated_tier", "associated_tier", SpecificTypeAndRelation, "feature-associated_tier")
	addNode(t, graph, "feature-tier_can_access", "tier_can_access", SpecificTypeAndRelation, "feature-tier_can_access")
	addNode(t, graph, "feature-can_access", "can_access", SpecificTypeAndRelation, "feature-can_access")
	addNode(t, graph, "feature-can_access-and", IntersectionOperator, OperatorNode, "feature-can_access")

	graph.AddEdge("company-member", "user", DirectEdge, "company-member", "", nil)
	graph.AddEdge("company-member", "employee", DirectEdge, "company-member", "", nil)
	graph.AddEdge("company-member", "user:*", DirectEdge, "company-member", "", nil)
	graph.AddEdge("company-facilitator", "company-member", ComputedEdge, "company-facilitator", "", nil)
	graph.AddEdge("company-user_in_context", "user", DirectEdge, "company-user_in_context", "", nil)
	graph.AddEdge("company-approved_member", "company-approved_member-and", RewriteEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member-and", "company-member", ComputedEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member-and", "company-user_in_context", ComputedEdge, "company-approved_member", "", nil)

	graph.AddEdge("group-member", "user", DirectEdge, "group-member", "", nil)
	graph.AddEdge("group-member", "user", DirectEdge, "group-member", "", nil)
	graph.AddEdge("group-user_in_context", "user", DirectEdge, "group-user_in_context", "", nil)
	graph.AddEdge("group-reader", "group-member", ComputedEdge, "group-reader", "", nil)
	graph.AddEdge("group-assignee", "group-reader", ComputedEdge, "group-assignee", "", nil)
	graph.AddEdge("group-approved_member", "group-approved_member-butnot", RewriteEdge, "group-approved_member", "", nil)
	graph.AddEdge("group-approved_member-butnot", "group-user_in_context", ComputedEdge, "group-approved_member", "", nil)
	graph.AddEdge("group-approved_member-butnot", "group-member", ComputedEdge, "group-approved_member", "", nil)

	graph.AddEdge("license-active_holder", "license-active_holder-or", RewriteEdge, "license-active_holder", "", nil)
	graph.AddEdge("license-active_holder-or", "license-holder", ComputedEdge, "license-active_holder", "", nil)
	graph.AddEdge("license-active_holder-or", "license-holder_member", ComputedEdge, "license-active_holder", "", nil)
	graph.AddEdge("license-owner", "company", DirectEdge, "license-owner", "", nil)
	graph.AddEdge("license-owner", "group", DirectEdge, "license-owner", "", nil)
	graph.AddEdge("license-owner", "group", DirectEdge, "license-owner", "x_condition", nil)
	graph.AddEdge("license-holder_member", "group-member", TTUEdge, "license-holder_member", "license-owner", nil)
	graph.AddEdge("license-holder_member", "company-member", TTUEdge, "license-holder_member", "license-owner", nil)
	graph.AddEdge("license-holder_member", "group-member", TTUEdge, "license-holder_member", "license-owner", nil)
	graph.AddEdge("license-holder_approved_member", "group-approved_member", TTUEdge, "license-holder_approved_member", "license-owner", nil)
	graph.AddEdge("license-holder_approved_member", "company-approved_member", TTUEdge, "license-holder_approved_member", "license-owner", nil)
	graph.AddEdge("license-holder_approved_member", "group-approved_member", TTUEdge, "license-holder_approved_member", "license-owner", nil)
	graph.AddEdge("license-holder", "user", DirectEdge, "license-holder", "", nil)
	graph.AddEdge("license-parent", "license", DirectEdge, "license-parent", "", nil)
	graph.AddEdge("license-trust_holder", "license-trust_holder-or", RewriteEdge, "license-trust_holder", "", nil)
	graph.AddEdge("license-trust_holder-or", "user", DirectEdge, "license-trust_holder", "", nil)
	graph.AddEdge("license-trust_holder-or", "user:*", DirectEdge, "license-trust_holder", "", nil)
	graph.AddEdge("license-trust_holder-or", "license-trust_holder", TTUEdge, "license-trust_holder", "license-parent", nil)

	graph.AddEdge("tier-subscriber", "company-facilitator", DirectEdge, "tier-subscriber", "", nil)
	graph.AddEdge("tier-assignee", "group-assignee", DirectEdge, "tier-assignee", "", nil)
	graph.AddEdge("tier-assignee", "group-user_in_context", DirectEdge, "tier-assignee", "", nil)
	graph.AddEdge("tier-assignee", "group-user_in_context", DirectEdge, "tier-assignee", "x_bigger_than", nil)
	graph.AddEdge("tier-subtier_owner", "user", DirectEdge, "tier-subtier_owner", "", nil)
	graph.AddEdge("tier-subtier_owner", "tier-subtier_owner", DirectEdge, "tier-subtier_owner", "", nil)
	graph.AddEdge("tier-assignee_sub", "tier-assignee_sub-and", RewriteEdge, "tier-assignee_sub", "", nil)
	graph.AddEdge("tier-assignee_sub-and", "tier-subscriber", ComputedEdge, "tier-assignee_sub", "", nil)
	graph.AddEdge("tier-assignee_sub-and", "tier-assignee", ComputedEdge, "tier-assignee_sub", "", nil)
	graph.AddEdge("tier-assignee_sub-and", "tier-subtier_owner", ComputedEdge, "tier-assignee_sub", "", nil)

	graph.AddEdge("module-associated_license", "license", DirectEdge, "module-associated_license", "", nil)
	graph.AddEdge("module-module_holder", "license-holder_member", TTUEdge, "module-module_holder", "module-associated_license", nil)
	graph.AddEdge("module-module_user", "license-active_holder", TTUEdge, "module-module_user", "module-associated_license", nil)

	graph.AddEdge("feature-associated_module", "module", DirectEdge, "feature-associated_module", "", nil)
	graph.AddEdge("feature-associated_tier", "tier", DirectEdge, "feature-associated_tier", "", nil)
	graph.AddEdge("feature-tier_can_access", "tier-subscriber", TTUEdge, "feature-tier_can_access", "feature-associated_tier", nil)
	graph.AddEdge("feature-can_access", "feature-can_access-and", RewriteEdge, "feature-can_access", "", nil)
	graph.AddEdge("feature-can_access-and", "module-module_user", TTUEdge, "feature-can_access", "feature-associated_module", nil)
	graph.AddEdge("feature-can_access-and", "tier-subscriber", TTUEdge, "feature-can_access", "feature-associated_tier", nil)

	Infinite := math.MaxInt32
	err := graph.AssignWeights()
	require.NoError(t, err)

	require.Equal(t, 1, graph.nodes["company-approved_member"].weights["user"])
	require.Equal(t, 1, graph.nodes["company-approved_member-and"].weights["user"])
	require.Equal(t, 1, graph.nodes["company-member"].weights["employee"])
	require.Equal(t, 1, graph.nodes["company-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["company-facilitator"].weights["user"])
	require.Equal(t, 1, graph.nodes["company-facilitator"].weights["employee"])
	require.Equal(t, 1, graph.nodes["company-user_in_context"].weights["user"])
	require.Equal(t, 1, graph.nodes["group-approved_member"].weights["user"])
	require.Equal(t, 1, graph.nodes["group-approved_member-butnot"].weights["user"])
	require.Equal(t, 1, graph.nodes["group-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["group-user_in_context"].weights["user"])
	require.Equal(t, 1, graph.nodes["group-reader"].weights["user"])
	require.Equal(t, 1, graph.nodes["group-assignee"].weights["user"])
	require.Equal(t, 2, graph.nodes["license-active_holder"].weights["user"])
	require.Equal(t, 2, graph.nodes["license-active_holder"].weights["employee"])
	require.Equal(t, 2, graph.nodes["license-active_holder-or"].weights["user"])
	require.Equal(t, 2, graph.nodes["license-active_holder-or"].weights["employee"])
	require.Equal(t, 2, graph.nodes["license-holder_member"].weights["user"])
	require.Equal(t, 2, graph.nodes["license-holder_member"].weights["employee"])
	require.Equal(t, 2, graph.nodes["license-holder_approved_member"].weights["user"])
	require.Equal(t, 1, graph.nodes["license-holder"].weights["user"])
	require.Equal(t, 1, graph.nodes["license-parent"].weights["license"])
	require.Equal(t, Infinite, graph.nodes["license-trust_holder"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["license-trust_holder-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["license-owner"].weights["company"])
	require.Equal(t, 1, graph.nodes["license-owner"].weights["group"])
	require.Equal(t, 2, graph.nodes["tier-subscriber"].weights["user"])
	require.Equal(t, 2, graph.nodes["tier-subscriber"].weights["employee"])
	require.Equal(t, 2, graph.nodes["tier-assignee"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["tier-subtier_owner"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["tier-assignee_sub"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["tier-assignee_sub-and"].weights["user"])
	require.Equal(t, 1, graph.nodes["module-associated_license"].weights["license"])
	require.Equal(t, 3, graph.nodes["module-module_holder"].weights["user"])
	require.Equal(t, 3, graph.nodes["module-module_holder"].weights["employee"])
	require.Equal(t, 3, graph.nodes["module-module_user"].weights["user"])
	require.Equal(t, 3, graph.nodes["module-module_user"].weights["employee"])
	require.Equal(t, 1, graph.nodes["feature-associated_module"].weights["module"])
	require.Equal(t, 1, graph.nodes["feature-associated_tier"].weights["tier"])
	require.Equal(t, 3, graph.nodes["feature-tier_can_access"].weights["user"])
	require.Equal(t, 3, graph.nodes["feature-tier_can_access"].weights["employee"])
	require.Equal(t, 4, graph.nodes["feature-can_access"].weights["user"])
	require.Equal(t, 4, graph.nodes["feature-can_access"].weights["employee"])
	require.Equal(t, 4, graph.nodes["feature-can_access-and"].weights["user"])
	require.Equal(t, 4, graph.nodes["feature-can_access-and"].weights["employee"])

	require.Len(t, graph.nodes["company-approved_member"].weights, 1)
	require.Len(t, graph.nodes["company-approved_member-and"].weights, 1)
	require.Len(t, graph.nodes["company-member"].weights, 2)
	require.Len(t, graph.nodes["company-facilitator"].weights, 2)
	require.Len(t, graph.nodes["company-user_in_context"].weights, 1)
	require.Len(t, graph.nodes["group-approved_member"].weights, 1)
	require.Len(t, graph.nodes["group-approved_member-butnot"].weights, 1)
	require.Len(t, graph.nodes["group-member"].weights, 1)
	require.Len(t, graph.nodes["group-user_in_context"].weights, 1)
	require.Len(t, graph.nodes["group-reader"].weights, 1)
	require.Len(t, graph.nodes["group-assignee"].weights, 1)
	require.Len(t, graph.nodes["license-active_holder"].weights, 2)
	require.Len(t, graph.nodes["license-active_holder-or"].weights, 2)
	require.Len(t, graph.nodes["license-holder_member"].weights, 2)
	require.Len(t, graph.nodes["license-holder_approved_member"].weights, 1)
	require.Len(t, graph.nodes["license-holder"].weights, 1)
	require.Len(t, graph.nodes["license-parent"].weights, 1)
	require.Len(t, graph.nodes["license-trust_holder"].weights, 1)
	require.Len(t, graph.nodes["license-trust_holder-or"].weights, 1)
	require.Len(t, graph.nodes["license-owner"].weights, 2)
	require.Len(t, graph.nodes["tier-subscriber"].weights, 2)
	require.Len(t, graph.nodes["tier-assignee"].weights, 1)
	require.Len(t, graph.nodes["tier-subtier_owner"].weights, 1)
	require.Len(t, graph.nodes["tier-assignee_sub"].weights, 1)
	require.Len(t, graph.nodes["tier-assignee_sub-and"].weights, 1)
	require.Len(t, graph.nodes["module-associated_license"].weights, 1)
	require.Len(t, graph.nodes["module-module_holder"].weights, 2)
	require.Len(t, graph.nodes["module-module_user"].weights, 2)
	require.Len(t, graph.nodes["feature-associated_module"].weights, 1)
	require.Len(t, graph.nodes["feature-associated_tier"].weights, 1)
	require.Len(t, graph.nodes["feature-tier_can_access"].weights, 2)
	require.Len(t, graph.nodes["feature-can_access"].weights, 2)
	require.Len(t, graph.nodes["feature-can_access-and"].weights, 2)

	require.Len(t, graph.nodes["company-approved_member"].wildcards, 1)
	require.Len(t, graph.nodes["company-approved_member-and"].wildcards, 1)
	require.Len(t, graph.nodes["company-member"].wildcards, 1)
	require.Len(t, graph.nodes["company-facilitator"].wildcards, 1)
	require.Len(t, graph.nodes["license-active_holder"].wildcards, 1)
	require.Len(t, graph.nodes["license-active_holder-or"].wildcards, 1)
	require.Len(t, graph.nodes["license-holder_member"].wildcards, 1)
	require.Len(t, graph.nodes["license-holder_approved_member"].wildcards, 1)
	require.Len(t, graph.nodes["license-trust_holder"].wildcards, 1)
	require.Len(t, graph.nodes["license-trust_holder-or"].wildcards, 1)
	require.Len(t, graph.nodes["tier-subscriber"].wildcards, 1)
	require.Len(t, graph.nodes["tier-assignee_sub"].wildcards, 1)
	require.Len(t, graph.nodes["tier-assignee_sub-and"].wildcards, 1)
	require.Len(t, graph.nodes["module-module_holder"].wildcards, 1)
	require.Len(t, graph.nodes["module-module_user"].wildcards, 1)
	require.Len(t, graph.nodes["feature-tier_can_access"].wildcards, 1)
	require.Len(t, graph.nodes["feature-can_access"].wildcards, 1)
	require.Len(t, graph.nodes["feature-can_access-and"].wildcards, 1)
	require.Empty(t, graph.nodes["company-user_in_context"].wildcards)
	require.Empty(t, graph.nodes["group-approved_member"].wildcards)
	require.Empty(t, graph.nodes["group-approved_member-butnot"].wildcards)
	require.Empty(t, graph.nodes["group-member"].wildcards)
	require.Empty(t, graph.nodes["group-user_in_context"].wildcards)
	require.Empty(t, graph.nodes["group-reader"].wildcards)
	require.Empty(t, graph.nodes["group-assignee"].wildcards)
	require.Empty(t, graph.nodes["license-holder"].wildcards)
	require.Empty(t, graph.nodes["license-parent"].wildcards)
	require.Empty(t, graph.nodes["license-owner"].wildcards)
	require.Empty(t, graph.nodes["tier-assignee"].wildcards)
	require.Empty(t, graph.nodes["tier-subtier_owner"].wildcards)
	require.Empty(t, graph.nodes["module-associated_license"].wildcards)
	require.Empty(t, graph.nodes["feature-associated_module"].wildcards)
	require.Empty(t, graph.nodes["feature-associated_tier"].wildcards)
}

/*
type user
type state

	relations
	    define member: ([user] or member from parent) or parent_member from parent
	    define parent: [state]
	    define parent_member: [user:*, state#member]
*/
func TestValidMixedRecursionWithTupleCycles(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	addNode(t, graph, "state-member", "member", SpecificTypeAndRelation, "state-member")
	addNode(t, graph, "state-member-or", UnionOperator, OperatorNode, "state-member")
	addNode(t, graph, "state-member-or-or", UnionOperator, OperatorNode, "state-member")
	addNode(t, graph, "state-parent", "parent", SpecificTypeAndRelation, "state-parent")
	addNode(t, graph, "state-parent_member", "parent_member", SpecificTypeAndRelation, "state-parent_member")
	addNode(t, graph, "user", "user", SpecificType, "")
	addNode(t, graph, "user:*", "user", SpecificTypeWildcard, "")
	addNode(t, graph, "state", "state", SpecificType, "")

	graph.AddEdge("state-member", "state-member-or", RewriteEdge, "state-member", "", nil)
	graph.AddEdge("state-member-or", "state-member-or-or", RewriteEdge, "state-member", "", nil)
	graph.AddEdge("state-member-or", "state-parent_member", TTUEdge, "state-member", "state-parent", nil)
	graph.AddEdge("state-member-or-or", "user", DirectEdge, "state-member", "", nil)
	graph.AddEdge("state-member-or-or", "state-member", TTUEdge, "state-member", "state-parent", nil)
	graph.AddEdge("state-parent_member", "user:*", DirectEdge, "state-parent_member", "", nil)
	graph.AddEdge("state-parent_member", "state-member", DirectEdge, "state-parent_member", "", nil)
	graph.AddEdge("state-parent", "state", DirectEdge, "state-parent", "", nil)

	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, Infinite, graph.nodes["state-member"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["state-member-or"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["state-member-or-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-parent"].weights["state"])
	require.Equal(t, Infinite, graph.nodes["state-parent_member"].weights["user"])

	require.Len(t, graph.nodes["state-member"].weights, 1)
	require.Len(t, graph.nodes["state-member-or"].weights, 1)
	require.Len(t, graph.nodes["state-member-or-or"].weights, 1)
	require.Len(t, graph.nodes["state-parent"].weights, 1)
	require.Len(t, graph.nodes["state-parent_member"].weights, 1)

	require.Len(t, graph.nodes["state-member"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-member"].wildcards[0])
	require.Len(t, graph.nodes["state-parent_member"].wildcards, 1)
	require.Equal(t, "user", graph.nodes["state-parent_member"].wildcards[0])

	require.Equal(t, "state-member", graph.nodes["state-member"].recursiveRelation)
	require.Equal(t, "state-member", graph.nodes["state-member-or"].recursiveRelation)
	require.Equal(t, "state-member", graph.nodes["state-member-or-or"].recursiveRelation)
	require.Equal(t, "", graph.nodes["state-parent"].recursiveRelation)
	require.Equal(t, "", graph.nodes["state-parent_member"].recursiveRelation)

	require.True(t, graph.nodes["state-member"].tupleCycle)
	require.True(t, graph.nodes["state-member-or"].tupleCycle)
	require.True(t, graph.nodes["state-member-or-or"].tupleCycle)
	require.False(t, graph.nodes["state-parent"].tupleCycle)
	require.True(t, graph.nodes["state-parent_member"].tupleCycle)
}

func TestAddNode(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	var authModelErr *errors.AuthorizationModelError
	// Test 1: Adding a node that doesn't exist should succeed
	node1, err := graph.AddNode("document", "document", SpecificType, "")
	require.NoError(t, err)
	require.NotNil(t, node1)
	require.Equal(t, "document", node1.uniqueLabel)
	require.Equal(t, "document", node1.label)
	require.Equal(t, SpecificType, node1.nodeType)
	require.Empty(t, node1.wildcards)

	// Test 2: Adding a node with the same unique label should fail
	node2, err := graph.AddNode("document", "document", SpecificType, "")
	require.Error(t, err)
	require.Nil(t, node2)
	require.ErrorAs(t, err, &authModelErr)
	require.Contains(t, err.Error(), "node with unique label document already exists")

	// Test 3: Adding a node with a different label but the same uniqueLabel should fail
	node3, err := graph.AddNode("document", "document_updated", SpecificType, "")
	require.Error(t, err)
	require.Nil(t, node3)
	require.ErrorAs(t, err, &authModelErr)

	// Test 4: Adding a node with a different unique label should succeed
	node4, err := graph.AddNode("folder", "folder", SpecificType, "")
	require.NoError(t, err)
	require.NotNil(t, node4)
	require.Equal(t, "folder", node4.uniqueLabel)
	require.Equal(t, "folder", node4.label)

	// Test 5: Test with wildcard type
	node5, err := graph.AddNode("user:*", "user:*", SpecificTypeWildcard, "")
	require.NoError(t, err)
	require.NotNil(t, node5)
	require.Equal(t, "user:*", node5.uniqueLabel)
	require.Equal(t, "user:*", node5.label)
	require.Equal(t, SpecificTypeWildcard, node5.nodeType)
	require.Len(t, node5.wildcards, 1)
	require.Equal(t, "user", node5.wildcards[0])
}

func TestCompareGetOrAddNodeWithAddNode(t *testing.T) {
	t.Parallel()
	var authModelErr *errors.AuthorizationModelError
	// Scenario 1: Using GetOrAddNode with new nodes
	graphA := NewWeightedAuthorizationModelGraph()
	node1 := graphA.GetOrAddNode("document", "document", SpecificType, "")
	require.NotNil(t, node1)
	require.Equal(t, "document", node1.label)

	// When adding the same node again with a different label, the original is returned
	node2 := graphA.GetOrAddNode("document", "document_updated", SpecificType, "")
	require.NotNil(t, node2)
	require.Equal(t, node1, node2)            // Same reference
	require.Equal(t, "document", node2.label) // Label did not change

	// Scenario 2: Using AddNode with new and existing nodes
	graphB := NewWeightedAuthorizationModelGraph()
	node3, err := graphB.AddNode("document", "document", SpecificType, "")
	require.NoError(t, err)
	require.NotNil(t, node3)
	require.Equal(t, "document", node3.label)

	// When trying to add the same node, it fails
	node4, err := graphB.AddNode("document", "document_updated", SpecificType, "")
	require.Error(t, err)
	require.Nil(t, node4)

	require.ErrorAs(t, err, &authModelErr)

	// We can still add different nodes
	node5, err := graphB.AddNode("folder", "folder", SpecificType, "")
	require.NoError(t, err)
	require.NotNil(t, node5)
	require.Equal(t, "folder", node5.label)
}
