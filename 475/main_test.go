package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	type ss struct {
		houses   []int
		heaters  []int
		expected int
	}

	cases := []ss{
		ss{houses: []int{31, 100}, heaters: []int{2, 30}, expected: 70},
		ss{houses: []int{1}, heaters: []int{2, 5, 30}, expected: 1},
		ss{houses: []int{1, 2, 3, 4}, heaters: []int{1, 4}, expected: 1},
		ss{houses: []int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}, heaters: []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}, expected: 161834419},
	}

	for _, c := range cases {
		actual := findRadius(c.houses, c.heaters)
		assert.Equal(t, c.expected, actual, "%+v", c)
	}
}
