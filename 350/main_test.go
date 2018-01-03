package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	out := intersect(nums1, nums2)
	fmt.Printf("%+v\n", out) // output for debug
}

func Test2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{1, 1}
	out := intersect(nums1, nums2)
	fmt.Printf("%+v\n", out) // output for debug
}
