package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func Test(t *testing.T) {
	type item struct {
		input    string
		expected int
	}

	cases := []item{
		item{"XXXVI", 36},
		item{"MMXII", 2012},
		item{"MCMXCVI", 1996},
		item{"MDLXX", 1570},
	}

	for _, c := range cases {
		actual := romanToInt(c.input)
		assert.Equal(t, c.expected, actual, "%+v", c)
	}
}
