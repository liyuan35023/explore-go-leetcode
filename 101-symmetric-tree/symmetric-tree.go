package _01_symmetric_tree

/*
	For example, this binary tree [1,2,2,3,4,4,3] is symmetric:


		1
	   / \
	  2   2
	 / \ / \
	3  4 4  3

	But the following [1,2,2,null,3,null,3] is not:

		1
	   / \
	  2   2
	   \   \
	   3    3

	Note:
	Bonus points if you could solve it both recursively and iteratively.
	题目大意 #
	这一题要求判断 2 颗树是否是左右对称的。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSymmetricRec(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return recHelper(root.Left, root.Right)
}

func recHelper(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil || left.Val != right.Val {
        return false
    }
    return recHelper(left.Left, right.Right) && recHelper(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	// 迭代
    if root == nil {
        return true
    }
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        n := len(queue)
        if n == 1 {
            queue = append(queue, queue[0].Left)
            queue = append(queue, queue[0].Right)
        }
        for i := 0; i < n-1; i=i+2 {
            if queue[i] == nil && queue[i+1] == nil {
                continue
            }
            if queue[i] == nil || queue[i+1] == nil || queue[i].Val != queue[i+1].Val {
                return false
            }
            queue = append(queue, queue[i].Left)
            queue = append(queue, queue[i+1].Right)
            queue = append(queue, queue[i].Right)
            queue = append(queue, queue[i+1].Left)
        }
        queue = queue[n:]
    }
    return true
}
