package leetcode

// https://leetcode.com/problems/replace-all-digits-with-characters/
// Runtime: 0 ms
// Memory Usage: 1.9 MB

func replaceDigits(s string) string {
	ans := make([]rune, len(s))
	for i, v := range s {
		if v >= 'a' && v <= 'z' {
			ans[i] = v
		} else {
			b := ans[i-1] + (v - '0')
			ans[i] = b
		}
	}

	return string(ans)
}
