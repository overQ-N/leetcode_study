package binary_tree_level_order_traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	nums := [][]int{{root.Val}}
	queue := [][]*TreeNode{}
	queue = append(queue, []*TreeNode{root})

	for len(queue) != 0 {
		levelNodes := queue[0]
		queue = queue[1:]
		prefix := []*TreeNode{}
		num := []int{}
		for i := 0; i < len(levelNodes); i++ {
			node := levelNodes[i]
			if node.Left != nil {
				prefix = append(prefix, node.Left)
				num = append(num, node.Left.Val)
			}
			if node.Right != nil {
				prefix = append(prefix, node.Right)
				num = append(num, node.Right.Val)
			}
		}
		if len(prefix) > 0 {
			queue = append(queue, prefix)
			nums = append(nums, num)
		}

	}
	return nums
}
