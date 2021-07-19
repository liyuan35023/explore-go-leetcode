package cn

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
// 每行中的整数从左到右按升序排列。 
// 每行的第一个整数大于前一行的最后一个整数。
//
// 示例 1： 
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
// 示例 2： 
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
// 提示：
// m == matrix.length
// n == matrix[i].length 
// 1 <= m, n <= 100 
// -104 <= matrix[i][j], target <= 104 


//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 一次二分
	// 坐标转化
	i, j := 0, m * n - 1
	for i <= j {
		mid := i + (j - i) / 2
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
