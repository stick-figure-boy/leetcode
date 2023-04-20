package leetcode

import "fmt"

// https://leetcode.com/problems/largest-3-same-digit-number-in-string/
// Runtime: 2 ms
// Memory Usage: 2 MB

func largestGoodInteger(num string) string {
	ans := -1
	for i := 0; i < len(num)-2; i++ {
		if num[i+1] == num[i] && num[i+2] == num[i] && int(num[i]-'0') > ans {
			ans = int(num[i] - '0')
		}
	}

	if ans == -1 {
		return ""
	}

	return fmt.Sprintf("%d%d%d", ans, ans, ans)
}
