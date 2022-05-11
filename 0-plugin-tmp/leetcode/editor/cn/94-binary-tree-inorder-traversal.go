package cn
//给定一个二叉树的根节点 root ，返回它的 中序 遍历。 
//
// 示例 1： 
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
// 示例 2： 
//
//输入：root = []
//输出：[]
// 
// 示例 3：
//
//输入：root = [1]
//输出：[1]
//
// 示例 4： 
//
//输入：root = [1,2]
//输出：[2,1]
//
// 示例 5： 
//
//输入：root = [1,null,2]
//输出：[1,2]
//
// 提示： 
//
// 树中节点数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	for root != nil {
		if root.Left == nil {
			ans = append(ans, root.Val)
			root = root.Right
		} else {
			x := root.Left
			for x.Right != nil && x.Right != root {
				x = x.Right
			}
			if x.Right == nil {
				x.Right = root
				root = root.Left
			} else {
				ans = append(ans, root.Val)
				x.Right = nil
				root = root.Right
			}
		}

	}
	return ans







}
//func inorderTraversal(root *TreeNode) []int {
//	ans := make([]int, 0)
//	for root != nil {
//		if root.Left == nil {
//			ans = append(ans, root.Val)
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
//				ans = append(ans, root.Val)
//				x.Right = nil
//				root = root.Right
//			}
//		}
//	}
//	return ans
//}
//leetcode submit region end(Prohibit modification and deletion)
