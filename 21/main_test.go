package leetcode

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test1(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	out := mergeTwoLists(l1, l2)
	spew.Dump(out)
}
