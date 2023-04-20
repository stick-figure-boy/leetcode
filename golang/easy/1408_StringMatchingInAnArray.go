package leetcode

// https://leetcode.com/problems/string-matching-in-an-array/
// Runtime: 3 ms
// Memory Usage: 2.3 MB

func stringMatching(words []string) []string {
	ans := make([]string, 0, len(words))
	m := map[string]string{}
	for i, w := range words {
		for j, s := range words {
			if i == j {
				continue
			}
			if helper(s, w) {
				m[w] = w
			}
		}
	}
	for k := range m {
		ans = append(ans, k)
	}
	return ans
}

func helper(str string, subStr string) bool {
	if len(str) < len(subStr) {
		return false
	}

	for i := 0; i < len(subStr); i++ {
		for j := 0; j < len(str)-len(subStr)+1; j++ {
			if str[j] == subStr[i] {
				if str[j:j+len(subStr)] == subStr {
					return true
				}
			}
		}
	}
	return false
}
