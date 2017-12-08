package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var long, short *ListNode
	var lenA, lenB int

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	for p := l1; p != nil; p = p.Next {
		lenA++
	}
	for p := l2; p != nil; p = p.Next {
		lenB++
	}

	long, short = l1, l2
	if lenB > lenA {
		long, short = l2, l1
	}

	p := long
	q := short
	var prev *ListNode
	var carry bool // 进位
	for q != nil {
		var tmp int
		if carry {
			tmp = p.Val + q.Val + 1
			carry = false
		} else {

			tmp = p.Val + q.Val
		}
		if tmp >= 10 {
			carry = true
			p.Val = tmp - 10
		} else {
			p.Val = tmp
		}
		prev = p
		p = p.Next
		q = q.Next
	}

	for carry {
		// long > short
		if p != nil {
			for p != nil {
				if !carry {
					break
				}

				tmp := p.Val + 1
				if tmp >= 10 {
					carry = true
					p.Val = tmp - 10
				} else {
					carry = false
					p.Val = tmp
				}
				prev = p
				p = p.Next
			}
		} else {
			// long == short
			prev.Next = &ListNode{
				Val: 1,
			}
			carry = false
		}

	}

	return long

}
