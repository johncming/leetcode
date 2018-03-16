package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func Test1(t *testing.T) {
	items := []struct {
		s        string
		n        int
		expected string
	}{
		{
			s:        "PAYPALISHIRING",
			n:        3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			s:        "AB",
			n:        1,
			expected: "AB",
		},
	}

	for _, item := range items {
		actual := convert(item.s, item.n)
		assert.Equal(t, item.expected, actual)
	}
}
