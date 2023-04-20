package leetcode

// https://leetcode.com/problems/path-sum/
// Runtime: 4 ms
// Memory Usage: 4.5 MB

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	targetSum = targetSum - root.Val

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
