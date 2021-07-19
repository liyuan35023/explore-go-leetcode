package cn
//给定一个二叉树，找出其最小深度。 
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。 
//
// 说明：叶子节点是指没有子节点的节点。 
//
// 示例 1： 
//
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：2
//
// 示例 2： 
//
//输入：root = [2,null,3,null,4,null,5,null,6]
//输出：5
//
// 提示： 
//
// 树中节点数的范围在 [0, 105] 内
// -1000 <= Node.val <= 1000 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	for len(queue) != 0 {
		n := len(queue)
		depth++
		for i := 0; i < n; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return depth
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	return depth
}

//leetcode submit region end(Prohibit modification and deletion)
