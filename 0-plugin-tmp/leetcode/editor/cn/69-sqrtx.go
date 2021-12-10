package cn

import "math"

//实现 int sqrt(int x) 函数。
//
// 计算并返回 x 的平方根，其中 x 是非负整数。 
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。 
//
// 示例 1: 
//
// 输入: 4
//输出: 2
// 
//
// 示例 2: 
//
// 输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842..., 
//    由于返回类型是整数，小数部分将被舍去。
//

//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
}
//func mySqrt(x int) int {
//	x0 := float64(x)
//	for x0 > 0 {
//		x1 := 0.5 * (x0 + float64(x) / x0)
//		if math.Abs(x0 - x1) < 1e-7 {
//			return int(x1)
//		}
//		x0 = x1
//	}
//	return int(x0)
//}
//func mySqrt(x int) int {
//	left, right := 0, x
//	for left < right {
//		mid := left + (right-left) / 2
//		if mid * mid == x {
//			return mid
//		}
//		if mid * mid > x && (mid-1) * (mid-1) <= x {
//			return mid - 1
//		}
//		if mid * mid < x && (mid+1) * (mid+1) > x {
//			return mid
//		}
//		if mid * mid > x {
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//
//	}
//	return left
//
//}
//leetcode submit region end(Prohibit modification and deletion)
