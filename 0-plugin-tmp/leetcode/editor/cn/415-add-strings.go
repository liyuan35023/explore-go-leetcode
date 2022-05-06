package cn

import "strconv"

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
// 提示： 
//
// num1 和num2 的长度都小于 5100 
// num1 和num2 都只包含数字 0-9 
// num1 和num2 都不包含任何前导零 
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	ans := make([]int, max(m, n) + 1)
	length := max(m, n) + 1
	i, j := m - 1, n - 1
	idx := len(ans) - 1
	for i >= 0 || j >= 0 {
		var v1, v2 int
		if i >= 0 {
			v1 = int(num1[i]-'0')
		}
		if j >= 0 {
			v2 = int(num2[j]-'0')
		}
		ans[idx] = v1 + v2
		i, j, idx = i - 1, j - 1, idx - 1
	}

	ansByte := make([]byte, length)
	carry := 0
	for i := len(ans) - 1; i >= 0; i-- {
		tmp := ans[i] + carry
		ansByte[i] = byte('0' + tmp % 10)
		carry = (ans[i] + carry) / 10
	}
	for len(ansByte) > 1 && ansByte[0] == '0' {
		ansByte = ansByte[1:]
	}
	return string(ansByte)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
