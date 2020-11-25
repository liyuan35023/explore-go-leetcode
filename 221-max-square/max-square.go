package _21_max_square

/*
	在一个由'0'和 '1'组成的二维矩阵中，找到只包含'1'的最大正方形，并返回其面积。

 */

func maximalSquare(matrix [][]byte) int {
	// dp[i][j] 表示以第i+1，j+1 为右下角的正方形的最大边长
	if len(matrix) < 1 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	maxSide := 0
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			maxSide = max(maxSide, dp[0][i])
		}
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxSide = max(maxSide, dp[i][0])
		}
	}
	for row := 1; row < m; row++ {
		for column := 1; column < n; column++ {
			if matrix[row][column] == '1' {
				dp[row][column] = min(dp[row-1][column-1], min(dp[row][column-1], dp[row-1][column])) + 1
				maxSide = max(maxSide, dp[row][column])
			}
		}
	}
	return maxSide * maxSide
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
