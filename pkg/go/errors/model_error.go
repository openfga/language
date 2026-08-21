package errors

import "fmt"

// The errors here name the part of a model a finding is about, and each declares
// only the fields its scope has: an ErrCondition has no relation, an
// ErrRelationCondition has all three. Cause holds the sentinel.
//
//	var relationErr *errors.ErrRelation
//	if errors.As(err, &relationErr) {
//	    fmt.Println(relationErr.ObjectType, relationErr.Relation)
//	}
//
// They are named for the scope rather than the problem because many sentinels
// share a scope, and the shapes match the server's pkg/typesystem. Every exported
// error name here is Err-prefixed, types as well as sentinel values, so each type
// below opts out of errname's XxxError rule.

// ModelError is the cause of a validation finding: a sentinel naming the problem,
// wrapped in the part of the model it was found in.
//
// A caller reaches the problem with errors.Is, and the part of the model either
// with errors.As on one of the concrete types, or through Kind and Scope when the
// concrete type does not matter:
//
//	var modelErr errors.ModelError
//	if errors.As(err, &modelErr) {
//	    fmt.Println(modelErr.Kind(), modelErr.Scope().Relation)
//	}
//
// The interface has an unexported method, so the five types below are its only
// implementations and a Kind always corresponds to one of them.
type ModelError interface {
	error

	// Kind reports which part of the model this finding is attached to.
	Kind() ModelErrorKind

	// Scope names that part. Which fields are set follows from Kind; the rest are
	// empty.
	Scope() ModelErrorScope

	// Unwrap returns the sentinel, so errors.Is reaches it through this error.
	Unwrap() error

	// withSentinel returns a copy reporting sentinel as the problem, leaving the
	// receiver alone. WithSentinel is the exported way in.
	withSentinel(sentinel error) ModelError
}

// ModelErrorScope names the part of a model a finding is attached to, for a caller
// that wants the names without switching on the concrete type. A finding about the
// model as a whole has none of the three set.
type ModelErrorScope struct {
	ObjectType string
	Relation   string
	Condition  string
}

// WithSentinel returns err reporting sentinel as the problem it names, leaving err
// unchanged. It lets the part of the model at fault and the problem be decided in
// different places: whoever finds the fault builds the scope, and whichever code is
// being raised supplies the sentinel.
//
// A nil sentinel yields nil rather than an error whose message reports nothing.
func WithSentinel(err ModelError, sentinel error) ModelError {
	if err == nil || sentinel == nil {
		return nil
	}

	return err.withSentinel(sentinel)
}

// Every scope type is a ModelError; a new one that forgets a method fails to build
// here rather than at whichever call site first needs it.
var (
	_ ModelError = (*ErrObjectType)(nil)
	_ ModelError = (*ErrRelation)(nil)
	_ ModelError = (*ErrRelationCondition)(nil)
	_ ModelError = (*ErrCondition)(nil)
	_ ModelError = (*ErrModel)(nil)
)

// ErrObjectType is a finding about an object type as a whole.
//
//nolint:errname // Err-prefixed by convention here; see the naming note above
type ErrObjectType struct {
	ObjectType string
	Cause      error
}

func (e *ErrObjectType) Error() string {
	return fmt.Sprintf("error in the definition of the object type '%s': %s", e.ObjectType, e.Cause)
}

func (e *ErrObjectType) Unwrap() error {
	return e.Cause
}

func (e *ErrObjectType) Kind() ModelErrorKind {
	return ErrorKindObjectType
}

func (e *ErrObjectType) Scope() ModelErrorScope {
	return ModelErrorScope{ObjectType: e.ObjectType}
}

func (e *ErrObjectType) withSentinel(sentinel error) ModelError {
	return &ErrObjectType{ObjectType: e.ObjectType, Cause: sentinel}
}

// ErrRelation is a finding about a relation on an object type.
//
//nolint:errname // Err-prefixed by convention here; see the naming note above
type ErrRelation struct {
	ObjectType string
	Relation   string
	Cause      error
}

func (e *ErrRelation) Error() string {
	if e.ObjectType == "" {
		return fmt.Sprintf("error in the definition of relation '%s': %s", e.Relation, e.Cause)
	}

	return fmt.Sprintf("error in the definition of relation '%s' of object type '%s': %s",
		e.Relation, e.ObjectType, e.Cause)
}

func (e *ErrRelation) Unwrap() error {
	return e.Cause
}

func (e *ErrRelation) Kind() ModelErrorKind {
	return ErrorKindRelation
}

func (e *ErrRelation) Scope() ModelErrorScope {
	return ModelErrorScope{ObjectType: e.ObjectType, Relation: e.Relation}
}

func (e *ErrRelation) withSentinel(sentinel error) ModelError {
	return &ErrRelation{ObjectType: e.ObjectType, Relation: e.Relation, Cause: sentinel}
}

// ErrRelationCondition is a finding about a condition as applied to one relation,
// rather than about the condition's own definition.
//
//nolint:errname // Err-prefixed by convention here; see the naming note above
type ErrRelationCondition struct {
	ObjectType string
	Relation   string
	Condition  string
	Cause      error
}

func (e *ErrRelationCondition) Error() string {
	return fmt.Sprintf("error in the definition of condition '%s' of relation '%s' in object type '%s': %s",
		e.Condition, e.Relation, e.ObjectType, e.Cause)
}

func (e *ErrRelationCondition) Unwrap() error {
	return e.Cause
}

func (e *ErrRelationCondition) Kind() ModelErrorKind {
	return ErrorKindRelationCondition
}

func (e *ErrRelationCondition) Scope() ModelErrorScope {
	return ModelErrorScope{ObjectType: e.ObjectType, Relation: e.Relation, Condition: e.Condition}
}

func (e *ErrRelationCondition) withSentinel(sentinel error) ModelError {
	return &ErrRelationCondition{
		ObjectType: e.ObjectType,
		Relation:   e.Relation,
		Condition:  e.Condition,
		Cause:      sentinel,
	}
}

// ErrCondition is a finding about a condition definition itself, independent of
// where it is applied.
//
//nolint:errname // Err-prefixed by convention here; see the naming note above
type ErrCondition struct {
	Condition string
	Cause     error
}

func (e *ErrCondition) Error() string {
	return fmt.Sprintf("error in the definition of condition '%s': %s", e.Condition, e.Cause)
}

func (e *ErrCondition) Unwrap() error {
	return e.Cause
}

func (e *ErrCondition) Kind() ModelErrorKind {
	return ErrorKindCondition
}

func (e *ErrCondition) Scope() ModelErrorScope {
	return ModelErrorScope{Condition: e.Condition}
}

func (e *ErrCondition) withSentinel(sentinel error) ModelError {
	return &ErrCondition{Condition: e.Condition, Cause: sentinel}
}

// ErrModel is a finding about the model as a whole, which cannot be attributed
// to a single type, relation or condition.
//
//nolint:errname // Err-prefixed by convention here; see the naming note above
type ErrModel struct {
	Cause error
}

func (e *ErrModel) Error() string {
	return fmt.Sprintf("error in authorization model: %s", e.Cause)
}

func (e *ErrModel) Unwrap() error {
	return e.Cause
}

func (e *ErrModel) Kind() ModelErrorKind {
	return ErrorKindInvalidModel
}

// Scope returns an empty scope: a finding about the model as a whole names no type,
// relation or condition, which is what distinguishes it from the other four.
func (e *ErrModel) Scope() ModelErrorScope {
	return ModelErrorScope{}
}

func (e *ErrModel) withSentinel(sentinel error) ModelError {
	return &ErrModel{Cause: sentinel}
}
