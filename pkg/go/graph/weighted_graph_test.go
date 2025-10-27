package graph

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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
	graph.AddNode("state-can_view", "can_view", SpecificTypeAndRelation)
	graph.AddNode("state-can_view-or", UnionOperator, OperatorNode)
	graph.AddNode("state-member", "member", SpecificTypeAndRelation)
	graph.AddNode("state-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("state-owner-and", IntersectionOperator, OperatorNode)
	graph.AddNode("state-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("state-can_apply", "can_apply", SpecificTypeAndRelation)
	graph.AddNode("state-can_apply-but", ExclusionOperator, OperatorNode)
	graph.AddNode("state-can_apply-or", UnionOperator, OperatorNode)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("user:*", "user:*", SpecificTypeWildcard)

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
	graph.AddNode("state-can_view", "can_view", SpecificTypeAndRelation)
	graph.AddNode("state-can_view-or", UnionOperator, OperatorNode)
	graph.AddNode("state-member", "member", SpecificTypeAndRelation)
	graph.AddNode("state-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("state-owner-and", IntersectionOperator, OperatorNode)
	graph.AddNode("user", "user", SpecificType)

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
	require.ErrorIs(t, err, ErrModelCycle)
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
	graph.AddNode("state-can_view", "can_view", SpecificTypeAndRelation)
	graph.AddNode("state-can_view-or", UnionOperator, OperatorNode)
	graph.AddNode("state-member", "member", SpecificTypeAndRelation)
	graph.AddNode("state-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("state-owner-and", IntersectionOperator, OperatorNode)
	graph.AddNode("state-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("state-can_apply", "can_apply", SpecificTypeAndRelation)
	graph.AddNode("state-can_apply-but", ExclusionOperator, OperatorNode)
	graph.AddNode("state-can_apply-or", UnionOperator, OperatorNode)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("employee", "employee", SpecificType)

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
	graph.AddNode("state-can_view", "can_view", SpecificTypeAndRelation)
	graph.AddNode("state-can_view-or", UnionOperator, OperatorNode)
	graph.AddNode("state-member", "member", SpecificTypeAndRelation)
	graph.AddNode("state-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("state-owner-and", IntersectionOperator, OperatorNode)
	graph.AddNode("state-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("employee", "employee", SpecificType)

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
	require.ErrorIs(t, err, ErrInvalidModel)
	require.True(t, strings.HasPrefix(err.Error(), "invalid model: not all paths return the same type for the node"))
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
	graph.AddNode("state-can_view", "can_view", SpecificTypeAndRelation)
	graph.AddNode("state-can_view-or", UnionOperator, OperatorNode)
	graph.AddNode("state-member", "member", SpecificTypeAndRelation)
	graph.AddNode("state-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("state-owner-and", IntersectionOperator, OperatorNode)
	graph.AddNode("state-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("state", "state", SpecificType)
	graph.AddNode("user:*", "user:*", SpecificTypeWildcard)
	graph.AddNode("transition-start", "start", SpecificTypeAndRelation)
	graph.AddNode("transition-end", "end", SpecificTypeAndRelation)
	graph.AddNode("transition-can_apply", "can_apply", SpecificTypeAndRelation)
	graph.AddNode("transition-can_apply-and", IntersectionOperator, OperatorNode)

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
	graph.AddNode("state-can_view", "can_view", SpecificTypeAndRelation)
	graph.AddNode("state-can_view-or", UnionOperator, OperatorNode)
	graph.AddNode("state-member", "member", SpecificTypeAndRelation)
	graph.AddNode("state-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("state-owner-and", IntersectionOperator, OperatorNode)
	graph.AddNode("state-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("state", "state", SpecificType)
	graph.AddNode("user:*", "user:*", SpecificTypeWildcard)
	graph.AddNode("transition-can_apply", "can_apply", SpecificTypeAndRelation)

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
	graph.AddNode("company-member", "member", SpecificTypeAndRelation)
	graph.AddNode("company-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("company-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("company-approved_member-or", UnionOperator, OperatorNode)
	graph.AddNode("group-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("employee", "employee", SpecificType)
	graph.AddNode("company", "company", SpecificType)
	graph.AddNode("group", "group", SpecificType)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("license-active_member", "active_member", SpecificTypeAndRelation)
	graph.AddNode("license-owner", "owner", SpecificTypeAndRelation)

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
	graph.AddNode("company-member", "member", SpecificTypeAndRelation)
	graph.AddNode("company-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("company-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("company-approved_member-or", UnionOperator, OperatorNode)
	graph.AddNode("user", "user", SpecificType)

	graph.AddEdge("company-member", "user", DirectEdge, "company-member", "", nil)
	graph.AddEdge("company-owner", "company-approved_member", ComputedEdge, "company-owner", "", nil)
	graph.AddEdge("company-approved_member", "company-approved_member-or", RewriteEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member-or", "company-member", DirectEdge, "company-approved_member", "", nil)
	graph.AddEdge("company-approved_member-or", "company-owner", ComputedEdge, "company-approved_member", "", nil)

	err := graph.AssignWeights()
	require.ErrorIs(t, err, ErrModelCycle)
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
	graph.AddNode("company-member", "member", SpecificTypeAndRelation)
	graph.AddNode("company-executive", "executive", SpecificTypeAndRelation)
	graph.AddNode("company-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("company-can_execute", "can_execute", SpecificTypeAndRelation)
	graph.AddNode("employee", "employee", SpecificType)
	graph.AddNode("company-can_execute-but", ExclusionOperator, OperatorNode)
	graph.AddNode("user", "user", SpecificType)

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
	graph.AddNode("document-rel4", "rel4", SpecificTypeAndRelation)
	graph.AddNode("document-rel4-or", UnionOperator, OperatorNode)
	graph.AddNode("document-rel6", "rel6", SpecificTypeAndRelation)
	graph.AddNode("document-rel7", "rel7", SpecificTypeAndRelation)
	graph.AddNode("document-rel1", "rel1", SpecificTypeAndRelation)
	graph.AddNode("document-rel1-or", UnionOperator, OperatorNode)
	graph.AddNode("document-rel2", "rel2", SpecificTypeAndRelation)
	graph.AddNode("document-rel3", "rel3", SpecificTypeAndRelation)
	graph.AddNode("document-rel5", "rel5", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("employee", "employee", SpecificType)

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
	graph.AddNode("document-rel4", "rel4", SpecificTypeAndRelation)
	graph.AddNode("document-rel4-or", UnionOperator, OperatorNode)
	graph.AddNode("document-rel4-and", IntersectionOperator, OperatorNode)
	graph.AddNode("document-rel6", "rel6", SpecificTypeAndRelation)
	graph.AddNode("document-rel7", "rel7", SpecificTypeAndRelation)
	graph.AddNode("document-rel1", "rel1", SpecificTypeAndRelation)
	graph.AddNode("document-rel1-or", UnionOperator, OperatorNode)
	graph.AddNode("document-rel2", "rel2", SpecificTypeAndRelation)
	graph.AddNode("document-rel3", "rel3", SpecificTypeAndRelation)
	graph.AddNode("document-rel5", "rel5", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("employee", "employee", SpecificType)

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
	require.ErrorIs(t, err, ErrContrainstTupleCycle)
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
	graph.AddNode("document-rel4", "rel4", SpecificTypeAndRelation)
	graph.AddNode("document-rel4-or", UnionOperator, OperatorNode)
	graph.AddNode("document-rel6", "rel6", SpecificTypeAndRelation)
	graph.AddNode("document-rel7", "rel7", SpecificTypeAndRelation)
	graph.AddNode("document-rel1", "rel1", SpecificTypeAndRelation)
	graph.AddNode("document-rel1-and", IntersectionOperator, OperatorNode)
	graph.AddNode("document-rel2", "rel2", SpecificTypeAndRelation)
	graph.AddNode("document-rel3", "rel3", SpecificTypeAndRelation)
	graph.AddNode("document-rel5", "rel5", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("employee", "employee", SpecificType)

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
	graph.AddNode("document-rel4", "rel4", SpecificTypeAndRelation)
	graph.AddNode("document-rel4-or", UnionOperator, OperatorNode)
	graph.AddNode("document-rel6", "rel6", SpecificTypeAndRelation)
	graph.AddNode("document-rel3", "rel3", SpecificTypeAndRelation)
	graph.AddNode("document-rel5", "rel5", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("employee", "employee", SpecificType)
	graph.AddNode("document", "document", SpecificType)

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
	graph.AddNode("state-can_view", "can_view", SpecificTypeAndRelation)
	graph.AddNode("state-can_view-or", UnionOperator, OperatorNode)
	graph.AddNode("state-member", "member", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("state", "state", SpecificType)
	graph.AddNode("transition-start", "start", SpecificTypeAndRelation)
	graph.AddNode("transition-end", "end", SpecificTypeAndRelation)
	graph.AddNode("transition-can_apply", "can_apply", SpecificTypeAndRelation)
	graph.AddNode("transition-can_apply-and", IntersectionOperator, OperatorNode)
	graph.AddNode("group-owner", "owner", SpecificTypeAndRelation)
	graph.AddNode("group-max_owner", "max_owner", SpecificTypeAndRelation)

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

	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("user:*", "user:*", SpecificTypeWildcard)
	graph.AddNode("employee", "employee", SpecificType)
	graph.AddNode("company", "company", SpecificType)
	graph.AddNode("group", "group", SpecificType)
	graph.AddNode("license", "license", SpecificType)
	graph.AddNode("tier", "tier", SpecificType)
	graph.AddNode("module", "module", SpecificType)
	graph.AddNode("feature", "feature", SpecificType)

	graph.AddNode("company-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("company-approved_member-and", IntersectionOperator, OperatorNode)
	graph.AddNode("company-member", "member", SpecificTypeAndRelation)
	graph.AddNode("company-facilitator", "facilitator", SpecificTypeAndRelation)
	graph.AddNode("company-user_in_context", "user_in_context", SpecificTypeAndRelation)

	graph.AddNode("group-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("group-approved_member-butnot", ExclusionOperator, OperatorNode)
	graph.AddNode("group-member", "member", SpecificTypeAndRelation)
	graph.AddNode("group-user_in_context", "user_in_context", SpecificTypeAndRelation)
	graph.AddNode("group-reader", "reader", SpecificTypeAndRelation)
	graph.AddNode("group-assignee", "assignee", SpecificTypeAndRelation)

	graph.AddNode("license-active_holder", "active_holder", SpecificTypeAndRelation)
	graph.AddNode("license-active_holder-or", UnionOperator, OperatorNode)
	graph.AddNode("license-holder_member", "holder_member", SpecificTypeAndRelation)
	graph.AddNode("license-holder_approved_member", "holder_approved_member", SpecificTypeAndRelation)
	graph.AddNode("license-holder", "holder", SpecificTypeAndRelation)
	graph.AddNode("license-parent", "parent", SpecificTypeAndRelation)
	graph.AddNode("license-trust_holder", "trust_holder", SpecificTypeAndRelation)
	graph.AddNode("license-trust_holder-or", UnionOperator, OperatorNode)
	graph.AddNode("license-owner", "owner", SpecificTypeAndRelation)

	graph.AddNode("tier-subscriber", "subscriber", SpecificTypeAndRelation)
	graph.AddNode("tier-assignee", "assignee", SpecificTypeAndRelation)
	graph.AddNode("tier-subtier_owner", "subtier_owner", SpecificTypeAndRelation)
	graph.AddNode("tier-assignee_sub", "assignee_sub", SpecificTypeAndRelation)
	graph.AddNode("tier-assignee_sub-and", IntersectionOperator, OperatorNode)

	graph.AddNode("module-associated_license", "associated_license", SpecificTypeAndRelation)
	graph.AddNode("module-module_holder", "module_holder", SpecificTypeAndRelation)
	graph.AddNode("module-module_user", "module_user", SpecificTypeAndRelation)

	graph.AddNode("feature-associated_module", "associated_module", SpecificTypeAndRelation)
	graph.AddNode("feature-associated_tier", "associated_tier", SpecificTypeAndRelation)
	graph.AddNode("feature-tier_can_access", "tier_can_access", SpecificTypeAndRelation)
	graph.AddNode("feature-can_access", "can_access", SpecificTypeAndRelation)
	graph.AddNode("feature-can_access-and", IntersectionOperator, OperatorNode)

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
	graph.AddNode("state-member", "member", SpecificTypeAndRelation)
	graph.AddNode("state-member-or", UnionOperator, OperatorNode)
	graph.AddNode("state-member-or-or", UnionOperator, OperatorNode)
	graph.AddNode("state-parent", "parent", SpecificTypeAndRelation)
	graph.AddNode("state-parent_member", "parent_member", SpecificTypeAndRelation)
	graph.AddNode("user", "user", SpecificType)
	graph.AddNode("user:*", "user", SpecificTypeWildcard)
	graph.AddNode("state", "state", SpecificType)

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
func TestGetEdgesFromNodeId(t *testing.T) {
	t.Parallel()

	t.Run("returns_edges_for_existing_node", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()

		// Add nodes
		graph.AddNode("state-member", "member", SpecificTypeAndRelation)
		graph.AddNode("state-owner", "owner", SpecificTypeAndRelation)
		graph.AddNode("user", "user", SpecificType)

		// Add edges from state-member
		graph.AddEdge("state-member", "user", DirectEdge, "state-member", "", nil)
		graph.AddEdge("state-member", "state-owner", ComputedEdge, "state-member", "", nil)

		// Get edges
		edges, found := graph.GetEdgesFromNodeId("state-member")

		require.True(t, found)
		require.Len(t, edges, 2)

		// Verify edge properties
		require.Equal(t, "user", edges[0].to.uniqueLabel)
		require.Equal(t, DirectEdge, edges[0].edgeType)
		require.Equal(t, "state-owner", edges[1].to.uniqueLabel)
		require.Equal(t, ComputedEdge, edges[1].edgeType)
	})

	t.Run("returns_empty_for_node_with_no_edges", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()

		// Add node without edges
		graph.AddNode("state-member", "member", SpecificTypeAndRelation)

		edges, found := graph.GetEdgesFromNodeId("state-member")

		require.False(t, found)
		require.Nil(t, edges)
	})

	t.Run("returns_empty_for_non_existent_node", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()

		edges, found := graph.GetEdgesFromNodeId("non-existent")

		require.False(t, found)
		require.Nil(t, edges)
	})

	t.Run("returns_multiple_edges_with_different_types", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()

		// Add nodes
		graph.AddNode("transition-can_apply", "can_apply", SpecificTypeAndRelation)
		graph.AddNode("transition-can_apply-and", IntersectionOperator, OperatorNode)
		graph.AddNode("state-can_view", "can_view", SpecificTypeAndRelation)
		graph.AddNode("user", "user", SpecificType)

		// Add different types of edges
		graph.AddEdge("transition-can_apply", "transition-can_apply-and", RewriteEdge, "transition-can_apply", "", nil)
		graph.AddEdge("transition-can_apply-and", "user", DirectEdge, "transition-can_apply", "", nil)
		graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-can_apply", "transition-start", nil)
		graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-can_apply", "transition-end", nil)

		// Get edges from transition-can_apply-and
		edges, found := graph.GetEdgesFromNodeId("transition-can_apply-and")

		require.True(t, found)
		require.Len(t, edges, 3)

		// Verify different edge types
		edgeTypes := make(map[EdgeType]int)
		for _, edge := range edges {
			edgeTypes[edge.edgeType]++
		}
		require.Equal(t, 1, edgeTypes[DirectEdge])
		require.Equal(t, 2, edgeTypes[TTUEdge])
	})

	t.Run("returns_edges_with_conditions", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()

		// Add nodes
		graph.AddNode("license-owner", "owner", SpecificTypeAndRelation)
		graph.AddNode("group", "group", SpecificType)

		// Add edges with conditions
		graph.AddEdge("license-owner", "group", DirectEdge, "license-owner", "", []string{"condition1", "condition2"})

		edges, found := graph.GetEdgesFromNodeId("license-owner")

		require.True(t, found)
		require.Len(t, edges, 1)
		require.Equal(t, []string{"condition1", "condition2"}, edges[0].conditions)
	})

	t.Run("handles_operator_nodes", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()

		// Add operator nodes
		graph.AddNode("state-owner-and", IntersectionOperator, OperatorNode)
		graph.AddNode("state-approved_member", "approved_member", SpecificTypeAndRelation)
		graph.AddNode("user", "user", SpecificType)

		// Add edges from operator node
		graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "state-owner", "", nil)
		graph.AddEdge("state-owner-and", "user", DirectEdge, "state-owner", "", nil)

		edges, found := graph.GetEdgesFromNodeId("state-owner-and")

		require.True(t, found)
		require.Len(t, edges, 2)
	})

	t.Run("preserves_edge_order", func(t *testing.T) {
		graph := NewWeightedAuthorizationModelGraph()

		// Add nodes
		graph.AddNode("parent", "parent", SpecificTypeAndRelation)
		graph.AddNode("child1", "child1", SpecificType)
		graph.AddNode("child2", "child2", SpecificType)
		graph.AddNode("child3", "child3", SpecificType)

		// Add edges in specific order
		graph.AddEdge("parent", "child1", DirectEdge, "parent", "", nil)
		graph.AddEdge("parent", "child2", DirectEdge, "parent", "", nil)
		graph.AddEdge("parent", "child3", DirectEdge, "parent", "", nil)

		edges, found := graph.GetEdgesFromNodeId("parent")

		require.True(t, found)
		require.Len(t, edges, 3)
		require.Equal(t, "child1", edges[0].to.uniqueLabel)
		require.Equal(t, "child2", edges[1].to.uniqueLabel)
		require.Equal(t, "child3", edges[2].to.uniqueLabel)
	})
}
