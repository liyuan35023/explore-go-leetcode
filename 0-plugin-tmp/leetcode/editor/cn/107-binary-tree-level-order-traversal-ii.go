package cn
//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历） 
//
// 例如： 
//给定二叉树 [3,9,20,null,null,15,7], 
//
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回其自底向上的层序遍历为： 
//
//[
//  [15,7],
//  [9,20],
//  [3]
//]
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

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > len(ans) {
			tmp := []int{node.Val}
			ans = append([][]int{tmp}, ans...)
		} else {
			ans[len(ans)-level] = append(ans[len(ans)-level], node.Val)
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 1)
	return ans

}
//leetcode submit region end(Prohibit modification and deletion)
