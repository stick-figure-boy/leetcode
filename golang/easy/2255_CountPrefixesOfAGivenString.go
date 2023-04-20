package leetcode

// https://leetcode.com/problems/count-prefixes-of-a-given-string/
// Runtime: 0 ms
// Memory Usage: 3.6 MB

func countPrefixes(words []string, s string) int {
	ans := 0
	for _, w := range words {
		if len(w) <= len(s) && s[:len(w)] == w {
			ans++
		}
	}
	return ans
}
