package cn

//给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bs
//t/ 相同 
//
// 示例 1： 
//
//输入：root = [4,2,6,1,3]
//输出：1
//
// 示例 2： 
//
//输入：root = [1,0,48,null,null,12,49]
//输出：1
//
// 提示： 
//
// 树中节点数目在范围 [2, 100] 内 
// 0 <= Node.val <= 105 
// 差值是一个正数，其数值等于两值之差的绝对值 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	ans := 1<<31 - 1
	var dfs func(node *TreeNode) (minVal, maxVal int)
	dfs = func(node *TreeNode) (minVal, maxVal int) {
		if node == nil {
			return -1, -1
		}
		lMin, lMax := dfs(node.Left)
		rMin, rMax := dfs(node.Right)
		if lMax >= 0 {
			ans = min(ans, abs(node.Val-lMax))
		}
		if rMin >= 0 {
			ans = min(ans, abs(node.Val-rMin))
		}
		if lMin < 0 {
			lMin = node.Val
		}
		if rMax < 0 {
			rMax = node.Val
		}
		return lMin, rMax
	}
	dfs(root)
	return ans
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}
//leetcode submit region end(Prohibit modification and deletion)
