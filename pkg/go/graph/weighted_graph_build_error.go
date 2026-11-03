package graph

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strings"
)

// RelationError is one reason Build rejected a model, with the relations it
// involves.
type RelationError struct {
	// Relations names the relations the error involves, as "type#relation",
	// sorted. A cycle names every relation in it. It is empty when the builder
	// cannot attribute the error to a relation.
	Relations []string

	// Err is the reason. It is, or wraps, one of ErrModelCycle,
	// ErrContrainstTupleCycle, ErrTupleCycle or ErrInvalidModel.
	Err error
}

func (e *RelationError) Error() string {
	if len(e.Relations) == 0 {
		return e.Err.Error()
	}

	return e.Err.Error() + ": " + strings.Join(e.Relations, ", ")
}

func (e *RelationError) Unwrap() error {
	return e.Err
}

// RelationMultipleError is the error Build returns when it rejects a model, so a
// caller can report every problem at once rather than the first.
type RelationMultipleError struct {
	// Errors holds one *RelationError per problem, sorted by the relations they
	// involve.
	Errors []error
}

func (e *RelationMultipleError) Unwrap() []error {
	return e.Errors
}

func (e *RelationMultipleError) Error() string {
	pluralS := ""
	if len(e.Errors) > 1 {
		pluralS = "s"
	}

	messages := make([]string, 0, len(e.Errors))
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}

	return fmt.Sprintf("%d error%s occurred:\n\t* %s\n\n", len(e.Errors), pluralS, strings.Join(messages, "\n\t* "))
}

// newRelationMultipleError sorts the relation errors, so the same model is
// rejected with the same error on every build.
func newRelationMultipleError(relationErrs []RelationError) *RelationMultipleError {
	slices.SortFunc(relationErrs, func(a, b RelationError) int {
		return cmp.Or(slices.Compare(a.Relations, b.Relations), cmp.Compare(a.Err.Error(), b.Err.Error()))
	})

	errs := make([]error, len(relationErrs))
	for i := range relationErrs {
		errs[i] = &relationErrs[i]
	}

	return &RelationMultipleError{Errors: errs}
}

// walkErrors turns a failed weight assignment into relation errors: every
// unresolvable cycle, read off the graph's structure, and the assignment's own
// error when it is not about those cycles.
func walkErrors(wg *WeightedAuthorizationModelGraph, err error) []RelationError {
	relationErrs := wg.unresolvableCycles()

	isCycleRefusal := errors.Is(err, ErrModelCycle) || errors.Is(err, ErrTupleCycle)
	if len(relationErrs) == 0 || !isCycleRefusal {
		relationErrs = append(relationErrs, RelationError{Err: err})
	}

	return relationErrs
}
