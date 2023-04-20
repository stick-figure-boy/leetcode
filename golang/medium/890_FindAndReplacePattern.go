package leetcode

// https://leetcode.com/problems/find-and-replace-pattern/
// Runtime: 2 ms
// Memory Usage: 2.4 MB

func findAndReplacePattern(words []string, pattern string) []string {
	patternMacher := matcher(pattern)
	ans := make([]string, 0)
	for _, w := range words {
		m := matcher(w)

		same := true
		for i, p := range patternMacher {
			if m[i] != p {
				same = false
				break
			}
		}

		if same {
			ans = append(ans, w)
		}
	}
	return ans
}

func matcher(word string) []int {
	wordMap := make(map[rune]int)
	m := make([]int, 0)

	for _, w := range word {
		_, ok := wordMap[w]
		if !ok {
			wordMap[w] = len(wordMap)
		}
		v, _ := wordMap[w]
		m = append(m, v)

	}

	return m
}
