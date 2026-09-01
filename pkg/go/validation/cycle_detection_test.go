package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
)

// entryOf runs the entry-point traversal for a single relation from a fresh
// visited set.
func entryOf(idx *index, typeName, relationName string) entryPointResult {
	return hasEntryPointOrLoop(idx, typeName, relationName,
		idx.userset(typeName, relationName), map[string]map[string]bool{})
}

func TestValidateEntryPoints(t *testing.T) {
	t.Run("Mutual computed-userset loop has no entry point", func(t *testing.T) {
		// viewer -> editor -> viewer, neither directly assignable.
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{Relation: "editor"},
							},
						},
						"editor": {
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"},
							},
						},
					},
				},
			},
		}

		findings := validateEntryPoints(newIndex(model), source{})

		// Each relation is impossible: one finding per relation, all RelationNoEntrypoint.
		assert.Len(t, findings, 2)
		for _, finding := range findings {
			assert.Equal(t, RelationNoEntrypoint, finding.Metadata.Kind)
			assert.Contains(t, finding.Message, "is an impossible relation")
			assert.Contains(t, finding.Message, "(potential loop)")
		}
	})

	t.Run("No findings when relations are reachable", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {DirectlyRelatedUserTypes: []*openfgav1.RelationReference{{Type: "user"}}},
							"editor": {DirectlyRelatedUserTypes: []*openfgav1.RelationReference{{Type: "user"}}},
						},
					},
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_Union{
								Union: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
										{Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "editor"}}},
									},
								},
							},
						},
						"editor": {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
					},
				},
				{Type: "user"},
			},
		}

		assert.Empty(t, validateEntryPoints(newIndex(model), source{}))
	})

	t.Run("Computed chain terminating in a direct assignment is reachable", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"owner": {DirectlyRelatedUserTypes: []*openfgav1.RelationReference{{Type: "user"}}},
						},
					},
					Relations: map[string]*openfgav1.Userset{
						"viewer": {Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "editor"}}},
						"editor": {Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "owner"}}},
						"owner":  {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
					},
				},
				{Type: "user"},
			},
		}

		// All three relations resolve to owner's direct assignment.
		assert.Empty(t, validateEntryPoints(newIndex(model), source{}))
	})
}

func TestHasEntryPointOrLoop(t *testing.T) {
	t.Run("Direct this assignment", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {DirectlyRelatedUserTypes: []*openfgav1.RelationReference{{Type: "user"}}},
						},
					},
					Relations: map[string]*openfgav1.Userset{
						"viewer": {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
					},
				},
				{Type: "user"},
			},
		}

		assert.True(t, entryOf(newIndex(model), "document", "viewer").hasEntry)
	})

	t.Run("Union with this has entry point", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {DirectlyRelatedUserTypes: []*openfgav1.RelationReference{{Type: "user"}}},
						},
					},
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_Union{
								Union: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
										{Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "editor"}}},
									},
								},
							},
						},
					},
				},
				{Type: "user"},
			},
		}

		assert.True(t, entryOf(newIndex(model), "document", "viewer").hasEntry)
	})

	t.Run("Self-referential computed userset is a loop", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"}}},
					},
				},
			},
		}

		result := entryOf(newIndex(model), "document", "viewer")
		assert.False(t, result.hasEntry)
		assert.True(t, result.loop)
	})
}

func TestHasEntryPointOrLoop_TupleToUserset(t *testing.T) {
	// ttuModel builds: document has `parent` (a tupleset assignable to folder) and
	// `viewer: viewer from parent`. folderViewer controls whether the resolved
	// folder#viewer relation exists / has an entry point.
	ttuModel := func(folderViewer *openfgav1.Userset) *openfgav1.AuthorizationModel {
		folder := &openfgav1.TypeDefinition{
			Type: "folder",
			Metadata: &openfgav1.Metadata{
				Relations: map[string]*openfgav1.RelationMetadata{
					"viewer": {DirectlyRelatedUserTypes: []*openfgav1.RelationReference{{Type: "user"}}},
				},
			},
			Relations: map[string]*openfgav1.Userset{},
		}
		if folderViewer != nil {
			folder.Relations["viewer"] = folderViewer
		} else {
			delete(folder.GetMetadata().GetRelations(), "viewer")
		}
		return &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"parent": {DirectlyRelatedUserTypes: []*openfgav1.RelationReference{{Type: "folder"}}},
						},
					},
					Relations: map[string]*openfgav1.Userset{
						"parent": {Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
						"viewer": {Userset: &openfgav1.Userset_TupleToUserset{
							TupleToUserset: &openfgav1.TupleToUserset{
								Tupleset:        &openfgav1.ObjectRelation{Relation: "parent"},
								ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"},
							},
						}},
					},
				},
				folder,
				{Type: "user"},
			},
		}
	}

	t.Run("TTU resolving to a direct assignment has an entry point", func(t *testing.T) {
		// folder#viewer is directly assignable to user, so document#viewer reaches it.
		idx := newIndex(ttuModel(&openfgav1.Userset{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}}))

		result := entryOf(idx, "document", "viewer")
		assert.True(t, result.hasEntry)
		assert.False(t, result.loop)
		// folder#viewer is itself directly assignable to user, so it also has an entry point.
		assert.True(t, entryOf(idx, "folder", "viewer").hasEntry)
		// document#parent is directly assignable to folder, so it also has an entry point.
		assert.True(t, entryOf(idx, "document", "parent").hasEntry)
	})

	t.Run("TTU whose computed relation is missing on the assignable type has no entry point", func(t *testing.T) {
		// folder has no `viewer` relation at all, so the computed lookup is nil and
		// the TTU branch skips it (the assignable == nil continue).
		result := entryOf(newIndex(ttuModel(nil)), "document", "viewer")
		assert.False(t, result.hasEntry)
		assert.False(t, result.loop)
	})

	t.Run("TTU through a self-looping computed relation has no entry point", func(t *testing.T) {
		// folder#viewer computes itself, so it never bottoms out at a concrete type.
		// The TTU branch swallows the looping sub-result and reports no entry point.
		idx := newIndex(ttuModel(&openfgav1.Userset{
			Userset: &openfgav1.Userset_ComputedUserset{
				ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"},
			},
		}))

		assert.False(t, entryOf(idx, "document", "viewer").hasEntry)
	})
}
