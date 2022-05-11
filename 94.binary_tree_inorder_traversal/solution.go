package binary_tree_inorder_traversal

// 94.二叉树的中序遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			traverse(node.Left)
		}
		res = append(res, node.Val)
		if node.Right != nil {
			traverse(node.Right)
		}
	}
	traverse(root)
	return res
}
