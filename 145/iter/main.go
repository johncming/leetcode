package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack []*TreeNode

func (s stack) Push(v *TreeNode) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, *TreeNode) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Top() *TreeNode {
	l := len(s)
	return s[l-1]
}

func postorderTraversal(root *TreeNode) []int {
	var arr []int

	if root == nil {
		return arr
	}

	s := make(stack, 0)
	s = s.Push(root)
	var top *TreeNode
	for len(s) > 0 {
		s, top = s.Pop()
		arr = append([]int{top.Val}, arr...)
		if top.Left != nil {
			s = s.Push(top.Left)
		}
		if top.Right != nil {
			s = s.Push(top.Right)
		}
	}

	return arr
}
