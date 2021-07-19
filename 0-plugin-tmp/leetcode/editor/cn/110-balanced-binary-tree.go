package cn
//给定一个二叉树，判断它是否是高度平衡的二叉树。 
//
// 本题中，一棵高度平衡二叉树定义为： 
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。 
//
// 示例 1： 
//
//输入：root = [3,9,20,null,null,15,7]
//输出：true
// 
// 示例 2：
//
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
//
// 示例 3： 
//
//输入：root = []
//输出：true
//
// 提示： 
//
// 树中的节点数在范围 [0, 5000] 内 
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
func isBalanced(root *TreeNode) bool {
	_, balance := getLevel(root)
	return balance
}

func getLevel(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, lBalance := getLevel(root.Left)
	r, rBalance := getLevel(root.Right)
	if !lBalance || !rBalance || l - r > 1 || r - l > 1 {
		return 0, false
	}
	return max(l, r) + 1, true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)