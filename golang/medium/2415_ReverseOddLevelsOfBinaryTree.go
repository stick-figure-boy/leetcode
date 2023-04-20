package leetcode

// https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/
// Runtime: 80 ms
// Memory Usage: 8.4 MB

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type vals struct {
	values []int
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	valueMap := map[int]*vals{}
	bfs(root.Left, root.Right, 1, valueMap)
	reverse(root.Left, root.Right, 1, valueMap)

	return root
}

func bfs(left, right *TreeNode, level int, valueMap map[int]*vals) {
	if left == nil {
		return
	}

	if level%2 != 0 {
		_, ok := valueMap[level]
		if ok {
			valueMap[level].values = append(valueMap[level].values, []int{left.Val, right.Val}...)
		} else {
			valueMap[level] = &vals{values: []int{left.Val, right.Val}}
		}

	}
	bfs(left.Left, left.Right, level+1, valueMap)
	bfs(right.Left, right.Right, level+1, valueMap)
}

func reverse(left, right *TreeNode, level int, valueMap map[int]*vals) {
	if left == nil {
		return
	}

	if level%2 != 0 {
		values := valueMap[level].values
		left.Val = valueMap[level].values[len(values)-1]
		right.Val = valueMap[level].values[len(values)-2]
		valueMap[level].values = values[:len(values)-2]
	}
	reverse(left.Left, left.Right, level+1, valueMap)
	reverse(right.Left, right.Right, level+1, valueMap)
}
