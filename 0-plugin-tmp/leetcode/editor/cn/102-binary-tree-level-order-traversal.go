package cn
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
// 示例： 
//二叉树：[3,9,20,null,null,15,7], 
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 
// 返回其层序遍历结果：
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	var helper func(level int, node *TreeNode)
	helper = func(level int, node *TreeNode) {
		if node == nil {
			return
		}
		if level > len(ans) {
			ans = append(ans, make([]int, 0))
		}
		ans[level-1] = append(ans[level-1], node.Val)
		helper(level+1, node.Left)
		helper(level+1, node.Right)
	}
	helper(1, root)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
