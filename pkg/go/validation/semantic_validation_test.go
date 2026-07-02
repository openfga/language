package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
)

func TestSemanticValidator(t *testing.T) {
	t.Run("NewSemanticValidator", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
				},
				{
					Type: "user",
				},
			},
		}

		validator := NewSemanticValidator(model)

		assert.NotNil(t, validator)
		assert.Equal(t, model, validator.model)
		assert.Len(t, validator.typeMap, 2)
		assert.Len(t, validator.relationMap, 1)
	})

	t.Run("TypeDefined", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{Type: "document"},
				{Type: "user"},
			},
		}

		validator := NewSemanticValidator(model)

		assert.True(t, validator.TypeDefined("document"))
		assert.True(t, validator.TypeDefined("user"))
		assert.False(t, validator.TypeDefined("group"))
		assert.False(t, validator.TypeDefined(""))
	})

	t.Run("RelationDefined", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
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
				},
				{
					Type: "user",
				},
			},
		}

		validator := NewSemanticValidator(model)

		assert.True(t, validator.RelationDefined("document", "viewer"))
		assert.True(t, validator.RelationDefined("document", "editor"))
		assert.False(t, validator.RelationDefined("document", "admin"))
		assert.False(t, validator.RelationDefined("user", "viewer"))
		assert.False(t, validator.RelationDefined("group", "viewer"))
	})

	t.Run("GetTypeDefinition", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{Type: "document"},
				{Type: "user"},
			},
		}

		validator := NewSemanticValidator(model)

		docType := validator.GetTypeDefinition("document")
		assert.NotNil(t, docType)
		assert.Equal(t, "document", docType.Type)

		userType := validator.GetTypeDefinition("user")
		assert.NotNil(t, userType)
		assert.Equal(t, "user", userType.Type)

		groupType := validator.GetTypeDefinition("group")
		assert.Nil(t, groupType)
	})
}

func TestValidateRelationReferences(t *testing.T) {
	t.Run("Valid references", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
						},
					},
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
				},
				{
					Type: "user",
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateRelationReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Empty(t, errors)
	})

	t.Run("Undefined type in restriction", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "undefined_type"},
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateRelationReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, InvalidType, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "undefined_type")
	})

	t.Run("Undefined relation in restriction", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user", RelationOrWildcard: &openfgav1.RelationReference_Relation{Relation: "undefined_relation"}},
								},
							},
						},
					},
				},
				{
					Type: "user",
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateRelationReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, InvalidRelationType, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "undefined_relation")
	})

	t.Run("Undefined relation in computed userset", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_ComputedUserset{
								ComputedUserset: &openfgav1.ObjectRelation{
									Relation: "undefined_relation",
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateRelationReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, MissingDefinition, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "undefined_relation")
	})

	t.Run("Complex userset validation", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_Union{
								Union: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{
											Userset: &openfgav1.Userset_ComputedUserset{
												ComputedUserset: &openfgav1.ObjectRelation{
													Relation: "editor",
												},
											},
										},
										{
											Userset: &openfgav1.Userset_ComputedUserset{
												ComputedUserset: &openfgav1.ObjectRelation{
													Relation: "undefined_relation",
												},
											},
										},
									},
								},
							},
						},
						"editor": {
							Userset: &openfgav1.Userset_This{
								This: &openfgav1.DirectUserset{},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateRelationReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, MissingDefinition, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "undefined_relation")
	})
}
