package __longest_palindromic_substring

/*
	给你一个字符串 s，找到 s 中最长的回文子串。
	示例 1：

	输入：s = "babad"
	输出："bab"
	解释："aba" 同样是符合题意的答案。
	示例 2：

	输入：s = "cbbd"
	输出："bb"
	示例 3：

	输入：s = "a"
	输出："a"
	示例 4：

	输入：s = "ac"
	输出："a"

*/

func longestPalindromeFinal(s string) string {
	// expand
	ans := s[:1]
	for center := 0; center < len(s) ; center++ {
		left, right := expandCenter(s, center, center)
		if right - left + 1 > len(ans) {
			ans = s[left:right+1]
		}
		left, right = expandCenter(s, center, center+1)
		if right - left + 1 > len(ans) {
			ans = s[left:right+1]
		}
	}
	return ans
}

func expandCenter(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
	}
	return i+1, j-1
}

















func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		l1, r1 := findExpandLongestSubString(s, i, i)
		l2, r2 := findExpandLongestSubString(s, i, i+1)
		if r1-l1 > start-end {
			start = l1
			end = r1
		}
		if r2-l2 > start-end {
			start = l2
			end = r2
		}
	}
	return s[start:end+1]

}

func findExpandLongestSubString(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left+1, right-1
}

func longestPalindromeDP(s string) string {
	// 动态规划
	// dp[i][j] 表示从第i+1个字母到第j+1个字母是不是回文子串
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	ans := s[:1]

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	for column := 1; column < len(s); column++ {
		for row := column-1; row >= 0; row-- {
			if s[row] == s[column] && (row == column-1 || dp[row+1][column-1]) {
				dp[row][column] = true
				if column - row + 1 > len(ans) {
					ans = s[row:column+1]
				}
			}
		}
	}
	return ans
}
