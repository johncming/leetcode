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

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	s := make(stack, 0)
	tmp := root
	for {
		for tmp != nil {
			s = s.Push(tmp)
			tmp = tmp.Left
		}
		if len(s) == 0 {
			break
		}
		s, tmp = s.Pop()
		res = append(res, tmp.Val)
		tmp = tmp.Right
	}

	return res
}
