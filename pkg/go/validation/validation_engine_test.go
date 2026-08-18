package validation

import (
	"errors"
	"fmt"
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
	"github.com/openfga/language/pkg/go/transformer"
)

// findingsFrom recovers the collection behind an error returned by a validation entry
// point. A nil error becomes an empty collection, so a test can read Count and
// GetErrors off the result either way.
func findingsFrom(err error) *ValidationErrors {
	var validationErrors *ValidationErrors
	if errors.As(err, &validationErrors) {
		return validationErrors
	}

	return NewValidationErrors(nil)
}

// TestValidationEngine_BasicIntegration tests the basic integration of all validation components.
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

		// A valid model reports nothing, from every entry point.
		assert.NoError(t, ValidateDSL(model, dslContent, DefaultEngineOptions()))
		assert.NoError(t, ValidateJSON(model, DefaultEngineOptions()))
		assert.NoError(t, ValidateModel(model, dslContent))
		assert.NoError(t, ValidateModelJSON(model))
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

		findings := findingsFrom(ValidateDSL(model, dslContent, DefaultEngineOptions()))
		assert.Positive(t, findings.Count())

		// Check that we have various types of errors
		errorList := findings.GetErrors()
		errorTypes := make(map[ValidationErrorType]bool)
		for _, err := range errorList {
			errorTypes[err.Metadata.ErrorType] = true
		}

		// Should have duplicate errors
		assert.NotEmpty(t, errorTypes, "Should have validation errors")
	})
}

// TestValidationEngine_OptionsConfiguration tests different validation options.
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
		normalErrorCount := findingsFrom(ValidateDSL(model, "", DefaultEngineOptions())).Count()

		// Skip semantic validation
		options := &EngineOptions{
			SkipSemanticValidation: true,
		}
		skippedErrorCount := findingsFrom(ValidateDSL(model, "", options)).Count()

		// Should have fewer errors when semantic validation is skipped
		assert.LessOrEqual(t, skippedErrorCount, normalErrorCount, "Skipping semantic validation should reduce or maintain error count")
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

		// Skipping complex-operation validation drops findings, never adds them, and
		// leaves the surrounding phases running.
		normalErrorCount := findingsFrom(ValidateDSL(model, "", DefaultEngineOptions())).Count()
		skippedErrorCount := findingsFrom(ValidateDSL(model, "", options)).Count()

		assert.LessOrEqual(t, skippedErrorCount, normalErrorCount,
			"Skipping complex operation validation should reduce or maintain error count")
	})
}

// TestValidationReport tests the comprehensive validation report functionality.
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
			assert.Positive(t, summary.TotalErrors)

			// Test GetErrorsByType functionality
			for errorType := range summary.ErrorsByType {
				errorsOfType := report.GetErrorsByType(errorType)
				assert.NotEmpty(t, errorsOfType, "Should find errors of type %s", errorType)
			}
		}
	})
}

// TestValidationEngine_RealWorldScenarios tests realistic authorization model scenarios.
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
										Tupleset:        &openfgav1.ObjectRelation{Relation: "owner"},
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
										Tupleset:        &openfgav1.ObjectRelation{Relation: "owner"},
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

		findings := findingsFrom(ValidateDSL(model, dslContent, DefaultEngineOptions()))

		// This complex model should pass validation
		if findings.Count() > 0 {
			t.Logf("Validation errors found: %d", findings.Count())
			for _, err := range findings.GetErrors() {
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

// TestValidationEngine_PerformanceBasics tests basic performance characteristics.
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
		findings := findingsFrom(ValidateDSL(model, "", DefaultEngineOptions()))

		// Should complete validation in reasonable time
		t.Logf("Large model validation completed with %d errors", findings.Count())

		// Test JSON validation performance
		jsonFindings := findingsFrom(ValidateJSON(model, DefaultEngineOptions()))
		t.Logf("Large model JSON validation completed with %d errors", jsonFindings.Count())
	})
}

// TestEntryPointsReportFindingsThroughTheError pins what the four entry points
// return: nil for a valid model, and otherwise an error carrying every finding with
// its sentinel and its scope still reachable, which is what findingsFrom relies on.
func TestEntryPointsReportFindingsThroughTheError(t *testing.T) {
	t.Parallel()

	const dsl = `model
  schema 1.1
type user
type document
  relations
    define viewer: [user, group]
`

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	validationErr := ValidateDSL(model, dsl, DefaultEngineOptions())
	require.Error(t, validationErr, "group is not defined, so this model must not validate")

	// errors.Is reaches each finding's sentinel through Unwrap() []error.
	require.ErrorIs(t, validationErr, fgaerrors.ErrInvalidType)

	// errors.As reaches the scope by the same path.
	var scoped *fgaerrors.ErrObjectType
	require.ErrorAs(t, validationErr, &scoped)
	assert.Equal(t, "group", scoped.ObjectType)

	// errors.As also recovers the collection itself, which is how a caller lists every
	// finding rather than the first one errors.As stops at.
	var collection *ValidationErrors
	require.ErrorAs(t, validationErr, &collection)
	assert.NotEmpty(t, collection.AllFindings())
}

// TestFindingOrderIsDeterministic checks that validating one model twice reports its
// findings in the same order.
//
// Relations and conditions reach every validation phase in a proto map, which has no
// order of its own. Ranging one directly ordered the findings by whatever the runtime
// handed back, so the same model produced four different orders across runs, and a
// caller printing the list or comparing it against a fixture saw it change under them.
func TestFindingOrderIsDeterministic(t *testing.T) {
	t.Parallel()

	// The relation names sort in a different order than the symbols they report, so a
	// run that happened to sort by message would not pass this.
	const dsl = `model
  schema 1.1
type user
type document
  relations
    define alpha: missing_a
    define beta: missing_b
    define gamma: missing_c
    define delta: missing_d
`

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	want := []string{
		"the relation `missing_a` does not exist.",
		"the relation `missing_b` does not exist.",
		"the relation `missing_d` does not exist.",
		"the relation `missing_c` does not exist.",
	}

	// Map iteration is randomized per range, so one passing run proves nothing; four
	// keys admit four orders, which 100 runs would not agree on by chance.
	for i := 0; i < 100; i++ {
		messages := make([]string, 0, len(want))
		for _, finding := range findingsFrom(ValidateDSL(model, dsl, nil)).AllFindings() {
			messages = append(messages, finding.Message)
		}

		require.Equal(t, want, messages, "run %d reported the findings in a different order", i)
	}
}

// TestValidateJSONDiffersFromValidateDSLOnlyInPosition checks the two entry points
// take the same parsed proto and report the same findings, and that position is the
// only difference: ValidateDSL resolves line and column from the source text, and
// ValidateJSON leaves both nil.
func TestValidateJSONDiffersFromValidateDSLOnlyInPosition(t *testing.T) {
	t.Parallel()

	const dsl = `model
  schema 1.1
type user
type document
  relations
    define viewer: [user, group]
`

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	fromDSL := findingsFrom(ValidateDSL(model, dsl, nil)).AllFindings()
	fromJSON := findingsFrom(ValidateJSON(model, nil)).AllFindings()

	require.NotEmpty(t, fromDSL, "the model must produce a finding, or this compares two empty lists")
	require.Len(t, fromJSON, len(fromDSL), "the two entry points must find the same problems")

	for i := range fromDSL {
		assert.Equal(t, fromDSL[i].Message, fromJSON[i].Message)
		assert.Equal(t, fromDSL[i].Severity, fromJSON[i].Severity)
		assert.Equal(t, fromDSL[i].Category, fromJSON[i].Category)
		assert.Equal(t, fromDSL[i].Metadata, fromJSON[i].Metadata)

		assert.NotNil(t, fromDSL[i].Line, "ValidateDSL has the source text, so it must resolve the line")
		assert.NotNil(t, fromDSL[i].Column)
		assert.Nil(t, fromJSON[i].Line, "ValidateJSON has no source text to resolve a line against")
		assert.Nil(t, fromJSON[i].Column)
	}
}
