package leetcode

// https://leetcode.com/problems/reverse-words-in-a-string-iii
// Runtime: 9 ms
// Memory Usage: 6.3 MB

func reverseWords(s string) string {
	ans := make([]byte, len(s))
	left := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			for l, r := left, len(s[:i])-1; l <= r; l, r = l+1, r-1 {
				ans[l], ans[r] = s[r], s[l]
			}
			ans[i] = ' '
			left = i + 1
		}
	}

	for l, r := left, len(s)-1; l <= r; l, r = l+1, r-1 {
		ans[l], ans[r] = s[r], s[l]
	}

	return string(ans)
}
