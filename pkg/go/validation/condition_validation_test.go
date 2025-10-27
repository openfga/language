package validation

import (
	"testing"

	fgaSdk "github.com/openfga/go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestNewConditionValidator(t *testing.T) {
	t.Run("Empty model", func(t *testing.T) {
		model := &fgaSdk.AuthorizationModel{}
		validator := NewConditionValidator(model)

		assert.NotNil(t, validator)
		assert.Equal(t, model, validator.model)
		assert.Empty(t, validator.definedConds)
		assert.Empty(t, validator.usedConds)
		assert.Empty(t, validator.conditionRefs)
	})

	t.Run("Model with conditions", func(t *testing.T) {
		conditionsMap := map[string]fgaSdk.Condition{
			"is_owner": {Name: "is_owner"},
			"is_admin": {Name: "is_admin"},
		}
		relationsMap := map[string]fgaSdk.RelationMetadata{
			"viewer": {
				DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
					{
						Type:      "user",
						Condition: fgaSdk.PtrString("is_owner"),
					},
				},
			},
		}
		model := &fgaSdk.AuthorizationModel{
			Conditions: &conditionsMap,
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &relationsMap,
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
	model := &fgaSdk.AuthorizationModel{
		Conditions: &map[string]fgaSdk.Condition{
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
	model := &fgaSdk.AuthorizationModel{
		Conditions: &map[string]fgaSdk.Condition{
			"used_condition":   {Name: "used_condition"},
			"unused_condition": {Name: "unused_condition"},
		},
		TypeDefinitions: []fgaSdk.TypeDefinition{
			{
				Type: "document",
				Metadata: &fgaSdk.Metadata{
					Relations: &map[string]fgaSdk.RelationMetadata{
						"viewer": {
							DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
								{
									Type:      "user",
									Condition: fgaSdk.PtrString("used_condition"),
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
	model := &fgaSdk.AuthorizationModel{
		Conditions: &map[string]fgaSdk.Condition{
			"test_condition": {Name: "test_condition"},
		},
		TypeDefinitions: []fgaSdk.TypeDefinition{
			{
				Type: "document",
				Metadata: &fgaSdk.Metadata{
					Relations: &map[string]fgaSdk.RelationMetadata{
						"viewer": {
							DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
								{
									Type:      "user",
									Condition: fgaSdk.PtrString("test_condition"),
								},
							},
						},
						"editor": {
							DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
								{
									Type:      "user",
									Condition: fgaSdk.PtrString("test_condition"),
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
		model := &fgaSdk.AuthorizationModel{
			Conditions: &map[string]fgaSdk.Condition{
				"used_condition": {Name: "used_condition"},
			},
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{
										Type:      "user",
										Condition: fgaSdk.PtrString("used_condition"),
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
		model := &fgaSdk.AuthorizationModel{
			Conditions: &map[string]fgaSdk.Condition{
				"unused_condition": {Name: "unused_condition"},
				"used_condition":   {Name: "used_condition"},
			},
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{
										Type:      "user",
										Condition: fgaSdk.PtrString("used_condition"),
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
		assert.Contains(t, errors[0].Message, "defined but not used")
	})

	t.Run("Multiple unused conditions", func(t *testing.T) {
		model := &fgaSdk.AuthorizationModel{
			Conditions: &map[string]fgaSdk.Condition{
				"unused1":        {Name: "unused1"},
				"unused2":        {Name: "unused2"},
				"used_condition": {Name: "used_condition"},
			},
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{
										Type:      "user",
										Condition: fgaSdk.PtrString("used_condition"),
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
		model := &fgaSdk.AuthorizationModel{
			Conditions: &map[string]fgaSdk.Condition{
				"valid_condition": {Name: "valid_condition"},
			},
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{
										Type:      "user",
										Condition: fgaSdk.PtrString("valid_condition"),
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
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{
										Type:      "user",
										Condition: fgaSdk.PtrString("undefined_condition"),
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
		model := &fgaSdk.AuthorizationModel{
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{
										Type:      "user",
										Condition: fgaSdk.PtrString("undefined1"),
									},
								},
							},
							"editor": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{
										Type:      "user",
										Condition: fgaSdk.PtrString("undefined2"),
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
		model := &fgaSdk.AuthorizationModel{
			Conditions: &map[string]fgaSdk.Condition{
				"valid_condition": {Name: "valid_condition"},
			},
		}

		collector := NewErrorCollector(nil)
		ValidateConditionConsistency(collector, model, nil)

		errors := collector.GetErrors()
		assert.Empty(t, errors)
	})

	t.Run("Anonymous condition detected", func(t *testing.T) {
		model := &fgaSdk.AuthorizationModel{
			Conditions: &map[string]fgaSdk.Condition{
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
		model := &fgaSdk.AuthorizationModel{
			Conditions: &map[string]fgaSdk.Condition{
				"condition1": {Name: "condition1"},
				"condition2": {Name: "condition2"},
				"condition3": {Name: "condition3"},
			},
			TypeDefinitions: []fgaSdk.TypeDefinition{
				{
					Type: "document",
					Metadata: &fgaSdk.Metadata{
						Relations: &map[string]fgaSdk.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: &[]fgaSdk.RelationReference{
									{
										Type:      "user",
										Condition: fgaSdk.PtrString("condition1"),
									},
									{
										Type:      "group",
										Condition: fgaSdk.PtrString("condition2"),
									},
								},
							},
						},
					},
					Relations: &map[string]fgaSdk.Userset{
						"editor": {
							Union: &fgaSdk.Usersets{
								Child: []fgaSdk.Userset{
									{This: &map[string]interface{}{}},
									{ComputedUserset: &fgaSdk.ObjectRelation{Relation: fgaSdk.PtrString("viewer")}},
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
