package _9_divide_two_integer

/*
	给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。返回被除数 dividend 除以除数 divisor 得到的商。

	说明:
	被除数和除数均为 32 位有符号整数。
	除数不为 0。
	假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。

	Example 1:
	Input: dividend = 10, divisor = 3
	Output: 3

	Example 2:
	Input: dividend = 7, divisor = -3
	Output: -2
 */

func Divide(dividend int, divisor int) int {
	signed := 1
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		signed = -1
	}
	absDividend := abs(dividend)
	absDivisor := abs(divisor)
	ans := 0
	for absDividend >= absDivisor {
		x := absDivisor
		nums := 1
		for x + x <= absDividend {
			x = x + x
			nums = nums + nums
		}
		absDividend = absDividend - x
		ans += nums
	}
	ans = ans * signed
	if ans > 1 << 31 - 1 || ans < -1 << 31 {
		return 1<<31 - 1
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

