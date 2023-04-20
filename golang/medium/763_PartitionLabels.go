package leetcode

// https://leetcode.com/problems/partition-labels/
// Runtime: 0 ms
// Memory Usage: 2.1 MB

func partitionLabels(s string) []int {
	lastIdx := make(map[rune]int)

	for idx, c := range s {
		lastIdx[c] = idx
	}

	start := 0
	end := 0
	ans := []int{}
	for i := range s {
		if end < lastIdx[rune(s[i])] {
			end = lastIdx[rune(s[i])]
		}
		if i == end {
			ans = append(ans, i-start+1)
			start = i + 1
		}
	}
	return ans
}
