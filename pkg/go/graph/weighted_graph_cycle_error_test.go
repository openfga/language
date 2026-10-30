package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/openfga/language/pkg/go/transformer"
)

func buildRefusalFor(t *testing.T, dsl string) *CycleError {
	t.Helper()

	model, err := transformer.TransformDSLToProto(dsl)
	require.NoError(t, err)

	_, err = NewWeightedAuthorizationModelGraphBuilder().Build(model)
	require.Error(t, err, "the builder accepted the model, so there is no refusal to read")

	var cycleErr *CycleError
	require.ErrorAs(t, err, &cycleErr, "a cycle refusal must carry a CycleError")

	return cycleErr
}

// TestBuildRefusalNamesCycleMembers checks that a refused build names the
// relations of the cycle it refused, with the sentinel and the message it
// carried before CycleError existed.
func TestBuildRefusalNamesCycleMembers(t *testing.T) {
	t.Parallel()

	t.Run("model cycle", func(t *testing.T) {
		t.Parallel()

		cycleErr := buildRefusalFor(t, `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
`)

		assert.ErrorIs(t, cycleErr, ErrModelCycle)
		assert.Equal(t, "model cycle", cycleErr.Error())
		assert.Equal(t, []string{"document#a", "document#b"}, cycleErr.Members)
	})

	t.Run("tuple cycle under an exclusion", func(t *testing.T) {
		t.Parallel()

		cycleErr := buildRefusalFor(t, `model
  schema 1.1
type user
type group
  relations
    define member: [user, group#member] but not blocked
    define blocked: [user, group#member]
`)

		assert.ErrorIs(t, cycleErr, ErrTupleCycle)
		assert.Equal(t, []string{"group#blocked", "group#member"}, cycleErr.Members)
	})

	t.Run("the same cycle is named on every run", func(t *testing.T) {
		t.Parallel()

		// Two independent cycles. The refusal is re-walked in sorted node
		// order, so which one is named does not depend on map iteration.
		dsl := `model
  schema 1.1
type user
type document
  relations
    define a: [user] or b
    define b: [user] or a
    define c: [user] or d
    define d: [user] or c
`
		for range 20 {
			assert.Equal(t, []string{"document#a", "document#b"}, buildRefusalFor(t, dsl).Members)
		}
	})
}
