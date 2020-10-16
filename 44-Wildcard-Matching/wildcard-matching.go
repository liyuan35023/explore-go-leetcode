package _4_Wildcard_Matching

/*
	给定一个字符串(s) 和一个字符模式(p) ，实现一个支持'?'和'*'的通配符匹配。

	'?' 可以匹配任何单个字符。
	'*' 可以匹配任意字符串（包括空字符串）。
	两个字符串完全匹配才算匹配成功。

	说明:

	s可能为空，且只包含从a-z的小写字母。
	p可能为空，且只包含从a-z的小写字母，以及字符?和*。
	示例1:

	输入:
	s = "aa"
	p = "a"
	输出: false
	解释: "a" 无法匹配 "aa" 整个字符串。


	示例2:
	输入:
	s = "aa"
	p = "*"
	输出: true
	解释:'*' 可以匹配任意字符串。

	示例3:
	输入:
	s = "cb"
	p = "?a"
	输出: false
	解释:'?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。

	示例4:
	输入:
	s = "adceb"
	p = "*a*b"
	输出: true
	解释:第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".

	示例5:
	输入:
	s = "acdcb"
	p = "a*c?b"
	输出: false
 */

func IsMatch(s string, p string) bool {
	// 动态规划
	// dp[i][j]表示字符串s的前i个 与 字符模式p的前j个是否匹配

	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true

	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}