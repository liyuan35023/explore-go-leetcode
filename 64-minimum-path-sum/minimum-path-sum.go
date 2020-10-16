package _4_minimum_path_sum

/*
	给定一个包含非负整数的 mxn网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

	说明：每次只能向下或者向右移动一步。

	示例:

	输入:
	[
	  [1,3,1],
	  [1,5,1],
	  [4,2,1]
	]
	输出: 7
	解释: 因为路径 1→3→1→1→1 的总和最小。
 */

func minPathSum(grid [][]int) int {
	if len(grid) < 1 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
