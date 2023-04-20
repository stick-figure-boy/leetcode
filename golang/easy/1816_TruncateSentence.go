package leetcode

// https://leetcode.com/problems/truncate-sentence/
// Runtime: 2 ms
// Memory Usage: 2 MB

func truncateSentence(s string, k int) string {
	lastIdx := 0
	found := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			found++
		}
		if found == k {
			lastIdx = i
			break
		}
	}

	if found < k {
		return s
	}

	return string([]byte(s)[0:lastIdx])
}
