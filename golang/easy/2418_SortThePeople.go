package leetcode

// https://leetcode.com/problems/sort-the-people/
// Runtime: 35 ms
// Memory Usage: 6.8 MB

func sortPeople(names []string, heights []int) []string {
	m := make(map[int]string, len(names))
	for i, n := range names {
		m[heights[i]] = n
	}

	talls := make([]int, len(names))
	counter := 0
	for k := range m {
		talls[counter] = k
		counter++
	}

	for i := 0; i < len(names)-1; i++ {
		for j := 0; j < len(names)-i-1; j++ {
			if talls[j] < talls[j+1] {
				talls[j], talls[j+1] = talls[j+1], talls[j]
			}
		}
	}

	ans := make([]string, len(names))
	for i, t := range talls {
		ans[i] = m[t]
	}

	return ans
}
