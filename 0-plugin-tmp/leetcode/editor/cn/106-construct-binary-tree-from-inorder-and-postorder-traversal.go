package cn

//根据一棵树的中序遍历与后序遍历构造二叉树。
//
// 注意: 
//你可以假设树中没有重复的元素。 
//
// 例如，给出 
//
// 中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3] 
//
// 返回如下的二叉树： 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
// Related Topics 树 数组 哈希表 分治 二叉树 
// 👍 531 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	inMap := make(map[int]int)
	for k, v := range inorder {
		inMap[v] = k
	}
	var buildHelper func(leftIn, rightIn int) *TreeNode
	buildHelper = func(leftIn, rightIn int) *TreeNode {
		if leftIn > rightIn {
			return nil
		}
		node := &TreeNode{Val: postorder[len(postorder)-1]}
		postorder = postorder[:len(postorder)-1]
		idx := inMap[node.Val]
		node.Right = buildHelper(idx+1, rightIn)
		node.Left = buildHelper(leftIn, idx-1)
		return node
	}

	return buildHelper(0, len(inorder)-1)
}
//func buildTree(inorder []int, postorder []int) *TreeNode {
//	inMap := make(map[int]int)
//	for k, v := range inorder {
//		inMap[v] = k
//	}
//	var helper func(in, post []int, former int) *TreeNode
//	helper = func(in, post []int, former int) *TreeNode {
//		if len(in) == 0 || len(post) == 0 {
//			return nil
//		}
//		root := &TreeNode{Val: post[len(post)-1]}
//		rootIdxInInorder := inMap[post[len(post)-1]]
//		realRootIdx := rootIdxInInorder - former
//		root.Left = helper(in[:realRootIdx], post[:realRootIdx], former)
//		root.Right = helper(in[realRootIdx+1:], post[realRootIdx:len(post)-1], former+realRootIdx+1)
//		return root
//	}
//	return helper(inorder, postorder, 0)
//}
//leetcode submit region end(Prohibit modification and deletion)
