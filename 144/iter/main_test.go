package leetcode

import "testing"

func Test1(t *testing.T) {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	res := preorderTraversal(root)
	t.Log(res)
}

func Test2(t *testing.T) {
	res := preorderTraversal(nil)
	t.Log(res)
}

func Test3(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	res := preorderTraversal(root)
	t.Log(res)
}
