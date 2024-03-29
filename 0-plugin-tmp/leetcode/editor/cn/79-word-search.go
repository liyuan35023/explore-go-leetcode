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
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	var dfs func(row, column int, idx int) bool
	dfs = func(row, column int, idx int) bool {
		if idx == len(word) {
			return true
		}
		if row < 0 || row >= m || column < 0 || column >= n {
			return false
		}
		if board[row][column] == word[idx] {
			tmp := board[row][column]
			board[row][column] = '0'
			for _, dir := range directions {
				if dfs(row+dir[0], column+dir[1], idx+1) {
					return true
				}
			}
			board[row][column] = tmp
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false


}
//leetcode submit region end(Prohibit modification and deletion)
