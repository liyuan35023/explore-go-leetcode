package cn
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。 
//
// 示例 1： 
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
// 
// 示例 2：
//
//输入：n = 1
//输出：["()"]
// 
// 提示：
// 1 <= n <= 8
// 

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	var dfs func(left, right int, brackets []byte)
	dfs = func(left, right int, brackets []byte) {
		if left == right && left == n {
			ans = append(ans, string(brackets))
			return
		}
		if left < right || left > n {
			return
		}
		dfs(left+1, right, append(brackets, '('))
		dfs(left, right+1, append(brackets, ')'))
	}
	dfs(1, 0, []byte{'('})
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
