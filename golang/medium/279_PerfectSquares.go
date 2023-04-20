package leetcode

import "fmt"

// https://leetcode.com/problems/perfect-squares/
// Runtime: 78 ms
// Memory Usage: 6.5 MB

func numSquares(n int) int {
	if n <= 0 {
		return 0
	}

	sumMap := make([]int, n+1)
	sumMap[0] = 0
	for i := 1; i <= n; i++ {
		sumMap[i] = sumMap[i-1] + 1
		for j := 1; j*j <= i; j++ {
			sumMap[i] = min(sumMap[i], sumMap[i-j*j]+1)
		}
	}
	fmt.Println(sumMap)
	return sumMap[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
