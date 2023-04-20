package leetcode

// https://leetcode.com/problems/intersection-of-two-arrays-ii/
// Runtime: 5 ms
// Memory Usage: 2.7 MB

func intersect(nums1 []int, nums2 []int) []int {
	var base []int
	var comparison []int

	if len(nums1) > len(nums2) {
		base = nums2
		comparison = nums1
	} else {
		base = nums1
		comparison = nums2
	}

	var ans []int
	for _, n := range base {
		for ci, c := range comparison {
			if c == n {
				ans = append(ans, n)
				comparison = comparison[:ci+copy(comparison[ci:], comparison[ci+1:])]
				break
			}
		}
	}
	return ans
}
