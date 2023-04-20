package leetcode

import "fmt"

// https://leetcode.com/problems/count-number-of-distinct-integers-after-reverse-operations/
// Runtime: 787 ms
// Memory Usage: 40.6 MB

func countDistinctIntegers(nums []int) int {
	m := map[string]int{}
	ns := make([]string, len(nums))
	reverse := make([]string, len(nums))
	for i, n := range nums {
		ns[i] = fmt.Sprint(n)

		runes := []rune(fmt.Sprint(n))
		for runes[len(runes)-1] == '0' {
			runes = runes[:len(runes)-1]
		}
		for ri, j := 0, len(runes)-1; ri < j; ri, j = ri+1, j-1 {
			runes[ri], runes[j] = runes[j], runes[ri]
		}
		reverse[i] = string(runes)
	}
	ns = append(ns, reverse...)
	for _, v := range ns {
		m[v] += 1
	}
	return len(m)
}
