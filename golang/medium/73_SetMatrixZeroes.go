package leetcode

// https://leetcode.com/problems/set-matrix-zeroes/
// Runtime: 12 ms
// Memory Usage: 6.4 MB

func setZeroes(matrix [][]int) {
	c := make([][]int, len(matrix))
	for i := range c {
		c[i] = make([]int, len(matrix[0]))
	}
	for i, m := range matrix {
		for j, v := range m {
			c[i][j] = v
		}
	}

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
			if c[i][j] == 0 {
				for k := 0; k < len(c[0]); k++ {
					matrix[i][k] = 0
				}
				for k := 0; k < len(c); k++ {
					matrix[k][j] = 0
				}
			}
		}
	}
}
