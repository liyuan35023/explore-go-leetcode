package _2_edit_distance

/*
	给你两个单词word1 和word2，请你计算出将word1转换成word2 所使用的最少操作数。
	你可以对一个单词进行如下三种操作：
	插入一个字符
	删除一个字符
	替换一个字符

	示例1：
	输入：word1 = "horse", word2 = "ros"
	输出：3
	解释：
	horse -> rorse (将 'h' 替换为 'r')
	rorse -> rose (删除 'r')
	rose -> ros (删除 'e')

	示例2：
	输入：word1 = "intention", word2 = "execution"
	输出：5
	解释：
	intention -> inention (删除 't')
	inention -> enention (将 'i' 替换为 'e')
	enention -> exention (将 'n' 替换为 'x')
	exention -> exection (将 'n' 替换为 'c')
	exection -> execution (插入 'u')

 */

func MinDistance(word1 string, word2 string) int {
	// dp[i][j]表示word1 的前i个字符转换为 word2的前j个字符所需的步数
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j]= 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func min(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		} else {
			return z
		}
	} else {
		if y < z {
			return y
		} else {
			return z
		}
	}
}