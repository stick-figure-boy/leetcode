package leetcode

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Runtime: 10 ms
// Memory Usage: 5.3 MB

func twoSum(numbers []int, target int) []int {
	first := 0
	second := len(numbers) - 1
	for first < second {
		if numbers[first]+numbers[second] == target {
			break
		} else if numbers[first]+numbers[second] < target {
			first += 1
		} else {
			second -= 1
		}
	}
	return []int{first + 1, second + 1}
}
