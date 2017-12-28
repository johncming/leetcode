package leetcode

import (
	"testing"
)

func Test1(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	ret := isValidBST(root)
	t.Log(ret)
}
