package _9_sqrt

/*
	Example 1:

	Input: 4
	Output: 2

	Example 2:

	Input: 8
	Output: 2
	Explanation: The square root of 8 is 2.82842..., and since
				 the decimal part is truncated, 2 is returned.
	题目大意 #
	实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，
	结果只保留整数的部分，小数部分将被舍去。

 */

func mySqrt(x int) int {
	left := 0
	right := x
	for left < right {
		half := (left + right) / 2
		tmp := half * half
		if tmp == x {
			return half
		} else if tmp < x {
			if (half+1) * (half+1) > x {
				return half
			}
			left = half+1
		} else if tmp > x {
			if (half-1) * (half-1) < x {
				return half-1
			}
			right = half-1
		}
	}
	return left
}

func mySqrt2(x int) int {
	left := 0
	right := x
	ans := -1
	for left <= right {
		half := (left + right) / 2
		tmp := half * half
		if tmp == x {
			return half
		} else if tmp < x {
			ans = half
			left = half+1
		} else if tmp > x {
			right = half-1
		}
	}
	return ans
}
