package leetcode

// https://leetcode.com/problems/pascals-triangle/
// Runtime: 1 ms
// Memory Usage: 2.2 MB

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1

		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
