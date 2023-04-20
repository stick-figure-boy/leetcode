package leetcode

// https://leetcode.com/problems/permutations/description/
// Runtime: 4 ms
// Memory Usage: 3.4 MB

func permute(nums []int) [][]int {
	var ans [][]int
	if len(nums) == 1 {
		return append(ans, nums)
	}

	for i, n := range nums {
		for _, v := range permute(remove(nums, i)) {
			ans = append(ans, append([]int{n}, v...))
		}
	}

	return ans
}

func remove(nums []int, i int) []int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	return append(tmp[0:i], tmp[i+1:]...)
}
