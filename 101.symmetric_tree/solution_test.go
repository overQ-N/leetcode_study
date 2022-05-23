package symmetric_tree

import "testing"

func TestIsSymmetric(t *testing.T) {
	test1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	if !isSymmetric(test1) {
		t.Error("error")
	}
	test2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	if isSymmetric(test2) {
		t.Error("error")
	}
	if !isSymmetric(&TreeNode{Val: 1}) {
		t.Error("error")
	}
}
