package leetcode

// https://leetcode.com/problems/possible-bipartition/
// Runtime: 85 ms
// Memory Usage: 7.8 MB

func possibleBipartition(n int, dislikes [][]int) bool {
	var i, a, b int
	adjlist := make([][]int, n+1)
	groups := make([]int, n+1)
	for ; i < len(dislikes); i++ {
		a, b = dislikes[i][0], dislikes[i][1]
		adjlist[a] = append(adjlist[a], b)
		adjlist[b] = append(adjlist[b], a)
	}
	for i = 1; i <= n; i++ {
		if groups[i] != 0 {
			continue
		}
		groups[i] = 1
		if !rec(i, groups, adjlist) {
			return false
		}
	}
	return true
}

func rec(i int, groups []int, adjlist [][]int) bool {
	var j int
	for _, j = range adjlist[i] {
		if groups[j] == groups[i] {
			return false
		} else if groups[j] != 0 {
			continue
		} else if groups[i] == 1 {
			groups[j] = 2
		} else {
			groups[j] = 1
		}

		if !rec(j, groups, adjlist) {
			return false
		}
	}
	return true
}
