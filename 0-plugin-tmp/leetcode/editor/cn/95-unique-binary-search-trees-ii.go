package cn
//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。 
//
// 示例 1：
//
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
// 
// 示例 2：
//
//输入：n = 1
//输出：[[1]]
// 
// 提示：
// 1 <= n <= 8

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	var dfs func(start, end int) []*TreeNode
	dfs = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		if start == end {
			return []*TreeNode{{Val: start}}
		}
		ans := make([]*TreeNode, 0)
		for i := start; i <= end; i++ {
			left := dfs(start, i-1)
			right := dfs(i+1, end)
			for _, l := range left {
				for _, r := range right {
					tmp := &TreeNode{Val: i, Left: l, Right: r}
					ans = append(ans, tmp)
				}
			}
		}
		return ans
	}
	return dfs(1, n)
}
//leetcode submit region end(Prohibit modification and deletion)
