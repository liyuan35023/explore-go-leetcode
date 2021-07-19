package cn
//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。 
//
// 示例 1： 
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：6
//解释：最大矩形如上图所示。
// 
// 示例 2：
//
//输入：matrix = []
//输出：0
//
// 示例 3： 
//
//输入：matrix = [["0"]]
//输出：0
//
// 示例 4： 
//
//输入：matrix = [["1"]]
//输出：1
//
// 示例 5： 
//
//输入：matrix = [["0","0"]]
//输出：0
//
// 提示： 
//
// rows == matrix.length
// cols == matrix[0].length 
// 0 <= row, cols <= 200 
// matrix[i][j] 为 '0' 或 '1' 


//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) < 1 {
		return 0
	}
	heights := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		heights[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			heights[0][i] = 1
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[i][j] = heights[i-1][j] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < len(matrix); i++ {
		ans = max(ans, maxArea(heights[i]))
	}
	return ans
}

func maxArea(heights []int) int {
	stack := []int{-1}
	heights = append(heights, 0)
	ans := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			area := heights[stack[len(stack)-1]] * (i - stack[len(stack)-2] - 1)
			ans = max(ans, area)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
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
