package graph

import (
	"cmp"
	"slices"
	"strings"
)

// RelationError is one problem Parse found, with the relations it involves.
type RelationError struct {
	Relations []string // "type#relation", sorted; empty if no relation can be named
	Err       error    // ErrModelCycle, ErrTupleCycle or ErrInvalidModel, possibly wrapped
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

// RelationErrors is the error Parse returns for a model it rejects, sorted by relations.
// Get it with errors.As; errors.Is matches each problem's sentinel.
type RelationErrors []*RelationError

// Error returns one line per problem.
func (e RelationErrors) Error() string {
	lines := make([]string, len(e))
	for i, relationErr := range e {
		lines[i] = relationErr.Error()
	}

	return strings.Join(lines, "\n")
}

// Unwrap returns each *RelationError, for errors.Is and errors.As.
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
