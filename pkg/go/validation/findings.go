// Package validation checks OpenFGA authorization models for semantic errors
// beyond what the DSL grammar enforces, mirroring the checks in pkg/js and
// pkg/java. The three implementations share the corpus under tests/data, which
// pins each finding's message, code, symbol and position.
package validation

import (
	"errors"
	"fmt"
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

// Error implements the error interface, formatting the finding with its
// position when it has one.
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

// joinFindings joins findings into a single error, or nil when there are none.
// A nil *Finding is dropped, so a check that found nothing can be appended
// without a guard. The findings are joined in the order given, and ExtractAllAs
// recovers them in that order.
func joinFindings(findings ...*Finding) error {
	errs := make([]error, 0, len(findings))
	for _, f := range findings {
		if f != nil {
			errs = append(errs, f)
		}
	}

	return errors.Join(errs...)
}

// ExtractAllAs walks an error tree and returns every error of type E, in the
// order errors.Join holds them: pre-order, left to right. It recovers the
// findings behind a validation error:
//
//	for _, finding := range validation.ExtractAllAs[*validation.Finding](err) {
//	    ...
//	}
//
// An error that is an E is collected and not descended into; anything else is
// expanded through its Unwrap() []error or Unwrap() error.
func ExtractAllAs[E error](err error) []E {
	var found []E

	var collect func(error)
	collect = func(err error) {
		if err == nil {
			return
		}

		if e, ok := err.(E); ok {
			found = append(found, e)

			return
		}

		switch unwrapped := err.(type) {
		case interface{ Unwrap() []error }:
			for _, child := range unwrapped.Unwrap() {
				collect(child)
			}
		case interface{ Unwrap() error }:
			collect(unwrapped.Unwrap())
		}
	}
	collect(err)

	return found
}
