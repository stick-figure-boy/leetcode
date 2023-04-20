package leetcode

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Runtime: 3 ms
// Memory Usage: 2.3 MB

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}
	ans := make([]bool, len(candies))
	for i, c := range candies {
		if c+extraCandies >= max {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
