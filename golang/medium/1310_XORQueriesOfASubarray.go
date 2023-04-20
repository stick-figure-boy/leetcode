package leetcode

// https://leetcode.com/problems/xor-queries-of-a-subarray/
// Runtime: 407 ms
// Memory Usage: 9.8 MB

func xorQueries(arr []int, queries [][]int) []int {
	targets := make([][]int, len(queries))
	for i, q := range queries {
		targets[i] = arr[q[0] : q[1]+1]
	}

	ans := []int{}
	for _, t := range targets {
		result := t[0]
		for i := 1; i < len(t); i++ {
			result = result ^ t[i]
		}
		ans = append(ans, result)
	}
	return ans
}
