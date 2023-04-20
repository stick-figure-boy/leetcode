package leetcode

// https://leetcode.com/problems/count-vowel-strings-in-ranges/
// Runtime: 110 ms
// Memory Usage: 19.4 MB

func vowelStrings(words []string, queries [][]int) []int {
	counts := make([]int32, len(words)+1)
	for i, word := range words {
		counts[i+1] = counts[i]
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			counts[i+1]++
		}
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = int(counts[query[1]+1] - counts[query[0]])
	}
	return ans
}

func isVowel(v byte) bool {
	switch v {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
