package cn


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
	digit := make([]int, max(m,n)+1)
	i, j := m-1, n-1
	idx := 0
	for i >= 0 || j >= 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}
		digit[idx] = n1 + n2
		idx++
	}
	ans := make([]byte, 0)
	carry := 0
	for i := 0; i < len(digit); i++ {
		d := (digit[i] + carry) % 10
		carry = (digit[i] + carry) / 10
		db := byte('0' + d)
		ans = append([]byte{db}, ans...)
	}
	if len(ans) > 1 && ans[0] == '0' {
		ans = ans[1:]
	}
	return string(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
