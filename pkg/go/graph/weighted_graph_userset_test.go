package graph

import (
	"testing"

	language "github.com/openfga/language/pkg/go/transformer"
	"github.com/stretchr/testify/require"
)

func TestUsersetWeightDirect(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
        type role
            relations
                define assignee: [user]
				define member: [user] or assignee
				define permission: [permission]
				define creator: member from permission
        type permission
            relations
                define assignee: [role#assignee, role#member]
				define member: [permission#assignee]
        type job
            relations
                define can_read: [user] or can_write
				define can_read_upgraded: [job#can_read]
                define permission: [permission]
                define cannot_read: [user] but not can_read
				define exceptional_read: can_read_upgraded but not can_read
				define can_write: assignee from permission
				define can_execute: [user] and can_read
				define exceptional_execute: [job#exceptional_read] and can_read
				define cannot_execute: ([user] and can_read) or can_write
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)
	weight, ok := graph.GetEdgeWeight(graph.edges["permission#assignee"][0], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 1, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_read"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 2, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#cannot_read"], "role#assignee")
	require.False(t, ok)
	require.Equal(t, 0, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["permission#assignee"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 1, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_write"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 2, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_execute"], "role#assignee")
	require.False(t, ok)
	require.Equal(t, 0, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#cannot_execute"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 2, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_read_upgraded"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 3, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#exceptional_read"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 3, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#exceptional_execute"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 4, weight)

	weight, ok = graph.GetEdgeWeight(graph.edges["permission#member"][0], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 2, weight)
	weight, ok = graph.GetEdgeWeight(graph.edges["role#creator"][0], "role#assignee")
	require.True(t, ok)
	require.Equal(t, 3, weight)
	_, ok = graph.GetEdgeWeight(graph.edges["role#member"][0], "role#assignee")
	require.False(t, ok)
}

func TestUsersetWeightRecursivePath(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
        type role
            relations
                define assignee: [user, role#assignee]
				define member: [user, role#assignee]
        type permission
            relations
                define assignee: [role#assignee, role#member] 
				define member: [role#member]
        type job
            relations
                define can_read: [user] or can_write
				define can_read_upgraded: [job#can_read]
                define permission: [permission]
                define cannot_read: [user] but not can_read
				define exceptional_read: can_read_upgraded but not can_read
				define can_write: assignee from permission
				define can_execute: [user] and can_read
				define exceptional_execute: [job#exceptional_read] and can_read
				define cannot_execute: ([user] and can_read) or can_write
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)
	weight, ok := graph.GetNodeWeight(graph.nodes["job#can_read"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#cannot_read"], "role#assignee")
	require.False(t, ok)
	require.Equal(t, 0, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["permission#assignee"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_write"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_execute"], "role#assignee")
	require.False(t, ok)
	require.Equal(t, 0, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#cannot_execute"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_read_upgraded"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#exceptional_read"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#exceptional_execute"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)

	weight, ok = graph.GetEdgeWeight(graph.edges["permission#member"][0], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetEdgeWeight(graph.edges["role#member"][1], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	_, ok = graph.GetEdgeWeight(graph.edges["role#member"][0], "role#assignee")
	require.False(t, ok)
}

func TestUsersetWeightTupleCyclePath(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
        type role
            relations
                define assignee: [user, permission#assignee]
				define member: [user]
        type permission
            relations
                define assignee: [role#assignee, role#member] 
				define member: [user, permission#assignee]
        type job
            relations
                define can_read: [user] or can_write
				define can_read_upgraded: [job#can_read]
                define permission: [permission]
                define cannot_read: [user] but not can_read
				define exceptional_read: can_read_upgraded but not can_read
				define can_write: assignee from permission
				define can_execute: [user] and can_read
				define exceptional_execute: [job#exceptional_read] and can_read
				define cannot_execute: ([user] and can_read) or can_write
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)
	weight, ok := graph.GetNodeWeight(graph.nodes["job#can_read"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#cannot_read"], "role#assignee")
	require.False(t, ok)
	require.Equal(t, 0, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["permission#assignee"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_write"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_execute"], "role#assignee")
	require.False(t, ok)
	require.Equal(t, 0, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#cannot_execute"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_read_upgraded"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#exceptional_read"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#exceptional_execute"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)

	weight, ok = graph.GetEdgeWeight(graph.edges["permission#member"][1], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	_, ok = graph.GetEdgeWeight(graph.edges["permission#member"][0], "role#assignee")
	require.False(t, ok)
}

func TestUsersetWeightDependsOnRecursive(t *testing.T) {
	t.Parallel()
	model := `
	model
  		schema 1.1
		type user
        type role
            relations
                define assignee: [user]
				define member: [user]
        type permission
            relations
                define assignee: [role#assignee, permission#assignee] 
        type job
            relations
                define can_read: [user] or can_write
				define can_read_upgraded: [job#can_read]
                define permission: [permission]
                define cannot_read: [user] but not can_read
				define exceptional_read: can_read_upgraded but not can_read
				define can_write: assignee from permission
				define can_execute: [user] and can_read
				define exceptional_execute: [job#exceptional_read] and can_read
				define cannot_execute: ([user] and can_read) or can_write
	`
	authorizationModel := language.MustTransformDSLToProto(model)
	wgb := NewWeightedAuthorizationModelGraphBuilder()
	graph, err := wgb.Build(authorizationModel)
	require.NoError(t, err)
	weight, ok := graph.GetNodeWeight(graph.nodes["job#can_read"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#cannot_read"], "role#assignee")
	require.False(t, ok)
	require.Equal(t, 0, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["permission#assignee"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_write"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_execute"], "role#assignee")
	require.False(t, ok)
	require.Equal(t, 0, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#cannot_execute"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#can_read_upgraded"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#exceptional_read"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
	weight, ok = graph.GetNodeWeight(graph.nodes["job#exceptional_execute"], "role#assignee")
	require.True(t, ok)
	require.Equal(t, Infinite, weight)
}
