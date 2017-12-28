package leetcode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTMore(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTMore(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValidBSTMore(root.Left, min, root.Val) && isValidBSTMore(root.Right, root.Val, max)
}
