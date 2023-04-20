package leetcode

// https://leetcode.com/problems/goal-parser-interpretation/
// Runtime: 0 ms
// Memory Usage: 1.9 MB

func interpret(command string) string {
	var ans []byte
	for i := 0; i < len(command); i++ {
		if command[i] == '(' {
			if command[i+1] == ')' {
				ans = append(ans, 'o')
				i++
			} else {
				ans = append(ans, []byte{'a', 'l'}...)
				i += 3
			}
		} else {
			ans = append(ans, 'G')
		}
	}
	return string(ans)
}
