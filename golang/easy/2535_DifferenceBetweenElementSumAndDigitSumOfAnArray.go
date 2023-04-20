package leetcode

import "fmt"

// https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/
// Runtime: 18 ms
// Memory Usage: 6.2 MB

func differenceOfSum(nums []int) int {
	var elementSum int
	var digitSum int

	for _, n := range nums {
		elementSum += n

		str := fmt.Sprint(n)
		for _, s := range str {
			digitSum += int(s - '0')
		}
	}

	return elementSum - digitSum
}
