package leetcode

// https://leetcode.com/problems/three-divisors/
// Runtime: 0 ms
// Memory Usage: 1.8 MB

func isThree(n int) bool {
	number := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			number++
		}
		if number > 3 {
			break
		}
	}
	if number == 3 {
		return true
	}
	return false
}
