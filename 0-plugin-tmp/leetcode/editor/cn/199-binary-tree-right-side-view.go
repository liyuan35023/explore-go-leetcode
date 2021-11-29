package cn
//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。 
//
// 示例 1: 
//
//输入:[1,2,3,null,5,null,4]
//输出:[1,3,4]
//
// 示例 2: 
//
//输入:[1,null,3]
//输出:[1,3]
//
// 示例 3: 
//
//输入:[]
//输出:[]
// 
// 提示:
//
// 二叉树的节点个数的范围是 [0,100]
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
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if i == n - 1 {
				ans = append(ans, queue[i].Val)
			}
		}
		queue = queue[n:]
	}
	return ans
}
//func rightSideView(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	ans := make([]int, 0)
//	var helper func(node *TreeNode, level int)
//	helper = func(node *TreeNode, level int) {
//		if node == nil {
//			return
//		}
//		if len(ans) < level {
//			ans = append(ans, node.Val)
//		}
//		helper(node.Right, level+1)
//		helper(node.Left, level+1)
//	}
//	helper(root, 1)
//	return ans
//}
//leetcode submit region end(Prohibit modification and deletion)
