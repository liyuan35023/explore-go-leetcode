package cn
//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。 
//
// 返回被除数 dividend 除以除数 divisor 得到的商。 
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2 
//
// 
//
// 示例 1: 
//
// 输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3 
//
// 示例 2: 
//
// 输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = truncate(-2.33333..) = -2 
//
// 
// 提示：
//
// 
// 被除数和除数均为 32 位有符号整数。 
// 除数不为 0。 
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。 
// 
// Related Topics 数学 二分查找 
// 👍 594 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func divide(dividend int, divisor int) int {
	signed := -1
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		signed = 1
	}
	absDividend := abs(dividend)
	absDivisor := abs(divisor)
	ans := 0
	curMulti := 1
	curDivisor := absDivisor
	for absDividend >= absDivisor {
		for absDividend >= curDivisor {
			absDividend -= curDivisor
			curDivisor += curDivisor
			if signed == 1 && ans > 1 << 31 - 1 - curMulti || signed == -1 && ans > 1 << 31 - curMulti {
				return 1 << 31 - 1
			}
			ans += curMulti
			curMulti += curMulti
		}
		curDivisor = absDivisor
		curMulti = 1
	}
	return signed * ans
}

//func abs(x int) int {
//	if x < 0 {
//		x = -x
//	}
//	return x
//}

//leetcode submit region end(Prohibit modification and deletion)
