package cn

import (
	"bytes"
	"strconv"
)

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
	i, j := m-1, n-1
	idx := 0
	for i >= 0 || j >= 0 {
		var digit1, digit2 int
		if i >= 0 {
			digit1 = int(num1[i] - '0')
		}
		if j >= 0 {
			digit2 = int(num2[j] - '0')
		}
		ans[idx] = digit1 + digit2
		idx++
		i, j = i - 1, j - 1
	}
	buf := make([]byte, 0)
	for i := 0; i < len(ans); i++ {
		tmp := ans[i]
		if i != len(ans) - 1 {
			ans[i+1] += tmp / 10
		} else {
			if ans[i] == 0 {
				break
			}
		}
		b := byte('0' + tmp % 10)
		buf = append([]byte{b}, buf...)
		//buf = append([]byte(strconv.Itoa(tmp%10)), buf...)
	}
	return string(buf)


}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
