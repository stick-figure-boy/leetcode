package leetcode

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Runtime: 6 ms
// Memory Usage: 4.1 MB

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
