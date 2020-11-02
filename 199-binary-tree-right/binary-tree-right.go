package _99_binary_tree_right

/*
	给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

	示例:

	输入:[1,2,3,null,5,null,4]
	输出:[1, 3, 4]
	解释:

	   1            <---
	 /   \
	2     3         <---
	 \     \
	  5     4       <---

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	rangeHelper(root, 1, &ans)
	return ans
}

func rangeHelper(root *TreeNode, depth int, ans *[]int) {
	if root == nil {
		return
	}
	if len(*ans) < depth {
		*ans = append(*ans, root.Val)
	}
	if root.Right != nil {
		rangeHelper(root.Right, depth+1, ans)
	}
	if root.Left != nil {
		rangeHelper(root.Left, depth+1, ans)
	}
}

func rightSideView2(root *TreeNode) []int {
	// bfs 广搜
	ans := make([]int, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		thisDepthNodeNum := len(queue)
		for i := 0; i < thisDepthNodeNum; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if i == thisDepthNodeNum-1 {
				ans = append(ans, queue[i].Val)
			}
		}
		queue = queue[thisDepthNodeNum:]
	}
	return ans
}
