package leetcode

import "fmt"

// https://leetcode.com/problems/fizz-buzz/
// Runtime: 3 ms
// Memory Usage: 3.4 MB

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i < n+1; i++ {
		if i%15 == 0 {
			ans[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			ans[i-1] = "Fizz"
		} else if i%5 == 0 {
			ans[i-1] = "Buzz"
		} else {
			ans[i-1] = fmt.Sprint(i)
		}
	}
	return ans
}
