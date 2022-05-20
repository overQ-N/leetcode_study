package sometree

import "testing"

func TestIsSameTree(t *testing.T) {
	if isSameTree(&TreeNode{Val: 0, Left: &TreeNode{Val: -5}}, &TreeNode{Val: 0, Left: &TreeNode{Val: -8}}) {
		t.Error("error")
	}
}
