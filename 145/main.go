package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var arr []int

	postorderTraversalMore(root, &arr)

	return arr
}

func postorderTraversalMore(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		postorderTraversalMore(root.Left, arr)
	}
	if root.Right != nil {
		postorderTraversalMore(root.Right, arr)
	}

	*arr = append(*arr, root.Val)
}
