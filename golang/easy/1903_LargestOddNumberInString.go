package leetcode

// https://leetcode.com/problems/largest-odd-number-in-string/
// Runtime: 8 ms
// Memory Usage: 6.5 MB

func largestOddNumber(num string) string {
	if num == "" {
		return ""
	}
	if int(num[len(num)-1])%2 != 0 {
		return num
	}

	nums := []byte(num)

	for i := len(nums) - 1; i > 0; i-- {
		if int(nums[i]-'0')%2 != 0 {
			return string(nums)
		}
		nums = nums[:i]
	}

	if int(nums[0])%2 == 0 {
		return ""
	}
	return string(nums)
}
