package tree

import "fmt"

/**
* leet code #95
*
/
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

func insert(node *TreeNode, newNode *TreeNode) {
	if newNode.Val < node.Val {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insert(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insert(node.Right, newNode)
		}
	}
}
func GenerateTrees(n int) (result []*TreeNode) {
	if n == 0 {
		return nil
	}
	return helper(1, n, "root")

}
func helper(start int, end int, flag string) []*TreeNode {
	fmt.Println("flag:", flag)
	if start > end {
		return []*TreeNode{nil}
	}
	var allTrees []*TreeNode
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1, "left")
		fmt.Println("leftTrees len", len(leftTrees))
		rightTrees := helper(i+1, end, "right")
		fmt.Println("rightTrees len", len(rightTrees))
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				rootNode := TreeNode{Val: i}
				rootNode.Left = leftTree
				rootNode.Right = rightTree
				allTrees = append(allTrees, &rootNode)
			}
		}
	}
	return allTrees
}

func init() {
	// dp(4)
	//  GenerateTrees(3)

}

type Node struct {
	left  *Node
	right *Node
	val   int
}

func dp(n int) int {
	G := make([]int, n+1)
	G[0] = 1
	G[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]

}

func numTrees(n int) int {
	res := 0
	if n == 0 {
		return res
	}
	res = len(buildTree(1, n))
	return res
}
func buildTree(start, end int) []*Node {
	if start > end {
		return []*Node{nil}
	}
	trees := []*Node{}

	for i := start; i <= end; i++ {
		leftTrees := buildTree(start, i-1)
		rightTrees := buildTree(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currentTree := Node{val: i}
				currentTree.left = left
				currentTree.right = right
				trees = append(trees, &currentTree)
			}
		}
	}
	return trees
}
