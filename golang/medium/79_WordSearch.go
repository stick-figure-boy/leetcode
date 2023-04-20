package leetcode

// https://leetcode.com/problems/word-search/
// Runtime: 228 ms
// Memory Usage: 2 MB

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, visited, word, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited [][]bool, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] || board[i][j] != word[0] {
		return false
	}
	visited[i][j] = true
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for _, d := range directions {
		if dfs(board, visited, word[1:], i+d[0], j+d[1]) {
			return true
		}
	}
	visited[i][j] = false
	return false
}
