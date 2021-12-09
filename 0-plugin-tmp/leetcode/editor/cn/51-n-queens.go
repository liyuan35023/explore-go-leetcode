package cn

import (
	"bytes"
)

//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。 
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。 
//
// 示例 1： 
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
// 示例 2： 
//
//输入：n = 1
//输出：[["Q"]]
//
// 提示： 
//
// 1 <= n <= 9 
// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。 

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	queen := make([]int, n)
	columnRecord := make([]bool, n)
	diag1Record := make([]bool, 2*n-1)
	diag2Record := make([]bool, 2*n-1)

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			solve := generateSolve(queen)
			ans = append(ans, solve)
			return
		}
		for i := 0; i < n; i++ {
			if columnRecord[i] || diag1Record[idx+i] || diag2Record[idx-i+n-1] {
				continue
			}
			queen[idx] = i
			columnRecord[i] = true
			diag1Record[idx+i] = true
			diag2Record[idx-i+n-1] = true
			dfs(idx+1)
			columnRecord[i] = false
			diag1Record[idx+i] = false
			diag2Record[idx-i+n-1] = false
		}
	}
	dfs(0)
	return ans
}

func generateSolve(queen []int) []string {
	ans := make([]string, 0)
	for i := 0; i < len(queen); i++ {
		buf := bytes.NewBuffer([]byte{})
		for j := 0; j < len(queen); j++ {
			if queen[i] == j {
				buf.WriteByte('Q')
			} else {
				buf.WriteByte('.')
			}
		}
		ans = append(ans, buf.String())
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
