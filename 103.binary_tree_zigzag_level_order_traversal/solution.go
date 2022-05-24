package binary_tree_zigzag_level_order_traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := [][]*TreeNode{{root}}
	res := [][]int{}
	flag := true

	for len(queue) > 0 {
		nodes := queue[0]
		queue = queue[1:]
		nums := []int{}
		children := []*TreeNode{}
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if node != nil {
				nums = append(nums, node.Val)
				if node.Left != nil {
					children = append(children, node.Left)
				}
				if node.Right != nil {
					children = append(children, node.Right)
				}
			}
		}

		if !flag {
			for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		flag = !flag
		res = append(res, nums)
		if len(children) > 0 {
			queue = append(queue, children)
		}
	}
	return res
}
