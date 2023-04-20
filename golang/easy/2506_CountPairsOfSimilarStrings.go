package leetcode

// https://leetcode.com/problems/count-pairs-of-similar-strings/
// Runtime: 4 ms
// Memory Usage: 4.1 MB

func similarPairs(words []string) int {
	var ans int
	m := map[int]byte{}
	for i := 0; i < len(words); i++ {
		c := 0
		for _, v := range words[i] {
			c |= (1 << (v - 'a'))
		}
		if a, ok := m[c]; ok {
			ans += int(a)
			m[c]++
		} else {
			m[c] = 1
		}

	}
	return ans
}
