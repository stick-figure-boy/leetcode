package leetcode

// https://leetcode.com/problems/find-the-duplicate-number/
// Runtime: 140 ms
// Memory Usage: 12.6 MB

func findDuplicate(nums []int) int {
	m := map[int]int{}
	ans := 0
	for _, n := range nums {
		m[n] += 1
		v, _ := m[n]
		if v >= 2 {
			ans = n
		}
	}
	return ans
}
