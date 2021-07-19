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
	columnMap := make(map[int]bool)
	diag1 := make(map[int]bool)
	diag2 := make(map[int]bool)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			ans++
			return
		}
		for column := 0; column < n; column++ {
			if columnMap[column] || diag1[idx+column] || diag2[idx-column] {
				continue
			}
			columnMap[column] = true
			diag1[idx+column] = true
			diag2[idx-column] = true
			dfs(idx+1)
			columnMap[column] = false
			diag1[idx+column] = false
			diag2[idx-column] = false
		}
	}
	dfs(0)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
