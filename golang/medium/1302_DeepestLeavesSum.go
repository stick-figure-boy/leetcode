package leetcode

// https://leetcode.com/problems/deepest-leaves-sum/
// Runtime: 43 ms
// Memory Usage: 7.7 MB

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	_, x := helper(root, 0, 0, 0)
	return x
}

func helper(root *TreeNode, level, highest, sum int) (int, int) {
	if root.Left != nil {
		highest, sum = helper(root.Left, level+1, highest, sum)
	}

	if root.Right != nil {
		highest, sum = helper(root.Right, level+1, highest, sum)
	}

	if root.Left == nil && root.Right == nil {
		if level > highest {
			highest = level
			sum = root.Val
		} else if level == highest {
			sum += root.Val
		}
	}

	return highest, sum
}
