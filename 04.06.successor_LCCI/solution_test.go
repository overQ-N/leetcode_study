package successor_LCCI

import (
	"fmt"
	"testing"
)

func TestInorderSuccessor(t *testing.T) {
	test1 := inorderSuccessor(&TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}, &TreeNode{Val: 1})
	fmt.Println("-------", test1)
}
