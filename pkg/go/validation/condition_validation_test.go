package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
)

func TestNewConditionValidator(t *testing.T) {
	t.Run("Empty model", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{}
		validator := NewConditionValidator(model)

		assert.NotNil(t, validator)
		assert.Equal(t, model, validator.model)
		assert.Empty(t, validator.definedConds)
		assert.Empty(t, validator.usedConds)
		assert.Empty(t, validator.conditionRefs)
	})

	t.Run("Model with conditions", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			Conditions: map[string]*openfgav1.Condition{
				"is_owner": {Name: "is_owner"},
				"is_admin": {Name: "is_admin"},
			},
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "is_owner",
									},
								},
							},
						},
					},
				},
			},
		}

		validator := NewConditionValidator(model)

		assert.NotNil(t, validator)
		assert.Len(t, validator.definedConds, 2)
		assert.Len(t, validator.usedConds, 1)
		assert.True(t, validator.IsConditionDefined("is_owner"))
		assert.True(t, validator.IsConditionDefined("is_admin"))
		assert.True(t, validator.IsConditionUsed("is_owner"))
		assert.False(t, validator.IsConditionUsed("is_admin"))
	})
}

func TestConditionValidator_GetDefinedConditions(t *testing.T) {
	model := &openfgav1.AuthorizationModel{
		Conditions: map[string]*openfgav1.Condition{
			"condition1": {Name: "condition1"},
			"condition2": {Name: "condition2"},
			"condition3": {Name: "condition3"},
		},
	}

	validator := NewConditionValidator(model)
	defined := validator.GetDefinedConditions()

	assert.Len(t, defined, 3)
	assert.Contains(t, defined, "condition1")
	assert.Contains(t, defined, "condition2")
	assert.Contains(t, defined, "condition3")
}

func TestConditionValidator_GetUsedConditions(t *testing.T) {
	model := &openfgav1.AuthorizationModel{
		Conditions: map[string]*openfgav1.Condition{
			"used_condition":   {Name: "used_condition"},
			"unused_condition": {Name: "unused_condition"},
		},
		TypeDefinitions: []*openfgav1.TypeDefinition{
			{
				Type: "document",
				Metadata: &openfgav1.Metadata{
					Relations: map[string]*openfgav1.RelationMetadata{
						"viewer": {
							DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
								{
									Type: "user",
									Condition: "used_condition",
								},
							},
						},
					},
				},
			},
		},
	}

	validator := NewConditionValidator(model)
	used := validator.GetUsedConditions()

	assert.Len(t, used, 1)
	assert.Contains(t, used, "used_condition")
	assert.NotContains(t, used, "unused_condition")
}

func TestConditionValidator_GetConditionReferences(t *testing.T) {
	model := &openfgav1.AuthorizationModel{
		Conditions: map[string]*openfgav1.Condition{
			"test_condition": {Name: "test_condition"},
		},
		TypeDefinitions: []*openfgav1.TypeDefinition{
			{
				Type: "document",
				Metadata: &openfgav1.Metadata{
					Relations: map[string]*openfgav1.RelationMetadata{
						"viewer": {
							DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
								{
									Type: "user",
									Condition: "test_condition",
								},
							},
						},
						"editor": {
							DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
								{
									Type: "user",
									Condition: "test_condition",
								},
							},
						},
					},
				},
			},
		},
	}

	validator := NewConditionValidator(model)
	refs := validator.GetConditionReferences("test_condition")

	assert.Len(t, refs, 2)

	// Check that we have references from both viewer and editor relations
	viewerFound := false
	editorFound := false
	for _, ref := range refs {
		if ref.RelationName == "viewer" {
			viewerFound = true
			assert.Equal(t, "document", ref.TypeName)
			assert.Equal(t, "type_restriction", ref.Context)
		}
		if ref.RelationName == "editor" {
			editorFound = true
			assert.Equal(t, "document", ref.TypeName)
			assert.Equal(t, "type_restriction", ref.Context)
		}
	}
	assert.True(t, viewerFound)
	assert.True(t, editorFound)
}

func TestValidateUnusedConditions(t *testing.T) {
	t.Run("No unused conditions", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			Conditions: map[string]*openfgav1.Condition{
				"used_condition": {Name: "used_condition"},
			},
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "used_condition",
									},
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateUnusedConditions(collector, model, nil)

		errors := collector.GetErrors()
		assert.Empty(t, errors)
	})

	t.Run("Unused condition detected", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			Conditions: map[string]*openfgav1.Condition{
				"unused_condition": {Name: "unused_condition"},
				"used_condition":   {Name: "used_condition"},
			},
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "used_condition",
									},
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateUnusedConditions(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, ConditionNotUsed, errors[0].Metadata.ErrorType)
		assert.Equal(t, "unused_condition", errors[0].Metadata.Symbol)
		assert.Contains(t, errors[0].Message, "unused_condition")
		assert.Contains(t, errors[0].Message, "is not used in the model")
	})

	t.Run("Multiple unused conditions", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			Conditions: map[string]*openfgav1.Condition{
				"unused1":        {Name: "unused1"},
				"unused2":        {Name: "unused2"},
				"used_condition": {Name: "used_condition"},
			},
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "used_condition",
									},
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateUnusedConditions(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 2)

		// Check that both unused conditions are reported
		unusedConditions := make([]string, 0)
		for _, err := range errors {
			assert.Equal(t, ConditionNotUsed, err.Metadata.ErrorType)
			unusedConditions = append(unusedConditions, err.Metadata.Symbol)
		}
		assert.Contains(t, unusedConditions, "unused1")
		assert.Contains(t, unusedConditions, "unused2")
	})
}

func TestValidateConditionReferences(t *testing.T) {
	t.Run("All referenced conditions defined", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			Conditions: map[string]*openfgav1.Condition{
				"valid_condition": {Name: "valid_condition"},
			},
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "valid_condition",
									},
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateConditionReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Empty(t, errors)
	})

	t.Run("Undefined condition referenced", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "undefined_condition",
									},
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateConditionReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, ConditionNotDefined, errors[0].Metadata.ErrorType)
		assert.Equal(t, "undefined_condition", errors[0].Metadata.Symbol)
		assert.Contains(t, errors[0].Message, "undefined_condition")
	})

	t.Run("Multiple undefined conditions", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "undefined1",
									},
								},
							},
							"editor": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "undefined2",
									},
								},
							},
						},
					},
				},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateConditionReferences(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 2)

		// Check that both undefined conditions are reported
		undefinedConditions := make([]string, 0)
		for _, err := range errors {
			assert.Equal(t, ConditionNotDefined, err.Metadata.ErrorType)
			undefinedConditions = append(undefinedConditions, err.Metadata.Symbol)
		}
		assert.Contains(t, undefinedConditions, "undefined1")
		assert.Contains(t, undefinedConditions, "undefined2")
	})
}

func TestValidateConditionConsistency(t *testing.T) {
	t.Run("Valid condition consistency", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			Conditions: map[string]*openfgav1.Condition{
				"valid_condition": {Name: "valid_condition"},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateConditionConsistency(collector, model, nil)

		errors := collector.GetErrors()
		assert.Empty(t, errors)
	})

	t.Run("Anonymous condition detected", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			Conditions: map[string]*openfgav1.Condition{
				"": {Name: ""}, // Anonymous condition
			},
		}

		collector := NewErrorCollector(nil)
		ValidateConditionConsistency(collector, model, nil)

		errors := collector.GetErrors()
		assert.Len(t, errors, 1)
		assert.Equal(t, DifferentNestedConditionName, errors[0].Metadata.ErrorType)
	})
}

func TestScanForConditionUsage(t *testing.T) {
	t.Run("Complex condition usage scanning", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			Conditions: map[string]*openfgav1.Condition{
				"condition1": {Name: "condition1"},
				"condition2": {Name: "condition2"},
				"condition3": {Name: "condition3"},
			},
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{
										Type: "user",
									Condition: "condition1",
									},
									{
										Type: "group",
									Condition: "condition2",
									},
								},
							},
						},
					},
					Relations: map[string]*openfgav1.Userset{
						"editor": {
							Userset: &openfgav1.Userset_Union{
								Union: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{
											Userset: &openfgav1.Userset_This{
												This: &openfgav1.DirectUserset{},
											},
										},
										{
											Userset: &openfgav1.Userset_ComputedUserset{
												ComputedUserset: &openfgav1.ObjectRelation{
													Relation: "viewer",
												},
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

		validator := NewConditionValidator(model)

		assert.True(t, validator.IsConditionUsed("condition1"))
		assert.True(t, validator.IsConditionUsed("condition2"))
		assert.False(t, validator.IsConditionUsed("condition3"))

		// Check condition references
		refs1 := validator.GetConditionReferences("condition1")
		refs2 := validator.GetConditionReferences("condition2")
		refs3 := validator.GetConditionReferences("condition3")

		assert.Len(t, refs1, 1)
		assert.Len(t, refs2, 1)
		assert.Empty(t, refs3)

		assert.Equal(t, "document", refs1[0].TypeName)
		assert.Equal(t, "viewer", refs1[0].RelationName)
		assert.Equal(t, "type_restriction", refs1[0].Context)
	})
}
