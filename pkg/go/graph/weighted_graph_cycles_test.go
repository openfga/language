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

// rejectionFor parses dsl, which Parse must reject, and returns the error.
func rejectionFor(t *testing.T, dsl string) error {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().Parse(model)
	require.Error(t, err, "the builder accepted the model, so there is no rejection to read")

	return err
}

// relationErrorsIn returns the RelationErrors errors.As finds in err, or nil.
func relationErrorsIn(err error) RelationErrors {
	var relationErrs RelationErrors
	if !errors.As(err, &relationErrs) {
		return nil
	}

	return relationErrs
}

// problemsOf returns the RelationErrors in err, failing if there are none.
func problemsOf(t *testing.T, err error) RelationErrors {
	t.Helper()

	relationErrs := relationErrorsIn(err)
	require.NotEmpty(t, relationErrs, "a rejection must hold relation errors")

	return relationErrs
}

// TestParseNamesCycleMembers checks that each rejected cycle is named with its relations
// and still matches its sentinel.
func TestParseNamesCycleMembers(t *testing.T) {
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

		// The builder accepts a relation that only rewrites itself, so it isn't named.
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

	t.Run("every rejected cycle is named", func(t *testing.T) {
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

	// Regression: cycle members read off the weight walk depended on relation names.
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
			_, err := NewWeightedAuthorizationModelGraphBuilder().Parse(model)

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

		_, err = NewWeightedAuthorizationModelGraphBuilder().Parse(model)
		require.NoError(t, err)
	})
}

// TestRejectedCyclesAgreeWithTheWalk checks that a model the walk rejects for a cycle
// gets a cycle error with the relations in it.
func TestRejectedCyclesAgreeWithTheWalk(t *testing.T) {
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

// TestParseErrorMatchesTheSentinels checks that errors.Is matches each problem's sentinel.
func TestParseErrorMatchesTheSentinels(t *testing.T) {
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

			_, err = NewWeightedAuthorizationModelGraphBuilder().Parse(model)
			require.Error(t, err)

			for _, sentinel := range test.want {
				require.ErrorIs(t, err, sentinel)
			}
		})
	}
}

// TestParseRejectsAModelTheSameWayOnEveryRun checks that a cycle doesn't hide an invalid
// relation, whichever sorts first, and that every run gives the same errors.
func TestParseRejectsAModelTheSameWayOnEveryRun(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl  string
		want string
	}{
		"the cycle sorts first": {
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
			want: "model cycle: document#a, document#b\n" +
				"invalid model: not all paths return the same type for the node document#x:intersection:0: document#x",
		},
		"the invalid relation sorts first": {
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
			want: "invalid model: not all paths return the same type for the node document#a:intersection:0: document#a\n" +
				"model cycle: document#x, document#y",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model, err := transformer.TransformDSLToProto(test.dsl)
			require.NoError(t, err)

			for range 200 {
				_, err := NewWeightedAuthorizationModelGraphBuilder().Parse(model)
				require.Error(t, err)
				require.Equal(t, test.want, err.Error())
			}
		})
	}
}

// TestOnlyParseReturnsRelationErrors checks that Parse's error is RelationErrors and Build's isn't.
func TestOnlyParseReturnsRelationErrors(t *testing.T) {
	t.Parallel()

	model, err := transformer.TransformDSLToProto(`model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().Parse(model)
	assert.IsType(t, RelationErrors{}, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.ErrorIs(t, err, ErrModelCycle)
	assert.Nil(t, relationErrorsIn(err))
}

// TestBuildStillStopsAtTheFirstProblem checks that Build still returns only the first problem.
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
	assert.Nil(t, relationErrorsIn(err))

	_, err = NewWeightedAuthorizationModelGraphBuilder().Parse(model)
	assert.Len(t, relationErrorsIn(err), 3)
}

// TestRelationErrorReadsAsItsReasonAndRelations checks how a RelationError reads and unwraps.
func TestRelationErrorReadsAsItsReasonAndRelations(t *testing.T) {
	t.Parallel()

	relationErr := &RelationError{Relations: []string{"document#a", "document#b"}, Err: ErrModelCycle}
	assert.Equal(t, "model cycle: document#a, document#b", relationErr.Error())
	require.ErrorIs(t, relationErr, ErrModelCycle)

	unattributed := &RelationError{Err: ErrInvalidModel}
	assert.Equal(t, "invalid model", unattributed.Error())
}

// TestRelationErrorsReadAndUnwrapAsTheirErrors checks how RelationErrors reads and unwraps.
func TestRelationErrorsReadAndUnwrapAsTheirErrors(t *testing.T) {
	t.Parallel()

	relationErrs := RelationErrors{
		{Relations: []string{"document#a", "document#b"}, Err: ErrModelCycle},
		{Relations: []string{"document#reader"}, Err: ErrInvalidModel},
	}
	assert.Equal(t, "model cycle: document#a, document#b\ninvalid model: document#reader", relationErrs.Error())
	require.ErrorIs(t, relationErrs, ErrModelCycle)
	require.ErrorIs(t, relationErrs, ErrInvalidModel)

	var relationErr *RelationError
	require.ErrorAs(t, relationErrs, &relationErr)
	assert.Same(t, relationErrs[0], relationErr)
}
