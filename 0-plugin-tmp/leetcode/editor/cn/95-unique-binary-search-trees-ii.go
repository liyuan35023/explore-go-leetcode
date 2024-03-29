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
	var dfs func(left, right int) []*TreeNode
	dfs = func(left, right int) []*TreeNode {
		if left > right {
			return []*TreeNode{nil}
		}
		ans := make([]*TreeNode, 0)
		for i := left; i <= right; i++ {
			l, r := dfs(left, i-1), dfs(i+1, right)
			for _, lNode := range l {
				for _, rNode := range r {
					node := &TreeNode{Val: i, Left: lNode, Right: rNode}
					ans = append(ans, node)
				}
			}

		}
		return ans
	}
	return dfs(1, n)

}
//leetcode submit region end(Prohibit modification and deletion)
