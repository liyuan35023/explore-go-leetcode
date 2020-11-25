package _26_invert_binary_tree

/*
	Example:

	Input:


		 4
	   /   \
	  2     7
	 / \   / \
	1   3 6   9

	Output:

		 4
	   /   \
	  7     2
	 / \   / \
	9   6 3   1

	Trivia:

	This problem was inspired by this original tweet by Max Howell:

	Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.

	题目大意 #
	“经典"的反转二叉树的问题。

 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// 递归
    if root == nil {
        return nil
    }
    l := invertTree(root.Left)
    r := invertTree(root.Right)
    root.Left = r
    root.Right = l
    return root
}

func invertTreeLoop(root *TreeNode) *TreeNode {
    // 循环
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        top := queue[0]
        queue = queue[1:]
        if top.Left != nil {
            queue = append(queue, top.Left)
        }
        if top.Right != nil {
            queue = append(queue, top.Right)
        }
        top.Left, top.Right = top.Right, top.Left
    }
    return root
}
