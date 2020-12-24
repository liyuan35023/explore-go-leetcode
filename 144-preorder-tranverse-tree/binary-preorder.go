package _44_preorder_tranverse_tree

/*
	Example:


	Input: [1,null,2,3]
	   1
		\
		 2
		/
	   3

	Output: [1,2,3]

	Follow up: Recursive solution is trivial, could you do it iteratively?

	题目大意 #
	先根遍历一颗树。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func preorderTraversal(root *TreeNode) []int {
    ans := make([]int, 0)
    var helper func(node *TreeNode)
    helper = func(node *TreeNode) {
        if node == nil {
            return
        }
        ans = append(ans, node.Val)
        helper(node.Left)
        helper(node.Right)
    }
    helper(root)
    return ans
}

func preorderTraversalLoop(root *TreeNode) []int {
    ans := make([]int, 0)
    if root == nil {
        return ans
    }
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    for len(stack) != 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = append(ans, top.Val)
        if top.Right != nil {
            stack = append(stack, top.Right)
        }
        if top.Left != nil {
            stack = append(stack, top.Left)
        }
    }
    return ans
}


func preorderTraversalMorrisLoop(root *TreeNode) []int {
    ans := make([]int, 0)
    x := root
    for x != nil {
        ans = append(ans, x.Val)
        if x.Left == nil {
            x = x.Right
        } else {
            pre := x.Left
            for pre.Right != nil && pre.Right != x.Right {
                pre = pre.Right
            }
            if pre.Right == nil {
                pre.Right = x.Right
                x = x.Left
                continue
            }
            if pre.Right == x.Right {
                pre.Right = nil
                x = x.Right
                continue
            }
        }
    }
    return ans
}
