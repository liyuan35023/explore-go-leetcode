package cn

import "database/sql"

//给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。 
//
// 示例 1：
//
//输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
//输出：4 
//解释：最长递增路径为[1, 2, 6, 9]。
//
// 示例 2： 
//
//输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
//输出：4 
//解释：最长递增路径是[3, 4, 5, 6]。注意不允许在对角线方向上移动。
// 
// 示例 3：
//
//输入：matrix = [[1]]
//输出：1
//
// 提示： 
//
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 200 
// 0 <= matrix[i][j] <= 231 - 1 

//leetcode submit region begin(Prohibit modification and deletion)
func longestIncreasingPath(matrix [][]int) int {
	ans := 0
	m, n := len(matrix), len(matrix[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var dfs func(row, column int) int
	dfs = func(row, column int) int {
		if dp[row][column] > 0 {
			return dp[row][column]
		}
		longest := 0
		for _, dir := range directions {
			newRow := row + dir[0]
			newColumn := column + dir[1]
			if newRow >= 0 && newRow < m && newColumn >=0 && newColumn < n && matrix[row][column] < matrix[newRow][newColumn] {
				longest = max(longest, dfs(newRow, newColumn))
			}
		}
		dp[row][column] = longest + 1
		return dp[row][column]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
