package cn
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。 
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。 
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。 
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。 
//
// 示例: 
//
// 输入: [2,1,5,6,2,3]
//输出: 10 

//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	// 单调栈 递增（不减）
	stack := make([]int, 0)
	stack = append(stack, -1)
	heights = append(heights, 0)
	ans := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			area := heights[stack[len(stack)-1]] * (i-stack[len(stack)-2]-1)
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
