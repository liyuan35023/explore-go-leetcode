package _7_add_binary

import "strconv"

/*
	给你两个二进制字符串，返回它们的和（用二进制表示）。输入为 非空 字符串且只包含数字 1 和 0。
	Example 1:

	Input: a = "11", b = "1"
	Output: "100"
	Example 2:
z
	Input: a = "1010", b = "1011"
	Output: "10101"
 */

func addBinary(a string, b string) string {
	ans := ""
	i, j := len(a)-1, len(b)-1
	carry := 0
	for ; i >= 0 || j >= 0 || carry != 0 ; i, j = i-1, j-1 {
		v1, v2 := 0, 0
		if i >= 0 {
			v1, _ = strconv.Atoi(a[i:i+1])
		}
		if j >= 0 {
			v2, _ = strconv.Atoi(b[j:j+1])
		}
		v := v1 + v2 + carry
		carry = v / 2
		ans = strconv.Itoa(v % 2) + ans
	}
	return ans
}
