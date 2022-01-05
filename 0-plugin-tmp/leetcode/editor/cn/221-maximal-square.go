package cn
//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。 
//
// 示例 1：
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：4
// 
// 示例 2：
//
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
// 
// 示例 3：
//
//输入：matrix = [["0"]]
//输出：0
// 
// 提示：
//
// m == matrix.length
// n == matrix[i].length 
// 1 <= m, n <= 300 
// matrix[i][j] 为 '0' 或 '1' 

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	ans := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ans = max(ans, dp[i][0])
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			ans = max(ans, dp[0][i])
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}




//func maximalSquare(matrix [][]byte) int {
//	m, n := len(matrix), len(matrix[0])
//	ans := 0
//	dp := make([][]int, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, n)
//	}
//	for i := 0; i < n; i++ {
//		if matrix[0][i] == '1' {
//			dp[0][i] = 1
//		}
//		ans = max(ans, dp[0][i])
//	}
//	for j := 0; j < m; j++ {
//		if matrix[j][0] == '1' {
//			dp[j][0] = 1
//		}
//		ans = max(ans, dp[j][0])
//	}
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			if matrix[i][j] == '1' {
//				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
//				ans = max(ans, dp[i][j])
//			}
//		}
//	}
//	return ans * ans
//
//
//}


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
	return  y
}

//leetcode submit region end(Prohibit modification and deletion)
