package _2_generate_parenttheses

/*
	Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

	For example, given n = 3, a solution set is:

	[
	  "((()))",
	  "(()())",
	  "(())()",
	  "()(())",
	  "()()()"
	]
 */
var ans []string

func GenerateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	backTrace("", '(', n, 0, 0)
	return ans
}

func backTrace(s string, c byte, n, left, right int) {
	switch c {
	case '(':
		left++
	case ')':
		right++
	}
	if left < right || left > n || right > n {
		return
	}
	if left == n && right == n {
		ans = append(ans, s+string(c))
		return
	}
	backTrace(s+string(c), '(', n, left, right)
	backTrace(s+string(c), ')', n, left, right)
}

