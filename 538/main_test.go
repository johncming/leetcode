package leetcode

import "testing"
import "github.com/davecgh/go-spew/spew"

func Test1(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 13},
	}

	convertBST(root)
	spew.Dump(root)
}

func Test2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  0,
			Left: &TreeNode{Val: -2},
		},
		Right: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 3},
		},
	}

	convertBST(root)
	spew.Dump(root)
}
