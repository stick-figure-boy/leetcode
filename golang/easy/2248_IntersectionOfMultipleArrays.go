package leetcode

// https://leetcode.com/problems/intersection-of-multiple-arrays/
// Runtime: 11 ms
// Memory Usage: 3.8 MB

func intersection(nums [][]int) []int {
	base := nums[0]

	var ans []int
	for _, b := range base {
		count := 1
		for _, num := range nums[1:] {
			for _, n := range num {
				if n == b {
					count++
					break
				}
			}
		}
		if count >= len(nums) {
			ans = append(ans, b)
		}
	}
	return sort(ans)
}

func sort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
