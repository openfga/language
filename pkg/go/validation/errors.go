package validation

import (
	"fmt"
	"strings"
)

// ValidationErrorType represents the different types of validation errors.
type ValidationErrorType string

const (
	SchemaVersionRequired     ValidationErrorType = "schema-version-required"
	SchemaVersionUnsupported  ValidationErrorType = "schema-version-unsupported"
	ReservedTypeKeywords      ValidationErrorType = "reserved-type-keywords"
	ReservedRelationKeywords  ValidationErrorType = "reserved-relation-keywords"
	SelfError                 ValidationErrorType = "self-error"
	InvalidName               ValidationErrorType = "invalid-name"
	MissingDefinition         ValidationErrorType = "missing-definition"
	InvalidRelationType       ValidationErrorType = "invalid-relation-type"
	InvalidRelationOnTupleset ValidationErrorType = "invalid-relation-on-tupleset"
	InvalidType               ValidationErrorType = "invalid-type"
	RelationNoEntrypoint      ValidationErrorType = "relation-no-entry-point"
	TuplesetNotDirect         ValidationErrorType = "tupleuserset-not-direct"
	DuplicatedError           ValidationErrorType = "duplicated-error"
	// Undefined reference errors.
	UndefinedType     ValidationErrorType = "undefined-type"
	UndefinedRelation ValidationErrorType = "undefined-relation"

	// Cycle and entry point errors.
	CyclicError ValidationErrorType = "cyclic-error"

	// Wildcard validation errors.
	InvalidWildcardError                         ValidationErrorType = "invalid-wildcard-error"
	AssignableRelationsMustHaveType              ValidationErrorType = "assignable-relation-must-have-type"
	InvalidSchema                                ValidationErrorType = "invalid-schema"
	InvalidSyntax                                ValidationErrorType = "invalid-syntax"
	TypeRestrictionCannotHaveWildcardAndRelation ValidationErrorType = "type-wildcard-relation"
	ConditionNotDefined                          ValidationErrorType = "condition-not-defined"
	ConditionNotUsed                             ValidationErrorType = "condition-not-used"
	DifferentNestedConditionName                 ValidationErrorType = "different-nested-condition-name"
	MultipleModulesInFile                        ValidationErrorType = "multiple-modules-in-file"
	CyclicRelation                               ValidationErrorType = "cyclic-relation"
	InvalidSchemaVersion                         ValidationErrorType = "invalid-schema-version"
)

// LineRange represents line start and end positions.
type LineRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// ColumnRange represents column start and end positions.
type ColumnRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// ErrorMetadata contains metadata about the validation error.
type ErrorMetadata struct {
	Symbol        string              `json:"symbol"`
	ErrorType     ValidationErrorType `json:"errorType"`
	Module        string              `json:"module,omitempty"`
	Type          string              `json:"type,omitempty"`
	Relation      string              `json:"relation,omitempty"`
	Condition     string              `json:"condition,omitempty"`
	OffendingType string              `json:"offendingType,omitempty"`
}

// ValidationError represents a single validation error.
type ValidationError struct {
	Message  string         `json:"msg"`
	Line     *LineRange     `json:"line,omitempty"`
	Column   *ColumnRange   `json:"column,omitempty"`
	File     string         `json:"file,omitempty"`
	Metadata *ErrorMetadata `json:"metadata,omitempty"`
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	location := ""
	if e.Line != nil && e.Column != nil {
		location = fmt.Sprintf(" at line=%d, column=%d", e.Line.Start, e.Column.Start)
	}
	return fmt.Sprintf("validation error%s: %s", location, e.Message)
}

// String returns a string representation of the error.
func (e *ValidationError) String() string {
	return e.Error()
}

// ValidationErrors represents a collection of validation errors.
type ValidationErrors struct {
	Errors []*ValidationError `json:"errors"`
}

// Error implements the error interface for ValidationErrors.
func (e *ValidationErrors) Error() string {
	if len(e.Errors) == 0 {
		return "no validation errors"
	}

	plural := ""
	if len(e.Errors) > 1 {
		plural = "s"
	}

	var errorStrings []string
	for _, err := range e.Errors {
		errorStrings = append(errorStrings, err.String())
	}

	return fmt.Sprintf("%d error%s occurred:\n\t* %s\n\n",
		len(e.Errors), plural, strings.Join(errorStrings, "\n\t* "))
}

// Add adds a validation error to the collection.
func (e *ValidationErrors) Add(err *ValidationError) {
	e.Errors = append(e.Errors, err)
}

// GetErrors returns a slice of all validation errors.
func (ve *ValidationErrors) GetErrors() []*ValidationError {
	return ve.Errors
}

// NewValidationErrors creates a new ValidationErrors instance from a slice of ValidationError
func NewValidationErrors(errors []*ValidationError) *ValidationErrors {
	if errors == nil {
		errors = make([]*ValidationError, 0)
	}
	return &ValidationErrors{
		Errors: errors,
	}
}

// HasErrors returns true if there are any errors.
func (e *ValidationErrors) HasErrors() bool {
	return len(e.Errors) > 0
}

// Count returns the number of errors.
func (e *ValidationErrors) Count() int {
	return len(e.Errors)
}

// Meta represents file and module metadata.
type Meta struct {
	File   string `json:"file,omitempty"`
	Module string `json:"module,omitempty"`
}

// ErrorCustomResolver is a function type for custom error position resolution.
type ErrorCustomResolver func(wordIndex int, rawLine string, symbol string) int
