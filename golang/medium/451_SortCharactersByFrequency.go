package leetcode

// https://leetcode.com/problems/sort-characters-by-frequency/
// Runtime: 4 ms
// Memory Usage: 6.7 MB

func frequencySort(s string) string {
	m := map[rune]int{}
	for _, v := range s {
		m[v] += 1
	}

	strs := make([][]rune, len(m))
	counter := 0
	for k, v := range m {
		str := make([]rune, v)
		for i := range str {
			str[i] = k
		}
		strs[counter] = str
		counter++
	}

	for i := 0; i < len(strs)-1; i++ {
		for j := 0; j < len(strs)-i-1; j++ {
			if len(strs[j]) < len(strs[j+1]) {
				strs[j], strs[j+1] = strs[j+1], strs[j]
			}
		}
	}

	var ans []rune
	for _, str := range strs {
		ans = append(ans, str...)
	}

	return string(ans)
}
