package binary_tree_inorder_traversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	expect := []int{1, 3, 2}
	if !reflect.DeepEqual(inorderTraversal(root), expect) {
		t.Errorf("result not equal to expect")
	}
}
