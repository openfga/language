package graph

import (
	"cmp"
	"errors"
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

// joinRelationErrors sorts the relation errors, so the same model is rejected
// with the same error on every build, and joins them. It is nil when there are
// none.
func joinRelationErrors(relationErrs []RelationError) error {
	slices.SortFunc(relationErrs, func(a, b RelationError) int {
		return cmp.Or(slices.Compare(a.Relations, b.Relations), cmp.Compare(a.Err.Error(), b.Err.Error()))
	})

	errs := make([]error, len(relationErrs))
	for i := range relationErrs {
		errs[i] = &relationErrs[i]
	}

	return errors.Join(errs...)
}

// RelationErrors returns every *RelationError in err, in the order Build
// reports them. It is nil for a nil error or one Build did not produce.
func RelationErrors(err error) []*RelationError {
	var relationErrs []*RelationError

	stack := []error{err}
	for len(stack) > 0 {
		ndx := len(stack) - 1
		current := stack[ndx]
		stack = stack[:ndx]

		if current == nil {
			continue
		}

		if relationErr, ok := current.(*RelationError); ok {
			relationErrs = append(relationErrs, relationErr)

			continue
		}

		if unwrapped, ok := current.(interface{ Unwrap() []error }); ok {
			// Push last to first so the stack pops them left to right.
			for _, child := range slices.Backward(unwrapped.Unwrap()) {
				stack = append(stack, child)
			}

			continue
		}

		stack = append(stack, errors.Unwrap(current))
	}

	return relationErrs
}

// assignWeightErrors turns a failed weight assignment into relation errors:
// every unresolvable cycle, read off the graph's structure, and the
// assignment's own error when it is not about those cycles.
//
// That error is still the first one the walk reaches, because the walk stops at
// its first error, so a cycle reached first can hide a later error that is not a
// cycle. The walk visits nodes in a fixed order, so the same model gives the same
// answer every time.
func assignWeightErrors(wg *WeightedAuthorizationModelGraph, err error) []RelationError {
	relationErrs := wg.unresolvableCycles()

	isCycleRefusal := errors.Is(err, ErrModelCycle) || errors.Is(err, ErrTupleCycle)
	if len(relationErrs) == 0 || !isCycleRefusal {
		relationErrs = append(relationErrs, RelationError{Err: err})
	}

	return relationErrs
}
