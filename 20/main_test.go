package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func Test1(t *testing.T) {
	items := []struct {
		input    string
		expected bool
	}{
		// {"([{}])", true},
		{"(", false},
	}

	for _, item := range items {
		res := isValid(item.input)
		assert.Equal(t, item.expected, res)
	}
}
