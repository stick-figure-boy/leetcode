package leetcode

// https://leetcode.com/problems/find-smallest-letter-greater-than-target
// Runtime: 0 ms
// Memory Usage: 2.6 MB

func nextGreatestLetter(letters []byte, target byte) byte {
	var smallest byte
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			if smallest == 0 {
				smallest = letters[i]
			} else {
				if letters[i] < smallest {
					smallest = letters[i]
				}
			}
		}
	}

	if smallest == 0 {
		return letters[0]
	}
	return smallest
}
