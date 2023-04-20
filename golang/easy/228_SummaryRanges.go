package leetcode

import "fmt"

// https://leetcode.com/problems/summary-ranges/
// Runtime: 1 ms
// Memory Usage: 1.9 MB

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{fmt.Sprint(nums[0])}
	}

	result := [][]int{}
	r := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] == r[len(r)-1]+1 {
			r = append(r, nums[i])
		} else {
			result = append(result, r)
			r = []int{nums[i]}
		}
	}
	result = append(result, r)

	ans := make([]string, len(result))
	for i, re := range result {
		if len(re) == 1 {
			ans[i] = fmt.Sprint(re[0])
		} else {
			ans[i] = fmt.Sprintf("%d->%d", re[0], re[len(re)-1])
		}
	}

	return ans
}
