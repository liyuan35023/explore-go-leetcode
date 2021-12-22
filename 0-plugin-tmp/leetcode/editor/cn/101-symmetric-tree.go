package cn
//给定一个二叉树，检查它是否是镜像对称的。 
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。 
//
//     1
//   / \
//  2   2
// / \ / \
//3  4 4  3
// 
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的: 
//
//     1
//   / \
//  2   2
//   \   \
//   3    3
//
// 进阶： 


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	var dfs func(left *TreeNode, right *TreeNode) bool
	dfs = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root, root)


}


//func isSymmetric(root *TreeNode) bool {
//	var dfs func(node1 *TreeNode, node2 *TreeNode) bool
//	dfs = func(node1 *TreeNode, node2 *TreeNode) bool {
//		if node1 == nil && node2 == nil {
//			return true
//		}
//		if node1 == nil || node2 == nil || node1.Val != node2.Val {
//			return false
//		}
//		return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
//	}
//	return dfs(root, root)
//}
//func isSymmetric(root *TreeNode) bool {
//	queue := make([]*TreeNode, 0)
//	queue = append(queue, root)
//	for len(queue) != 0 {
//		n := len(queue)
//		if n == 1 {
//			queue = append(queue, queue[0].Left)
//			queue = append(queue, queue[0].Right)
//		} else {
//			for i := 0; i < n - 1; i = i + 2 {
//				if queue[i] == nil && queue[i+1] == nil {
//					continue
//				}
//				if queue[i] == nil || queue[i+1] == nil || queue[i].Val != queue[i+1].Val {
//					return false
//				}
//				queue = append(queue, queue[i].Left)
//				queue = append(queue, queue[i+1].Right)
//				queue = append(queue, queue[i].Right)
//				queue = append(queue, queue[i+1].Left)
//			}
//		}
//		queue = queue[n:]
//	}
//	return true
//}
//leetcode submit region end(Prohibit modification and deletion)
