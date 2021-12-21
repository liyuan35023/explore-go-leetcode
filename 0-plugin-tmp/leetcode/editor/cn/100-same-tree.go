package cn
//给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。 
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。 
//
// 示例 1： 
//
//输入：p = [1,2,3], q = [1,2,3]
//输出：true
//
// 示例 2： 
//
//输入：p = [1,2], q = [1,null,2]
//输出：false
//
// 示例 3： 
//
//输入：p = [1,2,1], q = [1,1,2]
//输出：false
//
// 提示： 
// 两棵树上的节点数目都在范围 [0, 100] 内
// -104 <= Node.val <= 104 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func(node1, node2 *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		return dfs(node1.Left, node2.Left) && dfs(node1.Right, node2.Right)
	}

	return dfs(p, q)




}
//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	}
//	if p == nil || q == nil || p.Val != q.Val {
//		return false
//	}
//	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
//
//}
//leetcode submit region end(Prohibit modification and deletion)
