package cn
//给定一个二叉树，返回它的 后序 遍历。 
//
// 示例: 
//
// 输入: [1,null,2,3]  
//   1
//    \
//     2
//    /
//   3 
//
//输出: [3,2,1] 
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
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || pre == root.Right {
			ans = append(ans, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
