package _40_search_2d_matrix_II
/*
	Example:

	Consider the following matrix:

	[
	  [1,   4,  7, 11, 15],
	  [2,   5,  8, 12, 19],
	  [3,   6,  9, 16, 22],
	  [10, 13, 14, 17, 24],
	  [18, 21, 23, 26, 30]
	]
	Given target = 5, return true.

	Given target = 20, return false.

	题目大意 #
	编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

	每行的元素从左到右升序排列。
	每列的元素从上到下升序排列。
 */

func searchMatrix240(matrix [][]int, target int) bool {
	// 从右上角开始
	m := len(matrix)
	n := len(matrix[0])
	row, column := 0, n - 1
	for row >= 0 && row < m && column >= 0 && column < n {
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return false
}
