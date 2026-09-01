// Package validation checks OpenFGA authorization models for semantic errors
// beyond what the DSL grammar enforces, mirroring the checks in pkg/js and
// pkg/java. The three implementations share the corpus under tests/data, which
// pins each finding's message, code, symbol and position.
package validation

import (
	"fmt"
	"strings"
)

// Kind is a finding's machine-readable code, the wire `errorType`. The string
// values are shared with pkg/js and pkg/java, so they are fixed.
type Kind string

const (
	SchemaVersionRequired     Kind = "schema-version-required"
	SchemaVersionUnsupported  Kind = "schema-version-unsupported"
	ReservedTypeKeywords      Kind = "reserved-type-keywords"
	ReservedRelationKeywords  Kind = "reserved-relation-keywords"
	SelfError                 Kind = "self-error"
	InvalidName               Kind = "invalid-name"
	MissingDefinition         Kind = "missing-definition"
	InvalidRelationType       Kind = "invalid-relation-type"
	InvalidRelationOnTupleset Kind = "invalid-relation-on-tupleset"
	InvalidType               Kind = "invalid-type"
	RelationNoEntrypoint      Kind = "relation-no-entry-point"
	TuplesetNotDirect         Kind = "tupleuserset-not-direct"
	DuplicatedError           Kind = "duplicated-error"
	UndefinedType             Kind = "undefined-type"
	UndefinedRelation         Kind = "undefined-relation"
	CyclicError               Kind = "cyclic-error"
	InvalidWildcardError      Kind = "invalid-wildcard-error"

	AssignableRelationsMustHaveType              Kind = "assignable-relation-must-have-type"
	InvalidSchema                                Kind = "invalid-schema"
	InvalidSyntax                                Kind = "invalid-syntax"
	TypeRestrictionCannotHaveWildcardAndRelation Kind = "type-wildcard-relation"
	ConditionNotDefined                          Kind = "condition-not-defined"
	ConditionNotUsed                             Kind = "condition-not-used"
	DifferentNestedConditionName                 Kind = "different-nested-condition-name"
	MultipleModulesInFile                        Kind = "multiple-modules-in-file"
	CyclicRelation                               Kind = "cyclic-relation"
	InvalidSchemaVersion                         Kind = "invalid-schema-version"
)

// Range is a start and end position in the source text, used for both the line
// and the column a finding is at.
//
// The two are indexed differently: a line Range repeats the same zero-based
// index in Start and End, while a column Range is half-open, End being one past
// the symbol's last character.
type Range struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// Metadata is the structured half of a finding: the offending symbol, the code,
// and which part of the model the finding is about. The field names and JSON
// shape match pkg/js and pkg/java.
type Metadata struct {
	Symbol        string `json:"symbol"`
	Kind          Kind   `json:"errorType"`
	Module        string `json:"module,omitempty"`
	Type          string `json:"type,omitempty"`
	Relation      string `json:"relation,omitempty"`
	Condition     string `json:"condition,omitempty"`
	OffendingType string `json:"offendingType,omitempty"`
}

// Finding is one validation diagnostic. Its fields are the cross-language wire
// shape, so it marshals without any custom JSON code.
//
// Line and Column are nil when the finding could not be located in the source,
// which is always the case for a model validated from JSON.
//
//nolint:errname // named for what it is, as go/scanner.Error is
type Finding struct {
	Message  string   `json:"msg"`
	Line     *Range   `json:"line,omitempty"`
	Column   *Range   `json:"column,omitempty"`
	File     string   `json:"file,omitempty"`
	Metadata Metadata `json:"metadata"`
}

// Error implements the error interface, so a single finding recovered with
// errors.As prints like one.
func (f *Finding) Error() string {
	if f.Line != nil && f.Column != nil {
		return fmt.Sprintf("validation error at line=%d, column=%d: %s", f.Line.Start, f.Column.Start, f.Message)
	}

	return "validation error: " + f.Message
}

// in records the file and module a finding was raised about, read off the
// model's source info by whichever loop raised it. Chainable; nil-safe.
func (f *Finding) in(file, module string) *Finding {
	if f == nil {
		return nil
	}
	f.File, f.Metadata.Module = file, module

	return f
}

// Findings is every diagnostic raised for one model, in the order raised. It
// follows go/scanner.ErrorList: the slice is the collection and, when
// non-empty, the error.
//
//nolint:errname // named for what it holds, as go/scanner.ErrorList is
type Findings []*Finding

// add appends f when it is a finding; a nil *Finding means nothing was found.
func (fs Findings) add(f *Finding) Findings {
	if f == nil {
		return fs
	}

	return append(fs, f)
}

// Error implements the error interface.
func (fs Findings) Error() string {
	if len(fs) == 0 {
		return "no validation errors"
	}

	plural := ""
	if len(fs) > 1 {
		plural = "s"
	}

	messages := make([]string, 0, len(fs))
	for _, f := range fs {
		messages = append(messages, f.Error())
	}

	return fmt.Sprintf("%d error%s occurred:\n\t* %s\n\n", len(fs), plural, strings.Join(messages, "\n\t* "))
}

// Unwrap returns each finding, so errors.As reaches one through the collection.
func (fs Findings) Unwrap() []error {
	errs := make([]error, 0, len(fs))
	for _, f := range fs {
		errs = append(errs, f)
	}

	return errs
}

// Err returns the collection as an error, or nil when nothing was found. It is
// the one place a Findings becomes an error, so a caller never receives a
// non-nil error holding an empty collection.
func (fs Findings) Err() error {
	if len(fs) == 0 {
		return nil
	}

	return fs
}
