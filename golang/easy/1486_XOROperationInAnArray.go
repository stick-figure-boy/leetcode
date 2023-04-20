package leetcode

// https://leetcode.com/problems/xor-operation-in-an-array
// Runtime: 1 ms
// Memory Usage: 1.9 MB

func xorOperation(n int, start int) int {
	var ans int
	for i := 0; i < n; i++ {
		ans ^= start + 2*i
	}
	return ans
}
