package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func Test1(t *testing.T) {
	grid := [][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
	}
	out := islandPerimeter(grid)
	assert.Equal(t, 16, out)
}
