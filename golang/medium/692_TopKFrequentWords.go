package leetcode

// https://leetcode.com/problems/top-k-frequent-words/
// Runtime: 9 ms
// Memory Usage: 4.1 MB

type keys struct {
	word  string
	count int
}

func topKFrequent(words []string, k int) []string {
	wordArr := make([]keys, 0)

	for _, w := range words {
		idx := contains(wordArr, w)
		if idx == -1 {
			wordArr = append(wordArr, keys{word: w, count: 1})
		} else {
			wordArr[idx].count += 1
		}
	}

	for i := 0; i < len(wordArr)-1; i++ {
		for j := 0; j < len(wordArr)-i-1; j++ {
			if wordArr[j].count < wordArr[j+1].count {
				wordArr[j], wordArr[j+1] = wordArr[j+1], wordArr[j]
			}
			if wordArr[j].count == wordArr[j+1].count && wordArr[j].word > wordArr[j+1].word {
				wordArr[j], wordArr[j+1] = wordArr[j+1], wordArr[j]
			}
		}
	}

	ans := make([]string, 0)
	for _, w := range wordArr {
		for _, v := range wordArr[:k] {
			if w == v {
				ans = append(ans, w.word)
				if len(ans) == k {
					break
				}
			}
		}
	}

	return ans
}

func contains(words []keys, word string) int {
	for i, w := range words {
		if w.word == word {
			return i
		}
	}
	return -1
}
