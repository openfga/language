package graph

import (
	"cmp"
	"errors"
	"slices"
	"strings"
)

// BuildError is the refusal Build returns: every problem that stops the model
// from being built, each with the relations it is about, so a caller can report
// all of them at once rather than the first.
type BuildError struct {
	// Problems holds every problem found, sorted by the relations they are about.
	Problems []Problem
}

// Problem is one reason Build refused a model.
type Problem struct {
	// Kind is the sentinel for the problem: ErrModelCycle, ErrContrainstTupleCycle,
	// ErrTupleCycle or ErrInvalidModel.
	Kind error

	// Relations names the relations the problem is about ("type#relation"),
	// sorted. A cycle names every relation in it. It is empty when the builder
	// cannot attribute the problem to a relation.
	Relations []string

	// Message describes the problem.
	Message string
}

// Error joins every problem's message, so a log line or an API error carries all
// of them.
func (e *BuildError) Error() string {
	messages := make([]string, 0, len(e.Problems))
	for _, problem := range e.Problems {
		messages = append(messages, problem.Message)
	}

	return strings.Join(messages, "; ")
}

// Unwrap returns every problem's Kind, so errors.Is matches any of them.
func (e *BuildError) Unwrap() []error {
	kinds := make([]error, 0, len(e.Problems))
	for _, problem := range e.Problems {
		kinds = append(kinds, problem.Kind)
	}

	return kinds
}

// newBuildError sorts the problems so the same model reports the same refusal on
// every build.
func newBuildError(problems []Problem) *BuildError {
	slices.SortFunc(problems, func(a, b Problem) int {
		return cmp.Or(slices.Compare(a.Relations, b.Relations), cmp.Compare(a.Message, b.Message))
	})

	return &BuildError{Problems: problems}
}

// cycleProblems reports each unresolvable cycle as a problem naming its
// relations.
func cycleProblems(cycles []cycle) []Problem {
	problems := make([]Problem, 0, len(cycles))
	for _, c := range cycles {
		problems = append(problems, Problem{Kind: c.kind, Relations: c.members, Message: c.kind.Error() + ": " + strings.Join(c.members, ", ")})
	}

	return problems
}

// walkProblems turns the weight walk's refusal into problems: every unresolvable
// cycle, read off the graph's structure, and the walk's own error when it is not
// one of those cycles.
func walkProblems(wg *WeightedAuthorizationModelGraph, err error) []Problem {
	problems := cycleProblems(wg.unresolvableCycles())

	isCycleRefusal := errors.Is(err, ErrModelCycle) || errors.Is(err, ErrTupleCycle)
	if len(problems) == 0 || !isCycleRefusal {
		problems = append(problems, Problem{Kind: sentinelOf(err), Message: err.Error()})
	}

	return problems
}

// sentinelOf returns the sentinel err wraps, most specific first: the
// constrained tuple cycle wraps the plain one.
func sentinelOf(err error) error {
	for _, sentinel := range []error{ErrContrainstTupleCycle, ErrTupleCycle, ErrModelCycle, ErrInvalidModel} {
		if errors.Is(err, sentinel) {
			return sentinel
		}
	}

	return err
}
