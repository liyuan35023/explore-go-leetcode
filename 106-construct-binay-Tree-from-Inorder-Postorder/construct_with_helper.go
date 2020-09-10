package _06_construct_binay_Tree_from_Inorder_Postorder

var inorderMap map[int]int
var InOrder []int
var PostOrder []int
var rootIndexInPost int

func buildTreeWithHelper(inorder []int, postorder []int) *TreeNode {
	inorderMap = make(map[int]int, len(inorder))

	for k, v := range inorder {
		inorderMap[v] = k
	}
	InOrder = inorder
	PostOrder = postorder
	rootIndexInPost = len(postorder)-1

	return helper(0, rootIndexInPost-1)
}

func helper(leftIn int, rightIn int) *TreeNode {
	if leftIn > rightIn {
		return nil
	}
	root := new(TreeNode)
	root.Val = PostOrder[rootIndexInPost]
	rootIndexInIn := inorderMap[root.Val]
	rootIndexInPost--

	right := helper(rootIndexInIn+1, rightIn)
	left := helper(leftIn, rootIndexInIn-1)
	root.Right = right
	root.Left = left
	return root
}
