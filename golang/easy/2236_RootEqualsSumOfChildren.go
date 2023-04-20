package leetcode

// https://leetcode.com/problems/root-equals-sum-of-children/
// Runtime: 0 ms
// Memory Usage: 1.9 MB

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return root.Val == root.Left.Val+root.Right.Val
}
