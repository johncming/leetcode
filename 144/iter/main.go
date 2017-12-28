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

func preorderTraversal(root *TreeNode) []int {
	var arr []int
	s := make(stack, 0)
	tmp := root
	for tmp != nil {
		arr = append(arr, tmp.Val)
		if tmp.Right != nil {
			s = s.Push(tmp.Right)
		}

		tmp = tmp.Left
		if tmp == nil && len(s) != 0 {
			s, tmp = s.Pop()
		}
	}

	return arr
}
