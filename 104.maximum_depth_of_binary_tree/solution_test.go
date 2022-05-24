package maximum_depth_of_binary_tree

import "testing"

func TestMaxDepth(t *testing.T) {
	test1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	if maxDepth(test1) != 3 {
		t.Error("error")
	}
}
