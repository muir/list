package list_test

import (
	"testing"

	"github.com/muir/list"

	"github.com/stretchr/testify/assert"
)

func TestFilterEmptySlices(t *testing.T) {
	cases := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name:  "nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "empty",
			input: [][]int{},
			want:  [][]int{},
		},
		{
			name:  "no empty",
			input: [][]int{{1}, {2, 3}, {4}},
			want:  [][]int{{1}, {2, 3}, {4}},
		},
		{
			name:  "single empty middle",
			input: [][]int{{1}, {}, {2}},
			want:  [][]int{{1}, {2}},
		},
		{
			name:  "multiple empties",
			input: [][]int{{}, {1}, {}, {2}, {3}, {}},
			want:  [][]int{{1}, {2}, {3}},
		},
		{
			name:  "all empty",
			input: [][]int{{}, {}},
			want:  [][]int{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := list.FilterEmptySlices(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
