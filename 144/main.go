package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var arr []int

	preorderTraversalMore(root, &arr)

	return arr
}

func preorderTraversalMore(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	*arr = append(*arr, root.Val)

	if root.Left != nil {
		preorderTraversalMore(root.Left, arr)
	}

	if root.Right != nil {
		preorderTraversalMore(root.Right, arr)
	}

}
