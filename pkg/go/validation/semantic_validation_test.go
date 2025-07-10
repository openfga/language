package validation

import (
	"testing"

	fgaSdk "github.com/openfga/go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestSemanticValidator(t *testing.T) {
	t.Run("NewSemanticValidator", func(t *testing.T) {
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Relations: &map[string]fgaSdk.Userset{
						"viewer": {
							This: &map[string]interface{}{},
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
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
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
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Relations: &map[string]fgaSdk.Userset{
						"viewer": {
							This: &map[string]interface{}{},
						},
						"editor": {
							This: &map[string]interface{}{},
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
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
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
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{Type: "user"},
								},
							},
						},
					},
					Relations: &map[string]fgaSdk.Userset{
						"viewer": {
							This: &map[string]interface{}{},
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
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
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
		assert.Equal(t, UndefinedType, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "undefined_type")
	})

	t.Run("Undefined relation in restriction", func(t *testing.T) {
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{Type: "user", Relation: fgaSdk.PtrString("undefined_relation")},
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
		assert.Equal(t, UndefinedRelation, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "undefined_relation")
	})

	t.Run("Undefined relation in computed userset", func(t *testing.T) {
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Relations: &map[string]fgaSdk.Userset{
						"viewer": {
							ComputedUserset: &fgaSdk.ObjectRelation{
								Relation: fgaSdk.PtrString("undefined_relation"),
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
		assert.Equal(t, UndefinedRelation, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "undefined_relation")
	})

	t.Run("Complex userset validation", func(t *testing.T) {
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Relations: &map[string]fgaSdk.Userset{
						"viewer": {
							Union: &fgaSdk.Usersets{
								Child: []fgaSdk.Userset{
									{
										ComputedUserset: &fgaSdk.ObjectRelation{
											Relation: fgaSdk.PtrString("editor"),
										},
									},
									{
										ComputedUserset: &fgaSdk.ObjectRelation{
											Relation: fgaSdk.PtrString("undefined_relation"),
										},
									},
								},
							},
						},
						"editor": {
							This: &map[string]interface{}{},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateRelationReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, UndefinedRelation, errors[0].Metadata.ErrorType)
		assert.Contains(t, errors[0].Message, "undefined_relation")
	})
}
