package main

import (
	"strconv"
	"strings"
)

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	root *TreeNode
}

func Constructor() Codec {
	return Codec{root: nil}
}

func traverse(node *TreeNode, callback func(int)) {
	callback(node.Val)
	if node.Left != nil {
		traverse(node.Left, callback)
	}
	if node.Right != nil {
		traverse(node.Right, callback)
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := []int{}
	traverse(root, func(i int) {
		result = append(result, i)
	})
	strs := []string{}
	for i := 0; i < len(result); i++ {
		strs = append(strs, strconv.Itoa(result[i]))
	}
	return strings.Join(strs, ",")

}

func insertNode(node *TreeNode, val int) {
	if node.Val < val {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
			return
		}
		insertNode(node.Right, val)
	} else {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
			return
		}
		insertNode(node.Left, val)
	}
}

// func (c *Codec) insertRightNode(val int){
// 	if c.root.Right == nil{
// 		c.root.Right = &TreeNode{Val: val}
// 		return
// 	}
// 	c.insertRightNode(val)
// }

func (c *Codec) buildTree(val int) *TreeNode {
	if c.root == nil {
		c.root = &TreeNode{Val: val}
		return c.root
	}
	insertNode(c.root, val)
	return c.root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, ",")
	for i := 0; i < len(arr); i++ {
		if val, err := strconv.Atoi(arr[i]); err == nil {
			this.buildTree(val)
		}
	}
	return this.root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
