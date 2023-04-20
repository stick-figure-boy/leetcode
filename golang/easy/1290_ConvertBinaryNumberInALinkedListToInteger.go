package leetcode

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// Runtime: 0 ms
// Memory Usage: 1.9 MB

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var b string
	node := head
	for node != nil {
		b += fmt.Sprint(node.Val)
		node = node.Next
	}

	parsed, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return 0
	}

	return int(parsed)
}
