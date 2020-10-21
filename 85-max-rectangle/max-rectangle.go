package _5_max_rectangle

/*
	给定一个只包含0或者1的矩阵，找出只包含1的最大矩形， 并返回其面积

	Example:
	input:
	[
		["1", 0, 1, 0, 0]
		[1, 0, 1, 1, 1]
		[1, 1, 1, 1, 1]
 		[1, 0, 0, 1, 0]
	]
	output : 6
 */

func maximalRectangle(matrix [][]byte) int {
	// 转化为求最大矩形：单调递增栈
	if len(matrix) < 1 {
		return 0
	}
	heights := make([][]int, len(matrix))
	for k := range heights {
		heights[k] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 && matrix[i][j] != '0' {
				heights[i][j] = 1
				continue
			}
			if matrix[i][j] != '0' {
				heights[i][j] = heights[i-1][j] + 1
			}
		}
	}
	ans := 0
	for _, v := range heights {
		tmp := maxRectangleInHistogram(v)
		ans = max(ans, tmp)
	}
	return ans
}

func maxRectangleInHistogram(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	stack := make([]int, 1)
	// 哨兵
	heights = append(heights, 0)
	stack[0] = -1
	ans := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			tmpHeight := (i - stack[len(stack)-2] - 1) * heights[stack[len(stack)-1]]
			ans = max(ans, tmpHeight)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j

}


// dp :
func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) < 1 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	height := make([][]int, m)
	left := make([][]int, m)
	right := make([][]int, m)
	ans := 0

	for  i := 0; i < m; i++ {
		height[i] = make([]int, n)
		left[i] = make([]int, n)
		right[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		if matrix[0][i] != '0' {
			height[0][i] = 1
			k := i - 1
			for ; k >= 0; k-- {
				if height[0][k] < height[0][i] {
					break
				}
			}
			left[0][i] = k

			rk := i + 1
			for ; rk < n && matrix[0][rk] != '0'; rk++ {
			}
			right[0][i] = rk
			ans = max(ans, right[0][i] - left[0][i] - 1)
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != '0' {
				height[i][j] = height[i-1][j] + 1
				k := j - 1
				for ; k >= 0 && height[i][k] >= height[i][j]; k-- {}
				left[i][j] = k

				rk := j + 1
				for ; rk < n ; rk++ {
					if matrix[i][rk] == '0' {
						break
					}
					if height[i-1][rk] + 1 < height[i][j] {
						break
					}
				}
				right[i][j] = rk

				area := height[i][j] * (right[i][j] - left[i][j] - 1)
				ans = max(ans, area)
			}
		}
	}
	return ans
}
