package cn
//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。 
//
// 示例 1： 
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
// 
// 示例 2：
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//
// 示例 3： 
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25
//
//
// 提示： 
//
// -100.0 < x < 100.0 
// -231 <= n <= 231-1 
// -104 <= xn <= 104 

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	// 快速幂
	if n == 0 {
		return 1.0
	}
	reverse := false
	if n < 0 {
		reverse = true
		n = -n
	}
	ans := float64(1)
	con := x
	for n > 1 {
		if n % 2 == 1 {
			ans *= con
		}
		con *= con
		n = n / 2
	}
	ans = ans * con
	if reverse {
		return 1 / ans
	} else {
		return ans
	}
}

//leetcode submit region end(Prohibit modification and deletion)
