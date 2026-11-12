package graph

import (
	"cmp"
	"errors"
	"slices"
	"strings"
)

// RelationError is one reason Parse rejected a model, with the relations it
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

// RelationErrors is the error Parse returns for a model it rejects: every
// reason, sorted by relations, so the same model is rejected the same way on
// every run. Callers get it with errors.As, and errors.Is matches the sentinel
// of each error in it.
type RelationErrors []*RelationError

// Error is each error on its own line.
func (e RelationErrors) Error() string {
	lines := make([]string, len(e))
	for i, relationErr := range e {
		lines[i] = relationErr.Error()
	}

	return strings.Join(lines, "\n")
}

// Unwrap returns the errors, so errors.Is matches each one's sentinel and
// errors.As finds a single *RelationError.
func (e RelationErrors) Unwrap() []error {
	errs := make([]error, len(e))
	for i, relationErr := range e {
		errs[i] = relationErr
	}

	return errs
}

// sort orders the errors by their relations, then by their reason.
func (e RelationErrors) sort() {
	slices.SortFunc(e, func(a, b *RelationError) int {
		return cmp.Or(slices.Compare(a.Relations, b.Relations), cmp.Compare(a.Err.Error(), b.Err.Error()))
	})
}

// assignWeightErrors turns a failed weight assignment into relation errors:
// every unresolvable cycle, read off the graph's structure, and the
// assignment's own error when it is not about those cycles.
//
// That error is still the first one the walk reaches, because the walk stops at
// its first error, so a cycle reached first can hide a later error that is not a
// cycle. The walk Parse uses starts from the relations in a fixed order, so the
// same model gives the same answer every time.
func assignWeightErrors(wg *WeightedAuthorizationModelGraph, err error) RelationErrors {
	relationErrs := wg.unresolvableCycles()

	isCycleRefusal := errors.Is(err, ErrModelCycle) || errors.Is(err, ErrTupleCycle)
	if len(relationErrs) == 0 || !isCycleRefusal {
		relationErrs = append(relationErrs, &RelationError{Err: err})
	}

	return relationErrs
}
