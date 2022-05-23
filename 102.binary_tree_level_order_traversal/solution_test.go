package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
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
	res1 := levelOrder(test1)
	if !reflect.DeepEqual(res1, [][]int{
		{3},
		{9, 20},
		{15, 7},
	}) {
		t.Error("error")
	}

}
