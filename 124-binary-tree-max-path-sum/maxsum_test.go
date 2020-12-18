package _24_binary_tree_max_path_sum

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	root := &TreeNode{Val: 1, Left: node2, Right: node3}
	fmt.Println(maxPathSumII(root))
}

