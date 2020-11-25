package _8_validate_binary_search_tree

import "math"

/*
	Example 1:

		2
	   / \
	  1   3

	Input: [2,1,3]
	Output: true

	Example 2:
		5
	   / \
	  1   4
		 / \
		3   6

	Input: [5,1,4,null,null,3,6]
	Output: false
	Explanation: The root node's value is 5 but its right child's value is 4.
	题目大意 #
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。假设一个二叉搜索树具有如下特征：

	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    // 递归啊
    if root == nil {
        return true
    }
    bst, _, _ := bstHelper(root)
    return bst
}

func bstHelper(node *TreeNode) (bool, int, int) {
    var leftMin, leftMax, rightMin, rightMax int
    bstLeft, bstRight := true, true
    if node.Left != nil {
        bstLeft, leftMin, leftMax = bstHelper(node.Left)
    } else {
        leftMin, leftMax = node.Val, node.Val
    }
    if node.Right != nil {
        bstRight, rightMin, rightMax = bstHelper(node.Right)
    } else {
        rightMin, rightMax = node.Val, node.Val
    }
    if !bstLeft || !bstRight || (node.Left != nil && node.Val <= leftMax) || (node.Right != nil && node.Val >= rightMin) {
        return false, 0, 0
    }
    return true, leftMin, rightMax
}

func isValidBSTInorder(root *TreeNode) bool {
    //中序遍历
    //inorder := make([]int, 0)
    //return inOrderHelper(root, &inorder)
	stack := make([]*TreeNode, 0)
	pre := math.MinInt64
	for root != nil || len(stack) != 0 {
	    for root != nil {
	        stack = append(stack, root)
	        root = root.Left
        }
	    root = stack[len(stack)-1]
	    stack = stack[:len(stack)-1]
	    if root.Val <= pre {
	    	return false
		}
		pre = root.Val
	    root = root.Right
    }
	return true
}

func inOrderHelper(node *TreeNode, inorder *[]int) bool {
	if node == nil {
        return true
    }
    left := inOrderHelper(node.Left, inorder)
	if !left || len(*inorder) > 0 && node.Val <= (*inorder)[len(*inorder)-1] {
	    return false
    }
    *inorder = append(*inorder, node.Val)
    right := inOrderHelper(node.Right, inorder)
    if !right {
        return false
    }
    return true
}

