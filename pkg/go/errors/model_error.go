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
