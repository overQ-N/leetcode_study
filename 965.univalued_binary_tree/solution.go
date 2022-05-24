package univalued_binary_tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root}
	value := root.Val
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val != value {
			return false
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return true
}
