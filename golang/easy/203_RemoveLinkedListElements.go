package leetcode

// https://leetcode.com/problems/remove-linked-list-elements/
// Runtime: 6 ms
// Memory Usage: 4.6 MB

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	current := head

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	if head.Val == val {
		return head.Next
	}

	return head
}
