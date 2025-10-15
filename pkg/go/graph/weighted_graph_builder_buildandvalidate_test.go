package graph

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/openfga/language/pkg/go/errors"
	parser "github.com/openfga/language/pkg/go/transformer"
	"github.com/stretchr/testify/require"
)

const (
	// SchemaVersion1_2 for the authorization models.
	SchemaVersion1_2 string = "1.2"
)

// TestValidBuildAndValidate tests all valid model build scenarios
func TestValidBuildAndValidate(t *testing.T) {
	t.Run("successfully_builds_valid_model", func(t *testing.T) {
		t.Parallel()
		builder := NewWeightedAuthorizationModelGraphBuilder()

		// Create a simple model with different types of relations
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "user",
				},
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
						"editor": {
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{
									Relation: "viewer",
								},
							},
						},
					},
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									},
								},
							},
						},
					},
				},
			},
		}

		// Build the graph using BuildAndValidate
		graph, err := builder.ValidateAndBuild(model)
		require.NoError(t, err)
		require.NotNil(t, graph)

		// Verify nodes were created
		userNode, found := graph.GetNodeByID("user")
		require.True(t, found)
		require.Equal(t, "user", userNode.GetLabel())

		viewerNode, found := graph.GetNodeByID("document#viewer")
		require.True(t, found)
		require.Equal(t, "document#viewer", viewerNode.GetLabel())

		editorNode, found := graph.GetNodeByID("document#editor")
		require.True(t, found)
		require.Equal(t, "document#editor", editorNode.GetLabel())

		// Verify edges were created (after rewrites)
		edges, hasEdges := graph.GetEdgesFromNode(editorNode)
		require.True(t, hasEdges)
		require.Len(t, edges, 1)
		require.Equal(t, "document#viewer", edges[0].GetTo().GetUniqueLabel())
	})

	t.Run("creates_nodes_before_edges", func(t *testing.T) {
		t.Parallel()
		builder := NewWeightedAuthorizationModelGraphBuilder()

		// Create a model with relations referencing other types
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "user",
				},
				{
					Type: "group",
					Relations: map[string]*openfgav1.Userset{
						"member": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"member": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									},
								},
							},
						},
					},
				},
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
						"editor": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									},
									{
										Type: "group",
										RelationOrWildcard: &openfgav1.RelationReference_Relation{
											Relation: "member",
										},
									},
								},
							},
							"editor": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "group",
										RelationOrWildcard: &openfgav1.RelationReference_Relation{
											Relation: "member",
										},
									},
								},
							},
						},
					},
				},
			},
		}

		// Build the graph using BuildAndValidate
		graph, err := builder.ValidateAndBuild(model)
		require.NoError(t, err)
		require.NotNil(t, graph)

		// Verify nodes were created for all types and relations
		_, found := graph.GetNodeByID("user")
		require.True(t, found)

		_, found = graph.GetNodeByID("group")
		require.True(t, found)

		_, found = graph.GetNodeByID("document")
		require.True(t, found)

		_, found = graph.GetNodeByID("group#member")
		require.True(t, found)

		_, found = graph.GetNodeByID("document#viewer")
		require.True(t, found)

		_, found = graph.GetNodeByID("document#editor")
		require.True(t, found)

		// Verify edges were created correctly
		documentViewerNode, _ := graph.GetNodeByID("document#viewer")
		edges, hasEdges := graph.GetEdgesFromNode(documentViewerNode)
		require.True(t, hasEdges)
		// document#viewer should have edges to user and group#member
		require.Len(t, edges, 2)

		documentEditorNode, _ := graph.GetNodeByID("document#editor")
		edges, hasEdges = graph.GetEdgesFromNode(documentEditorNode)
		require.True(t, hasEdges)
		// document#editor should have edge to group#member
		require.Len(t, edges, 1)
		require.Equal(t, "group#member", edges[0].GetTo().GetUniqueLabel())
	})

	t.Run("direct_self_reference_is_allowed", func(t *testing.T) {
		t.Parallel()
		builder := NewWeightedAuthorizationModelGraphBuilder()

		model := &openfgav1.AuthorizationModel{
			SchemaVersion: SchemaVersion1_2,
			TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user

				type document
				relations
					define editor: [user]
					define viewer: [document#viewer] or editor`).GetTypeDefinitions(),
		}

		graph, err := builder.ValidateAndBuild(model)
		require.NoError(t, err)
		require.NotNil(t, graph)
	})

	t.Run("valid_relations_with_entry_points", func(t *testing.T) {
		validTests := map[string]struct {
			model string
		}{
			`this_has_entrypoints_to_same_type`: {
				model: `
				model
					schema 1.2
				type document
					relations
						define viewer: [document]`,
			},
			`this_has_entrypoints_through_user_wildcard`: {
				model: `
				model
					schema 1.2
				type document
					relations
						define viewer: [document:*]`,
			},
			`this_has_entrypoints_through_userset`: {
				model: `
				model
					schema 1.2
				type user
				type org
					relations
						define member: [user]
				type folder
					relations
						define parent: [org#member]`,
			},
			`this_with_two_assignable_types_has_entrypoints_through_first`: {
				model: `
				model
					schema 1.2
				type user
				type folder
					relations
						define parent: [user, folder#parent]`,
			},
			`this_with_two_assignable_types_has_entrypoints_through_second`: {
				model: `
				model
					schema 1.2
				type user
				type folder
					relations
						define editor: [user]
						define parent: [folder#parent, folder#editor]`,
			},
			`computed_relation_has_entrypoint_through_user`: {
				model: `
				model
					schema 1.2
				type user
				type document
					relations
						define editor: [user]
						define viewer: editor`,
			},
			`computed_relation_has_entrypoint_through_userset`: {
				model: `
				model
					schema 1.2
				type user
				type org
					relations
					define member: [user]
				type folder
					relations
					define a2: [org#member]
					define a1: a2`,
			},
			`union_has_entrypoint_through_user`: {
				model: `
				model
					schema 1.2
				type user

				type document
					relations
						define editor: [user]
						define viewer: [document#viewer] or editor`,
			},
			`ttu_has_entrypoint_through_user`: {
				model: `
				model
					schema 1.2
				type user
				type org
					relations
						define viewer: [user]
				type folder
					relations
						define parent: [org]
						define viewer: viewer from parent`,
			},
			`intersection_has_entrypoint_and_no_cycle`: {
				model: `
				model
					schema 1.2
				type user

				type document
					relations
						define action1: admin and editor
						define admin: [user]
						define editor: [user]`,
			},
			`difference_has_entrypoints_and_no_cycle`: {
				model: `
				model
					schema 1.2
				type user

				type document
					relations
						define action1: admin but not editor
						define admin: [user]
						define editor: [user]`,
			},
			`difference_has_entrypoints_and_no_cycle_2`: {
				model: `
				model
					schema 1.2
				type user

				type document
					relations
						define restricted: [user]
						define editor: [user]
						define viewer: [document#viewer] or editor
						define can_view: viewer but not restricted
						define can_view_actual: can_view`,
			},
			`issue_1385`: {
				model: `
				model
					schema 1.2

				type user

				type entity
					relations
						define member : [user]
						define contextual_user: [user]
						define contextual_member : member and contextual_user
						define has_logging_product: [entity]
						define block_logging : [user] and contextual_user
						define has_access_to_logging : contextual_member from has_logging_product but not block_logging from has_logging_product
						define can_enable_logging : has_access_to_logging
			`,
			},
			`issue_1260_parallel_edges_mean_entrypoints`: {
				model: `
				model
					schema 1.2

				type user

				type state
					relations
						define can_view: [user]
						define associated_transition: [transition]
						define can_transition_with: can_apply from associated_transition

				type transition
					relations
						define start: [state]
						define end: [state]
						define can_apply: [user] and can_view from start and can_view from end
			`,
			},
			`ttu_has_entrypoint_through_second_tupleset`: {
				model: `
				model
					schema 1.2
				type user
				type group
					relations
						define viewer: [user]
				type folder
					relations
						define parent: [folder, group]
						define viewer: viewer from parent`,
			},
			`revisited_direct_has_entrypoints`: {
				model: `
				model
					schema 1.2

				type user

				type document
					relations
						define a: [user]
						define b: a
						define c: a
						define d: b and c
			`,
			},
		}

		builder := NewWeightedAuthorizationModelGraphBuilder()
		for name, test := range validTests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				model := parser.MustTransformDSLToProto(test.model)
				graph, err := builder.ValidateAndBuild(model)
				require.NoError(t, err)
				require.NotNil(t, graph)
			})
		}
	})
}

// TestErrorDuplicateTypes tests for duplicate type definition errors
func TestErrorDuplicateTypes(t *testing.T) {
	t.Run("detects_duplicate_type_definitions", func(t *testing.T) {
		t.Parallel()
		builder := NewWeightedAuthorizationModelGraphBuilder()

		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "user",
				},
				{
					Type: "user", // Duplicate type
				},
			},
		}

		_, err := builder.ValidateAndBuild(model)
		require.Error(t, err)
		require.ErrorIs(t, err, errors.ErrDuplicateTypes)
	})

	t.Run("duplicate_types_is_invalid", func(t *testing.T) {
		t.Parallel()
		builder := NewWeightedAuthorizationModelGraphBuilder()

		model := &openfgav1.AuthorizationModel{
			SchemaVersion: SchemaVersion1_2,
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type:      "repo",
					Relations: map[string]*openfgav1.Userset{},
				},
				{
					Type:      "repo",
					Relations: map[string]*openfgav1.Userset{},
				},
			},
		}

		_, err := builder.ValidateAndBuild(model)
		require.ErrorIs(t, err, errors.ErrDuplicateTypes)
	})
}

// TestErrorRelationUndefined tests for undefined relation errors
func TestErrorRelationUndefined(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "empty_rewrites",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"reader": {},
						},
					},
				},
			},
		},
		{
			name: "invalid_relation:_computedUserset_to_relation_which_does_not_exist",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader: writer`).GetTypeDefinitions(),
			},
		},
		{
			name: "invalid_relation:_computedUserset_in_a_union",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader: [user] or writer`).GetTypeDefinitions(),
			},
		},
		{
			name: "invalid_relation:_computedUserset_in_a_intersection",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader:[user] and writer`).GetTypeDefinitions(),
			},
		},
		{
			name: "invalid_relation:_computedUserset_in_a_difference_base",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader: [user] but not writer`).GetTypeDefinitions(),
			},
		},
		{
			name: "invalid_relation:_tupleToUserset_where_tupleset_is_not_valid",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader: member from parent`).GetTypeDefinitions(),
			},
		},
		{
			name: "invalid_relation:_tupleToUserset_where_computed_userset_is_not_valid",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user
					type group
					type document
						relations
							define reader: notavalidrelation from writer
							define writer: [group]`).GetTypeDefinitions(),
			},
		},
		{
			name: "assignable_relation_with_no_type:_this",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"reader": {
								Userset: &openfgav1.Userset_This{},
							},
						},
					},
				},
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.ErrorIs(t, err, errors.ErrRelationUndefined)
		})
	}
}

// TestErrorInvalidUsersetRewrite tests for invalid userset rewrite errors
func TestErrorInvalidUsersetRewrite(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "invalid_relation:_self_reference_in_computedUserset",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader: reader`).GetTypeDefinitions(),
			},
		},
		{
			name: "invalid_relation:_self_reference_in_union",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader: [user] or reader`).GetTypeDefinitions(),
			},
		},
		{
			name: "invalid_relation:_self_reference_in_intersection",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader: [user] and reader`).GetTypeDefinitions(),
			},
		},
		{
			name: "invalid_relation:_self_reference_in_difference_substract",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define reader: [user] but not reader`).GetTypeDefinitions(),
			},
		},
		{
			name: "Fails_no_relation_defined_in_userset",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
                    model
                        schema 1.2
                    type user
                    type group
                        relations
                            define member: [user]
                    type folder
                        relations
                            define viewer: [group#computed_member]`).GetTypeDefinitions(),
			},
		},
		{
			name: "undefined_type_in_assignable_type",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type document
					relations
						define viewer: [unknown#editor]`).GetTypeDefinitions(),
			},
		},
		{
			name: "undefined_relation_in_assignable_type",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type document
					relations
						define viewer: [document#unknown]`).GetTypeDefinitions(),
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.ErrorIs(t, err, errors.ErrInvalidUsersetRewrite)
		})
	}
}

// TestErrorCycle tests for cycle detection errors
func TestErrorCycle(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "computed_userset_1",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type resource
					relations
						define x: y
						define y: x`).GetTypeDefinitions(),
			},
		},
		{
			name: "computed_userset_2",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type resource
						relations
							define x: y
							define y: z
							define z: x`).GetTypeDefinitions(),
			},
		},
		{
			name: "union_1",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user
					type resource
						relations
							define x: [user] or y
							define y: [user] or z
							define z: [user] or x`).GetTypeDefinitions(),
			},
		},
		{
			name: "intersection_and_union",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user
					type resource
						relations
							define x: [user] and y
							define y: [user] and z
							define z: [user] or x`).GetTypeDefinitions(),
			},
		},
		{
			name: "exclusion_and_union",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user
					type resource
						relations
							define x: [user] but not y
							define y: [user] but not z
							define z: [user] or x`).GetTypeDefinitions(),
			},
		},
		{
			name: "union_3",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type group
						relations
							define member: [user] or memberA or memberB or memberC
							define memberA: [user] or member or memberB or memberC
							define memberB: [user] or member or memberA or memberC
							define memberC: [user] or member or memberA or memberB`).GetTypeDefinitions(),
			},
		},
		{
			name: "union_4",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type account
						relations
							define admin: [user] or member or super_admin or owner
							define member: [user] or owner or admin or super_admin
							define super_admin: [user] or admin or member or owner
							define owner: [user]`).GetTypeDefinitions(),
			},
		},
		{
			name: "union_5",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type account
					relations
						define admin: [user] or member or super_admin or owner
						define member: [user] or owner or admin or super_admin
						define super_admin: [user] or admin or member or owner
						define owner: [user]`).GetTypeDefinitions(),
			},
		},
		{
			name: "many_circular_computed_relations",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type canvas
					relations
						define can_edit: editor or owner
						define editor: [user, account#member]
						define owner: [user]
						define viewer: [user, account#member]

				type account
					relations
						define admin: [user] or member or super_admin or owner
						define member: [user] or owner or admin or super_admin
						define owner: [user]
						define super_admin: [user] or admin or member`).GetTypeDefinitions(),
			},
		},
		{
			name: "computed_relation_has_no_entrypoints_because_no_direct_relationships",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type folder
					relations
						define a2: a1
						define a1: a2`).GetTypeDefinitions(),
			},
		},
		{
			name: "difference_has_no_entrypoint_and_has_cycle",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user

				type document
					relations
						define admin: [user]
						define action1: admin but not action2
						define action2: admin but not action3
						define action3: admin but not action1`).GetTypeDefinitions(),
			},
		},
		{
			name: "intersection_has_no_entrypoint_and_has_cycle_2",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user

				type document
					relations
						define admin: [user]
						define action1: admin and action2 and action3
						define action2: admin and action1 and action3
						define action3: admin and action1 and action2`).GetTypeDefinitions(),
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.ErrorIs(t, err, errors.ErrCycle)
		})
	}
}

// TestErrorNoEntrypoints tests for lack of entrypoints errors
func TestErrorNoEntrypoints(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "this_has_no_entrypoints_through_userset",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type folder
					relations
						define parent: [folder#parent]`).GetTypeDefinitions(),
			},
		},
		{
			name: "this_has_no_entrypoints_through_recursive_userset",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type group
					relations
						define member: [group#member]

				type folder
					relations
						define parent: [group#member]`).GetTypeDefinitions(),
			},
		},
		{
			name: "computed_relation_has_no_entrypoint_through_usersets",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user
				type document
					relations
						define editor: [document#viewer]
						define viewer: [document#editor]`).GetTypeDefinitions(),
			},
		},
		{
			name: "computed_relation_has_no_entrypoints_through_ttu",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user

				type folder
					relations
						define parent: [document]
						define viewer: editor from parent

				type document
					relations
						define parent: [folder]
						define editor: viewer
						define viewer: viewer from parent`).GetTypeDefinitions(),
			},
		},
		{
			name: "union_has_no_entrypoint",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user

				type document
					relations
						define editor: [document#viewer]
						define viewer: [document#viewer] or editor`).GetTypeDefinitions(),
			},
		},
		{
			name: "ttu_has_no_entrypoint",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type folder
					relations
						define parent: [folder]
						define viewer: viewer from parent`).GetTypeDefinitions(),
			},
		},
		{
			name: "no_entrypoint_4",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user

				type folder
					relations
						define parent: [document]
						define editor: editor from parent

				type document
					relations
						define parent: [folder]
						define editor: viewer
						define viewer: editor from parent
			`).GetTypeDefinitions(),
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.ErrorIs(t, err, errors.ErrNoEntrypoints)
		})
	}
}

// TestErrorObjectTypeUndefined tests for undefined object type errors
func TestErrorObjectTypeUndefined(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "undefined_type",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type document
					relations
						define viewer: [folder]`).GetTypeDefinitions(),
			},
		},
		{
			name: "relational_type_which_does_not_exist",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type document
					relations
						define reader: [group]`).GetTypeDefinitions(),
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.ErrorIs(t, err, errors.ErrObjectTypeUndefined)
		})
	}
}

// TestErrorReservedKeywords tests for reserved keywords errors
func TestErrorReservedKeywords(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "Fails_If_Using_This_As_Relation_Name",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define this: [user]`).GetTypeDefinitions(),
			},
		},
		{
			name: "Fails_If_Using_Self_As_Relation_Name",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type document
						relations
							define self: [user]`).GetTypeDefinitions(),
			},
		},
		{
			name: "Fails_If_Using_This_As_Type_Name",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
					model
						schema 1.2
					type user

					type this
						relations
							define member: [user]`).GetTypeDefinitions(),
			},
		},
		{
			name: "Fails_If_Using_Self_As_Type_Name",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
                    model
                        schema 1.2
                    type user

                    type self
                        relations
                            define member: [user]`).GetTypeDefinitions(),
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.ErrorIs(t, err, errors.ErrReservedKeywords)
		})
	}
}

// TestErrorConstraintTupleCycle tests for constraint tuple cycle errors
func TestErrorConstraintTupleCycle(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "intersection_has_no_entrypoint_and_no_cycle",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user

				type document
					relations
						define action1: [document#action1] and editor
						define editor: [user]`).GetTypeDefinitions(),
			},
		},
		{
			name: "difference_has_no_entrypoint_and_no_cycle",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user

				type document
					relations
						define action1: [document#action1] but not editor
						define editor: [user]`).GetTypeDefinitions(),
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.ErrorIs(t, err, errors.ErrConstraintTupleCycle)
		})
	}
}

// TestErrorCondition tests for undefined condition errors
func TestErrorCondition(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "relation_references_undefined_condition",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "user",
					},
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"viewer": {
								Userset: &openfgav1.Userset_This{
									This: &openfgav1.DirectUserset{},
								},
							},
						},
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{
											Type:      "user",
											Condition: "is_active", // This condition is not defined
										},
									},
								},
							},
						},
					},
				},
				// Note: No conditions are defined here
			},
		},
		{
			name: "multiple_relations_reference_undefined_conditions",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "user",
					},
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"viewer": {
								Userset: &openfgav1.Userset_This{
									This: &openfgav1.DirectUserset{},
								},
							},
							"editor": {
								Userset: &openfgav1.Userset_This{
									This: &openfgav1.DirectUserset{},
								},
							},
						},
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{
											Type:      "user",
											Condition: "is_active", // This condition is not defined
										},
									},
								},
								"editor": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{
											Type:      "user",
											Condition: "is_admin", // This condition is not defined
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.Error(t, err)
			require.ErrorIs(t, err, errors.ErrCondition)
		})
	}
}

// TestErrorConditionUnReferenced tests for unused condition errors
func TestErrorConditionUnReferenced(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "condition_defined_but_not_referenced",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "user",
					},
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"viewer": {
								Userset: &openfgav1.Userset_This{
									This: &openfgav1.DirectUserset{},
								},
							},
						},
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{
											Type: "user",
											// Note: No condition is referenced here
										},
									},
								},
							},
						},
					},
				},
				Conditions: map[string]*openfgav1.Condition{
					"is_active": {
						Name:       "is_active",
						Expression: "user.is_active == true",
					},
				},
			},
		},
		{
			name: "multiple_unused_conditions",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: []*openfgav1.TypeDefinition{
					{
						Type: "user",
					},
					{
						Type: "document",
						Relations: map[string]*openfgav1.Userset{
							"viewer": {
								Userset: &openfgav1.Userset_This{
									This: &openfgav1.DirectUserset{},
								},
							},
							"editor": {
								Userset: &openfgav1.Userset_This{
									This: &openfgav1.DirectUserset{},
								},
							},
						},
						Metadata: &openfgav1.Metadata{
							Relations: map[string]*openfgav1.RelationMetadata{
								"viewer": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{
											Type: "user",
										},
									},
								},
								"editor": {
									DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
										{
											Type: "user",
										},
									},
								},
							},
						},
					},
				},
				Conditions: map[string]*openfgav1.Condition{
					"is_active": {
						Name:       "is_active",
						Expression: "user.is_active == true",
					},
					"is_admin": {
						Name:       "is_admin",
						Expression: "user.role == 'admin'",
					},
					"is_owner": {
						Name:       "is_owner",
						Expression: "user.id == object.owner_id",
					},
				},
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.Error(t, err)
			require.ErrorIs(t, err, errors.ErrConditionUnReferenced)
		})
	}
}

// TestErrorInvalidRelationOnTupleset tests for invalid relation on tupleset errors
func TestErrorInvalidRelationOnTupleset(t *testing.T) {
	tests := []struct {
		name  string
		model *openfgav1.AuthorizationModel
	}{
		{
			name: "userset_specified_as_allowed_type_but_the_relation_is_used_in_a_TTU_rewrite",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
                model
                    schema 1.2
                type folder
                    relations
                        define member: [user]
                type user
                type document
                    relations
                        define reader: member from parent
                        define parent: [folder#member]`).GetTypeDefinitions(),
			},
		},
		{
			name: "userset_specified_as_allowed_type_but_the_relation_is_used_in_a_TTU_rewrite_included_in_a_union",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
                model
                    schema 1.2
                type folder
                    relations
                        define viewer: [user]
                        define parent: [folder]
                type user
                type document
                    relations
                        define reader: [user] or viewer from parent
                        define parent: [folder#parent]`).GetTypeDefinitions(),
			},
		},
		{
			name: "WildcardNotAllowedInTheTuplesetPartOfTTU",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
                model
                    schema 1.2
                type folder
                    relations
                        define viewer: [user]
                        define parent: [folder]
                type user
                type document
                    relations
                        define reader: [user] or viewer from parent
                        define parent: [folder:*]`).GetTypeDefinitions(),
			},
		},
		{
			name: "ttu_has_entrypoint_through_userset",
			model: &openfgav1.AuthorizationModel{
				SchemaVersion: SchemaVersion1_2,
				TypeDefinitions: parser.MustTransformDSLToProto(`
				model
					schema 1.2
				type user
				type org
					relations
						define viewer: [user]
						define member: [user]
				type folder
					relations
						define parent: [org#member]
						define viewer: viewer from parent`).GetTypeDefinitions(),
			},
		},
	}

	builder := NewWeightedAuthorizationModelGraphBuilder()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := builder.ValidateAndBuild(test.model)
			require.ErrorIs(t, err, errors.ErrInvalidRelationOnTupleset)
		})
	}
}
