package leetcode

// https://leetcode.com/problems/count-of-matches-in-tournament/
// Runtime: 2 ms
// Memory Usage: 1.8 MB

func numberOfMatches(n int) int {
	var ans int
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			ans += n
		} else {
			n = ((n - 1) / 2)
			ans += n
			n += 1
		}
	}
	return ans
}
