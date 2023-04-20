package leetcode

// https://leetcode.com/problems/invert-binary-tree/
// Runtime: 0 ms
// Memory Usage: 2.1 MB

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	bfs(root, 0)

	return root
}

func bfs(node *TreeNode, level int) {
	if node == nil {
		return
	}

	left := node.Left
	node.Left = node.Right
	node.Right = left

	bfs(node.Left, level+1)
	bfs(node.Right, level+1)
}
