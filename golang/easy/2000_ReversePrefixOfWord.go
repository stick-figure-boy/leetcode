package leetcode

// https://leetcode.com/problems/reverse-prefix-of-word/
// Runtime: 0 ms
// Memory Usage: 1.9 MB

func reversePrefix(word string, ch byte) string {
	ans := make([]byte, len(word))
	found := false
	for i := 0; i < len(word); i++ {
		if !found && word[i] == ch {
			found = true
			for l, r := 0, len(word[:i]); l <= r; l, r = l+1, r-1 {
				ans[l], ans[r] = word[r], word[l]
			}
		} else {
			ans[i] = word[i]
		}
	}

	return string(ans)
}
