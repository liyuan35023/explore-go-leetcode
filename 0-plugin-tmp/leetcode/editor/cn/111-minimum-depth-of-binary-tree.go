package cn
//给定一个二叉树，找出其最小深度。 
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。 
//
// 说明：叶子节点是指没有子节点的节点。 
//
// 示例 1： 
//
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：2
//
// 示例 2： 
//
//输入：root = [2,null,3,null,4,null,5,null,6]
//输出：5
//
// 提示： 
//
// 树中节点数的范围在 [0, 105] 内
// -1000 <= Node.val <= 1000 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return 1
		}
		if node.Left == nil {
			return dfs(node.Right) + 1
		} else if node.Right == nil {
			return dfs(node.Left) + 1
		} else {
			return min(dfs(node.Left), dfs(node.Right)) + 1
		}
	}
	return dfs(root)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

//leetcode submit region end(Prohibit modification and deletion)
