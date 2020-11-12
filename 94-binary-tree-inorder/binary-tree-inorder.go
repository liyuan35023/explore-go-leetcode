package _4_binary_tree_inorder

/*
	Example:

	Input: [1,null,2,3]
	   1
		\
		 2
		/
	   3

	Output: [1,3,2]

	Follow up: Recursive solution is trivial, could you do it iteratively?

	题目大意 #
	中根遍历一颗树。

 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    ans := make([]int, 0)
    inorderHelper(root, &ans)
    return ans
}

func inorderHelper(root *TreeNode, ans *[]int) {
    // 递归实现中序遍历
    if root == nil {
        return
    }
    inorderHelper(root.Left, ans)
    *ans = append(*ans, root.Val)
    inorderHelper(root.Right, ans)
}

func InorderTraversalLoop(root *TreeNode) []int {
    ans := make([]int, 0)
    if root == nil {
        return ans
    }
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)

    for len(stack) != 0 {
        n := len(stack)
        top := stack[n-1]
        if top.Left == nil {
            ans = append(ans, top.Val)
            stack = stack[:n-1]
            if top.Right != nil {
                stack = append(stack, top.Right)
            }
        } else {
            stack = append(stack, top.Left)
            top.Left = nil
        }
    }
    return ans
}
