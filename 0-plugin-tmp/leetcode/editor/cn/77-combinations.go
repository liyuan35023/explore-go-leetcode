package cn
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。 
//
// 示例: 
//
// 输入:n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	var dfs func(idx int, solve []int)
	dfs = func(idx int, solve []int) {
		if len(solve) == k {
			ans = append(ans, append([]int{}, solve...))
			return
		}
		for i := idx; i <= n && k - len(solve) <= n - i + 1; i++ {
			solve = append(solve, i)
			dfs(i+1, solve)
			solve = solve[:len(solve)-1]
		}
	}
	dfs(1, []int{})
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
