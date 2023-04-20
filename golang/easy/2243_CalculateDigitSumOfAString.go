package leetcode

import "fmt"

// https://leetcode.com/problems/calculate-digit-sum-of-a-string/
// Runtime: 2 ms
// Memory Usage: 1.9 MB

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	} else {
		var next string
		var sum int
		for i, b := range s {
			sum += int(b - '0')
			if (i+1)%k == 0 || i == len(s)-1 {
				next += fmt.Sprint(sum)
				sum = 0
			}
		}
		if len(s) < k {
			return next
		} else {
			return digitSum(next, k)
		}
	}
}
