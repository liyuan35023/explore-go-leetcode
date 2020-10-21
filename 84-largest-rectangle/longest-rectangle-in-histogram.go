package _4_largest_rectangle

/*
	给出每个直方图的高度，要求在这些直方图之中找到面积最大的矩形，输出矩形的面积。

	Example:
	Input: [2,1,5,6,2,3]
	Output: 10

*/
func LargestRectangleArea(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	ans := 0
	// 单调递增栈,保存的index
	stack := make([]int, 0)
	for i := 0; i < len(heights); i++ {
		n := len(stack) - 1
		for j := n; j >= 0 && heights[stack[j]] > heights[i]; j-- {
			tmp := 0
			if j > 0 {
				tmp = (i - stack[j-1] - 1) * heights[stack[j]]
			} else {
				tmp = i * heights[stack[j]]
			}
			ans = max(ans, tmp)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	// 清空栈
	for i := len(stack)-1; i >= 0; i-- {
		tmp := 0
		if i > 0 {
			tmp = heights[stack[i]] * (stack[len(stack)-1] - stack[i-1])
		} else {
			tmp = heights[stack[i]] * (stack[len(stack)-1] + 1)
		}
		ans = max(ans, tmp)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j

}
