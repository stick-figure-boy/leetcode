package leetcode

// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/
// Runtime: 27 ms
// Memory Usage: 6.8 MB

func findMaxK(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n] += 1
	}

	mm := map[int]int{}
	for k := range m {
		if k < 0 {
			mm[k*-1] += 1
		} else {
			mm[k] += 1
		}
	}

	ans := -1
	for k, v := range mm {
		if v >= 2 && ans < k {
			ans = k
		}
	}

	return ans
}
