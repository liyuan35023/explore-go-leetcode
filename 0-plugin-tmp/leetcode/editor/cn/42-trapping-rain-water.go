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
	if len(height) < 1 {
		return 0
	}
	// 单调不增栈
	ans := 0
	// stack 保存的是idx
	stack := []int{0}
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			if len(stack) > 1 {
				leftHeight := height[stack[len(stack)-2]]
				ans += (min(leftHeight, height[i]) - height[stack[len(stack)-1]]) * (i - stack[len(stack)-2] - 1)
			}
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
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

func trapII(height []int) int {
	// 左侧最高
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	left, right := -1, -1
	for i := 0; i < len(height); i++ {
		if height[i] < left {
			leftMax[i] = left
		} else {
			leftMax[i] = -1
			left = max(left, height[i])
		}
	}
	for i := len(height)-1; i >=0; i-- {
		if height[i] < right {
			rightMax[i] = right
		} else {
			rightMax[i] = -1
			right = max(right, height[i])
		}
	}
	ans := 0
	for i := 0; i < len(height); i++ {
		minHeight := min(leftMax[i], rightMax[i])
		if minHeight > height[i] {
			ans += minHeight - height[i]
		}
	}
	return ans
}
