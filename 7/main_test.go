package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func Test(t *testing.T) {
	type item struct {
		input    int
		expected int
	}

	cases := []item{
		item{123, 321},
		item{-123, -321},
		item{120, 21},
		item{1534236469, 0},
	}

	for _, c := range cases {
		actual := reverse(c.input)
		assert.Equal(t, c.expected, actual, "%+v", c)
	}
}
