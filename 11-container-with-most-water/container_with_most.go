package _1_container_with_most_water

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
 */

func maxAreaFinal(height []int) int {
	ans := 0
	left, right := 0, len(height)
	for left < right {
		if height[left] < height[right] {
			ans = max(ans, height[left] * (right - left))
			left++
		} else {
			ans = max(ans, height[right] * (right - left))
			right--
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




















func maxArea(height []int) int {
	left, right := 0, len(height)-1
	lVal, rVal, distance := 0, 0, 0
	cur, max := 0, 0

	for left < right {
		lVal = height[left]
		rVal = height[right]
		distance = right-left
		if lVal < rVal {
			cur = lVal * distance
			left++
		} else {
			cur = rVal * distance
			right--
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
