package _0_regular_expression_matching

/*
	给你一个字符串s和一个字符规律p，请你来实现一个支持 '.'和'*'的正则表达式匹配。

	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。

	说明:

	s可能为空，且只包含从a-z的小写字母。
	p可能为空，且只包含从a-z的小写字母，以及字符.和*。
	示例 1:

	输入:
	s = "aa"
	p = "a"
	输出: false
	解释: "a" 无法匹配 "aa" 整个字符串。
	示例 2:

	输入:
	s = "aa"
	p = "a*"
	输出: true
	解释:因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
	示例3:

	输入:
	s = "ab"
	p = ".*"
	输出: true
	解释:".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
	示例 4:

	输入:
	s = "aab"
	p = "c*a*b"
	输出: true
	解释:因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
	示例 5:

	输入:
	s = "mississippi"
	p = "mis*is*p*."
	输出: false
 */

func isMatchFinal(s string, p string) bool {
	match := func(i, j int) bool {
		return s[i] == p[j] || p[j] == '.'
	}
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 2; i < n+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if match(i-1, j-1) {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				if match(i-1, j-2) {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}



func isMatch(s string, p string) bool {
	// 动态规划
	// dp[i][j] 表示 字符串前i个字符可不可以被 pattern的前j个字符所匹配
	// * 要与前面的字符看做一个整体
	if p == "" &&  s != "" {
		return false
	}
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 2; i < n+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}
