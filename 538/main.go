package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Right != nil {
		dfs(root.Right)
	}

	root.Val += sum
	sum = root.Val

	if root.Left != nil {
		dfs(root.Left)
	}
}
