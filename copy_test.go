package list_test

import (
	"testing"

	"github.com/muir/list"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	ints := []int{3, 4, 5}
	c := list.Copy(ints)
	assert.Equal(t, ints, c)
	ints[0] = 1
	assert.NotEqual(t, ints, c)
	var n []string
	assert.Equal(t, n, list.Copy(n))
}
