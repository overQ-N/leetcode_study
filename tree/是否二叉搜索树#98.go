/**
* leetCode #98 验证二叉搜索树
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
package tree

func insertNode(root, node *TreeNode) {
	if root.Val == 0 {
		root.Val = node.Val
		return
	}
	if node.Val < root.Val {
		if root.Left == nil {
			root.Left = node
		} else {
			insertNode(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			insertNode(root.Right, node)
		}
	}
}
func init() {
	// arr := []int{5, 4, 6, 1, 0, 3, 7}
	// var root TreeNode
	// for _, v := range arr {
	// 	node := TreeNode{Val: v}
	// 	insertNode(&root, &node)
	// }
	// fmt.Println(isValidBST(&root))
}

/*
* 解题思路：中序遍历为升序
* tips: 1.去重
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var isValidBST bool = true
	var inorderNodes []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node.Left != nil {
			traverse(node.Left)
		}
		inorderNodes = append(inorderNodes, node.Val)
		if node.Right != nil {
			traverse(node.Right)
		}
	}
	if root != nil {
		traverse(root)
	}
	// 校验是否升序，去重
	l := len(inorderNodes)
	for i := 0; i < l; i++ {
		if i != l-1 {
			if inorderNodes[i] >= inorderNodes[i+1] {
				isValidBST = false
				break
			}
		}
	}
	return isValidBST
}
