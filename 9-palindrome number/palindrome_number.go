package __palindrome_number

import (
	"strconv"
)

/*
    判断是否为回文数字
    第二种解法：不把数字转换为字符串，翻转字符串然后判断？（有可能溢出）
              所以只翻转一半的数字，然后判断
 */
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
