package leetcode

// https://leetcode.com/problems/top-k-frequent-elements/
// Runtime: 23 ms
// Memory Usage: 5.4 MB

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, n := range nums {
		m[n] += 1
	}

	frequents := make([][]int, len(m))
	counter := 0
	for k, v := range m {
		frequents[counter] = []int{k, v}
		counter++
	}

	for i := 0; i < len(frequents)-1; i++ {
		for j := 0; j < len(frequents)-i-1; j++ {
			if frequents[j][1] < frequents[j+1][1] {
				frequents[j], frequents[j+1] = frequents[j+1], frequents[j]
			}
		}
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = frequents[i][0]
	}

	return ans
}
