package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		error    *ValidationError
		expected string
	}{
		{
			name: "error with line and column",
			error: &ValidationError{
				Message: "test error message",
				Line:    &LineRange{Start: 5, End: 5},
				Column:  &ColumnRange{Start: 10, End: 15},
				Metadata: &ErrorMetadata{
					Symbol:    "test_symbol",
					ErrorType: InvalidName,
				},
			},
			expected: "validation error at line=5, column=10: test error message",
		},
		{
			name: "error without line/column",
			error: &ValidationError{
				Message: "test error message",
				Metadata: &ErrorMetadata{
					Symbol:    "test_symbol",
					ErrorType: InvalidType,
				},
			},
			expected: "validation error: test error message",
		},
		{
			name: "error with file",
			error: &ValidationError{
				Message: "test error message",
				File:    "test.fga",
				Line:    &LineRange{Start: 3, End: 3},
				Column:  &ColumnRange{Start: 0, End: 4},
			},
			expected: "validation error at line=3, column=0: test error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.error.Error()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestValidationError_String(t *testing.T) {
	error := &ValidationError{
		Message: "test error",
		Line:    &LineRange{Start: 1, End: 1},
		Column:  &ColumnRange{Start: 5, End: 10},
	}

	// String() should return the same as Error()
	assert.Equal(t, error.Error(), error.String())
}

func TestValidationErrors_Error(t *testing.T) {
	tests := []struct {
		name     string
		errors   *ValidationErrors
		expected string
	}{
		{
			name:     "no errors",
			errors:   &ValidationErrors{Errors: []*ValidationError{}},
			expected: "no validation errors",
		},
		{
			name: "single error",
			errors: &ValidationErrors{
				Errors: []*ValidationError{
					{Message: "first error"},
				},
			},
			expected: "1 error occurred:\n\t* validation error: first error\n\n",
		},
		{
			name: "multiple errors",
			errors: &ValidationErrors{
				Errors: []*ValidationError{
					{Message: "first error"},
					{Message: "second error"},
				},
			},
			expected: "2 errors occurred:\n\t* validation error: first error\n\t* validation error: second error\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.errors.Error()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestValidationErrors_Add(t *testing.T) {
	errors := &ValidationErrors{}

	// Initially no errors
	assert.False(t, errors.HasErrors())
	assert.Equal(t, 0, errors.Count())

	// Add first error
	error1 := &ValidationError{Message: "first error"}
	errors.Add(error1)

	assert.True(t, errors.HasErrors())
	assert.Equal(t, 1, errors.Count())
	assert.Equal(t, error1, errors.Errors[0])

	// Add second error
	error2 := &ValidationError{Message: "second error"}
	errors.Add(error2)

	assert.Equal(t, 2, errors.Count())
	assert.Equal(t, error2, errors.Errors[1])
}

func TestValidationErrors_HasErrors(t *testing.T) {
	errors := &ValidationErrors{}
	assert.False(t, errors.HasErrors())

	errors.Add(&ValidationError{Message: "test"})
	assert.True(t, errors.HasErrors())
}

func TestValidationErrors_Count(t *testing.T) {
	errors := &ValidationErrors{}
	assert.Equal(t, 0, errors.Count())

	errors.Add(&ValidationError{Message: "test1"})
	assert.Equal(t, 1, errors.Count())

	errors.Add(&ValidationError{Message: "test2"})
	assert.Equal(t, 2, errors.Count())
}

func TestErrorMetadata(t *testing.T) {
	metadata := &ErrorMetadata{
		Symbol:        "test_symbol",
		ErrorType:     InvalidName,
		Module:        "test_module",
		Type:          "test_type",
		Relation:      "test_relation",
		Condition:     "test_condition",
		OffendingType: "offending_type",
	}

	assert.Equal(t, "test_symbol", metadata.Symbol)
	assert.Equal(t, InvalidName, metadata.ErrorType)
	assert.Equal(t, "test_module", metadata.Module)
	assert.Equal(t, "test_type", metadata.Type)
	assert.Equal(t, "test_relation", metadata.Relation)
	assert.Equal(t, "test_condition", metadata.Condition)
	assert.Equal(t, "offending_type", metadata.OffendingType)
}

func TestLineRange(t *testing.T) {
	line := &LineRange{Start: 5, End: 10}
	assert.Equal(t, 5, line.Start)
	assert.Equal(t, 10, line.End)
}

func TestColumnRange(t *testing.T) {
	column := &ColumnRange{Start: 15, End: 25}
	assert.Equal(t, 15, column.Start)
	assert.Equal(t, 25, column.End)
}

func TestValidationErrorTypes(t *testing.T) {
	// Test that all validation error types are defined as expected
	errorTypes := []ValidationErrorType{
		SchemaVersionRequired,
		SchemaVersionUnsupported,
		ReservedTypeKeywords,
		ReservedRelationKeywords,
		SelfError,
		InvalidName,
		MissingDefinition,
		InvalidRelationType,
		InvalidRelationOnTupleset,
		InvalidType,
		RelationNoEntrypoint,
		TuplesetNotDirect,
		DuplicatedError,
		AssignableRelationsMustHaveType,
		InvalidSchema,
		InvalidSyntax,
		TypeRestrictionCannotHaveWildcardAndRelation,
		ConditionNotDefined,
		ConditionNotUsed,
		DifferentNestedConditionName,
		MultipleModulesInFile,
	}

	// Ensure all error types have string values
	for _, errorType := range errorTypes {
		assert.NotEmpty(t, string(errorType))
		assert.Contains(t, string(errorType), "-")
	}

	// Test specific error type values match JS implementation
	assert.Equal(t, "schema-version-required", string(SchemaVersionRequired))
	assert.Equal(t, "missing-definition", string(MissingDefinition))
	assert.Equal(t, "reserved-type-keywords", string(ReservedTypeKeywords))
	assert.Equal(t, "relation-no-entry-point", string(RelationNoEntrypoint))
}

func TestMeta(t *testing.T) {
	meta := &Meta{
		File:   "test.fga",
		Module: "test_module",
	}

	assert.Equal(t, "test.fga", meta.File)
	assert.Equal(t, "test_module", meta.Module)
}
