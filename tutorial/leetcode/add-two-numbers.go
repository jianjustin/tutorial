package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail, head *ListNode
	var m = 0
	for l1 != nil || l2 != nil {
		var n1, n2 = 0, 0

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		var n = (n1 + n2 + m) % 10
		m = (n1 + n2 + m) / 10

		if head == nil {
			head = &ListNode{Val: n}
			tail = head
		} else {
			tail.Next = &ListNode{Val: n}
			tail = tail.Next
		}
	}

	if m > 0 {
		tail.Next = &ListNode{Val: m}
	}

	return head
}
