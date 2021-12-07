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
	if len(height) == 0 {
		return 0
	}
	ans := 0
	lMax, rMax := height[0], height[len(height) - 1]
	l, r := 0, len(height) - 1
	for l < r {
		if lMax < rMax {
			ans += lMax - height[l]
			l++
			lMax = max(lMax, height[l])
		} else {
			ans += rMax - height[r]
			r--
			rMax = max(rMax, height[r])
		}
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
