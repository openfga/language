package graph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// const Infinite = math.MaxInt32

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

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "")
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "")
	graph.AddEdge("state-member", "user", DirectEdge, "")
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "")
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "")
	graph.AddEdge("state-owner-and", "user", DirectEdge, "")
	graph.AddEdge("state-approved_member", "user:*", DirectEdge, "")
	graph.AddEdge("state-can_apply", "state-can_apply-or", RewriteEdge, "")
	graph.AddEdge("state-can_apply-or", "state-can_apply-but", RewriteEdge, "")
	graph.AddEdge("state-can_apply-or", "state-owner", ComputedEdge, "")
	graph.AddEdge("state-can_apply-but", "state-approved_member", ComputedEdge, "")
	graph.AddEdge("state-can_apply-but", "state-can_view", ComputedEdge, "")

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

	require.Equal(t, 1, len(graph.nodes["state-can_view"].weights))
	require.Equal(t, 1, len(graph.nodes["state-member"].weights))
	require.Equal(t, 1, len(graph.nodes["state-owner"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_view-or"].weights))
	require.Equal(t, 1, len(graph.nodes["state-owner-and"].weights))
	require.Equal(t, 1, len(graph.nodes["state-approved_member"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_apply"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_apply-but"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_apply-or"].weights))
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

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "")
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "")
	graph.AddEdge("state-member", "user", DirectEdge, "")
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "")
	graph.AddEdge("state-owner-and", "state-can_view", ComputedEdge, "")

	err := graph.AssignWeights()
	require.Equal(t, "model cycle", err.Error())
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

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "")
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "")
	graph.AddEdge("state-member", "user", DirectEdge, "")
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "")
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "")
	graph.AddEdge("state-owner-and", "employee", DirectEdge, "")
	graph.AddEdge("state-approved_member", "employee", DirectEdge, "")
	graph.AddEdge("state-can_apply", "state-can_apply-or", RewriteEdge, "")
	graph.AddEdge("state-can_apply-or", "state-can_apply-but", RewriteEdge, "")
	graph.AddEdge("state-can_apply-or", "state-member", ComputedEdge, "")
	graph.AddEdge("state-can_apply-but", "state-approved_member", ComputedEdge, "")
	graph.AddEdge("state-can_apply-but", "state-can_view", ComputedEdge, "")

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

	require.Equal(t, 2, len(graph.nodes["state-can_view"].weights))
	require.Equal(t, 1, len(graph.nodes["state-member"].weights))
	require.Equal(t, 1, len(graph.nodes["state-owner"].weights))
	require.Equal(t, 2, len(graph.nodes["state-can_view-or"].weights))
	require.Equal(t, 1, len(graph.nodes["state-owner-and"].weights))
	require.Equal(t, 1, len(graph.nodes["state-approved_member"].weights))
	require.Equal(t, 2, len(graph.nodes["state-can_apply"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_apply-but"].weights))
	require.Equal(t, 2, len(graph.nodes["state-can_apply-or"].weights))
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

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "")
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "")
	graph.AddEdge("state-member", "user", DirectEdge, "")
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "")
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "")
	graph.AddEdge("state-owner-and", "employee", DirectEdge, "")
	graph.AddEdge("state-approved_member", "user", DirectEdge, "")

	err := graph.AssignWeights()
	require.Equal(t, "not all paths return the same type for the node", err.Error())
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

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "")
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "")
	graph.AddEdge("state-member", "user:*", DirectEdge, "")
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "")
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "")
	graph.AddEdge("state-owner-and", "user", DirectEdge, "")
	graph.AddEdge("state-approved_member", "user", DirectEdge, "")
	graph.AddEdge("transition-start", "state", DirectEdge, "")
	graph.AddEdge("transition-end", "state", DirectEdge, "")
	graph.AddEdge("transition-can_apply", "transition-can_apply-and", RewriteEdge, "")
	graph.AddEdge("transition-can_apply-and", "user", DirectEdge, "")
	graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-start")
	graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-end")

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

	require.Equal(t, 1, len(graph.nodes["state-can_view"].weights))
	require.Equal(t, 1, len(graph.nodes["state-member"].weights))
	require.Equal(t, 1, len(graph.nodes["state-owner"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_view-or"].weights))
	require.Equal(t, 1, len(graph.nodes["state-owner-and"].weights))
	require.Equal(t, 1, len(graph.nodes["state-approved_member"].weights))
	require.Equal(t, 1, len(graph.nodes["transition-end"].weights))
	require.Equal(t, 1, len(graph.nodes["transition-start"].weights))
	require.Equal(t, 1, len(graph.nodes["transition-can_apply"].weights))
	require.Equal(t, 1, len(graph.nodes["transition-can_apply-and"].weights))
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

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "")
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "state-owner", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "")
	graph.AddEdge("state-member", "user:*", DirectEdge, "")
	graph.AddEdge("state-owner", "state-owner-and", RewriteEdge, "")
	graph.AddEdge("state-owner-and", "state-approved_member", ComputedEdge, "")
	graph.AddEdge("state-owner-and", "user", DirectEdge, "")
	graph.AddEdge("state-approved_member", "user", DirectEdge, "")
	graph.AddEdge("transition-can_apply", "user", DirectEdge, "")
	graph.AddEdge("transition-can_apply", "state-can_view", DirectEdge, "")

	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, 1, graph.nodes["state-can_view"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-member"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-can_view-or"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-owner-and"].weights["user"])
	require.Equal(t, 1, graph.nodes["state-approved_member"].weights["user"])
	require.Equal(t, 2, graph.nodes["transition-can_apply"].weights["user"])

	require.Equal(t, 1, len(graph.nodes["state-can_view"].weights))
	require.Equal(t, 1, len(graph.nodes["state-member"].weights))
	require.Equal(t, 1, len(graph.nodes["state-owner"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_view-or"].weights))
	require.Equal(t, 1, len(graph.nodes["state-owner-and"].weights))
	require.Equal(t, 1, len(graph.nodes["state-approved_member"].weights))
	require.Equal(t, 1, len(graph.nodes["transition-can_apply"].weights))
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

	graph.AddEdge("company-member", "user", DirectEdge, "")
	graph.AddEdge("company-owner", "user", DirectEdge, "")
	graph.AddEdge("company-approved_member", "company-approved_member-or", RewriteEdge, "")
	graph.AddEdge("company-approved_member-or", "company-member", DirectEdge, "")
	graph.AddEdge("company-approved_member-or", "company-owner", ComputedEdge, "")
	graph.AddEdge("group-approved_member", "employee", DirectEdge, "")
	graph.AddEdge("license-active_member", "company-approved_member", TTUEdge, "company-owner")
	graph.AddEdge("license-active_member", "group-approved_member", TTUEdge, "company-owner")
	graph.AddEdge("license-owner", "company", DirectEdge, "")
	graph.AddEdge("license-owner", "group", DirectEdge, "")

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

	require.Equal(t, 1, len(graph.nodes["company-member"].weights))
	require.Equal(t, 1, len(graph.nodes["company-owner"].weights))
	require.Equal(t, 1, len(graph.nodes["company-approved_member"].weights))
	require.Equal(t, 1, len(graph.nodes["company-approved_member-or"].weights))
	require.Equal(t, 1, len(graph.nodes["group-approved_member"].weights))
	require.Equal(t, 2, len(graph.nodes["license-active_member"].weights))
	require.Equal(t, 2, len(graph.nodes["license-owner"].weights))
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

	graph.AddEdge("company-member", "user", DirectEdge, "")
	graph.AddEdge("company-owner", "company-approved_member", ComputedEdge, "")
	graph.AddEdge("company-approved_member", "company-approved_member-or", RewriteEdge, "")
	graph.AddEdge("company-approved_member-or", "company-member", DirectEdge, "")
	graph.AddEdge("company-approved_member-or", "company-owner", ComputedEdge, "")

	err := graph.AssignWeights()
	require.Equal(t, "model cycle", err.Error())
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
func TestInvalidWeight2ButNotMistmatchType(t *testing.T) {
	t.Parallel()
	graph := NewWeightedAuthorizationModelGraph()
	graph.AddNode("company-member", "member", SpecificTypeAndRelation)
	graph.AddNode("company-executive", "executive", SpecificTypeAndRelation)
	graph.AddNode("company-approved_member", "approved_member", SpecificTypeAndRelation)
	graph.AddNode("company-can_execute", "can_execute", SpecificTypeAndRelation)
	graph.AddNode("employee", "employee", SpecificType)
	graph.AddNode("company-can_execute-but", ExclusionOperator, OperatorNode)
	graph.AddNode("user", "user", SpecificType)

	graph.AddEdge("company-member", "user", DirectEdge, "")
	graph.AddEdge("company-executive", "employee", DirectEdge, "")
	graph.AddEdge("company-approved_member", "user", DirectEdge, "")
	graph.AddEdge("company-approved_member", "company-member", DirectEdge, "")
	graph.AddEdge("company-can_execute", "company-can_execute-but", RewriteEdge, "")
	graph.AddEdge("company-can_execute-but", "company-executive", ComputedEdge, "")
	graph.AddEdge("company-can_execute-but", "company-approved_member", ComputedEdge, "")

	err := graph.AssignWeights()
	require.Equal(t, "not all paths return the same type for the node", err.Error())
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

	graph.AddEdge("document-rel4", "document-rel4-or", RewriteEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel6", ComputedEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel7", ComputedEdge, "")
	graph.AddEdge("document-rel6", "document-rel1", DirectEdge, "")
	graph.AddEdge("document-rel7", "document-rel4", DirectEdge, "")
	graph.AddEdge("document-rel1", "document-rel1-or", RewriteEdge, "")
	graph.AddEdge("document-rel1-or", "document-rel2", ComputedEdge, "")
	graph.AddEdge("document-rel1-or", "document-rel3", ComputedEdge, "")
	graph.AddEdge("document-rel1-or", "document-rel4", ComputedEdge, "")
	graph.AddEdge("document-rel2", "user", DirectEdge, "")
	graph.AddEdge("document-rel3", "document-rel5", ComputedEdge, "")
	graph.AddEdge("document-rel5", "employee", DirectEdge, "")

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

	require.Equal(t, 2, len(graph.nodes["document-rel1"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel2"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel3"].weights))
	require.Equal(t, 2, len(graph.nodes["document-rel4"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel5"].weights))
	require.Equal(t, 2, len(graph.nodes["document-rel6"].weights))
	require.Equal(t, 2, len(graph.nodes["document-rel7"].weights))
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

	graph.AddEdge("document-rel4", "document-rel4-and", RewriteEdge, "")
	graph.AddEdge("document-rel4-and", "document-rel4-or", RewriteEdge, "")
	graph.AddEdge("document-rel4-and", "document-rel5", ComputedEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel6", ComputedEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel7", ComputedEdge, "")
	graph.AddEdge("document-rel6", "document-rel1", DirectEdge, "")
	graph.AddEdge("document-rel7", "document-rel4", DirectEdge, "")
	graph.AddEdge("document-rel1", "document-rel1-or", RewriteEdge, "")
	graph.AddEdge("document-rel1-or", "document-rel2", ComputedEdge, "")
	graph.AddEdge("document-rel1-or", "document-rel3", ComputedEdge, "")
	graph.AddEdge("document-rel1-or", "document-rel4", ComputedEdge, "")
	graph.AddEdge("document-rel2", "user", DirectEdge, "")
	graph.AddEdge("document-rel3", "document-rel5", ComputedEdge, "")
	graph.AddEdge("document-rel5", "employee", DirectEdge, "")

	err := graph.AssignWeights()
	require.Equal(t, "AND or BUT NOT cannot be involved in a cycle", err.Error())

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

	graph.AddEdge("document-rel4", "document-rel4-or", RewriteEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel6", ComputedEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel7", ComputedEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel5", ComputedEdge, "")
	graph.AddEdge("document-rel6", "employee", DirectEdge, "")
	graph.AddEdge("document-rel7", "document-rel4", DirectEdge, "")
	graph.AddEdge("document-rel1", "document-rel1-and", RewriteEdge, "")
	graph.AddEdge("document-rel1-and", "document-rel2", ComputedEdge, "")
	graph.AddEdge("document-rel1-and", "document-rel3", ComputedEdge, "")
	graph.AddEdge("document-rel1-and", "document-rel4", ComputedEdge, "")
	graph.AddEdge("document-rel2", "user", DirectEdge, "")
	graph.AddEdge("document-rel3", "user", DirectEdge, "")
	graph.AddEdge("document-rel3", "employee", DirectEdge, "")
	graph.AddEdge("document-rel5", "user", DirectEdge, "")
	graph.AddEdge("document-rel6", "employee", DirectEdge, "")

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

	require.Equal(t, 1, len(graph.nodes["document-rel1"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel1-and"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel2"].weights))
	require.Equal(t, 2, len(graph.nodes["document-rel3"].weights))
	require.Equal(t, 2, len(graph.nodes["document-rel4"].weights))
	require.Equal(t, 2, len(graph.nodes["document-rel4-or"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel5"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel6"].weights))
	require.Equal(t, 2, len(graph.nodes["document-rel7"].weights))
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

	graph.AddEdge("document-rel4", "document-rel4-or", RewriteEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel5", ComputedEdge, "")
	graph.AddEdge("document-rel4-or", "document-rel4", TTUEdge, "document-rel6")
	graph.AddEdge("document-rel6", "document", DirectEdge, "")
	graph.AddEdge("document-rel5", "user", DirectEdge, "")
	graph.AddEdge("document-rel3", "user", DirectEdge, "")
	graph.AddEdge("document-rel3", "employee", DirectEdge, "")
	graph.AddEdge("document-rel3", "document-rel3", DirectEdge, "")

	Infinite := math.MaxInt32
	err := graph.AssignWeights()
	require.NoError(t, err)
	require.Equal(t, Infinite, graph.nodes["document-rel3"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel3"].weights["employee"])
	require.Equal(t, 1, graph.nodes["document-rel5"].weights["user"])
	require.Equal(t, 1, graph.nodes["document-rel6"].weights["document"])
	require.Equal(t, Infinite, graph.nodes["document-rel4"].weights["user"])
	require.Equal(t, Infinite, graph.nodes["document-rel4-or"].weights["user"])

	require.Equal(t, 2, len(graph.nodes["document-rel3"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel4"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel4-or"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel5"].weights))
	require.Equal(t, 1, len(graph.nodes["document-rel6"].weights))
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

	graph.AddEdge("state-can_view", "state-can_view-or", RewriteEdge, "")
	graph.AddEdge("state-can_view-or", "state-member", ComputedEdge, "")
	graph.AddEdge("state-can_view-or", "user", DirectEdge, "")
	graph.AddEdge("state-member", "user", DirectEdge, "")
	graph.AddEdge("transition-start", "state", DirectEdge, "")
	graph.AddEdge("transition-end", "state", DirectEdge, "")
	graph.AddEdge("transition-can_apply", "transition-can_apply-and", RewriteEdge, "")
	graph.AddEdge("transition-can_apply-and", "user", DirectEdge, "")
	graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-start")
	graph.AddEdge("transition-can_apply-and", "state-can_view", TTUEdge, "transition-end")
	graph.AddEdge("group-owner", "user", DirectEdge, "")
	graph.AddEdge("group-owner", "transition-can_apply", DirectEdge, "")
	graph.AddEdge("group-max_owner", "group-owner", DirectEdge, "")
	graph.AddEdge("group-max_owner", "group-max_owner", DirectEdge, "")

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

	require.Equal(t, 1, len(graph.nodes["group-owner"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_view"].weights))
	require.Equal(t, 1, len(graph.nodes["state-can_view-or"].weights))
	require.Equal(t, 1, len(graph.nodes["transition-start"].weights))
	require.Equal(t, 1, len(graph.nodes["transition-end"].weights))
	require.Equal(t, 1, len(graph.nodes["transition-end"].weights))
	require.Equal(t, 1, len(graph.nodes["group-owner"].weights))
	require.Equal(t, 1, len(graph.nodes["group-max_owner"].weights))

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

	graph.AddEdge("company-member", "user", DirectEdge, "")
	graph.AddEdge("company-member", "employee", DirectEdge, "")
	graph.AddEdge("company-member", "user:*", DirectEdge, "")
	graph.AddEdge("company-facilitator", "company-member", ComputedEdge, "")
	graph.AddEdge("company-user_in_context", "user", DirectEdge, "x_less_than")
	graph.AddEdge("company-approved_member", "company-approved_member-and", RewriteEdge, "")
	graph.AddEdge("company-approved_member-and", "company-member", ComputedEdge, "")
	graph.AddEdge("company-approved_member-and", "company-user_in_context", ComputedEdge, "")

	graph.AddEdge("group-member", "user", DirectEdge, "")
	graph.AddEdge("group-member", "user", DirectEdge, "x_greater_than")
	graph.AddEdge("group-user_in_context", "user", DirectEdge, "")
	graph.AddEdge("group-reader", "group-member", ComputedEdge, "")
	graph.AddEdge("group-assignee", "group-reader", ComputedEdge, "")
	graph.AddEdge("group-approved_member", "group-approved_member-butnot", RewriteEdge, "")
	graph.AddEdge("group-approved_member-butnot", "group-user_in_context", ComputedEdge, "")
	graph.AddEdge("group-approved_member-butnot", "group-member", ComputedEdge, "")

	graph.AddEdge("license-active_holder", "license-active_holder-or", RewriteEdge, "")
	graph.AddEdge("license-active_holder-or", "license-holder", ComputedEdge, "")
	graph.AddEdge("license-active_holder-or", "license-holder_member", ComputedEdge, "")
	graph.AddEdge("license-owner", "company", DirectEdge, "")
	graph.AddEdge("license-owner", "group", DirectEdge, "")
	graph.AddEdge("license-owner", "group", DirectEdge, "x_condition")
	graph.AddEdge("license-holder_member", "group-member", TTUEdge, "license-owner")
	graph.AddEdge("license-holder_member", "company-member", TTUEdge, "license-owner")
	graph.AddEdge("license-holder_member", "group-member", TTUEdge, "license-owner")
	graph.AddEdge("license-holder_approved_member", "group-approved_member", TTUEdge, "license-owner")
	graph.AddEdge("license-holder_approved_member", "company-approved_member", TTUEdge, "license-owner")
	graph.AddEdge("license-holder_approved_member", "group-approved_member", TTUEdge, "license-owner")
	graph.AddEdge("license-holder", "user", DirectEdge, "")
	graph.AddEdge("license-parent", "license", DirectEdge, "")
	graph.AddEdge("license-trust_holder", "license-trust_holder-or", RewriteEdge, "")
	graph.AddEdge("license-trust_holder-or", "user", DirectEdge, "")
	graph.AddEdge("license-trust_holder-or", "user:*", DirectEdge, "")
	graph.AddEdge("license-trust_holder-or", "license-trust_holder", TTUEdge, "license-parent")

	graph.AddEdge("tier-subscriber", "company-facilitator", DirectEdge, "")
	graph.AddEdge("tier-assignee", "group-assignee", DirectEdge, "")
	graph.AddEdge("tier-assignee", "group-user_in_context", DirectEdge, "")
	graph.AddEdge("tier-assignee", "group-user_in_context", DirectEdge, "x_bigger_than")
	graph.AddEdge("tier-subtier_owner", "user", DirectEdge, "")
	graph.AddEdge("tier-subtier_owner", "tier-subtier_owner", DirectEdge, "")
	graph.AddEdge("tier-assignee_sub", "tier-assignee_sub-and", RewriteEdge, "")
	graph.AddEdge("tier-assignee_sub-and", "tier-subscriber", ComputedEdge, "")
	graph.AddEdge("tier-assignee_sub-and", "tier-assignee", ComputedEdge, "")
	graph.AddEdge("tier-assignee_sub-and", "tier-subtier_owner", ComputedEdge, "")

	graph.AddEdge("module-associated_license", "license", DirectEdge, "")
	graph.AddEdge("module-module_holder", "license-holder_member", TTUEdge, "module-associated_license")
	graph.AddEdge("module-module_user", "license-active_holder", TTUEdge, "module-associated_license")

	graph.AddEdge("feature-associated_module", "module", DirectEdge, "")
	graph.AddEdge("feature-associated_tier", "tier", DirectEdge, "")
	graph.AddEdge("feature-tier_can_access", "tier-subscriber", TTUEdge, "feature-associated_tier")
	graph.AddEdge("feature-can_access", "feature-can_access-and", RewriteEdge, "")
	graph.AddEdge("feature-can_access-and", "module-module_user", TTUEdge, "feature-associated_module")
	graph.AddEdge("feature-can_access-and", "tier-subscriber", TTUEdge, "feature-associated_tier")

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

	require.Equal(t, 1, len(graph.nodes["company-approved_member"].weights))
	require.Equal(t, 1, len(graph.nodes["company-approved_member-and"].weights))
	require.Equal(t, 2, len(graph.nodes["company-member"].weights))
	require.Equal(t, 2, len(graph.nodes["company-facilitator"].weights))
	require.Equal(t, 1, len(graph.nodes["company-user_in_context"].weights))
	require.Equal(t, 1, len(graph.nodes["group-approved_member"].weights))
	require.Equal(t, 1, len(graph.nodes["group-approved_member-butnot"].weights))
	require.Equal(t, 1, len(graph.nodes["group-member"].weights))
	require.Equal(t, 1, len(graph.nodes["group-user_in_context"].weights))
	require.Equal(t, 1, len(graph.nodes["group-reader"].weights))
	require.Equal(t, 1, len(graph.nodes["group-assignee"].weights))
	require.Equal(t, 2, len(graph.nodes["license-active_holder"].weights))
	require.Equal(t, 2, len(graph.nodes["license-active_holder-or"].weights))
	require.Equal(t, 2, len(graph.nodes["license-holder_member"].weights))
	require.Equal(t, 1, len(graph.nodes["license-holder_approved_member"].weights))
	require.Equal(t, 1, len(graph.nodes["license-holder"].weights))
	require.Equal(t, 1, len(graph.nodes["license-parent"].weights))
	require.Equal(t, 1, len(graph.nodes["license-trust_holder"].weights))
	require.Equal(t, 1, len(graph.nodes["license-trust_holder-or"].weights))
	require.Equal(t, 2, len(graph.nodes["license-owner"].weights))
	require.Equal(t, 2, len(graph.nodes["tier-subscriber"].weights))
	require.Equal(t, 1, len(graph.nodes["tier-assignee"].weights))
	require.Equal(t, 1, len(graph.nodes["tier-subtier_owner"].weights))
	require.Equal(t, 1, len(graph.nodes["tier-assignee_sub"].weights))
	require.Equal(t, 1, len(graph.nodes["tier-assignee_sub-and"].weights))
	require.Equal(t, 1, len(graph.nodes["module-associated_license"].weights))
	require.Equal(t, 2, len(graph.nodes["module-module_holder"].weights))
	require.Equal(t, 2, len(graph.nodes["module-module_user"].weights))
	require.Equal(t, 1, len(graph.nodes["feature-associated_module"].weights))
	require.Equal(t, 1, len(graph.nodes["feature-associated_tier"].weights))
	require.Equal(t, 2, len(graph.nodes["feature-tier_can_access"].weights))
	require.Equal(t, 2, len(graph.nodes["feature-can_access"].weights))
	require.Equal(t, 2, len(graph.nodes["feature-can_access-and"].weights))

}

// super test
