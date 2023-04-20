package leetcode

import "fmt"

// https://leetcode.com/problems/alternating-digit-sum/
// Runtime: 1 ms
// Memory Usage: 1.8 MB

func alternateDigitSum(n int) int {
	num := fmt.Sprint(n)
	var ans int
	for i, n := range num {
		if i%2 == 0 {
			ans += int(n - '0')
		} else {
			ans -= int(n - '0')
		}
	}
	return ans
}
