package graph

import (
	"cmp"
	"slices"
	"strings"
)

// RelationError is one reason Parse rejected a model, with the relations it
// involves.
type RelationError struct {
	// Relations names the relations the error involves, as "type#relation",
	// sorted. A cycle names every relation in it, and an error found while
	// assigning weights names the relations of the strongly connected component
	// it was found in. It is empty only when Parse finds no problem in a model
	// Build rejects, and returns Build's error on its own.
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
