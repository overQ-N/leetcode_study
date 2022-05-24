package univalued_binary_tree

import "testing"

func TestIsUnivalTree(t *testing.T) {
	test1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	if !isUnivalTree(test1) {
		t.Error("error")
	}
}
