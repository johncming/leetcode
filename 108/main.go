package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = preOrder(nums[:mid])
	root.Right = preOrder(nums[mid+1:])
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	return preOrder(nums)
}
