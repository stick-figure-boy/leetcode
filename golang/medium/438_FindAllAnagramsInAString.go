package leetcode

// https://leetcode.com/problems/find-all-anagrams-in-a-string/
// Runtime: 3970 ms
// Memory Usage: 7.1 MB

func findAnagrams(s string, p string) []int {
	if s == p {
		return []int{0}
	}

	lenP := len(p)
	mapP := map[byte]int{}
	for i := 0; i < lenP; i++ {
		mapP[p[i]] += 1
	}

	var ans []int
	for i := 0; i < len(s)-lenP+1; i++ {
		target := s[i : i+lenP]
		m := map[byte]int{}
		for j := 0; j < len(target); j++ {
			m[target[j]] += 1
		}

		result := true
		for key, val := range mapP {
			v, ok := m[key]
			if !ok || v != val {
				result = false
			}
		}
		if result {
			ans = append(ans, i)
		}
	}

	return ans
}
