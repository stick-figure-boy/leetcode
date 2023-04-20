package leetcode

// https://leetcode.com/problems/isomorphic-strings/
// Runtime: 4 ms
// Memory Usage: 2.7 MB

func isIsomorphic(s string, t string) bool {
	m := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		m[s[i]] = t[i]
	}

	a := map[byte]int{}

	for _, v := range m {
		a[v] += 1
		if a[v] > 1 {
			return false
		}
	}

	var ans []byte
	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if !ok {
			return false
		}
		ans = append(ans, v)
	}
	return string(ans) == t
}
