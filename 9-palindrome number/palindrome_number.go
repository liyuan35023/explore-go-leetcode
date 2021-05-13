package __palindrome_number

import (
	"strconv"
)

/*
    判断是否为回文数字
	给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

	示例 1：

	输入：x = 121
	输出：true
	示例2：

	输入：x = -121
	输出：false
	解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
	示例 3：

	输入：x = 10
	输出：false
	解释：从右向左读, 为 01 。因此它不是一个回文数。
	示例 4：

	输入：x = -101
	输出：false

	提示：
	-231<= x <= 231- 1

    第二种解法：不把数字转换为字符串，翻转字符串然后判断？（有可能溢出）
              所以只翻转一半的数字，然后判断
 */


func isPalindromeII(x int) bool {
	if x < 0 || x % 10 == 0 && x != 0 {
		return false
	}
	tmp := 0
	for x > tmp {
		tmp = tmp * 10 + x % 10
		x = x / 10
	}
	return x == tmp || x == tmp / 10
}




















func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	// convert to string
	xStr := strconv.Itoa(x)
	if len(xStr) % 2 == 0 {
		return palindrome(xStr, len(xStr)/2 - 1, len(xStr)/2)
	} else {
		return palindrome(xStr, len(xStr)/2, len(xStr)/2)
	}
}

func palindrome(s string, i, j int) bool {
	for ; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] == s[j] {
			continue
		}
		return false
	}
	return true
}

// 不把数字转换为string
func isPalindromeWithNoConvert(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	tmp := 0
	//for float64(x) >= math.Pow(float64(10), float64(numOfTmp)) {
	for x > tmp {
		tmp = tmp * 10 + x % 10
		x = x / 10
	}
	//if float64(x) < math.Pow(float64(10), float64(numOfTmp-1)) {
	//	tmp = tmp / 10
	//}
	//if tmp == x {
	//	return true
	//} else {
	//	return false
	//}
	return x == tmp || x == tmp/10

}
