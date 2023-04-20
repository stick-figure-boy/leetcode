package leetcode

// https://leetcode.com/problems/next-greater-element-i/
// Runtime: 6 ms
// Memory Usage: 2.8 MB

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var ans []int
	for _, n := range nums1 {
		for j, v := range nums2 {
			if v == n {
				if j == len(nums2)-1 {
					ans = append(ans, -1)
					break
				}

				found := false
				target := 0
				for _, c := range nums2[j:] {
					if v < c {
						found = true
						target = c
						break
					}
				}
				if found {
					ans = append(ans, target)
				} else {
					ans = append(ans, -1)
				}
			}
		}
	}
	return ans
}
