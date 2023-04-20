package leetcode

// https://leetcode.com/problems/find-triangular-sum-of-an-array/
// Runtime: 297 ms
// Memory Usage: 19.2 MB

func triangularSum(nums []int) int {
	count := len(nums) - 1
	current := nums
	for count > 0 {
		sum := []int{}
		for i := 0; i < len(current)-1; i++ {
			v := current[i] + current[i+1]
			sum = append(sum, v%10)
		}
		current = sum
		count--
	}

	return current[0]
}
