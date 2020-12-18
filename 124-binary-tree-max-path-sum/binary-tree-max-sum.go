package _24_binary_tree_max_path_sum

import "math"

/*
	给定一个非空二叉树，返回其最大路径和。
	本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
	Example 1:

	Input: [1,2,3]

		   1
		  / \
		 2   3

	Output: 6

	Example 2:
	Input: [-10,9,20,null,null,15,7]

	   -10
	   / \
	  9  20
		/  \
	   15   7

	Output: 42
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// 递归的方法
	ans, _ := getMaxPathSumAndDepthPathSum(root)
    return ans
}

func getMaxPathSumAndDepthPathSum(root *TreeNode) (maxPathSum int, maxDepthPathSum int) {
    if root == nil {
        return math.MinInt32, 0
    }
    leftChildPathSum, leftDepthSum := getMaxPathSumAndDepthPathSum(root.Left)
    rightChildPathSum, rightDepthSum := getMaxPathSumAndDepthPathSum(root.Right)

    maxPathOverRoot := max(max(max(root.Val, leftDepthSum + root.Val), rightDepthSum + root.Val), leftDepthSum + root.Val + rightDepthSum)
    maxPathSum = max(max(leftChildPathSum, rightChildPathSum), maxPathOverRoot)
    maxDepthPathSum = max(max(leftDepthSum + root.Val, rightDepthSum + root.Val), root.Val)
    return
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func maxPathSumII(root *TreeNode) int {
    ans := math.MinInt32
    var maxGain func(node *TreeNode) int
    maxGain = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftGain := maxGain(node.Left)
        rightGain := maxGain(node.Right)
        pathSum := leftGain + rightGain + node.Val
        ans = max(ans, pathSum)
        return max(max(leftGain, rightGain) + node.Val, 0)
    }
    maxGain(root)
    return ans
}
