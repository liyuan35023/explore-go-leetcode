package cn

import "strings"

//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
// 输入为 非空 字符串且只包含数字 1 和 0。 
//
// 示例 1: 
//
// 输入: a = "11", b = "1"
//输出: "100" 
//
// 示例 2: 
//
// 输入: a = "1010", b = "1011"
//输出: "10101" 
//
// 提示： 
//
// 每个字符串仅由字符 '0' 或 '1' 组成。 
// 1 <= a.length, b.length <= 10^4 
// 字符串如果不是 "0" ，就都不含前导零。 


//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	ans := make([]int, max(m, n) + 1)
	i, j := m-1, n-1
	idx := 0
	for i >= 0 || j >= 0 {
		var a0, b0 int
		if i >= 0 {
			a0 = int(a[i]-'0')
		}
		if j >= 0 {
			b0 = int(b[j]-'0')
		}
		i, j = i-1, j-1
		ans[idx] = a0 + b0
		idx++
	}
	carry := 0
	ansStr := ""
	for i := 0; i < len(ans); i++ {
		if i == len(ans) - 1 && carry == 0 {
			break
		}
		ans[i] += carry
		switch ans[i] {
		case 0:
			ansStr = "0" + ansStr
			carry = 0
		case 1:
			ansStr = "1" + ansStr
			carry = 0
		case 2:
			ansStr = "0" + ansStr
			carry = 1
		case 3:
			ansStr = "1" + ansStr
			carry = 1
		}
	}
	return ansStr
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}
//leetcode submit region end(Prohibit modification and deletion)
