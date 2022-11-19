package list_test

import (
	"fmt"
	"testing"

	"github.com/muir/list"

	"github.com/stretchr/testify/assert"
)

func TestSplice(t *testing.T) {
	cases := []struct {
		name    string
		input   []int
		start   int
		end     int
		replace []int
		want    []int
	}{
		{
			name:    "append",
			input:   []int{1, 2, 3},
			start:   3,
			end:     5,
			replace: []int{4, 5},
			want:    []int{1, 2, 3, 4, 5},
		},
		{
			name:    "append w/gap",
			input:   []int{1, 2, 3},
			start:   4,
			end:     6,
			replace: []int{4, 5},
			want:    []int{1, 2, 3, 0, 4, 5},
		},
		{
			name:    "shrink within",
			input:   []int{1, 2, 3, 4, 5},
			start:   1,
			end:     4,
			replace: []int{10},
			want:    []int{1, 10, 5},
		},
		{
			name:    "shrink within",
			input:   []int{1, 2, 3, 4, 5},
			start:   1,
			end:     3,
			replace: []int{10, 11, 12},
			want:    []int{1, 10, 11, 12, 4, 5},
		},
		{
			name:    "replace across end",
			input:   []int{1, 2, 3, 4},
			start:   2,
			end:     6,
			replace: []int{10, 11, 12, 13},
			want:    []int{1, 2, 10, 11, 12, 13},
		},
		{
			name:    "shrink across end",
			input:   []int{1, 2, 3, 4},
			start:   2,
			end:     6,
			replace: []int{10, 11},
			want:    []int{1, 2, 10, 11},
		},
		{
			name:    "expand across end",
			input:   []int{1, 2, 3, 4},
			start:   2,
			end:     6,
			replace: []int{10, 11, 12, 13, 14, 15},
			want:    []int{1, 2, 10, 11, 12, 13, 14, 15},
		},
		{
			name:    "shrink beyond end",
			input:   []int{1, 2, 3},
			start:   4,
			end:     14,
			replace: []int{10, 11},
			want:    []int{1, 2, 3, 0, 10, 11},
		},
		{
			name:    "insert middle",
			input:   []int{1, 2, 3},
			start:   1,
			end:     1,
			replace: []int{10, 11},
			want:    []int{1, 10, 11, 2, 3},
		},
		{
			name:    "insert start",
			input:   []int{1, 2, 3},
			start:   0,
			end:     0,
			replace: []int{10, 11},
			want:    []int{10, 11, 1, 2, 3},
		},
		{
			name:    "insert end",
			input:   []int{1, 2, 3},
			start:   3,
			end:     3,
			replace: []int{10, 11},
			want:    []int{1, 2, 3, 10, 11},
		},
		{
			name:    "delete middle",
			input:   []int{1, 2, 3, 4},
			start:   1,
			end:     3,
			replace: []int{},
			want:    []int{1, 4},
		},
		{
			name:    "delete start",
			input:   []int{1, 2, 3, 4},
			start:   0,
			end:     2,
			replace: []int{},
			want:    []int{3, 4},
		},
		{
			name:    "delete end",
			input:   []int{1, 2, 3, 4},
			start:   2,
			end:     4,
			replace: []int{},
			want:    []int{1, 2},
		},
	}

	for _, tc := range cases {
		for _, extend := range []bool{false, true} {
			t.Run(tc.name, func(t *testing.T) {
				t.Run(fmt.Sprintf("extend %v", extend), func(t *testing.T) {
					getInput := func() []int {
						var input []int
						if extend {
							input = make([]int, len(tc.input), len(tc.input)+40)
						} else {
							input = make([]int, len(tc.input))
						}
						copy(input, tc.input)
						return input
					}
					t.Log("splice beyond")
					got := list.SpliceBeyond(getInput(), tc.start, tc.end, tc.replace...)
					assert.Equal(t, tc.want, got)

					if len(tc.replace) == tc.end-tc.start {
						t.Log("replace beyond")
						got := list.ReplaceBeyond(getInput(), tc.start, tc.replace...)
						assert.Equal(t, tc.want, got)

						if tc.start <= len(tc.input) {
							t.Log("replace")
							got := list.Replace(getInput(), tc.start, tc.replace...)
							assert.Equal(t, tc.want, got)
						}
					}
					if tc.end <= len(tc.input) {
						t.Log("splice")
						got := list.Splice(getInput(), tc.start, tc.end, tc.replace...)
						assert.Equal(t, tc.want, got)
					}
				})
			})
		}
	}
}

func TestPanicSplice(t *testing.T) {
	assert.Panics(t, func() {
		_ = list.Replace([]int{1, 2, 3}, 20, 4, 5)
	})
	assert.Panics(t, func() {
		_ = list.Splice([]int{1, 2, 3}, 2, 5, 10, 11, 12)
	})
	assert.Panics(t, func() {
		_ = list.Splice([]int{1, 2, 3}, 2, 1, 10)
	})
}
