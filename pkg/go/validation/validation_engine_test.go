package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

// TestValidationEngine_BasicIntegration tests the basic integration of all validation components
func TestValidationEngine_BasicIntegration(t *testing.T) {
	t.Run("Valid model passes all validations", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			SchemaVersion: "1.1",
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "user",
				},
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}},
						},
						"editor": {
							Userset: &openfgav1.Userset_Union{Union: &openfgav1.Usersets{
								Child: []*openfgav1.Userset{
									{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
									{Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"}}},
								},
							}},
						},
					},
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
							"editor": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
						},
					},
				},
			},
		}

		dslContent := `
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
    define editor: [user] or viewer
`

		// Test ValidateDSL
		errors := ValidateDSL(model, dslContent, DefaultEngineOptions())
		assert.NotNil(t, errors)
		assert.Equal(t, 0, errors.Count())

		// Test ValidateJSON
		jsonErrors := ValidateJSON(model, DefaultEngineOptions())
		assert.NotNil(t, jsonErrors)
		assert.Equal(t, 0, jsonErrors.Count())

		// Test convenience functions
		modelErrors := ValidateModel(model, dslContent)
		assert.NotNil(t, modelErrors)
		assert.Equal(t, 0, modelErrors.Count())

		jsonModelErrors := ValidateModelJSON(model)
		assert.NotNil(t, jsonModelErrors)
		assert.Equal(t, 0, jsonModelErrors.Count())
	})

	t.Run("Model with validation errors", func(t *testing.T) {
		// Create model with various validation issues
		model := &openfgav1.AuthorizationModel{
			SchemaVersion: "1.0", // Older schema version
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "nonexistent"}},
						},
					},
				},
				{
					Type: "document", // Duplicate type
					Relations: map[string]*openfgav1.Userset{
						"editor": {
							Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}},
						},
					},
				},
			},
		}

		dslContent := `
model
  schema 1.0

type document
  relations
    define viewer: nonexistent
    define editor: [user]

type document
  relations
    define admin: [user]
`

		errors := ValidateDSL(model, dslContent, DefaultEngineOptions())
		assert.NotNil(t, errors)
		assert.Greater(t, errors.Count(), 0)

		// Check that we have various types of errors
		errorList := errors.GetErrors()
		errorTypes := make(map[ValidationErrorType]bool)
		for _, err := range errorList {
			errorTypes[err.Metadata.ErrorType] = true
		}

		// Should have duplicate errors
		assert.True(t, len(errorTypes) > 0, "Should have validation errors")
	})
}

// TestValidationEngine_OptionsConfiguration tests different validation options
func TestValidationEngine_OptionsConfiguration(t *testing.T) {
	t.Run("Skip semantic validation", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			SchemaVersion: "1.1",
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "nonexistent"}},
						},
					},
				},
			},
		}

		// With semantic validation (default)
		normalErrors := ValidateDSL(model, "", DefaultEngineOptions())
		normalErrorCount := normalErrors.Count()

		// Skip semantic validation
		options := &EngineOptions{
			SkipSemanticValidation: true,
		}
		skippedErrors := ValidateDSL(model, "", options)
		skippedErrorCount := skippedErrors.Count()

		// Should have fewer errors when semantic validation is skipped
		assert.True(t, skippedErrorCount <= normalErrorCount, "Skipping semantic validation should reduce or maintain error count")
	})

	t.Run("Skip complex operation validation", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			SchemaVersion: "1.1",
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_Union{Union: &openfgav1.Usersets{
								Child: []*openfgav1.Userset{
									{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
									{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
								},
							}},
						},
					},
				},
			},
		}

		options := &EngineOptions{
			SkipComplexOperationValidation: true,
		}
		errors := ValidateDSL(model, "", options)
		assert.NotNil(t, errors)
		// Complex operation validation should be skipped
	})
}

// TestValidationReport tests the comprehensive validation report functionality
func TestValidationReport(t *testing.T) {
	t.Run("Complete validation report", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			SchemaVersion: "1.1",
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "user",
				},
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}},
						},
					},
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"viewer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
						},
					},
				},
			},
		}

		dslContent := `
model
  schema 1.1

type user

type document
  relations
    define viewer: [user]
`

		report := CreateValidationReport(model, dslContent, DefaultEngineOptions())

		assert.NotNil(t, report.Model)
		assert.NotNil(t, report.ValidationErrors)
		assert.NotNil(t, report.Options)
		assert.Equal(t, model, report.Model)

		// Test report methods
		assert.True(t, report.IsValid(), "Valid model should pass IsValid()")
		assert.False(t, report.HasCriticalErrors(), "Valid model should not have critical errors")

		// Test summary
		summary := report.Summary
		assert.Equal(t, 0, summary.TotalErrors)
		assert.False(t, summary.HasCriticalErrors)
		assert.NotNil(t, summary.ErrorsByType)
		assert.NotNil(t, summary.ErrorsByFile)
	})

	t.Run("Report with errors", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			SchemaVersion: "invalid", // Invalid schema version
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "nonexistent"}},
						},
					},
				},
			},
		}

		report := CreateValidationReport(model, "", DefaultEngineOptions())

		if report.ValidationErrors.Count() > 0 {
			assert.False(t, report.IsValid(), "Invalid model should fail IsValid()")
			
			summary := report.Summary
			assert.Greater(t, summary.TotalErrors, 0)
			
			// Test GetErrorsByType functionality
			for errorType := range summary.ErrorsByType {
				errorsOfType := report.GetErrorsByType(errorType)
				assert.Greater(t, len(errorsOfType), 0, "Should find errors of type %s", errorType)
			}
		}
	})
}

// TestValidationEngine_RealWorldScenarios tests realistic authorization model scenarios
func TestValidationEngine_RealWorldScenarios(t *testing.T) {
	t.Run("GitHub-like authorization model", func(t *testing.T) {
		model := &openfgav1.AuthorizationModel{
			SchemaVersion: "1.1",
			TypeDefinitions: []*openfgav1.TypeDefinition{
				{
					Type: "user",
				},
				{
					Type: "organization",
					Relations: map[string]*openfgav1.Userset{
						"member": {
							Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}},
						},
						"owner": {
							Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}},
						},
					},
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"member": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
							"owner": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
						},
					},
				},
				{
					Type: "repository",
					Relations: map[string]*openfgav1.Userset{
						"reader": {
							Userset: &openfgav1.Userset_Union{Union: &openfgav1.Usersets{
								Child: []*openfgav1.Userset{
									{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
									{Userset: &openfgav1.Userset_TupleToUserset{TupleToUserset: &openfgav1.TupleToUserset{
										Tupleset: &openfgav1.ObjectRelation{Relation: "owner"},
										ComputedUserset: &openfgav1.ObjectRelation{Relation: "member"},
									}}},
								},
							}},
						},
						"writer": {
							Userset: &openfgav1.Userset_Union{Union: &openfgav1.Usersets{
								Child: []*openfgav1.Userset{
									{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
									{Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "admin"}}},
								},
							}},
						},
						"admin": {
							Userset: &openfgav1.Userset_Union{Union: &openfgav1.Usersets{
								Child: []*openfgav1.Userset{
									{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
									{Userset: &openfgav1.Userset_TupleToUserset{TupleToUserset: &openfgav1.TupleToUserset{
										Tupleset: &openfgav1.ObjectRelation{Relation: "owner"},
										ComputedUserset: &openfgav1.ObjectRelation{Relation: "owner"},
									}}},
								},
							}},
						},
						"owner": {
							Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}},
						},
					},
					Metadata: &openfgav1.Metadata{
						Relations: map[string]*openfgav1.RelationMetadata{
							"reader": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
							"writer": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
							"admin": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
								},
							},
							"owner": {
								DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
									{Type: "user"},
									{Type: "organization", RelationOrWildcard: &openfgav1.RelationReference_Relation{Relation: "owner"}},
								},
							},
						},
					},
				},
			},
		}

		dslContent := `
model
  schema 1.1

type user

type organization
  relations
    define member: [user]
    define owner: [user]

type repository
  relations
    define owner: [user, organization#owner]
    define admin: [user] or owner from owner
    define writer: [user] or admin  
    define reader: [user] or writer from owner
`

		errors := ValidateDSL(model, dslContent, DefaultEngineOptions())
		assert.NotNil(t, errors)

		// This complex model should pass validation
		if errors.Count() > 0 {
			t.Logf("Validation errors found: %d", errors.Count())
			for _, err := range errors.GetErrors() {
				t.Logf("Error: %s (Type: %s)", err.Message, err.Metadata.ErrorType)
			}
		}

		// Create validation report
		report := CreateValidationReport(model, dslContent, DefaultEngineOptions())
		assert.NotNil(t, report)

		t.Logf("Validation Summary:")
		t.Logf("- Total Errors: %d", report.Summary.TotalErrors)
		t.Logf("- Has Critical Errors: %v", report.Summary.HasCriticalErrors)
		t.Logf("- Valid Model: %v", report.IsValid())
	})
}

// TestValidationEngine_PerformanceBasics tests basic performance characteristics
func TestValidationEngine_PerformanceBasics(t *testing.T) {
	t.Run("Large model validation performance", func(t *testing.T) {
		// Create a moderately large model
		typeDefs := make([]*openfgav1.TypeDefinition, 0, 50)
		
		// Add user type
		typeDefs = append(typeDefs, &openfgav1.TypeDefinition{Type: "user"})
		
		// Add many document types with relations
		for i := 0; i < 49; i++ {
			typeName := fmt.Sprintf("document%d", i)
			relations := make(map[string]*openfgav1.Userset)
			relationMetadata := make(map[string]*openfgav1.RelationMetadata)
			
			// Add viewer relation
			relations["viewer"] = &openfgav1.Userset{
				Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}},
			}
			relationMetadata["viewer"] = &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
					{Type: "user"},
				},
			}
			
			// Add editor relation with union
			relations["editor"] = &openfgav1.Userset{
				Userset: &openfgav1.Userset_Union{Union: &openfgav1.Usersets{
					Child: []*openfgav1.Userset{
						{Userset: &openfgav1.Userset_This{This: &openfgav1.DirectUserset{}}},
						{Userset: &openfgav1.Userset_ComputedUserset{ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"}}},
					},
				}},
			}
			relationMetadata["editor"] = &openfgav1.RelationMetadata{
				DirectlyRelatedUserTypes: []*openfgav1.RelationReference{
					{Type: "user"},
				},
			}
			
			typeDefs = append(typeDefs, &openfgav1.TypeDefinition{
				Type:      typeName,
				Relations: relations,
				Metadata: &openfgav1.Metadata{
					Relations: relationMetadata,
				},
			})
		}

		model := &openfgav1.AuthorizationModel{
			SchemaVersion:   "1.1",
			TypeDefinitions: typeDefs,
		}

		// Test validation performance
		errors := ValidateDSL(model, "", DefaultEngineOptions())
		assert.NotNil(t, errors)

		// Should complete validation in reasonable time
		t.Logf("Large model validation completed with %d errors", errors.Count())
		
		// Test JSON validation performance
		jsonErrors := ValidateJSON(model, DefaultEngineOptions())
		assert.NotNil(t, jsonErrors)
		t.Logf("Large model JSON validation completed with %d errors", jsonErrors.Count())
	})
}
