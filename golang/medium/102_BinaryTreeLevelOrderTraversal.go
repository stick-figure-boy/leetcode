package leetcode

// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Runtime: 3 ms
// Memory Usage: 3.1 MB

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type vals struct {
	values []int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	m := map[int]*vals{}
	m[0] = &vals{[]int{root.Val}}
	bfs(root, 0, m)

	ans := make([][]int, len(m))
	for k, v := range m {
		ans[k] = append(ans[k], v.values...)
	}

	return ans
}

func bfs(root *TreeNode, level int, m map[int]*vals) {
	if root.Left == nil && root.Right == nil {
		return
	}

	l := level + 1
	_, ok := m[l]
	if !ok {
		m[l] = &vals{}
	}

	if root.Left != nil {
		m[l].values = append(m[l].values, root.Left.Val)
		bfs(root.Left, l, m)
	}
	if root.Right != nil {
		m[l].values = append(m[l].values, root.Right.Val)
		bfs(root.Right, l, m)
	}
}
