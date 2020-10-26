package offer

import "strconv"

/*
	给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
	一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

	Example:
	输入: 12258
	输出: 5
	解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

	提示：

	0 <= num < 2^31
 */

func TranslateNum2(num int) int {
	if num == 0 {
		return 1
	}
	last2, remain := getAndExcludeLast2(num)
	if last2 >= 10 && last2 <= 25 {
		return TranslateNum(remain) + TranslateNum(num / 10)
	} else {
		return TranslateNum(num / 10)
	}
}

func getAndExcludeLast2(x int) (last2 int, remain int) {
	return x % 100, x / 100
}

func TranslateNum(num int) int {
	// use dp
	numStr := strconv.Itoa(num)
	n := len(numStr)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		tmp, _  := strconv.Atoi(string(numStr[i-1:i+1]))
		if tmp >= 10 && tmp <= 25 {
			if i == 1 {
				dp[i] = dp[i-1] + 1
				continue
			}
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}
