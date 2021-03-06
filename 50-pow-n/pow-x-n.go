package _0_pow_n

/*
	Example 1:
	Input: 2.00000, 10
	Output: 1024.00000

	Example 2:
	Input: 2.10000, 3
	Output: 9.26100

	Example 3:
	Input: 2.00000, -2
	Output: 0.25000
	Explanation: 2-2 = 1/22 = 1/4 = 0.25

	Note:
	-100.0 < x < 100.0
	n is a 32-bit signed integer, within the range [−2^31, 2^31− 1]

	题目大意 #
	实现 pow(x, n) ，即计算 x 的 n 次幂函数。
 */
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	} else {
		return 1 / quickMul(x, -n)
	}
}

func powHelper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	return powHelper(x * x, n / 2) * powHelper(x, n % 2)
}

func quickMul(x float64, n int) float64 {
	// 迭代
	ans := 1.0
	contributor := x
	for n > 0 {
		if n % 2 == 1 {
			ans *= contributor
		}
		contributor *= contributor
		n = n / 2
	}
	return ans
}

