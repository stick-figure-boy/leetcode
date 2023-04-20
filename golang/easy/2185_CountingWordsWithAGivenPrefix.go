package leetcode

// https://leetcode.com/problems/counting-words-with-a-given-prefix/
// Runtime: 5 ms
// Memory Usage: 3.5 MB

func prefixCount(words []string, pref string) int {
	ans := 0
	for _, w := range words {
		if helper(w, pref) {
			ans++
		}
	}
	return ans
}

func helper(str string, subStr string) bool {
	if len(str) < len(subStr) {
		return false
	}
	if str[0:len(subStr)] == subStr {
		return true
	}
	return false
}
