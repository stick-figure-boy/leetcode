package leetcode

// https://leetcode.com/problems/excel-sheet-column-number/
// Runtime: 7 ms
// Memory Usage: 2 MB

func titleToNumber(columnTitle string) int {
	ans := 0
	for _, s := range columnTitle {
		ans = ans*26 + (int(s) - int('A') + 1)
	}
	return ans
}
