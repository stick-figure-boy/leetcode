package leetcode

import "fmt"

// https://leetcode.com/problems/separate-the-digits-in-an-array/
// Runtime: 6 ms
// Memory Usage: 5.4 MB

func separateDigits(nums []int) []int {
	ans := make([]int, 0)
	for _, num := range nums {
		n := fmt.Sprint(num)
		if len(n) == 1 {
			ans = append(ans, num)
		} else {
			for _, v := range n {
				ans = append(ans, int(v-'0'))
			}
		}
	}
	return ans
}
