package cn

//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不
//一定经过根节点。 
//
// 路径和 是路径中各节点值的总和。 
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。 
//
// 示例 1： 
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6 
//
// 示例 2： 
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
// 提示： 
//
// 树中节点数目范围是 [1, 3 * 104] 
// -1000 <= Node.val <= 1000 
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
func maxPathSum(root *TreeNode) int {
	ans := -1 << 31
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lGain := maxGain(node.Left)
		rGain := maxGain(node.Right)
		ans = max(ans, lGain+node.Val+rGain)
		return max(0, max(lGain, rGain) + node.Val)
	}

	maxGain(root)
	return ans












}
//func maxPathSum(root *TreeNode) int {
//	ans := -1 << 31
//	var maxGain func(node *TreeNode) int
//	maxGain = func(node *TreeNode) int {
//		if node == nil {
//			return 0
//		}
//		l := maxGain(node.Left)
//		r := maxGain(node.Right)
//		ans = max(ans, l+r+node.Val)
//		return max(0, max(l, r) + node.Val)
//	}
//	maxGain(root)
//	return ans
//}
//
func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

//leetcode submit region end(Prohibit modification and deletion)
