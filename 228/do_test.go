package leetcode

import (
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}
	res := summaryRanges(nums)

	t.Logf("input: %#v\n%#v", nums, res)
}

// First failure
// Input:
// [-1]
// Output:
// ["-1->-1"]
// Expected:
// ["-1"]
func Test2(t *testing.T) {
	nums := []int{-1}
	res := summaryRanges(nums)

	t.Logf("input: %#v\n%#v", nums, res)
}
