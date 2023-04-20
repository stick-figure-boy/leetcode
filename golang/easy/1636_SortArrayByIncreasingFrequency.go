package leetcode

// https://leetcode.com/problems/sort-array-by-increasing-frequency/
// Runtime: 7 ms
// Memory Usage: 4 MB

func frequencySort(nums []int) []int {
	m := map[int][]int{}
	for _, n := range nums {
		m[n] = append(m[n], n)
	}

	var sortedNums [][]int
	for _, v := range m {
		sortedNums = append(sortedNums, v)
	}

	for i := 0; i < len(sortedNums)-1; i++ {
		for j := 0; j < len(sortedNums)-i-1; j++ {
			if len(sortedNums[j]) > len(sortedNums[j+1]) {
				sortedNums[j], sortedNums[j+1] = sortedNums[j+1], sortedNums[j]
			}
			if len(sortedNums[j]) == len(sortedNums[j+1]) {
				if sortedNums[j][0] < sortedNums[j+1][0] {
					sortedNums[j], sortedNums[j+1] = sortedNums[j+1], sortedNums[j]
				}
			}
		}
	}

	var ans []int
	for _, num := range sortedNums {
		ans = append(ans, num...)
	}

	return ans
}
