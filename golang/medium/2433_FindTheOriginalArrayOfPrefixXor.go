package leetcode

// https://leetcode.com/problems/find-the-original-array-of-prefix-xor/
// Runtime: 115 ms
// Memory Usage: 9.3 MB

func findArray(pref []int) []int {
	if len(pref) <= 1 {
		return pref
	}

	pre := pref[0]
	for i := 1; i < len(pref); i++ {
		pref[i] ^= pre
		pre ^= pref[i]
	}

	return pref
}
