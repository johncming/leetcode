package leetcode

import "testing"
import "github.com/johncming/go-spew/spew"

func Test1(t *testing.T) {

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	res := addTwoNumbers(l1, l2)
	spew.Dump(res)

}

func Test2(t *testing.T) {
	l1 := &ListNode{
		Val: 5,
	}
	l2 := &ListNode{
		Val: 5,
	}
	res := addTwoNumbers(l1, l2)
	spew.Dump(res)

}

func Test3(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
		},
	}
	res := addTwoNumbers(l1, l2)
	spew.Dump(res)

}
