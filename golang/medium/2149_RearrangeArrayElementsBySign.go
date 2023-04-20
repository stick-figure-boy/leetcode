package leetcode

// https://leetcode.com/problems/rearrange-array-elements-by-sign/
// Runtime: 225 ms
// Memory Usage: 12.5 MB

func rearrangeArray(nums []int) []int {
	ans := make([]int, len(nums))
	i, j := 0, 1
	for _, num := range nums {
		if num > 0 {
			ans[i] = num
			i += 2
		} else {
			ans[j] = num
			j += 2
		}
	}
	return ans
}
