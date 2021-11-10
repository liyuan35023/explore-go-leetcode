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
	var dfs func(solve string, left, right int)
	dfs = func(solve string, left, right int) {
		if left == n && right == n {
			ans = append(ans, solve)
			return
		}
		if left > n || right > n || left < right {
			return
		}
		dfs(solve+"(", left+1, right)
		dfs(solve+")", left, right+1)
	}
	dfs("", 0, 0)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
