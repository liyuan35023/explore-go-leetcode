package __reverse_integer

import "math"

/*
	给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
	如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31− 1] ，就返回 0。
	假设环境不允许存储 64 位整数（有符号或无符号）。

	示例 1：
	输入：x = 123
	输出：321

	示例 2：
	输入：x = -123
	输出：-321

	示例 3：
	输入：x = 120
	输出：21

	示例 4：
	输入：x = 0
	输出：0

	提示：
	-231 <= x <= 231 - 1
 */

func reverseFinal(x int) int {
	ans := 0
	for x != 0 {
		tmp := x % 10
		if ans > (math.MaxInt32 - tmp) / 10 || ans < (math.MinInt32 - tmp) / 10 {
			return 0
		}
		ans = ans * 10 + tmp
		x /= 10
	}
	return ans
}

func reverse(x int) int {
	tmp := 0
	for x != 0 {
		if tmp > (1<<31 - 1 - x %10) / 10 || tmp < (-(1<<31)-x%10) / 10 {
			return 0
		}
		tmp = tmp*10 + x % 10
		x = x / 10
	}
	return tmp
}
