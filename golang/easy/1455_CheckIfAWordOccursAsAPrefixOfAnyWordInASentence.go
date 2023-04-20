package leetcode

// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
// Runtime: 1 ms
// Memory Usage: 2 MB

func isPrefixOfWord(sentence string, searchWord string) int {
	var words []string
	var word []rune
	for _, s := range sentence {
		if s != ' ' {
			word = append(word, s)
		} else {
			words = append(words, string(word))
			word = []rune{}
		}
	}
	if len(word) > 0 {
		words = append(words, string(word))
	}

	for i, w := range words {
		if helper(w, searchWord) {
			return i + 1
		}
	}

	return -1
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
