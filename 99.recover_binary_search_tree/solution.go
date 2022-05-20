package recover_binary_search_tree

// 99.恢复一颗搜索二叉树

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	nums := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findTwoSwapNum(nums)
	recover(root, 2, x, y)
}

func findTwoSwapNum(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	x, y := nums[index1], nums[index2]
	return x, y
}

func recover(node *TreeNode, count, x, y int) {
	if node == nil {
		return
	}
	if node.Val == x || node.Val == y {
		if node.Val == x {
			node.Val = y
		} else {
			node.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(node.Left, count, x, y)
	recover(node.Right, count, x, y)
}
