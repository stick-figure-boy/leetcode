package leetcode

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
// Runtime: 8 ms
// Memory Usage: 3.1 MB

func smallerNumbersThanCurrent(nums []int) []int {
	ans := make([]int, len(nums))
	for i, n := range nums {
		count := 0
		for j, v := range nums {
			if i == j {
				continue
			}
			if v < n {
				count++
			}
		}
		ans[i] = count
	}
	return ans
}
