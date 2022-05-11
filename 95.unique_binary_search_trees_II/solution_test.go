package unique_binary_search_trees_II

import (
	"reflect"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	res1 := generateTrees(1)
	expect1 := []*TreeNode{{Val: 1}}
	if !reflect.DeepEqual(res1, expect1) {
		t.Errorf("res1 not equal to expect1")
	}

	res2 := generateTrees(3)
	if len(res2) != 5 {
		t.Errorf("error")
	}
}
