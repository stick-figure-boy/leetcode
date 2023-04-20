package leetcode

// https://leetcode.com/problems/find-closest-number-to-zero/
// Runtime: 25 ms
// Memory Usage: 6.6 MB

func findClosestNumber(nums []int) int {
	var min int
	var ans int

	for i, n := range nums {
		d := n
		if n < 0 {
			d = n * -1
		}

		if i == 0 {
			min = d
			ans = n
			continue
		}

		if d == min && n > ans {
			min = d
			ans = n
			continue
		}

		if d < min {
			min = d
			ans = n
		}
	}

	return ans
}
