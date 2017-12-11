package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func Test1(t *testing.T) {

	type item struct {
		input    string
		expected int
	}

	cases := []item{
		item{"abcabcbb", 3},
		item{"bbbbb", 1},
		item{"pwwkew", 3},
		item{"dvdf", 3},
	}

	for _, c := range cases {
		actual := lengthOfLongestSubstring(c.input)
		assert.Equal(t, c.expected, actual, "%+v", c)

	}

}
