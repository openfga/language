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

func buildErrorFor(t *testing.T, dsl string) *BuildError {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.Error(t, err, "the builder accepted the model, so there is no refusal to read")

	var buildErr *BuildError
	require.ErrorAs(t, err, &buildErr, "a refusal must be a BuildError")

	return buildErr
}

// TestBuildRefusalNamesCycleMembers checks that a refused build names every
// relation of every unresolvable cycle, and still matches the sentinel it
// carried before BuildError existed.
func TestBuildRefusalNamesCycleMembers(t *testing.T) {
	t.Parallel()

	t.Run("model cycle", func(t *testing.T) {
		t.Parallel()

		buildErr := buildErrorFor(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`)

		require.ErrorIs(t, buildErr, ErrModelCycle)
		assert.Equal(t, "model cycle: document#a, document#b", buildErr.Error())
		require.Len(t, buildErr.Problems, 1)
		require.ErrorIs(t, buildErr.Problems[0].Kind, ErrModelCycle)
		assert.Equal(t, []string{"document#a", "document#b"}, buildErr.Problems[0].Relations)
	})

	t.Run("a relation rewriting itself is not a cycle", func(t *testing.T) {
		t.Parallel()

		// The builder accepts a bare self-rewrite, leaving it with no weights, so
		// it must not be named when a real cycle elsewhere triggers the refusal.
		buildErr := buildErrorFor(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
    define a: [user] or b
    define b: [user] or a
`)

		require.Len(t, buildErr.Problems, 1)
		assert.Equal(t, []string{"document#a", "document#b"}, buildErr.Problems[0].Relations)
	})

	t.Run("every unresolvable cycle is named", func(t *testing.T) {
		t.Parallel()

		buildErr := buildErrorFor(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
    define c: [user] or d
    define d: [user] or c
`)

		require.Len(t, buildErr.Problems, 2)
		assert.Equal(t, []string{"document#a", "document#b"}, buildErr.Problems[0].Relations)
		assert.Equal(t, []string{"document#c", "document#d"}, buildErr.Problems[1].Relations)
	})

	t.Run("tuple cycle under an exclusion", func(t *testing.T) {
		t.Parallel()

		buildErr := buildErrorFor(t, `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not blocked
    define blocked: [user, group#member]
`)

		require.ErrorIs(t, buildErr, ErrTupleCycle)
		require.Len(t, buildErr.Problems, 1)
		require.ErrorIs(t, buildErr.Problems[0].Kind, ErrContrainstTupleCycle)
		assert.Equal(t, []string{"group#blocked", "group#member"}, buildErr.Problems[0].Relations)
	})

	// The two cases below are the regression the structural read exists for:
	// membership read off the weight walk depended on relation names, because
	// the walk recorded only the nodes it happened to re-enter.
	t.Run("membership does not depend on relation names", func(t *testing.T) {
		t.Parallel()

		buildErr := buildErrorFor(t, `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not zblocked
    define zblocked: [user, group#member]
`)

		require.Len(t, buildErr.Problems, 1)
		assert.Equal(t, []string{"group#member", "group#zblocked"}, buildErr.Problems[0].Relations)
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

			buildErr := buildErrorFor(t, dsl)
			require.Len(t, buildErr.Problems, 1)
			assert.Equal(t, expected, buildErr.Problems[0].Relations, "under the name %q", viewer)
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
			_, err := NewWeightedAuthorizationModelGraphBuilder().Build(model)

			var buildErr *BuildError
			require.ErrorAs(t, err, &buildErr)
			require.Len(t, buildErr.Problems, 1)
			assert.Equal(t, []string{"group#member", "group#zblocked"}, buildErr.Problems[0].Relations)
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

		_, err = NewWeightedAuthorizationModelGraphBuilder().Build(model)
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
		buildErr := buildErrorFor(t, dsl)
		assert.NotEmptyf(t, buildErr.Problems, "%s: refused for a cycle but none was named", name)

		for _, problem := range buildErr.Problems {
			assert.NotEmptyf(t, problem.Relations, "%s: a cycle with no relations is unreportable", name)
			assert.Truef(t, errors.Is(problem.Kind, ErrModelCycle) || errors.Is(problem.Kind, ErrTupleCycle),
				"%s: problem kind %v is not a cycle sentinel", name, problem.Kind)
		}
	}
}

// TestBuildErrorMatchesTheSentinels pins what callers of Build keep: errors.Is
// still matches the sentinel for every problem, so code that checks the kind of
// refusal is unaffected. Only code that compares the error with == would notice
// that it is now a *BuildError.
func TestBuildErrorMatchesTheSentinels(t *testing.T) {
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

			_, err = NewWeightedAuthorizationModelGraphBuilder().Build(model)

			var buildErr *BuildError
			require.ErrorAs(t, err, &buildErr)

			for _, sentinel := range test.want {
				require.ErrorIs(t, err, sentinel)
			}
		})
	}
}
