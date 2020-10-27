package _4_spiral_matrix

/*
	给定一个包含m x n个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

	示例1:
	输入:
	[
	 [ 1, 2, 3 ],
	 [ 4, 5, 6 ],
	 [ 7, 8, 9 ]
	]
	输出: [1,2,3,6,9,8,7,4,5]

	示例2:
	输入:
	[
	  [1, 2, 3, 4],
	  [5, 6, 7, 8],
	  [9,10,11,12]
	]
	输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 {
		return nil
	}
	ans := make([]int, 0)
	m := len(matrix)
	n := len(matrix[0])

	top, bottom, left, right := 0, m-1, 0, n-1

	for top < bottom && left < right {
		row, column := top, left
		for ; column < right; column++ {
			ans = append(ans, matrix[row][column])
		}
		for ; row < bottom; row++ {
			ans = append(ans, matrix[row][column])
		}
		for ; column > left; column-- {
			ans = append(ans, matrix[row][column])
		}
		for ; row > top; row-- {
			ans = append(ans, matrix[row][column])
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	if top == bottom {
		for column := left; column <= right; column++ {
			ans = append(ans, matrix[top][column])
		}
	} else if left == right {
		for row := top; row <= bottom; row++ {
			ans = append(ans, matrix[row][left])
		}
	}
	return ans
}
