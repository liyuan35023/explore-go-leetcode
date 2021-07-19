package cn
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
	i, j := len(a) - 1, len(b) - 1
	carry := 0
	ans := ""
	for ; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		varA, varB := 0, 0
		if i >= 0 {
			varA = int(a[i]-'0')
		}
		if j >= 0 {
			varB = int(b[j]-'0')
		}
		switch varA + varB + carry {
		case 0:
			ans = "0" + ans
			carry = 0
		case 1:
			ans = "1" + ans
			carry = 0
		case 2:
			ans = "0" + ans
			carry = 1
		case 3:
			ans = "1" + ans
			carry = 1
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
