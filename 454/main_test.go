package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func Test(t *testing.T) {
	type ss struct {
		A        []int
		B        []int
		C        []int
		D        []int
		expected int
	}

	cases := []ss{
		ss{A: []int{1, 2}, B: []int{-2, -1}, C: []int{-1, 2}, D: []int{0, 2}, expected: 2},
		ss{A: []int{0, 1, -1}, B: []int{-1, 1, 0}, C: []int{0, 0, 1}, D: []int{-1, 1, 1}, expected: 17},
	}

	for _, c := range cases {
		actual := fourSumCount(c.A, c.B, c.C, c.D)
		assert.Equal(t, c.expected, actual, "%+v", c)
	}
}
