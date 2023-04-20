package leetcode

// https://leetcode.com/problems/divide-array-into-equal-pairs/
// Runtime: 14 ms
// Memory Usage: 5.7 MB

func divideArray(nums []int) bool {
	m := map[int]int{}
	for _, n := range nums {
		m[n] += 1
	}

	ans := len(nums) / 2
	for _, v := range m {
		ans = ans - v/2
	}

	return ans == 0
}
