package leetcode

// https://leetcode.com/problems/first-unique-character-in-a-string/
// Runtime: 32 ms
// Memory Usage: 5.4 MB

func firstUniqChar(s string) int {
	m := map[rune]int{}

	for _, v := range s {
		m[v] += 1
	}

	for i, v := range s {
		v, _ := m[v]
		if v == 1 {
			return i
		}
	}

	return -1
}
