package leetcode

// https://leetcode.com/problems/unique-morse-code-words/
// Runtime: 3 ms
// Memory Usage: 2.3 MB

func uniqueMorseRepresentations(words []string) int {
	code := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int)
	for _, w := range words {
		var c string
		for _, v := range w {
			c += code[v-'a']
		}
		m[c] = 1
	}
	return len(m)
}
