package _3_multipy_strings

import "strconv"

/*
	给定两个以字符串形式表示的非负整数num1和num2，返回num1和num2的乘积，它们的乘积也表示为字符串形式。

	示例 1:
	输入: num1 = "2", num2 = "3"
	输出: "6"

	示例2:
	输入: num1 = "123", num2 = "456"
	输出: "56088"

	说明：
	num1和num2的长度小于110。
	num1 和num2 只包含数字0-9。
	num1 和num2均不以零开头，除非是数字 0 本身。
	不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

 */
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := "0"
	for i := len(num1)-1; i >= 0; i-- {
		cur := ""
		carry := 0
		v1 := int(num1[i] - '0')
		// 补0
		for k := i; k < len(num1) -1 ; k++ {
			cur += "0"
		}
		for j := len(num2)-1; j >= 0; j-- {
			v2 := int(num2[j] - '0')
			tmp := v1 * v2 + carry
			cur = strconv.Itoa(tmp % 10) + cur
			carry = tmp / 10
		}
		if carry != 0 {
			cur = strconv.Itoa(carry) + cur
		}
		ans = strSum(ans, cur)
	}
	return ans
}

func strSum(num1, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	i, j := len(num1) - 1, len(num2) - 1
	v1, v2 := 0, 0
	ans := ""
	carry := 0
	for ; i >= 0 || j >= 0; i, j = i-1, j-1 {
		if i < 0 {
			v1 = 0
		} else {
			v1 = int(num1[i] - '0')
		}
		if j < 0 {
			v2 = 0
		} else {
			v2 = int(num2[j] - '0')
		}
		tmp := v1 + v2 + carry
		ans = strconv.Itoa(tmp % 10) + ans
		carry = tmp / 10
	}
	if carry != 0 {
		ans = strconv.Itoa(carry) + ans
	}
	return ans
}
