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

func cycleRefusalFor(t *testing.T, dsl string) *CycleError {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().BuildWithCycles(model)
	require.Error(t, err, "the builder accepted the model, so there is no refusal to read")

	var cycleErr *CycleError
	require.ErrorAs(t, err, &cycleErr, "a cycle refusal must carry a CycleError")

	return cycleErr
}

// TestBuildRefusalNamesCycleMembers checks that a refused build names every
// relation of every unresolvable cycle, with the sentinel and message it
// carried before CycleError existed.
func TestBuildRefusalNamesCycleMembers(t *testing.T) {
	t.Parallel()

	t.Run("model cycle", func(t *testing.T) {
		t.Parallel()

		cycleErr := cycleRefusalFor(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`)

		require.ErrorIs(t, cycleErr, ErrModelCycle)
		assert.Equal(t, "model cycle", cycleErr.Error())
		require.Len(t, cycleErr.Cycles, 1)
		require.ErrorIs(t, cycleErr.Cycles[0].Kind, ErrModelCycle)
		assert.Equal(t, []string{"document#a", "document#b"}, cycleErr.Cycles[0].Members)
	})

	t.Run("a relation rewriting itself is not a cycle", func(t *testing.T) {
		t.Parallel()

		// The builder resolves a bare self-rewrite through its recursive-edge
		// special case, so it must not be named when a real cycle elsewhere
		// triggers the refusal.
		cycleErr := cycleRefusalFor(t, `model
  schema 1.1
type user
type document
  relations
    define viewer: viewer
    define a: [user] or b
    define b: [user] or a
`)

		require.Len(t, cycleErr.Cycles, 1)
		assert.Equal(t, []string{"document#a", "document#b"}, cycleErr.Cycles[0].Members)
	})

	t.Run("every unresolvable cycle is named", func(t *testing.T) {
		t.Parallel()

		cycleErr := cycleRefusalFor(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
    define c: [user] or d
    define d: [user] or c
`)

		require.Len(t, cycleErr.Cycles, 2)
		assert.Equal(t, []string{"document#a", "document#b"}, cycleErr.Cycles[0].Members)
		assert.Equal(t, []string{"document#c", "document#d"}, cycleErr.Cycles[1].Members)
	})

	t.Run("tuple cycle under an exclusion", func(t *testing.T) {
		t.Parallel()

		cycleErr := cycleRefusalFor(t, `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not blocked
    define blocked: [user, group#member]
`)

		require.ErrorIs(t, cycleErr, ErrTupleCycle)
		require.Len(t, cycleErr.Cycles, 1)
		require.ErrorIs(t, cycleErr.Cycles[0].Kind, ErrContrainstTupleCycle)
		assert.Equal(t, []string{"group#blocked", "group#member"}, cycleErr.Cycles[0].Members)
	})

	// The two cases below are the regression the structural read exists for:
	// membership read off the weight walk depended on relation names, because
	// the walk recorded only the nodes it happened to re-enter.
	t.Run("membership does not depend on relation names", func(t *testing.T) {
		t.Parallel()

		cycleErr := cycleRefusalFor(t, `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not zblocked
    define zblocked: [user, group#member]
`)

		require.Len(t, cycleErr.Cycles, 1)
		assert.Equal(t, []string{"group#member", "group#zblocked"}, cycleErr.Cycles[0].Members)
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

			cycleErr := cycleRefusalFor(t, dsl)
			require.Len(t, cycleErr.Cycles, 1)
			assert.Equal(t, expected, cycleErr.Cycles[0].Members, "under the name %q", viewer)
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
			_, buildErr := NewWeightedAuthorizationModelGraphBuilder().BuildWithCycles(model)

			var cycleErr *CycleError
			require.ErrorAs(t, buildErr, &cycleErr)
			require.Len(t, cycleErr.Cycles, 1)
			assert.Equal(t, []string{"group#member", "group#zblocked"}, cycleErr.Cycles[0].Members)
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

		_, buildErr := NewWeightedAuthorizationModelGraphBuilder().BuildWithCycles(model)
		require.NoError(t, buildErr)
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
		cycleErr := cycleRefusalFor(t, dsl)
		assert.NotEmptyf(t, cycleErr.Cycles, "%s: refused for a cycle but none was named", name)

		for _, cycle := range cycleErr.Cycles {
			assert.NotEmptyf(t, cycle.Members, "%s: a cycle with no relations is unreportable", name)
			assert.Truef(t, errors.Is(cycle.Kind, ErrModelCycle) || errors.Is(cycle.Kind, ErrTupleCycle),
				"%s: cycle kind %v is not a cycle sentinel", name, cycle.Kind)
		}
	}
}

// TestBuildRefusalIsUnchanged pins that Build itself is untouched: a cycle
// refusal is still the bare sentinel, not a *CycleError. Only BuildWithCycles
// names the cycles.
func TestBuildRefusalIsUnchanged(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		dsl  string
		want error
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
			want: ErrModelCycle,
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
			want: ErrContrainstTupleCycle,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			model, err := transformer.TransformDSLToProto(test.dsl)
			require.NoError(t, err)

			_, err = NewWeightedAuthorizationModelGraphBuilder().Build(model)
			require.Equal(t, test.want, err, "Build must return the sentinel itself")
		})
	}
}
