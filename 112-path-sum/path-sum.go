package _12_path_sum
/*
	Example:
	Given the below binary tree and sum = 22,

		  5
		 / \
		4   8
	   /   / \
	  11  13  4
	 /  \      \
	7    2      1
	return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

	题目大意 #
	给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。说明: 叶子节点是指没有子节点的节点。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func hasPathSumLoop(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	nodeQueue := []*TreeNode{root}
	sumQueue := []int{sum}
	for len(nodeQueue) != 0 {
		n := len(nodeQueue)
		for i := 0; i < n; i++ {
			node := nodeQueue[i]
			if node.Left == nil && node.Right == nil && sumQueue[i] == node.Val {
				return true
			}
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
				sumQueue = append(sumQueue, sumQueue[i]-node.Val)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
				sumQueue = append(sumQueue, sumQueue[i]-node.Val)
			}
		}
		nodeQueue = nodeQueue[n:]
		sumQueue = sumQueue[n:]
	}
	return false
}

