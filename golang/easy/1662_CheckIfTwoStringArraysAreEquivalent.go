package leetcode

// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent
// Runtime: 0 ms
// Memory Usage: 2.6 MB

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var w1 string
	for _, w := range word1 {
		w1 += w
	}
	var w2 string
	for _, w := range word2 {
		w2 += w
	}

	return w1 == w2
}
