package __longest_palindromic_substring


// 最长回文子串：https://leetcode-cn.com/problems/longest-palindromic-substring/
// 1.  暴力法，把所有的回文子串找出来
// 2.  中心扩展算法

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
