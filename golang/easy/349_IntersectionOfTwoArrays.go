package leetcode

// https://leetcode.com/problems/intersection-of-two-arrays/
// Runtime: 0 ms
// Memory Usage: 3.4 MB

func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	m2 := map[int]int{}

	for _, n := range nums1 {
		m1[n] = n
	}
	for _, n := range nums2 {
		m2[n] = n
	}

	baseMap := map[int]int{}
	comparisonMap := map[int]int{}
	if len(m1) > len(m2) {
		baseMap = m2
		comparisonMap = m1
	} else {
		baseMap = m1
		comparisonMap = m2
	}

	var ans []int
	for k := range baseMap {
		_, ok := comparisonMap[k]
		if ok {
			ans = append(ans, k)
		}
	}
	return ans
}
