package leetcode

// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
// Runtime: 0 ms
// Memory Usage: 2.3 MB

func removeOccurrences(s string, part string) string {
	if len(s) < len(part) {
		return s
	}

	found := true
	for found {
		for i := 0; i < len(s)-len(part)+1; i++ {
			if s[i] == part[0] {
				if s[i:i+len(part)] == part {
					found = true
					h := s[:i]
					l := s[i+len(part):]
					s = h + l
					if s == "" || len(s) < len(part) {
						found = false
					}
					break
				}
			} else {
				found = false
			}
		}
	}

	return s
}
