package _2_trapping_rain_water

/*
	Example:


	Input: [0,1,0,2,1,0,1,3,2,1,2,1]
	Output: 6

	题目大意 #
	从 x 轴开始，给出一个数组，数组里面的数字代表从 (0,0) 点开始，宽度为 1 个单位，高度为数组元素的值。如果下雨了，问这样一个容器能装多少单位的水？
 */

func trap(height []int) int {
	// 使用栈解决
	if len(height) < 3 {
		return 0
	}

	// stack用来存储索引. 单调递减栈
	stack := make([]int, 0)
	ans := 0

	for index, h := range height {
		if len(stack) == 0 {
			stack = append(stack, index)
			continue
		}

		for height[stack[len(stack)-1]] < h {
			if len(stack) == 1 {
				stack = stack[:len(stack)-1]
				break
			}
			dis := index - stack[len(stack)-2] - 1
			minHeight := min(height[stack[len(stack)-2]], h)
			ans += dis * (minHeight-height[stack[len(stack)-1]])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	return ans
}



func TrapDoublePointer(height []int) int {
	// 双指针法
	if len(height) < 3 {
		return 0
	}
	ans := 0
	//leftMax, rightMax := height[0], height[len(height)-1]
	leftMax, rightMax := 0, 0
	i, j := 0, len(height)-1
	for i <= j {
		//if leftMax < rightMax {
		//	if leftMax >= height[i] {
		//		ans += leftMax - height[i]
		//	} else {
		//		leftMax = height[i]
		//	}
		//	i++
		//} else {
		//	if rightMax >= height[j] {
		//		ans += rightMax - height[j]
		//	} else {
		//		rightMax = height[j]
		//	}
		//	j--
		//}

		if height[i] < height[j] {
			if height[i] >= leftMax {
				leftMax = height[i]
			} else {
				ans += leftMax - height[i]
			}
			i++
		} else {
			if height[j] >= rightMax {
				rightMax = height[j]
			} else {
				ans += rightMax - height[j]
			}
			j--
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

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
