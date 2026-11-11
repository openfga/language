package graph

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/transformer"
)

// rejectionFor builds dsl, which the builder must reject, and returns the error
// it was rejected with.
func rejectionFor(t *testing.T, dsl string) error {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().BuildForValidation(model)
	require.Error(t, err, "the builder accepted the model, so there is no rejection to read")

	return err
}

// problemsOf returns the relation errors a rejection holds, failing if it holds
// none.
func problemsOf(t *testing.T, err error) []*RelationError {
	t.Helper()

	relationErrs := RelationErrors(err)
	require.NotEmpty(t, relationErrs, "a rejection must hold relation errors")

	return relationErrs
}

// TestBuildForValidationNamesCycleMembers checks that a refused build names every
// relation of every unresolvable cycle, and still matches the sentinel it
// carried before it named relations.
func TestBuildForValidationNamesCycleMembers(t *testing.T) {
	t.Parallel()

	t.Run("model cycle", func(t *testing.T) {
		t.Parallel()

		err := rejectionFor(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`)

		require.ErrorIs(t, err, ErrModelCycle)
		assert.Equal(t, "model cycle: document#a, document#b", err.Error())

		relationErrs := problemsOf(t, err)
		require.Len(t, relationErrs, 1)
		require.ErrorIs(t, relationErrs[0], ErrModelCycle)
		assert.Equal(t, []string{"document#a", "document#b"}, relationErrs[0].Relations)

		var relationErr *RelationError
		require.ErrorAs(t, err, &relationErr, "errors.As must find a single relation error")
		assert.Equal(t, []string{"document#a", "document#b"}, relationErr.Relations)
	})

	t.Run("a relation rewriting itself is not a cycle", func(t *testing.T) {
		t.Parallel()

		// The builder accepts a bare self-rewrite, leaving it with no weights, so
		// it must not be named when a real cycle elsewhere triggers the refusal.
		relationErrs := problemsOf(t, rejectionFor(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
    define a: [user] or b
    define b: [user] or a
`))

		require.Len(t, relationErrs, 1)
		assert.Equal(t, []string{"document#a", "document#b"}, relationErrs[0].Relations)
	})

	t.Run("every unresolvable cycle is named", func(t *testing.T) {
		t.Parallel()

		relationErrs := problemsOf(t, rejectionFor(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
    define c: [user] or d
    define d: [user] or c
`))

		require.Len(t, relationErrs, 2)
		assert.Equal(t, []string{"document#a", "document#b"}, relationErrs[0].Relations)
		assert.Equal(t, []string{"document#c", "document#d"}, relationErrs[1].Relations)
	})

	t.Run("tuple cycle under an exclusion", func(t *testing.T) {
		t.Parallel()

		err := rejectionFor(t, `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not blocked
    define blocked: [user, group#member]
`)

		require.ErrorIs(t, err, ErrTupleCycle)

		relationErrs := problemsOf(t, err)
		require.Len(t, relationErrs, 1)
		require.ErrorIs(t, relationErrs[0], ErrContrainstTupleCycle)
		assert.Equal(t, []string{"group#blocked", "group#member"}, relationErrs[0].Relations)
	})

	// The two cases below are the regression the structural read exists for:
	// membership read off the weight walk depended on relation names, because
	// the walk recorded only the nodes it happened to re-enter.
	t.Run("membership does not depend on relation names", func(t *testing.T) {
		t.Parallel()

		relationErrs := problemsOf(t, rejectionFor(t, `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not zblocked
    define zblocked: [user, group#member]
`))

		require.Len(t, relationErrs, 1)
		assert.Equal(t, []string{"group#member", "group#zblocked"}, relationErrs[0].Relations)
	})

	t.Run("a relation between the cycle start and the operator is named", func(t *testing.T) {
		t.Parallel()

		for _, viewer := range []string{"viewer", "aviewer"} {
			dsl := fmt.Sprintf(`model
  schema 1.1
type user
type group
  relations
    define member: [user, group#%[1]s]
    define %[1]s: member but not blocked
    define blocked: [user]
`, viewer)

			expected := []string{"group#" + viewer, "group#member"}
			slices.Sort(expected)

			relationErrs := problemsOf(t, rejectionFor(t, dsl))
			require.Len(t, relationErrs, 1)
			assert.Equal(t, expected, relationErrs[0].Relations, "under the name %q", viewer)
		}
	})

	t.Run("the same answer on every run", func(t *testing.T) {
		t.Parallel()

		dsl := `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not zblocked
    define zblocked: [user, group#member]
`
		model, err := transformer.TransformDSLToProto(dsl)
		require.NoError(t, err)

		for range 20 {
			_, err := NewWeightedAuthorizationModelGraphBuilder().BuildForValidation(model)

			relationErrs := problemsOf(t, err)
			require.Len(t, relationErrs, 1)
			assert.Equal(t, []string{"group#member", "group#zblocked"}, relationErrs[0].Relations)
		}
	})

	t.Run("legal recursion is not refused", func(t *testing.T) {
		t.Parallel()

		model, err := transformer.TransformDSLToProto(`model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member]
`)
		require.NoError(t, err)

		_, err = NewWeightedAuthorizationModelGraphBuilder().BuildForValidation(model)
		require.NoError(t, err)
	})
}

// TestUnresolvableCyclesAgreeWithTheWalk holds the structural read to the
// weight walk on refusals the walk makes for a cycle: each one names at least
// one cycle, with a cycle sentinel and the relations in it.
func TestUnresolvableCyclesAgreeWithTheWalk(t *testing.T) {
	t.Parallel()

	refusals := map[string]string{
		"intersection": `model
  schema 1.1
type user
type document
  relations
    define admin: [user]
    define viewer: admin and editor
    define editor: viewer
`,
		"exclusion ttu": `model
  schema 1.1
type user
type folder
  relations
    define parent: [folder]
    define viewer: [user] but not banned
    define banned: viewer from parent
`,
	}

	for name, dsl := range refusals {
		relationErrs := problemsOf(t, rejectionFor(t, dsl))
		assert.NotEmptyf(t, relationErrs, "%s: refused for a cycle but none was named", name)

		for _, relationErr := range relationErrs {
			assert.NotEmptyf(t, relationErr.Relations, "%s: a cycle with no relations is unreportable", name)
			assert.Truef(t, errors.Is(relationErr, ErrModelCycle) || errors.Is(relationErr, ErrTupleCycle),
				"%s: %v is not a cycle sentinel", name, relationErr.Err)
		}
	}
}

// TestBuildForValidationErrorMatchesTheSentinels pins that errors.Is matches
// the sentinel for every problem in BuildForValidation's error, as it does for
// Build's, so code that checks the kind of refusal works with either.
func TestBuildForValidationErrorMatchesTheSentinels(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl  string
		want []error
	}{
		"model cycle": {
			dsl: `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`,
			want: []error{ErrModelCycle},
		},
		"tuple cycle under an exclusion": {
			dsl: `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not blocked
    define blocked: [user, group#member]
`,
			want: []error{ErrContrainstTupleCycle, ErrTupleCycle},
		},
		"a model cycle and a missing from target": {
			dsl: `model
  schema 1.1
type user
type org
type folder
  relations
    define viewer: [user]
type document
  relations
    define parent: [folder, org]
    define viewer: viewer from parent
    define a: [user] or b
    define b: [user] or a
`,
			want: []error{ErrModelCycle, ErrInvalidModel},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model, err := transformer.TransformDSLToProto(test.dsl)
			require.NoError(t, err)

			_, err = NewWeightedAuthorizationModelGraphBuilder().BuildForValidation(model)
			require.Error(t, err)

			for _, sentinel := range test.want {
				require.ErrorIs(t, err, sentinel)
			}
		})
	}
}

// TestBuildForValidationRejectsAModelTheSameWayOnEveryRun pins that its weight
// walk visits nodes in a fixed order. The walk stops at its first error, so
// which of a cycle and an invalid relation it reaches first decides whether the
// invalid relation is reported, and that must not change from one build to the
// next.
func TestBuildForValidationRejectsAModelTheSameWayOnEveryRun(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl      string
		contains []string
		excludes []string
	}{
		"the cycle is reached first and hides the invalid relation": {
			dsl: `model
  schema 1.1
type user
type employee
type document
  relations
    define a: [user] or b
    define b: [user] or a
    define x: [user] and y
    define y: [employee]
`,
			contains: []string{"model cycle: document#a, document#b"},
			excludes: []string{"invalid model"},
		},
		"the invalid relation is reached first and both are reported": {
			dsl: `model
  schema 1.1
type user
type employee
type document
  relations
    define x: [user] or y
    define y: [user] or x
    define a: [user] and b
    define b: [employee]
`,
			contains: []string{"model cycle: document#x, document#y", "invalid model"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model, err := transformer.TransformDSLToProto(test.dsl)
			require.NoError(t, err)

			_, err = NewWeightedAuthorizationModelGraphBuilder().BuildForValidation(model)
			require.Error(t, err)

			first := err.Error()
			for _, want := range test.contains {
				assert.Contains(t, first, want)
			}

			for _, unwanted := range test.excludes {
				assert.NotContains(t, first, unwanted)
			}

			for range 200 {
				_, err := NewWeightedAuthorizationModelGraphBuilder().BuildForValidation(model)
				require.Error(t, err)
				require.Equal(t, first, err.Error())
			}
		})
	}
}

// TestRelationErrorsIsNilForOtherErrors pins that RelationErrors finds nothing
// in a nil error, in one that holds no relation error, or in Build's error.
func TestRelationErrorsIsNilForOtherErrors(t *testing.T) {
	t.Parallel()

	assert.Nil(t, RelationErrors(nil))
	assert.Nil(t, RelationErrors(errors.New("x")))

	model, err := transformer.TransformDSLToProto(`model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.ErrorIs(t, err, ErrModelCycle)
	assert.Nil(t, RelationErrors(err))
}

// TestBuildStillStopsAtTheFirstProblem pins that Build is unaffected by
// BuildForValidation: on a model with a relation that fails to build and two
// broken cycles, it returns the construction error, as it always has, with no
// relation errors in it.
func TestBuildStillStopsAtTheFirstProblem(t *testing.T) {
	t.Parallel()

	model, err := transformer.TransformDSLToProto(`model
  schema 1.1
type user
type folder
  relations
    define viewer: [user]
type document
  relations
    define parent: [folder]
    define owner: [user]
    define editor: [user] or writer
    define writer: [user] or editor
    define member: [user, document#member] but not blocked
    define blocked: [user, document#member]
    define reader: owner from parent
`)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.ErrorIs(t, err, ErrInvalidModel)
	assert.Equal(t, "invalid model: Model cannot be parsed. folder type does not have defined owner relation", err.Error())
	assert.Nil(t, RelationErrors(err))

	_, err = NewWeightedAuthorizationModelGraphBuilder().BuildForValidation(model)
	assert.Len(t, RelationErrors(err), 3)
}

// TestRelationErrorReadsAsItsReasonAndRelations pins how a relation error reads
// in a log line: its reason, then the relations it involves, and that it
// unwraps to the reason.
func TestRelationErrorReadsAsItsReasonAndRelations(t *testing.T) {
	t.Parallel()

	relationErr := &RelationError{Relations: []string{"document#a", "document#b"}, Err: ErrModelCycle}
	assert.Equal(t, "model cycle: document#a, document#b", relationErr.Error())
	require.ErrorIs(t, relationErr, ErrModelCycle)

	unattributed := &RelationError{Err: ErrInvalidModel}
	assert.Equal(t, "invalid model", unattributed.Error())
}
