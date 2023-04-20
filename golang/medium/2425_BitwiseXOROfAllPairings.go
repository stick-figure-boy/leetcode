package leetcode

// https://leetcode.com/problems/bitwise-xor-of-all-pairings/
// Runtime: 99 ms
// Memory Usage: 10.8 MB

func xorAllNums(nums1 []int, nums2 []int) int {
	var ans int = 0
	if len(nums1)%2 == 1 {
		for _, n := range nums2 {
			ans ^= n
		}
	}
	if len(nums2)%2 == 1 {
		for _, n := range nums1 {
			ans ^= n
		}
	}
	return ans
}
