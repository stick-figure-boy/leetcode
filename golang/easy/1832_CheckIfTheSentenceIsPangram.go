package leetcode

// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// Runtime: 0 ms
// Memory Usage: 2 MB

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}

	m := map[rune]int{}
	for _, s := range sentence {
		m[s] += 1
	}

	return len(m) > 25
}
