package cn
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。 
//
// 示例 1： 
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
// 
// 示例 2：
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
// 提示： 
//
// n == height.length 
// 0 <= n <= 3 * 104 
// 0 <= height[i] <= 105 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	ans := 0
	// 单调不增栈
	stack := make([]int, 0)
	for i := 0; i < len(height); i++ {
		for len(stack) > 1 && height[i] > height[stack[len(stack)-1]] {
			dis := i - stack[len(stack)-2]
			h := min(height[i], height[stack[len(stack)-2]]) - height[stack[len(stack)-1]]
			area := dis * h
			ans += area
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}
//leetcode submit region end(Prohibit modification and deletion)
