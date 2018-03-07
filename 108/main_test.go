package leetcode

import "testing"
import "github.com/davecgh/go-spew/spew"

func Test1(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	spew.Dump(root)
}
