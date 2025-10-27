package validation

import (
	"strings"
	"testing"

	fgaSdk "github.com/openfga/go-sdk"
	"github.com/stretchr/testify/assert"
)

func TestNewErrorCollector(t *testing.T) {
	lines := []string{"line 1", "line 2", "line 3"}
	collector := NewErrorCollector(lines)

	assert.NotNil(t, collector)
	assert.Equal(t, lines, collector.lines)
	assert.Equal(t, 0, collector.Count())
	assert.False(t, collector.HasErrors())
}

func TestErrorCollector_GetErrors(t *testing.T) {
	collector := NewErrorCollector(nil)

	// Initially no errors
	errors := collector.GetErrors()
	assert.Empty(t, errors)

	// Add an error
	collector.RaiseInvalidName("test", "rule", nil, nil, nil)

	errors = collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Contains(t, errors[0].Message, "test")
}

func TestErrorCollector_HasErrors(t *testing.T) {
	collector := NewErrorCollector(nil)

	assert.False(t, collector.HasErrors())

	collector.RaiseInvalidName("test", "rule", nil, nil, nil)

	assert.True(t, collector.HasErrors())
}

func TestErrorCollector_Count(t *testing.T) {
	collector := NewErrorCollector(nil)

	assert.Equal(t, 0, collector.Count())

	collector.RaiseInvalidName("test1", "rule", nil, nil, nil)
	assert.Equal(t, 1, collector.Count())

	collector.RaiseInvalidName("test2", "rule", nil, nil, nil)
	assert.Equal(t, 2, collector.Count())
}

func TestErrorCollector_RaiseInvalidName(t *testing.T) {
	tests := []struct {
		name         string
		symbol       string
		clause       string
		typeName     *string
		lineIndex    *int
		meta         *Meta
		expectedMsg  string
		expectedType ValidationErrorType
	}{
		{
			name:         "type invalid name",
			symbol:       "invalid-type",
			clause:       "[a-zA-Z]+",
			typeName:     nil,
			expectedMsg:  "type 'invalid-type' does not match naming rule: '[a-zA-Z]+'.",
			expectedType: InvalidName,
		},
		{
			name:         "relation invalid name",
			symbol:       "invalid-relation",
			clause:       "[a-zA-Z]+",
			typeName:     fgaSdk.PtrString("document"),
			expectedMsg:  "relation 'invalid-relation' of type 'document' does not match naming rule: '[a-zA-Z]+'.",
			expectedType: InvalidName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewErrorCollector(nil)
			collector.RaiseInvalidName(tt.symbol, tt.clause, tt.typeName, tt.lineIndex, tt.meta)

			errors := collector.GetErrors()
			assert.Len(t, errors, 1)
			assert.Equal(t, tt.expectedMsg, errors[0].Message)
			assert.Equal(t, tt.expectedType, errors[0].Metadata.ErrorType)
			assert.Equal(t, tt.symbol, errors[0].Metadata.Symbol)
		})
	}
}

func TestErrorCollector_RaiseReservedTypeName(t *testing.T) {
	collector := NewErrorCollector(nil)
	lineIndex := 5
	meta := &Meta{File: "test.fga", Module: "test"}

	collector.RaiseReservedTypeName("self", &lineIndex, meta)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "a type cannot be named 'self' or 'this'.", errors[0].Message)
	assert.Equal(t, ReservedTypeKeywords, errors[0].Metadata.ErrorType)
	assert.Equal(t, "self", errors[0].Metadata.Symbol)
	assert.Equal(t, "test.fga", errors[0].File)
}

func TestErrorCollector_RaiseReservedRelationName(t *testing.T) {
	collector := NewErrorCollector(nil)
	lineIndex := 3
	meta := &Meta{File: "test.fga", Module: "test"}

	collector.RaiseReservedRelationName("this", &lineIndex, meta)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "a relation cannot be named 'self' or 'this'.", errors[0].Message)
	assert.Equal(t, ReservedRelationKeywords, errors[0].Metadata.ErrorType)
	assert.Equal(t, "this", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseTupleUsersetRequiresDirect(t *testing.T) {
	lines := []string{
		"type document",
		"  relations",
		"    define viewer: user from parent",
		"    define admin: [user]",
	}
	collector := NewErrorCollector(lines)
	lineIndex := 2
	meta := &Meta{File: "test.fga"}

	collector.RaiseTupleUsersetRequiresDirect("user", "document", "viewer", meta, &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "`user` relation used inside from allows only direct relation.", errors[0].Message)
	assert.Equal(t, TuplesetNotDirect, errors[0].Metadata.ErrorType)
	assert.Equal(t, "user", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseDuplicateTypeName(t *testing.T) {
	collector := NewErrorCollector(nil)
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 10

	collector.RaiseDuplicateTypeName("document", meta, &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "the type definition 'document' is defined more than once.", errors[0].Message)
	assert.Equal(t, DuplicatedError, errors[0].Metadata.ErrorType)
	assert.Equal(t, "document", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseDuplicateTypeRestriction(t *testing.T) {
	collector := NewErrorCollector(nil)
	meta := &Meta{File: "test.fga"}
	lineIndex := 5

	collector.RaiseDuplicateTypeRestriction("user", "viewer", "document", meta, &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "the type restriction 'user' in relation 'viewer' of type 'document' is defined more than once.", errors[0].Message)
	assert.Equal(t, DuplicatedError, errors[0].Metadata.ErrorType)
	assert.Equal(t, "user", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseNoEntryPointLoop(t *testing.T) {
	collector := NewErrorCollector(nil)
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 8

	collector.RaiseNoEntryPointLoop("viewer", "document", meta, &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "`viewer` is an impossible relation for `document` (potential loop).", errors[0].Message)
	assert.Equal(t, RelationNoEntrypoint, errors[0].Metadata.ErrorType)
	assert.Equal(t, "viewer", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseNoEntryPoint(t *testing.T) {
	collector := NewErrorCollector(nil)
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 12

	collector.RaiseNoEntryPoint("viewer", "document", meta, &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "`viewer` is an impossible relation for `document`.", errors[0].Message)
	assert.Equal(t, RelationNoEntrypoint, errors[0].Metadata.ErrorType)
	assert.Equal(t, "viewer", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseInvalidType(t *testing.T) {
	collector := NewErrorCollector(nil)
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 3

	collector.RaiseInvalidType("unknown_type", "document", "viewer", meta, &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "type 'unknown_type' is not defined.", errors[0].Message)
	assert.Equal(t, InvalidType, errors[0].Metadata.ErrorType)
	assert.Equal(t, "unknown_type", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseAssignableRelationMustHaveTypes(t *testing.T) {
	collector := NewErrorCollector(nil)
	lineIndex := 6

	collector.RaiseAssignableRelationMustHaveTypes("viewer", &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "the assignable relation 'viewer' must have at least one assignable type.", errors[0].Message)
	assert.Equal(t, AssignableRelationsMustHaveType, errors[0].Metadata.ErrorType)
	assert.Equal(t, "viewer", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseInvalidRelationError(t *testing.T) {
	collector := NewErrorCollector(nil)
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 4
	validRelations := []string{"admin", "viewer"}

	collector.RaiseInvalidRelationError("unknown", "document", "relation", validRelations, &lineIndex, meta)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "the relation `unknown` does not exist.", errors[0].Message)
	assert.Equal(t, MissingDefinition, errors[0].Metadata.ErrorType)
	assert.Equal(t, "unknown", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseSchemaVersionRequired(t *testing.T) {
	collector := NewErrorCollector(nil)
	lineIndex := 0

	collector.RaiseSchemaVersionRequired("", &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "a schema version is required in the model.", errors[0].Message)
	assert.Equal(t, SchemaVersionRequired, errors[0].Metadata.ErrorType)
}

func TestErrorCollector_RaiseInvalidSchemaVersion(t *testing.T) {
	collector := NewErrorCollector(nil)
	lineIndex := 1

	collector.RaiseInvalidSchemaVersion("2.0", &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "the schema version '2.0' is not supported.", errors[0].Message)
	assert.Equal(t, SchemaVersionUnsupported, errors[0].Metadata.ErrorType)
	assert.Equal(t, "2.0", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseUnusedCondition(t *testing.T) {
	collector := NewErrorCollector(nil)
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 15

	collector.RaiseUnusedCondition("unused_condition", meta, &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "condition 'unused_condition' is defined but not used.", errors[0].Message)
	assert.Equal(t, ConditionNotUsed, errors[0].Metadata.ErrorType)
	assert.Equal(t, "unused_condition", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseDifferentNestedConditionName(t *testing.T) {
	collector := NewErrorCollector(nil)

	collector.RaiseDifferentNestedConditionName("condition1", "condition2")

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "the 'condition1' condition has a different nested condition name ('condition2').", errors[0].Message)
	assert.Equal(t, DifferentNestedConditionName, errors[0].Metadata.ErrorType)
	assert.Equal(t, "condition1", errors[0].Metadata.Symbol)
}

func TestErrorCollector_RaiseMultipleModulesInSingleFile(t *testing.T) {
	collector := NewErrorCollector(nil)
	modules := []string{"module1", "module2", "module3"}

	collector.RaiseMultipleModulesInSingleFile("test.fga", modules)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)
	assert.Equal(t, "file 'test.fga' contains multiple modules: module1, module2, module3.", errors[0].Message)
	assert.Equal(t, MultipleModulesInFile, errors[0].Metadata.ErrorType)
	assert.Equal(t, "test.fga", errors[0].Metadata.Symbol)
}

func TestErrorCollector_LineAndColumnResolution(t *testing.T) {
	lines := []string{
		"model",
		"  schema 1.1",
		"type document",
		"  relations",
		"    define viewer: [user]",
	}
	collector := NewErrorCollector(lines)
	lineIndex := 4

	collector.RaiseInvalidName("viewer", "rule", nil, &lineIndex, nil)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)

	// Check line information
	assert.NotNil(t, errors[0].Line)
	assert.Equal(t, 4, errors[0].Line.Start)
	assert.Equal(t, 4, errors[0].Line.End)

	// Check column information (should find "viewer" in the line)
	assert.NotNil(t, errors[0].Column)
	line := lines[4]
	expectedStart := strings.Index(line, "viewer")
	assert.Equal(t, expectedStart, errors[0].Column.Start)
	assert.Equal(t, expectedStart+len("viewer"), errors[0].Column.End)
}

func TestErrorCollector_CustomResolver(t *testing.T) {
	lines := []string{
		"type document",
		"  relations",
		"    define viewer: user from parent",
	}
	collector := NewErrorCollector(lines)
	lineIndex := 2
	meta := &Meta{File: "test.fga"}

	collector.RaiseTupleUsersetRequiresDirect("user", "document", "viewer", meta, &lineIndex)

	errors := collector.GetErrors()
	assert.Len(t, errors, 1)

	// The custom resolver should position the error after "from" keyword
	assert.NotNil(t, errors[0].Column)
	line := lines[2]
	fromIndex := strings.Index(line, "from")
	expectedStart := fromIndex + len("from") + strings.Index(line[fromIndex+len("from"):], "user")
	assert.Equal(t, expectedStart, errors[0].Column.Start)
}
