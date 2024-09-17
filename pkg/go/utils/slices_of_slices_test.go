package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSlicesOfSlices_Less(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input SlicesOfSlices
		want  bool
	}{
		{
			name: "first_slice_longer",
			input: SlicesOfSlices{
				{"a", "b"},
				{},
			},
			want: false,
		},
		{
			name: "second_slice_longer",
			input: SlicesOfSlices{
				{},
				{"a", "b"},
			},
			want: true,
		},
		{
			name: "equal_length_first_slice_smaller_element",
			input: SlicesOfSlices{
				{"a", "b"},
				{"x", "b"},
			},
			want: true,
		},
		{
			name: "equal_length_first_slice_larger_element",
			input: SlicesOfSlices{
				{"x", "b"},
				{"a", "b"},
			},
			want: false,
		},
		{
			name: "equal_length_first_slice_smaller_last_element",
			input: SlicesOfSlices{
				{"a", "b"},
				{"a", "c"},
			},
			want: true,
		},
		{
			name: "equal_length_second_slice_smaller_last_element",
			input: SlicesOfSlices{
				{"a", "c"},
				{"a", "b"},
			},
			want: false,
		},
		{
			name: "equal_everything",
			input: SlicesOfSlices{
				{"a", "b"},
				{"a", "b"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			output := tt.input.Less(0, 1)
			require.Equal(t, tt.want, output)
		})
	}
}
