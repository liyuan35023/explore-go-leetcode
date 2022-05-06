package cn

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
// 示例 1： 
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
// 示例 2： 
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
// 提示： 
//
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 10 
// -100 <= matrix[i][j] <= 100 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	m, n := len(matrix), len(matrix[0])
	top, bottom := 0, m - 1
	left, right := 0, n - 1
	for top < bottom && left < right {
		i, j := top, left
		for ; j < right; j++ {
			ans = append(ans, matrix[top][j])
		}
		for ; i < bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		for ; j > left; j-- {
			ans = append(ans, matrix[bottom][j])
		}
		for ; i > top; i-- {
			ans = append(ans, matrix[i][left])
		}
		top, bottom = top + 1, bottom - 1
		left, right = left + 1, right - 1
	}
	if top == bottom {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][left])
		}

	}


	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
