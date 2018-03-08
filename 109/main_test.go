package leetcode

import "testing"
import "github.com/davecgh/go-spew/spew"

func Test1(t *testing.T) {
	head := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	}

	root := sortedListToBST(head)
	spew.Dump(root)
}
