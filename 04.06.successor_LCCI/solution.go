package successor_LCCI

// 面试题04.06. 后继者
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverseAndSearch(node *TreeNode, p *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if p.Val >= node.Val {
		return traverseAndSearch(node.Right, p)
	}
	targetNode := traverseAndSearch(node.Left, p)
	if targetNode == nil {
		return node
	} else {
		return targetNode
	}
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return traverseAndSearch(root, p)
}
