package leetcode

// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// Runtime: 13 ms
// Memory Usage: 6.4 MB

func minPartitions(n string) int {
	max := 0
	for i := 0; i < len(n); i++ {
		if int(n[i]) > max {
			max = int(n[i])
		}
	}
	return max - '0'
}
