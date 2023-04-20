package leetcode

// https://leetcode.com/problems/partition-array-according-to-given-pivot/
// Runtime: 302 ms
// Memory Usage: 12.9 MB

func pivotArray(nums []int, pivot int) []int {
	small := []int{}
	big := []int{}
	pivots := []int{}
	for _, n := range nums {
		if n == pivot {
			pivots = append(pivots, n)
			continue
		}
		if n < pivot {
			small = append(small, n)
		} else {
			big = append(big, n)
		}
	}
	ans := append(small, pivots...)

	return append(ans, big...)
}
