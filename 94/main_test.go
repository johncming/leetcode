package leetcode

import "testing"

func Test1(t *testing.T) {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	res := inorderTraversal(root)
	t.Log(res)
}

func Test2(t *testing.T) {
	res := inorderTraversal(nil)
	t.Log(res)
}
