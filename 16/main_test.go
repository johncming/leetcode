package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	out := threeSumClosest(nums, target)
	assert.Equal(t, 2, out)
}
