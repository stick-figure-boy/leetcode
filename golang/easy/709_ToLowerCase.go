package leetcode

// https://leetcode.com/problems/to-lower-case/
// Runtime: 2 ms
// Memory Usage: 1.9 MB

func toLowerCase(s string) string {
	results := make([]rune, len(s))
	for i, v := range s {
		if v >= 'A' && v <= 'Z' {
			results[i] = v + 32
		} else {
			results[i] = v
		}
	}

	return string(results)
}
