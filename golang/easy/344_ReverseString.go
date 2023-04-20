package leetcode

// https://leetcode.com/problems/reverse-string/
// Runtime: 39 ms
// Memory Usage: 6.6 MB

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		a := s[i]
		b := s[len(s)-1-i]
		s[i] = b
		s[len(s)-1-i] = a
	}
}
