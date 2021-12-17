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
	ans := 0
	m, n := len(matrix), len(matrix[0])
	heightsArr := make([][]int, m)
	for i := 0; i < m; i++ {
		heightsArr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			heightsArr[0][i] = 1
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heightsArr[i][j] = 0
			} else {
				heightsArr[i][j] = heightsArr[i-1][j] + 1
			}
		}
	}
	for _, heights := range heightsArr {
		ans = max(ans, maxArea(heights))
	}
	return ans
}

func maxArea(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	heights = append(heights, 0)
	ans := 0
	for k, h := range heights {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > h {
			area := (k - stack[len(stack)-2] - 1) * heights[stack[len(stack)-1]]
			ans = max(ans, area)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, k)
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
