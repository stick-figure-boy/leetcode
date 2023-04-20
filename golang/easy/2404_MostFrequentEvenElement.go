package leetcode

// https://leetcode.com/problems/most-frequent-even-element/
// Runtime: 52 ms
// Memory Usage: 7 MB

func mostFrequentEven(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		if n%2 == 0 {
			m[n] += 1
		}
	}

	if len(m) == 0 {
		return -1
	}

	count := 0
	ans := 0
	for k, v := range m {
		if v > count {
			count = v
			ans = k
		} else if v == count {
			if k < ans {
				ans = k
			}
		}
	}
	return ans
}
