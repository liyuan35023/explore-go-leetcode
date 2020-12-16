package _04_maximum_depth_binary_tree

/*
    Example:

    Given binary tree [3,9,20,null,null,15,7],


        3
       / \
      9  20
        /  \
       15   7

    return its depth = 3.

    题目大意 #
    要求输出一棵树的最大高度。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l := maxDepth(root.Left)
    r := maxDepth(root.Right)
    return max(l, r) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func maxDepthLoop(root *TreeNode) int {
    ans := 0
    if root == nil {
        return ans
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) != 0 {
        ans++
        n := len(queue)
        for i := 0; i < n; i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        queue = queue[n:]
    }
    return ans
}


