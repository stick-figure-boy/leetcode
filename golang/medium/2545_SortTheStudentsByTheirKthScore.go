package leetcode

// https://leetcode.com/problems/sort-the-students-by-their-kth-score/
// Runtime: 62 ms
// Memory Usage: 8.1 MB

func sortTheStudents(score [][]int, k int) [][]int {
	m := make(map[int]int, len(score))
	for i, s := range score {
		m[s[k]] = i
	}

	mv := make([]int, len(score))
	count := 0
	for k := range m {
		mv[count] = k
		count++
	}

	for i := 0; i < len(mv)-1; i++ {
		for j := 0; j < len(mv)-i-1; j++ {
			if mv[j] < mv[j+1] {
				mv[j], mv[j+1] = mv[j+1], mv[j]
			}
		}
	}

	ans := make([][]int, len(score))
	count = 0
	for _, v := range mv {
		ans[count] = score[m[v]]
		count++
	}

	return ans
}
