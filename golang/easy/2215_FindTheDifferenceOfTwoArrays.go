package leetcode

// https://leetcode.com/problems/find-the-difference-of-two-arrays/
// Runtime: 31 ms
// Memory Usage: 7 MB

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := map[int]int{}
	m2 := map[int]int{}
	for _, n := range nums1 {
		m1[n] = n
	}
	for _, n := range nums2 {
		m2[n] = n
	}

	ans := make([][]int, 2)
	for k := range m1 {
		_, ok := m2[k]
		if !ok {
			ans[0] = append(ans[0], k)
		}
	}
	for k := range m2 {
		_, ok := m1[k]
		if !ok {
			ans[1] = append(ans[1], k)
		}
	}
	return ans
}
