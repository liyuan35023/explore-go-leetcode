package cn
//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。 
//
// 示例 1： 
//
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
// 示例 2： 
//
//输入：n = 1
//输出：1
// 提示：
//
// 1 <= n <= 9 
// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	ans := 0
	columnRecord := make(map[int]bool)
	diag1Record := make(map[int]bool)
	diag2Record := make(map[int]bool)
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			ans++
			return
		}
		for i := 0; i < n; i++ {
			if columnRecord[i] || diag1Record[row+i] || diag2Record[row-i] {
				continue
			}
			columnRecord[i] = true
			diag1Record[row+i] = true
			diag2Record[row-i] = true
			dfs(row+1)
			columnRecord[i] = false
			diag1Record[row+i] = false
			diag2Record[row-i] = false
		}
	}
	dfs(0)
	return ans

}
//leetcode submit region end(Prohibit modification and deletion)
