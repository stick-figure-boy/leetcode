package leetcode

// https://leetcode.com/problems/valid-anagram/
// Runtime: 9 ms
// Memory Usage: 2.7 MB

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[byte]int{}
	tMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]] += 1
		tMap[t[i]] += 1
	}

	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true
}
