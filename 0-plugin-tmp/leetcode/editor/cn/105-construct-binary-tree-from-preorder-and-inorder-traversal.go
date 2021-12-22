package cn
//给定一棵树的前序遍历 preorder 与中序遍历 inorder。请构造二叉树并返回其根节点。 
//
// 示例 1: 
//
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
//
// 示例 2: 
//
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
//
// 提示: 
//
// 1 <= preorder.length <= 3000 
// inorder.length == preorder.length 
// -3000 <= preorder[i], inorder[i] <= 3000 
// preorder 和 inorder 均无重复元素 
// inorder 均出现在 preorder 
// preorder 保证为二叉树的前序遍历序列 
// inorder 保证为二叉树的中序遍历序列 
//

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for k, v := range inorder {
		inorderMap[v] = k
	}
	idx := 0
	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if idx >= len(preorder) || left > right {
			return nil
		}
		mid := inorderMap[preorder[idx]]
		node := &TreeNode{Val: preorder[idx]}
		idx++
		node.Left = dfs(left, mid-1)
		node.Right = dfs(mid+1, right)
		return node
	}
	return dfs(0, len(preorder)-1)

}
//leetcode submit region end(Prohibit modification and deletion)
