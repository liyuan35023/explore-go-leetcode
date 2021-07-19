package cn

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。 
//
// 示例 1： 
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//
// 示例 2： 
//
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//
// 示例 3： 
//
//输入：root = [1,2], targetSum = 0
//输出：[]
//
// 提示： 
// 树中节点总数在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000 
// -1000 <= targetSum <= 1000 

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	target := []int{targetSum}
	parent := make(map[*TreeNode]*TreeNode)
	parent[root] = nil
	generatePath := func(node *TreeNode) []int {
		path := make([]int, 0)
		for node != nil {
			path = append([]int{node.Val}, path...)
			node = parent[node]
		}
		return path
	}

	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				if queue[i].Val == target[i] {
					path := generatePath(queue[i])
					ans = append(ans, path)
				}
				continue
			}
			if queue[i].Left != nil {
				parent[queue[i].Left] = queue[i]
				target = append(target, target[i]-queue[i].Val)
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				parent[queue[i].Right] = queue[i]
				target = append(target, target[i]-queue[i].Val)
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		target = target[n:]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
