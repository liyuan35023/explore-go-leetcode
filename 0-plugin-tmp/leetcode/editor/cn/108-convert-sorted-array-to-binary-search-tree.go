package cn
//给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。 
//
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。 
//
// 示例 1： 
//
//输入：nums = [-10,-3,0,5,9]
//输出：[0,-3,9,-10,null,5]
//解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
//
// 示例 2： 
//
//输入：nums = [1,3]
//输出：[3,1]
//解释：[1,3] 和 [3,1] 都是高度平衡二叉搜索树。
//
// 提示： 
//
// 1 <= nums.length <= 104 
// -104 <= nums[i] <= 104 
// nums 按 严格递增 顺序排列 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
	idx := 0
	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right - left + 1) / 2
		l := dfs(left, mid-1)
		node := &TreeNode{Val: nums[idx]}
		idx++
		node.Left = l
		node.Right = dfs(mid+1, right)
		return node
	}
	return dfs(0, len(nums)-1)

}
//func sortedArrayToBST(nums []int) *TreeNode {
//	if len(nums) <= 0 {
//		return nil
//	}
//	mid := len(nums) / 2
//	root := &TreeNode{Val: nums[mid]}
//	root.Left = sortedArrayToBST(nums[:mid])
//	root.Right = sortedArrayToBST(nums[mid+1:])
//	return root
//}
//leetcode submit region end(Prohibit modification and deletion)
