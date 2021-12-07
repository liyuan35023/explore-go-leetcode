package cn

import "strconv"

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 示例 1: 
//
// 输入: num1 = "2", num2 = "3"
//输出: "6" 
//
// 示例 2: 
//
// 输入: num1 = "123", num2 = "456"
//输出: "56088" 
//
// 说明： 
//
// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。 
// num1 和 num2 均不以零开头，除非是数字 0 本身。 
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。 


//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	result := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			realPos := (m - 1) - i + (n - 1) - j
			tmp := int(num1[i] - '0') * int(num2[j] - '0')
			result[realPos] += tmp
		}
	}
	ans := ""
	carry := 0
	for i := 0; i < m+n; i++ {
		tmp := result[i] + carry
		result[i] = tmp % 10
		carry = tmp / 10
		if i == m + n - 1 && result[i] == 0 {
			break
		}
		ans = strconv.Itoa(result[i]) + ans
	}
	return ans


}

//leetcode submit region end(Prohibit modification and deletion)
