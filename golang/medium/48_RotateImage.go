package leetcode

// https://leetcode.com/problems/rotate-image/
// Runtime: 2 ms
// Memory Usage: 2.2 MB

func rotate(matrix [][]int) {
	length := len(matrix[0])
	c := make([][]int, len(matrix))
	for i := range c {
		c[i] = make([]int, length)
	}
	for i, m := range matrix {
		for j, v := range m {
			c[i][j] = v
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < length; j++ {
			matrix[i][j] = c[length-1-j][i]
		}
	}
}
