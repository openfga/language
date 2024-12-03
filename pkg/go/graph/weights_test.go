package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWeightsMapToString(t *testing.T) {
	t.Parallel()
	wm := WeightMap{"user": 1, "group": 2}

	require.Equal(t, "weights:[group=2,user=1]", wm.String())

	wm = WeightMap{"user": 1, "group": Infinite}

	require.Equal(t, "weights:[group=+∞,user=1]", wm.String())
}
