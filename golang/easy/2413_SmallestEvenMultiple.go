package leetcode

// https://leetcode.com/problems/smallest-even-multiple/
// Runtime: 0 ms
// Memory Usage: 1.8 MB

func smallestEvenMultiple(n int) int {
	if n&1 == 0 {
		return n
	}
	return n * 2
}
