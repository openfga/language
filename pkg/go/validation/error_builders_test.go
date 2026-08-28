package validation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

func TestWordIndex(t *testing.T) {
	tests := []struct {
		name    string
		rawLine string
		symbol  string
		want    int
	}{
		{"empty symbol returns 0", "define viewer: [user]", "", 0},
		{"not found returns 0", "define viewer: [user]", "missing", 0},
		{"word-boundary match", "define viewer: [user]", "user", 16},
		{"prefers boundary over earlier substring", "define ownerx: owner", "owner", 15},
		{"falls back to substring when no boundary", "type usergroup", "user", 5},
		{"non-word symbol falls back to substring", "define x: [user:*]", "user:*", 11},
		{"first occurrence wins on boundary", "a or a", "a", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, wordIndex(tt.rawLine, tt.symbol))
		})
	}
}

func TestValidationErrors_AllFindings(t *testing.T) {
	errs := NewValidationErrors(nil)

	// Initially no errors
	assert.Empty(t, errs.AllFindings())

	// Add an error
	errs.Add(newInvalidNameError(nil, "test", "rule", nil, nil, nil))

	findings := errs.AllFindings()
	assert.Len(t, findings, 1)
	assert.Contains(t, findings[0].Message, "test")
}

func TestValidationErrors_HasErrorsAfterAdd(t *testing.T) {
	errs := NewValidationErrors(nil)

	assert.False(t, errs.HasErrors())

	errs.Add(newInvalidNameError(nil, "test", "rule", nil, nil, nil))

	assert.True(t, errs.HasErrors())
}

func TestValidationErrors_CountAfterAdd(t *testing.T) {
	errs := NewValidationErrors(nil)

	assert.Equal(t, 0, errs.Count())

	errs.Add(newInvalidNameError(nil, "test1", "rule", nil, nil, nil))
	assert.Equal(t, 1, errs.Count())

	errs.Add(newInvalidNameError(nil, "test2", "rule", nil, nil, nil))
	assert.Equal(t, 2, errs.Count())
}

func TestNewInvalidNameError(t *testing.T) {
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
			typeName:     ptrString("document"),
			expectedMsg:  "relation 'invalid-relation' of type 'document' does not match naming rule: '[a-zA-Z]+'.",
			expectedType: InvalidName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := NewValidationErrors(nil)
			errs.Add(newInvalidNameError(nil, tt.symbol, tt.clause, tt.typeName, tt.lineIndex, tt.meta))

			findings := errs.AllFindings()
			assert.Len(t, findings, 1)
			assert.Equal(t, tt.expectedMsg, findings[0].Message)
			assert.Equal(t, tt.expectedType, findings[0].Metadata.ErrorType)
			assert.Equal(t, tt.symbol, findings[0].Metadata.Symbol)
		})
	}
}

func TestNewInvalidConditionNameError(t *testing.T) {
	lineIndex := 5
	meta := &Meta{File: "test.fga", Module: "test"}

	err := newInvalidConditionNameError(nil, "bad name", "[a-zA-Z]+", &lineIndex, meta)

	assert.Equal(t, "condition 'bad name' does not match naming rule: '[a-zA-Z]+'.", err.Message)
	assert.Equal(t, InvalidName, err.Metadata.ErrorType)
	assert.Equal(t, "bad name", err.Metadata.Symbol)
	assert.Equal(t, fgaerrors.ErrorKindCondition, err.Category)
	assert.Equal(t, "bad name", err.Metadata.Condition)
	assert.Empty(t, err.Metadata.Type)

	var scoped *fgaerrors.ErrCondition
	require.ErrorAs(t, err, &scoped)
	assert.Equal(t, "bad name", scoped.Condition)
}

func TestNewReservedTypeNameError(t *testing.T) {
	lineIndex := 5
	meta := &Meta{File: "test.fga", Module: "test"}

	err := newReservedTypeNameError(nil, "self", &lineIndex, meta)

	assert.Equal(t, "a type cannot be named 'self' or 'this'.", err.Message)
	assert.Equal(t, ReservedTypeKeywords, err.Metadata.ErrorType)
	assert.Equal(t, "self", err.Metadata.Symbol)
	assert.Equal(t, "test.fga", err.File)
}

func TestNewReservedRelationNameError(t *testing.T) {
	lineIndex := 3
	meta := &Meta{File: "test.fga", Module: "test"}

	err := newReservedRelationNameError(nil, "this", "document", &lineIndex, meta)

	assert.Equal(t, "a relation cannot be named 'self' or 'this'.", err.Message)
	assert.Equal(t, ReservedRelationKeywords, err.Metadata.ErrorType)
	assert.Equal(t, "this", err.Metadata.Symbol)
	assert.Equal(t, "document", err.Metadata.Type)
}

func TestNewTupleUsersetRequiresDirectError(t *testing.T) {
	lines := []string{
		"type document",
		"  relations",
		"    define viewer: user from parent",
		"    define admin: [user]",
	}
	lineIndex := 2
	meta := &Meta{File: "test.fga"}

	err := newTupleUsersetRequiresDirectError(lines, "user", "document", "viewer", meta, &lineIndex)

	assert.Equal(t, "`user` relation used inside from allows only direct relation.", err.Message)
	assert.Equal(t, TuplesetNotDirect, err.Metadata.ErrorType)
	assert.Equal(t, "user", err.Metadata.Symbol)
}

func TestNewDuplicateTypeNameError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 10

	err := newDuplicateTypeNameError(nil, "document", meta, &lineIndex)

	assert.Equal(t, "the type `document` is a duplicate.", err.Message)
	assert.Equal(t, DuplicatedError, err.Metadata.ErrorType)
	assert.Equal(t, "document", err.Metadata.Symbol)
}

func TestNewDuplicateTypeRestrictionError(t *testing.T) {
	meta := &Meta{File: "test.fga"}
	lineIndex := 5

	err := newDuplicateTypeRestrictionError(nil, "user", "viewer", "document", meta, &lineIndex)

	assert.Equal(t, "the type restriction `user` is a duplicate in the relation `viewer`.", err.Message)
	assert.Equal(t, DuplicatedError, err.Metadata.ErrorType)
	assert.Equal(t, "user", err.Metadata.Symbol)
}

func TestNewNoEntryPointLoopError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 8

	err := newNoEntryPointLoopError(nil, "viewer", "document", meta, &lineIndex)

	assert.Equal(t, "`viewer` is an impossible relation for `document` (potential loop).", err.Message)
	assert.Equal(t, RelationNoEntrypoint, err.Metadata.ErrorType)
	assert.Equal(t, "viewer", err.Metadata.Symbol)
}

func TestNewNoEntryPointError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 12

	err := newNoEntryPointError(nil, "viewer", "document", meta, &lineIndex)

	assert.Equal(t, "`viewer` is an impossible relation for `document` (no entrypoint).", err.Message)
	assert.Equal(t, RelationNoEntrypoint, err.Metadata.ErrorType)
	assert.Equal(t, "viewer", err.Metadata.Symbol)
}

func TestNewInvalidTypeError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 3

	err := newInvalidTypeError(nil, "unknown_type", meta, &lineIndex)

	assert.Equal(t, "`unknown_type` is not a valid type.", err.Message)
	assert.Equal(t, InvalidType, err.Metadata.ErrorType)
	assert.Equal(t, "unknown_type", err.Metadata.Symbol)
}

func TestNewAssignableRelationMustHaveTypesError(t *testing.T) {
	lineIndex := 6

	err := newAssignableRelationMustHaveTypesError(nil, "viewer", &lineIndex)

	assert.Equal(t, "the assignable relation 'viewer' must have at least one assignable type.", err.Message)
	assert.Equal(t, AssignableRelationsMustHaveType, err.Metadata.ErrorType)
	assert.Equal(t, "viewer", err.Metadata.Symbol)
}

func TestNewInvalidRelationError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 4

	err := newInvalidRelationError(nil, "unknown", "document", "relation", &lineIndex, meta)

	assert.Equal(t, "the relation `unknown` does not exist.", err.Message)
	assert.Equal(t, MissingDefinition, err.Metadata.ErrorType)
	assert.Equal(t, "unknown", err.Metadata.Symbol)
}

func TestNewSchemaVersionRequiredError(t *testing.T) {
	lineIndex := 0

	err := newSchemaVersionRequiredError(nil, &lineIndex)

	assert.Equal(t, "schema version required", err.Message)
	assert.Equal(t, SchemaVersionRequired, err.Metadata.ErrorType)
}

func TestNewInvalidSchemaVersionError(t *testing.T) {
	lineIndex := 1

	err := newInvalidSchemaVersionError(nil, "2.0", &lineIndex)

	assert.Equal(t, "invalid schema 2.0", err.Message)
	assert.Equal(t, InvalidSchema, err.Metadata.ErrorType)
	assert.Equal(t, "2.0", err.Metadata.Symbol)
}

func TestNewSchemaVersionUnsupportedError(t *testing.T) {
	lineIndex := 1

	err := newSchemaVersionUnsupportedError(nil, "1.0", &lineIndex)

	assert.Equal(t, "schema version no longer supported", err.Message)
	assert.Equal(t, SchemaVersionUnsupported, err.Metadata.ErrorType)
	assert.Equal(t, "1.0", err.Metadata.Symbol)
}

func TestNewUnusedConditionError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 15

	err := newUnusedConditionError(nil, "unused_condition", meta, &lineIndex)

	assert.Equal(t, "`unused_condition` condition is not used in the model.", err.Message)
	assert.Equal(t, ConditionNotUsed, err.Metadata.ErrorType)
	assert.Equal(t, "unused_condition", err.Metadata.Symbol)
}

func TestNewDifferentNestedConditionNameError(t *testing.T) {
	err := newDifferentNestedConditionNameError("condition1", "condition2")

	assert.Equal(t, "condition key is `condition1` but nested name property is condition2", err.Message)
	assert.Equal(t, DifferentNestedConditionName, err.Metadata.ErrorType)
	assert.Equal(t, "condition2", err.Metadata.Symbol)
}

func TestNewMultipleModulesInSingleFileError(t *testing.T) {
	modules := []string{"module1", "module2", "module3"}

	err := newMultipleModulesInSingleFileError("test.fga", modules)

	assert.Equal(t, "file test.fga would contain multiple module definitions (module1, module2, module3) "+
		"when transforming to DSL. Only one module can be defined per file.", err.Message)
	assert.Equal(t, MultipleModulesInFile, err.Metadata.ErrorType)
	assert.Equal(t, "test.fga", err.Metadata.Symbol)
}

func TestLineAndColumnResolution(t *testing.T) {
	lines := []string{
		"model",
		"  schema 1.1",
		"type document",
		"  relations",
		"    define viewer: [user]",
	}
	lineIndex := 4

	err := newInvalidNameError(lines, "viewer", "rule", nil, &lineIndex, nil)

	// Check line information
	assert.NotNil(t, err.Line)
	assert.Equal(t, 4, err.Line.Start)
	assert.Equal(t, 4, err.Line.End)

	// Check column information (should find "viewer" in the line)
	assert.NotNil(t, err.Column)
	line := lines[4]
	expectedStart := strings.Index(line, "viewer")
	assert.Equal(t, expectedStart, err.Column.Start)
	assert.Equal(t, expectedStart+len("viewer"), err.Column.End)
}

func TestCustomResolver(t *testing.T) {
	lines := []string{
		"type document",
		"  relations",
		"    define viewer: user from parent",
	}
	lineIndex := 2
	meta := &Meta{File: "test.fga"}

	err := newTupleUsersetRequiresDirectError(lines, "user", "document", "viewer", meta, &lineIndex)

	// The custom resolver should position the error after the "from" keyword
	assert.NotNil(t, err.Column)
	line := lines[2]
	fromIndex := strings.Index(line, "from")
	expectedStart := fromIndex + len("from") + strings.Index(line[fromIndex+len("from"):], "user")
	assert.Equal(t, expectedStart, err.Column.Start)
}

func TestNewUndefinedRelationError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 4

	err := newUndefinedRelationError(nil, "viewer", "document", "can_view", "folder", meta, &lineIndex)

	assert.Equal(t, "Relation 'viewer' is not defined on type 'document' (referenced in relation 'can_view' of type 'folder')", err.Message)
	assert.Equal(t, UndefinedRelation, err.Metadata.ErrorType)
	assert.Equal(t, "viewer", err.Metadata.Symbol)
	assert.Equal(t, "document", err.Metadata.Type)
	assert.Equal(t, "viewer", err.Metadata.Relation)
}

func TestNewDuplicateRelationshipDefinitionError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 7

	err := newDuplicateRelationshipDefinitionError(nil, "viewer", meta, &lineIndex)

	assert.Equal(t, "the relation 'viewer' is defined more than once.", err.Message)
	assert.Equal(t, DuplicatedError, err.Metadata.ErrorType)
	assert.Equal(t, "viewer", err.Metadata.Symbol)
	assert.Equal(t, "viewer", err.Metadata.Relation)
}

func TestNewAssignableTypeWildcardRelationError(t *testing.T) {
	meta := &Meta{File: "test.fga", Module: "test"}
	lineIndex := 9

	err := newAssignableTypeWildcardRelationError(nil, "user", "document", "viewer", meta, &lineIndex)

	assert.Equal(t, "the type restriction 'user' on relation 'viewer' of type 'document' is not allowed to have both a wildcard and a relation.", err.Message)
	assert.Equal(t, TypeRestrictionCannotHaveWildcardAndRelation, err.Metadata.ErrorType)
	assert.Equal(t, "user", err.Metadata.Symbol)
	assert.Equal(t, "document", err.Metadata.Type)
	assert.Equal(t, "viewer", err.Metadata.Relation)
}

func TestNewMaximumOneDirectRelationshipError(t *testing.T) {
	lineIndex := 11

	err := newMaximumOneDirectRelationshipError(nil, "viewer", &lineIndex)

	assert.Equal(t, "the relation 'viewer' can have at most one direct relationship.", err.Message)
	assert.Equal(t, DuplicatedError, err.Metadata.ErrorType)
	assert.Equal(t, "viewer", err.Metadata.Symbol)
	assert.Equal(t, "viewer", err.Metadata.Relation)
}
