package leetcode

// https://leetcode.com/problems/percentage-of-letter-in-string
// Runtime: 2 ms
// Memory Usage: 1.9 MB

func percentageLetter(s string, letter byte) int {
	b := []byte(s)
	m := map[byte]int{}
	for _, v := range b {
		m[v] += 1
	}

	v, ok := m[letter]
	if !ok {
		return 0
	}

	return int(float64(v) / float64(len(b)) * 100)
}
