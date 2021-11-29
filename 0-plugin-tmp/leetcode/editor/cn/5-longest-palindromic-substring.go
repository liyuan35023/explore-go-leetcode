package cn
//给你一个字符串 s，找到 s 中最长的回文子串。 
//
// 示例 1： 
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
// 
// 示例 2：
//
//输入：s = "cbbd"
//输出："bb"
//
// 示例 3： 
//
//输入：s = "a"
//输出："a"
// 
// 示例 4：
//
//输入：s = "ac"
//输出："a"
//
// 提示： 
//
// 1 <= s.length <= 1000 
// s 仅由数字和英文字母（大写和/或小写）组成 
// 

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		l, r := expandCenter(s, i, i)
		if r - l + 1 > len(ans) {
			ans = s[l:r+1]
		}
		l, r = expandCenter(s, i, i+1)
		if r - l + 1 > len(ans) {
			ans = s[l:r+1]
		}
	}
	return ans
}

func expandCenter(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1{
	}
	return i+1, j-1
}

//leetcode submit region end(Prohibit modification and deletion)
