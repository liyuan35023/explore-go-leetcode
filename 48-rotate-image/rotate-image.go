package _8_rotate_image

/*
Example 1:
Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],
rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]


Example 2:
Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
题目大意 #
给定一个 n × n 的二维矩阵表示一个图像。将图像顺时针旋转 90 度。说明：你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

 */
func rotate(matrix [][]int) {
	if len(matrix) < 2 {
		return
	}
	n := len(matrix)
	for i := 0; i <= n/2; i++ {
		for j := i; j < n-1-i; j++ {
			row, column := i, j
			targetRow, targetColumn := 0, 0
			pre := matrix[row][column]
			for {
				targetRow, targetColumn = column, n-row-1
				tmp := matrix[targetRow][targetColumn]
				matrix[targetRow][targetColumn] = pre
				pre = tmp
				row, column = targetRow, targetColumn
				if row == i && column == j {
					break
				}
			}
		}
	}
}

