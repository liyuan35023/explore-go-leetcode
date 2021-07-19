package cn
//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。 
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。 
//
// 示例 1：
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "AB
//CCED"
//输出：true
//
// 示例 2： 
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SE
//E"
//输出：true
//
// 示例 3： 
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "AB
//CB"
//输出：false
//
// 提示： 
//
// m == board.length 
// n = board[i].length 
// 1 <= m, n <= 6 
// 1 <= word.length <= 15 
// board 和 word 仅由大小写英文字母组成
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	//used := make([][]bool, len(board))
	//for i := 0; i < len(board); i++ {
	//	used[i] = make([]bool, len(board[0]))
	//}
	var dfs func(idx int, row, column int) bool
	dfs = func(idx int, row, column int) bool {
		if idx > len(word) - 1 {
			return true
		}
		if row < 0 || row > len(board)-1 || column < 0 || column > len(board[0])-1 {
			return false
		}
		if board[row][column] == word[idx] {
			board[row][column] = '0'
			ret := dfs(idx+1, row-1, column) ||
				dfs(idx+1, row+1, column) ||
				dfs(idx+1, row, column-1) ||
				dfs(idx+1, row, column+1)
			board[row][column] = word[idx]
			return ret
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
