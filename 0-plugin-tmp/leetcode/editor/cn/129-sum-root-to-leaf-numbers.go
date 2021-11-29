package cn
//给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//
// 每条从根节点到叶节点的路径都代表一个数字： 
//
// 
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。 
// 
//
// 计算从根节点到叶节点生成的 所有数字之和 。 
//
// 叶节点 是指没有子节点的节点。 
//
// 示例 1：
//
//输入：root = [1,2,3]
//输出：25
//解释：
//从根到叶子节点路径 1->2 代表数字 12
//从根到叶子节点路径 1->3 代表数字 13
//因此，数字总和 = 12 + 13 = 25 
//
// 示例 2： 
//
//输入：root = [4,9,0,5,1]
//输出：1026
//解释：
//从根到叶子节点路径 4->9->5 代表数字 495
//从根到叶子节点路径 4->9->1 代表数字 491
//从根到叶子节点路径 4->0 代表数字 40
//因此，数字总和 = 495 + 491 + 40 = 1026
//
// 提示： 
//
// 
// 树中节点的数目在范围 [1, 1000] 内 
// 0 <= Node.val <= 9 
// 树的深度不超过 10 
// 
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode, value int) int
	dfs = func(node *TreeNode, value int) int {
		if node.Left == nil && node.Right == nil {
			return value*10 + node.Val
		}
		var l, r int
		if node.Left != nil {
			l = dfs(node.Left, value * 10 + node.Val)
		}
		if node.Right != nil {
			r = dfs(node.Right, value * 10 + node.Val)
		}
		return l + r
	}
	return dfs(root, 0)
}
//func sumNumbers(root *TreeNode) int {
//	ans := 0
//	if root == nil {
//		return ans
//	}
//	queue := make([]*TreeNode, 0)
//	queue = append(queue, root)
//	valQueue := make([]int, 0)
//	valQueue = append(valQueue, root.Val)
//	for len(queue) != 0 {
//		node := queue[0]
//		val := valQueue[0]
//		queue = queue[1:]
//		valQueue = valQueue[1:]
//		if node.Left == nil && node.Right == nil {
//			ans += val
//		}
//		if node.Left != nil {
//			queue = append(queue, node.Left)
//			valQueue = append(valQueue, val * 10 + node.Left.Val)
//		}
//		if node.Right != nil {
//			queue = append(queue, node.Right)
//			valQueue = append(valQueue, val * 10 + node.Right.Val)
//		}
//	}
//	return ans
//}
//func sumNumbers(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var dfs func(value int, node *TreeNode) int
//	dfs = func(value int, node *TreeNode) int {
//		var l, r int
//		if node == nil {
//			return 0
//		}
//		if node.Left == nil && node.Right == nil {
//			return value * 10 + node.Val
//		}
//		l = dfs(value * 10 + node.Val, node.Left)
//		r = dfs(value * 10 + node.Val, node.Right)
//		return l + r
//	}
//	return dfs(0, root)
//}
//leetcode submit region end(Prohibit modification and deletion)
