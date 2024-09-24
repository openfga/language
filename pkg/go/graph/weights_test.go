package graph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWeightsMapToString(t *testing.T) {
	t.Parallel()
	wm := WeightMap{"user": 1, "group": 2}

	require.Equal(t, "weights:[group=2,user=1]", wm.String())

	wm = WeightMap{"user": 1, "group": math.MaxInt}

	require.Equal(t, "weights:[group=+∞,user=1]", wm.String())
}

func TestIntersectionOfWeightsMap(t *testing.T) {
	t.Parallel()
	_, err := IntersectionOfKeys()
	require.ErrorContains(t, err, "no maps given to compute intersection")

	testcases := map[string]struct {
		map1                 WeightMap
		map2                 WeightMap
		expectedIntersection []string
	}{
		`second_map_is_nil`: {
			map1:                 WeightMap{"user": 1},
			expectedIntersection: []string{"user"},
		},
		`first_subset_of_second`: {
			map1:                 WeightMap{"user": 1},
			map2:                 WeightMap{"user": 1, "group": 2},
			expectedIntersection: []string{"user"},
		},
		`second_subset_of_first`: {
			map1:                 WeightMap{"user": 1, "group": 2},
			map2:                 WeightMap{"user": 1},
			expectedIntersection: []string{"user"},
		},
		`no_intersection`: {
			map1:                 WeightMap{"user": 1},
			map2:                 WeightMap{"group": 1},
			expectedIntersection: []string{},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			intersection, err := IntersectionOfKeys(tc.map1, tc.map2)
			require.NoError(t, err)
			require.Equal(t, tc.expectedIntersection, intersection)
		})
	}
}
