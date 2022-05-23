package symmetric_tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
* 对称：
* 左树的左节点=右树的右节点
* 左树的右节点=左树的左节点
 */
func isSymmetric(root *TreeNode) bool {
	return traverseChildren(root.Left, root.Right)
}

func traverseChildren(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return traverseChildren(left.Left, right.Right) && traverseChildren(left.Right, right.Left)
}
