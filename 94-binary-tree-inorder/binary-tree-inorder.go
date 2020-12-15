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
            // 这里修改了树的结构
            top.Left = nil
        }
    }
    return ans
}
func InorderTraversalLoopII(root *TreeNode) []int {
    ans := make([]int, 0)
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = append(ans, root.Val)
        root = root.Right
    }
	return ans
}

func InorderTraversalLoopMorris(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
	    return ans
    }
    x := root
    for x != nil {
        if x.Left == nil {
            ans = append(ans, x.Val)
            x = x.Right
        } else {
            // find pre
            pre := x.Left
            for pre.Right != nil && pre.Right != x {
                pre = pre.Right
            }
            if pre.Right == x {
                pre.Right = nil
                ans = append(ans, x.Val)
                x = x.Right
            } else {
                pre.Right = x
                x = x.Left
            }
        }
    }
    return ans
}
