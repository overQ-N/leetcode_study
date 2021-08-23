package main

import (
	"fmt"

	_ "example.com/day0717/dp"
	_ "example.com/day0717/other"
	_ "example.com/day0717/tree"
)

type Node struct {
	left  *Node
	right *Node
	val   int
}

func (node *Node) String() int {
	return node.val
}
func (node *Node) insertNode(newNode *Node, rootNode *Node) {
	if newNode.val < rootNode.val {
		if rootNode.left == nil {
			rootNode.left = newNode
		} else {
			node.insertNode(newNode, rootNode.left)
		}
	} else {
		if rootNode.right == nil {
			rootNode.right = newNode
		} else {
			node.insertNode(newNode, rootNode.right)
		}
	}

}
func (node *Node) insert(insertNode *Node) {
	// if node is rootNode
	if node.val == 0 {
		node.val = insertNode.val
		return
	}
	// if node.left == nil {
	// 	node.left = newNode
	// } else if node.right == nil {
	// 	node.right = newNode
	// } else {
	// 	node.left.insert(newNode)
	// }
	node.insertNode(insertNode, node)
}

// 树的遍历
func (node *Node) traverse() {

	if node.left != nil {
		node.left.traverse()
	}

	if node.right != nil {
		node.right.traverse()
	}
	fmt.Println(node.val)
}

// 树的广度遍历
func (node *Node) levelTraverse() {
	var tempNodeSlice = []*Node{node}
	for len(tempNodeSlice) > 0 {
		var len = len(tempNodeSlice)
		var tempNode = tempNodeSlice[0]
		tempNodeSlice = tempNodeSlice[1:len]
		fmt.Println(tempNode.val)
		if tempNode.left != nil {
			tempNodeSlice = append(tempNodeSlice, tempNode.left)
		}
		if tempNode.right != nil {
			tempNodeSlice = append(tempNodeSlice, tempNode.right)
		}

	}
}
func (node *Node) getNode(val int) *Node {
	if node.val == val {
		return node
	}
	if val < node.val {
		return node.left.getNode(val)
	}
	if val >= node.val {
		return node.right.getNode(val)
	}
	return nil
}
func (node *Node) searchNode(val int) (resultNode *Node) {

	n := node
	for n != nil {
		if val < n.val {
			n = n.left
		} else if val > n.val {
			n = n.right
		} else {

			resultNode = n
			n = nil
		}
	}
	return
}
func (node *Node) max() *Node {
	n := node.right
	for n.right != nil {
		n = n.right
	}
	return n
}
func (node *Node) min() *Node {
	n := node.left
	for n.left != nil {
		n = n.left
	}
	return n
}
func main() {
	// nodes := tree.GenerateTrees(3)
	// fmt.Println("nodes:", nodes)
	return
	var binaryTree = new(Node)
	var testTree = new(Node)
	testTree.val = 999999
	// []int{1, 3, 10, 6, 5, 2}
	arr := []int{7, 10, 6, 7, 4, 5, 9, 11, 12}
	for _, v := range arr {
		tNode := Node{
			val: v,
		}
		binaryTree.insert(&tNode)
		// fmt.Printf("%p  ========== \n ", binaryTree)
	}
	// fmt.Println(binaryTree, binaryTree.left, binaryTree.right, binaryTree.left.left)
	// binaryTree.traverse()
	fmt.Println(binaryTree.String(), "+=======")
	binaryTree.max()
	// n := binaryTree.searchNode(13)
	fmt.Printf("getNode:%#v,%#v\n", binaryTree.max(), binaryTree.min())
	binaryTree.levelTraverse()
}
