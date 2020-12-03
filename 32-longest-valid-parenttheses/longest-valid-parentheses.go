package _2_longest_valid_parenttheses

/*
	给定一个只包含 '('和 ')'的字符串，找出最长的包含有效括号的子串的长度。

	示例1:

	输入: "(()"
	输出: 2
	解释: 最长有效括号子串为 "()"
	示例 2:

	输入: ")()())"
	输出: 4
	解释: 最长有效括号子串为 "()()"

 */

func longestValidParenthesesDp(s string) int {
	ans := 0
	n := len(s)
	if n < 2 {
		return 0
	}
	// dp[i]表示以第i个(i为index)括号为结尾的最长有效子串长度
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				if i - dp[i-1] -1 >= 0 && s[i-dp[i-1]-1] == '(' {
					if i - dp[i-1] - 2 >= 0 {
						dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}

			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func longestValidParentheses(s string) int {
	// stack 的方法
	// 栈底的元素表示字符串中最后一个没有被匹配的右括号的下标
	ans := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
