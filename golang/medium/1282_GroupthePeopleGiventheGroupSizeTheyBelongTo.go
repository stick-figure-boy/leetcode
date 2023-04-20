package leetcode

// https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
// Runtime: 4 ms
// Memory Usage: 5.1 MB

func groupThePeople(groupSizes []int) [][]int {
	m := map[int][]int{}
	for i, g := range groupSizes {
		m[g] = append(m[g], i)
	}

	ans := make([][]int, 0, len(groupSizes))
	for key, val := range m {
		for i := 0; i < len(val)/key; i++ {
			ans = append(ans, val[i*key:i*key+key])
		}
	}

	return ans
}
