package leetcode

// https://leetcode.com/problems/pascals-triangle-ii/
// Runtime: 1 ms
// Memory Usage: 2 MB

func getRow(rowIndex int) []int {
	rows := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		rows[i] = make([]int, i+1)
		rows[i][0] = 1
		rows[i][i] = 1

		for j := 1; j < i; j++ {
			rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
		}
	}
	return rows[rowIndex]
}
