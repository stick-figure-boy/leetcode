package leetcode

import "fmt"

// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
// Runtime: 1 ms
// Memory Usage: 1.9 MB

func subtractProductAndSum(n int) int {
	s := fmt.Sprint(n)
	pro, sum := int(s[0]-'0'), int(s[0]-'0')
	for i := 1; i < len(s); i++ {
		pro = pro * (int(s[i] - '0'))
		sum = sum + (int(s[i] - '0'))
	}
	return int(pro) - int(sum)
}
