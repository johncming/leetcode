package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	inorderTraversalMore(root, &res)

	return res
}

func inorderTraversalMore(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorderTraversalMore(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorderTraversalMore(root.Right, arr)
}
