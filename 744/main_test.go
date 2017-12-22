package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func Test(t *testing.T) {
	type ss struct {
		input    []byte
		target   byte
		expected byte
	}

	cases := []ss{
		ss{input: []byte{'c', 'f', 'j'}, target: 'a', expected: 'c'},
		ss{input: []byte{'c', 'f', 'j'}, target: 'c', expected: 'f'},
		ss{input: []byte{'c', 'f', 'j'}, target: 'd', expected: 'f'},
		ss{input: []byte{'c', 'f', 'j'}, target: 'g', expected: 'j'},
		ss{input: []byte{'c', 'f', 'j'}, target: 'j', expected: 'c'},
		ss{input: []byte{'c', 'f', 'j'}, target: 'k', expected: 'c'},
	}

	for _, c := range cases {
		actual := nextGreatestLetter(c.input, c.target)
		assert.Equal(t, c.expected, actual, "%+v", c)
	}
}
