package cn
//给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。 
//
// 例如： 
//给定二叉树 [3,9,20,null,null,15,7], 
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回锯齿形层序遍历如下： 
//
//[
//  [3],
//  [20,9],
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	left := true
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		n := len(queue)
		level := make([]int, n)

		for i := 0; i < n; i++ {
			if left {
				level[i] = queue[i].Val
			} else {
				level[n-i-1] = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		ans = append(ans, level)
		queue = queue[n:]
		left = !left
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
