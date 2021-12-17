package cn

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征： 
//
// 节点的左子树只包含小于当前节点的数。 
// 节点的右子树只包含大于当前节点的数。 
// 所有左子树和右子树自身必须也是二叉搜索树。 
//
// 示例 1: 
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
// 示例 2: 
//
// 输入:
//    5
//   / \
//  1   4
//    / \
//   3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//    根节点的值为 5 ，但是其右子节点值为 4 。

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {








}
//func isValidBST(root *TreeNode) bool {
//	// morris inorder
//	preVal := -1<<31 - 1
//	for root != nil {
//		if root.Left == nil {
//			if root.Val <= preVal {
//				return false
//			}
//			preVal = root.Val
//			root = root.Right
//		} else {
//			x := root.Left
//			for x.Right != nil && x.Right != root {
//				x = x.Right
//			}
//			if x.Right == nil {
//				x.Right = root
//				root = root.Left
//			} else {
//				if root.Val <= preVal {
//					return false
//				}
//				preVal = root.Val
//				root = root.Right
//				x.Right = nil
//			}
//		}
//	}
//	return true
//}
//func isValidBST(root *TreeNode) bool {
//	// inorder
//	stack := make([]*TreeNode, 0)
//	pre := math.MinInt32 - 1
//	for root != nil || len(stack) != 0 {
//		for root != nil {
//			stack = append(stack, root)
//			root = root.Left
//		}
//		root = stack[len(stack)-1]
//		if pre >= root.Val {
//			return false
//		}
//		pre = root.Val
//		stack = stack[:len(stack)-1]
//		root = root.Right
//	}
//	return true
//}
//leetcode submit region end(Prohibit modification and deletion)
