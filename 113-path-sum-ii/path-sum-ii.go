package _13_path_sum_ii

/*
	Example:

	Given the below binary tree and sum = 22,

		  5
		 / \
		4   8
	   /   / \
	  11  13  4
	 /  \    / \
	7    2  5   1
	Return:

	[
	   [5,4,11,2],
	   [5,8,4,5]
	]
	题目大意 #
	给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。说明: 叶子节点是指没有子节点的节点。
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 解法一 递归 dfs
func pathSum(root *TreeNode, sum int) [][]int {
    ans := make([][]int, 0)
    if root == nil {
        return ans
    }
    path := make([]int, 0)
    recurHelper(root, sum, path, &ans)
    return ans
}

func recurHelper(root *TreeNode, sum int, path []int, ans *[][]int) {
    if root == nil {
        return
    }
    path = append(path, root.Val)
    if root.Left == nil && root.Right == nil {
        if sum == root.Val {
            tmp := append([]int{}, path...)
            *ans = append(*ans, tmp)
        }
        return
    }
    recurHelper(root.Left, sum-root.Val, path, ans)
    recurHelper(root.Right, sum-root.Val, path, ans)
}

// bfs
func pathSum2(root *TreeNode, sum int) [][]int {
    ans := make([][]int, 0)
    if root == nil {
        return ans
    }

    nodeQueue := []*TreeNode{root}
    sumQueue := []int{sum}
    parent := make(map[*TreeNode]*TreeNode)
    appendPath := func(leaf *TreeNode) {
        path := make([]int, 0)
        for leaf != nil {
            path = append([]int{leaf.Val}, path...)
            leaf = parent[leaf]
        }
        ans = append(ans, path)
    }
    for len(nodeQueue) != 0 {
    	first := nodeQueue[0]
    	nodeQueue = nodeQueue[1:]
    	sum := sumQueue[0]
    	sumQueue = sumQueue[1:]

        if first.Left == nil && first.Right == nil {
            if sum == first.Val {
            	appendPath(first)
            }
        }
        if first.Left != nil {
            parent[first.Left] = first
            nodeQueue = append(nodeQueue, first.Left)
            sumQueue = append(sumQueue, sum-first.Val)
        }
        if first.Right != nil {
            parent[first.Right] = first
            nodeQueue = append(nodeQueue, first.Right)
            sumQueue = append(sumQueue, sum-first.Val)
        }
    }
    return ans
}
