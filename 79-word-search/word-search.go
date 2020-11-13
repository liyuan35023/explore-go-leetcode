package _9_word_search

/*
	Example:

	board =
	[
	  ['A','B','C','E'],
	  ['S','F','C','S'],
	  ['A','D','E','E']
	]

	Given word = "ABCCED", return true.
	Given word = "SEE", return true.
	Given word = "ABCB", return false.
	题目大意 #
	给定一个二维网格和一个单词，找出该单词是否存在于网格中。单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
	同一个单元格内的字母不允许被重复使用。
 */

func exist(board [][]byte, word string) bool {
	if len(board) < 1 {
		return false
	}
	m := len(board)
	n := len(board[0])

	for row := 0; row < m; row++ {
		for column := 0; column < n; column++ {
			if backTrace(board, row, column, word, 0) {
				return true
			}
		}
	}
	return false
}

func backTrace(board [][]byte, row int, column int, word string, idx int) bool {
	if idx == len(word) {
		return true
	}
	if row < 0 || row >= len(board) || column < 0 || column >= len(board[0]) {
		return false
	}
	if board[row][column] == word[idx] && board[row][column] != '0' {
		board[row][column] = '0'
		if backTrace(board, row-1, column, word, idx+1) {
			// 向上
			return true
		}
		if backTrace(board, row+1, column, word, idx+1) {
			// 向下
			return true
		}
		if backTrace(board, row, column-1, word, idx+1) {
			// 向左
			return true
		}
		if backTrace(board, row, column+1, word, idx+1) {
			// 向右
			return true
		}
		board[row][column] = word[idx]
	}
	return false
}
