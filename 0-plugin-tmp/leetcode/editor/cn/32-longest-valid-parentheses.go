package cn

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
// 示例 1： 
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
// 示例 2：
// 
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
// 示例 3： 
//
//输入：s = ""
//输出：0
//
// 提示：
// 
// 0 <= s.length <= 3 * 104 
// s[i] 为 '(' 或 ')' 
//

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	ans := 0
	dp := make([]int, len(s)+1)
	for i := 2; i < len(s)+1; i++ {
		if s[i-1] == ')' {
			if s[i-2] == '(' {
				dp[i] = dp[i-2] + 2
			} else {
				leftParentIdx := i - 2 - dp[i-1]
				if leftParentIdx >= 0 && s[leftParentIdx] == '(' {
					dp[i] = dp[leftParentIdx] + dp[i-1] + 2
				}
			}
		}
		ans = max(ans, dp[i])
	}

	return ans

}
//func longestValidParentheses(s string) int {
//	// stack 表示上一个没有匹配的右括号的idx
//	stack := make([]int, 0)
//	stack = append(stack, -1)
//	ans := 0
//	for i := 0; i < len(s); i++ {
//		if s[i] == '(' {
//			stack = append(stack, i)
//		} else {
//			if len(stack) > 1 {
//				length := i - stack[len(stack) - 2]
//				ans = Max(ans, length)
//				stack = stack[:len(stack)-1]
//			} else {
//				stack[0] = i
//			}
//		}
//	}
//	return ans
//}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
