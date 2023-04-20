package leetcode

// https://leetcode.com/problems/delete-greatest-value-in-each-row/
// Runtime: 16 ms
// Memory Usage: 4.6 MB

func deleteGreatestValue(grid [][]int) int {
	for _, g := range grid {
		for i := 0; i < len(g)-1; i++ {
			for j := 0; j < len(g)-i-1; j++ {
				if g[j] < g[j+1] {
					g[j], g[j+1] = g[j+1], g[j]
				}
			}
		}
	}

	var ans int
	for len(grid[0]) != 0 {
		max := grid[0][0]
		for j, g := range grid {
			if max < g[0] {
				max = g[0]
			}
			grid[j] = g[1:]
		}
		ans += max
	}

	return ans
}
