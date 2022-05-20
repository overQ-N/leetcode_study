package recover_binary_search_tree

import (
	"reflect"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	node1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	recoverTree(node1)
	result1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	if !reflect.DeepEqual(node1, result1) {
		t.Error("error")
	}
}
