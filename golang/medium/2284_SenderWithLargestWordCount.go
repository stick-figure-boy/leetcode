package leetcode

import "sort"

// https://leetcode.com/problems/sender-with-largest-word-count/
// Runtime: 82 ms
// Memory Usage: 9.7 MB

func largestWordCount(messages []string, senders []string) string {
	m := make(map[string]int)

	for i, s := range senders {
		msg := messages[i]
		count := 0
		for _, ms := range msg {
			if ms == ' ' {
				count++
			}
		}
		m[s] += count + 1
	}

	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	sender := keys[0]
	max := m[keys[0]]
	for _, key := range keys {
		v, _ := m[key]
		if max < v {
			max = v
			sender = key
		}
	}

	return sender
}
