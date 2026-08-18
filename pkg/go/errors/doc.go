// Package errors holds the error types and sentinels that validation findings are
// built from, so callers match on values rather than on message text.
//
// A finding has two parts. The sentinel says what the problem is and is matched
// with errors.Is; see sentinels.go. The scope says which part of a model the
// problem is in, is the type the sentinel arrives wrapped in, and is matched with
// errors.As: ErrObjectType, ErrRelation, ErrRelationCondition, ErrCondition, or
// ErrModel when no single part is responsible.
//
// For a consumer that sees only serialised output, ModelErrorKind is the scope as
// a name and Severity is whether the finding blocks. Both reserve zero for "not
// set" and serialise as their name, so the names are API and the numbers are not.
package errors
