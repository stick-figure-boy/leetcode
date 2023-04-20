package leetcode

// https://leetcode.com/problems/words-within-two-edits-of-dictionary/
// Runtime: 4 ms
// Memory Usage: 3.2 MB

func twoEditWords(queries []string, dictionary []string) []string {
	var ans []string
	for _, query := range queries {
		for _, dict := range dictionary {
			if query == dict {
				ans = append(ans, query)
				break
			}
			count := 0
			for i := 0; i < len(query); i++ {
				if query[i] != dict[i] {
					count++
				}
			}
			if count < 3 {
				ans = append(ans, query)
				break
			}
		}
	}

	return ans
}
