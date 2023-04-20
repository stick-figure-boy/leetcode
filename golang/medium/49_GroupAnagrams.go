package leetcode

// https://leetcode.com/problems/group-anagrams/
// Runtime: 18 ms
// Memory Usage: 7.7 MB

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]int{}
	for i, s := range strs {
		sortedS := sort(s)
		hashMap[sortedS] = append(hashMap[sortedS], i)
	}

	var ans [][]string
	for _, v := range hashMap {
		var group []string
		for _, idx := range v {
			group = append(group, strs[idx])
		}
		ans = append(ans, group)
	}
	return ans
}

func sort(str string) string {
	b := []byte(str)
	for i := 0; i < len(b)-1; i++ {
		for j := 0; j < len(b)-i-1; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
	return string(b)
}
