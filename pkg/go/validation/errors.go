package validation

import (
	"fmt"
	"slices"
	"strings"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
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

	// Weighted graph errors.
	GraphModelUnbuildable ValidationErrorType = "graph-model-unbuildable"
)

// Range is a start and end position in the source text, used for both the line and
// the column a finding is at.
//
// The two are indexed differently: a line Range repeats the same zero-based index
// in Start and End, while a column Range is half-open, End being one past the
// symbol's last character.
type Range struct {
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
	Message string `json:"msg"`

	// Severity states whether this finding makes the model invalid. Findings that
	// do not block are reported without failing validation.
	Severity fgaerrors.Severity `json:"severity,omitempty"`

	// Category is the part of the model this finding is about.
	Category fgaerrors.ModelErrorKind `json:"category,omitempty"`

	Line     *Range         `json:"line,omitempty"`
	Column   *Range         `json:"column,omitempty"`
	File     string         `json:"file,omitempty"`
	Metadata *ErrorMetadata `json:"metadata,omitempty"`

	// Cause is the scoped error this finding wraps, and what Unwrap returns:
	// errors.Is identifies the condition, errors.As the part of the model.
	//
	// It stays off the wire because an error field has no concrete type to decode
	// into, which would leave ValidationError unable to round-trip. The message,
	// severity and metadata carry the same information in JSON.
	Cause error `json:"-"`
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	location := ""
	if e.Line != nil && e.Column != nil {
		location = fmt.Sprintf(" at line=%d, column=%d", e.Line.Start, e.Column.Start)
	}
	return fmt.Sprintf("validation error%s: %s", location, e.Message)
}

// Unwrap returns Cause, which is nil for an error built directly rather than
// through the collector.
func (e *ValidationError) Unwrap() error {
	return e.Cause
}

// Blocks reports whether this finding makes the model invalid. A directly-constructed
// error has no severity set and blocks; see Severity.Blocks. A nil finding is not a
// finding, so it blocks nothing.
func (e *ValidationError) Blocks() bool {
	if e == nil {
		return false
	}

	return e.Severity.Blocks()
}

// String returns a string representation of the error.
func (e *ValidationError) String() string {
	return e.Error()
}

// ValidationErrors represents a collection of validation errors.
//
//nolint:errname // plural name intentionally describes a collection of errors
type ValidationErrors struct {
	// Errors holds every finding in the order it was raised, blocking or not.
	// len(Errors) is CountAll, not Count; HasErrors and GetErrors are the
	// blocking-only views.
	Errors []*ValidationError `json:"errors"`
}

// findings is the slice every read method below goes through, so a nil collection is
// an empty one in one place rather than in nine. A nil *ValidationErrors reaches these
// methods through a zero ValidationReport, among other paths.
//
// A nil *ValidationError is dropped, because it is not a finding: counting one would
// have CountAll and HasFindings disagree with Blocks and Unwrap, and would put an
// entry in the slice AllFindings hands out that dereferences nil on Severity or
// String. The collector never appends one; a collection built through
// NewValidationErrors, Add or the exported Errors field can hold one.
//
// The scan returns the slice untouched when there is nothing to drop, so the usual
// case does not allocate.
func (e *ValidationErrors) findings() []*ValidationError {
	if e == nil {
		return nil
	}

	if !slices.Contains(e.Errors, nil) {
		return e.Errors
	}

	held := make([]*ValidationError, 0, len(e.Errors))

	for _, err := range e.Errors {
		if err != nil {
			held = append(held, err)
		}
	}

	return held
}

// Error implements the error interface for ValidationErrors.
//
// It reports the blocking findings only, so the count in the message agrees with
// Count.
func (e *ValidationErrors) Error() string {
	blocking := e.GetErrors()
	if len(blocking) == 0 {
		return "no validation errors"
	}

	plural := ""
	if len(blocking) > 1 {
		plural = "s"
	}

	var errorStrings []string
	for _, err := range blocking {
		errorStrings = append(errorStrings, err.String())
	}

	return fmt.Sprintf("%d error%s occurred:\n\t* %s\n\n",
		len(blocking), plural, strings.Join(errorStrings, "\n\t* "))
}

// Unwrap returns every finding, so errors.Is and errors.As reach each sentinel and
// scope through the collection. Non-blocking findings are included, since errors.Is
// asks whether a condition was reported, not whether it blocks.
//
// Because errors.As stops at the first match, enumerating every finding of one
// scope means walking AllFindings.
func (e *ValidationErrors) Unwrap() []error {
	held := e.findings()
	if len(held) == 0 {
		return nil
	}

	// findings drops nil entries, so none reaches errors.Is here: handing it a nil
	// *ValidationError as a non-nil error panics.
	unwrapped := make([]error, 0, len(held))
	for _, err := range held {
		unwrapped = append(unwrapped, err)
	}

	return unwrapped
}

// ErrorOrNil returns e as an error, or nil when no finding blocks.
//
// The validation entry points return this, so err != nil means the model is invalid
// rather than that something was reported: a model whose only findings are warnings
// or advisories yields nil. Non-blocking findings alongside a blocking one stay
// reachable through errors.As and AllFindings.
//
// Findings from a model that stays valid are only reachable off this path:
// CreateValidationReport returns the collection itself, and AllFindings on it lists
// everything raised.
func (e *ValidationErrors) ErrorOrNil() error {
	if !e.HasErrors() {
		return nil
	}

	return e
}

// Add adds a validation error to the collection.
func (e *ValidationErrors) Add(err *ValidationError) {
	e.Errors = append(e.Errors, err)
}

// GetErrors returns the findings that make the model invalid.
//
// Non-blocking findings are excluded; AllFindings returns everything.
func (e *ValidationErrors) GetErrors() []*ValidationError {
	held := e.findings()

	blocking := make([]*ValidationError, 0, len(held))
	for _, err := range held {
		if err.Blocks() {
			blocking = append(blocking, err)
		}
	}
	return blocking
}

// AllFindings returns every finding, blocking or not, in the order raised. Each
// finding's Severity says how to present it.
func (e *ValidationErrors) AllFindings() []*ValidationError {
	return e.findings()
}

// NewValidationErrors creates a new ValidationErrors instance from a slice of ValidationError.
func NewValidationErrors(errors []*ValidationError) *ValidationErrors {
	if errors == nil {
		errors = make([]*ValidationError, 0)
	}
	return &ValidationErrors{
		Errors: errors,
	}
}

// HasErrors reports whether any finding makes the model invalid. A model with only
// warnings or advisories is valid, so this is false; HasFindings covers everything
// raised.
func (e *ValidationErrors) HasErrors() bool {
	for _, err := range e.findings() {
		if err.Blocks() {
			return true
		}
	}
	return false
}

// HasFindings reports whether anything at all was reported, blocking or not.
func (e *ValidationErrors) HasFindings() bool {
	return len(e.findings()) > 0
}

// Count returns the number of findings that make the model invalid.
func (e *ValidationErrors) Count() int {
	count := 0
	for _, err := range e.findings() {
		if err.Blocks() {
			count++
		}
	}
	return count
}

// CountAll returns the total number of findings, blocking or not.
func (e *ValidationErrors) CountAll() int {
	return len(e.findings())
}

// Meta represents file and module metadata.
type Meta struct {
	File   string `json:"file,omitempty"`
	Module string `json:"module,omitempty"`
}

// ErrorCustomResolver is a function type for custom error position resolution.
type ErrorCustomResolver func(wordIndex int, rawLine string, symbol string) int
