package leetcode

// https://leetcode.com/problems/third-maximum-number/
// Runtime: 96 ms
// Memory Usage: 4.9 MB

func thirdMax(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n] = n
	}

	var distinctNums []int
	for k := range m {
		distinctNums = append(distinctNums, k)
	}
	for i := 0; i < len(distinctNums)-1; i++ {
		for j := 0; j < len(distinctNums)-i-1; j++ {
			if distinctNums[j] < distinctNums[j+1] {
				distinctNums[j], distinctNums[j+1] = distinctNums[j+1], distinctNums[j]
			}
		}
	}
	if len(distinctNums) < 3 {
		return distinctNums[0]
	} else {
		return distinctNums[2]
	}
}
