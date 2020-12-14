package _29_longest_increasing_path_in_matrix

import "math"

/*
	Example 1:

	Input: nums =
	[
	  [9,9,4],
	  [6,6,8],
	  [2,1,1]
	]
	Output: 4
	Explanation: The longest increasing path is [1, 2, 6, 9].

	Example 2:
	Input: nums =
	[
	  [3,4,5],
	  [3,2,6],
	  [2,2,1]
	]
	Output: 4
	Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
	题目大意 #
	给定一个整数矩阵，找出最长递增路径的长度。
	对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
 */
func longestIncreasingPath(matrix [][]int) int {
	// 记忆优化
	if len(matrix) < 1 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	ans := 1
	visitedPath := make([][]int, m)
	for i := 0; i < m; i++ {
		visitedPath[i] = make([]int, n)
	}
	var dfs func(row int, column int, preVal int) int
	dfs = func(row int, column int, preVal int) int {
		if row < 0 || row >= m || column < 0 || column >= n {
			return 0
		}
		if preVal >= matrix[row][column] {
			return 0
		}
		if visitedPath[row][column] != 0 {
			return visitedPath[row][column]
		}

		cur := matrix[row][column]
		v := dfs(row-1, column, cur)
		v = max(v, dfs(row+1, column, cur))
		v = max(v, dfs(row, column-1, cur))
		v = max(v, dfs(row, column+1, cur))
		visitedPath[row][column] = v + 1
		return v + 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j, math.MinInt32))
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

func longestIncreasingPathII(matrix [][]int) int {
	// 纯dfs
	// 会超时
	if len(matrix) < 1 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	ans := 1
	var dfs func(row int, column int, preVal int) int
	dfs = func(row int, column int, preVal int) int {
		if row < 0 || row >= m || column < 0 || column >= n {
			return 0
		}
		if preVal >= matrix[row][column] {
			return 0
		}
		cur := matrix[row][column]
		v := dfs(row-1, column, cur)
		v = max(v, dfs(row+1, column, cur))
		v = max(v, dfs(row, column-1, cur))
		v = max(v, dfs(row, column+1, cur))
		return v + 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j, math.MinInt32))
		}
	}
	return ans
}