package serialize_and_deserialize_BST

import (
	"reflect"
	"testing"
)

func TestTreeNodeSerialize(t *testing.T) {
	root := &TreeNode{Val: 2, Left: &TreeNode{
		Val: 1,
	}, Right: &TreeNode{Val: 3}}
	ser := Constructor()
	deser := Constructor()
	tree := ser.serialize(root)
	ans := deser.deserialize(tree)
	if !reflect.DeepEqual(ans, root) {
		t.Errorf("error")
	}
}
