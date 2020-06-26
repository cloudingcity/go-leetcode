package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	title string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var b strings.Builder
	this.buildString(root, &b)
	return b.String()
}

func (this *Codec) buildString(node *TreeNode, b *strings.Builder) {
	if node == nil {
		b.WriteString("nil")
		b.WriteString(",")
		return
	}
	b.WriteString(strconv.Itoa(node.Val))
	b.WriteString(",")
	this.buildString(node.Left, b)
	this.buildString(node.Right, b)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return this.buildTree(&vals)
}

func (this *Codec) buildTree(vals *[]string) *TreeNode {
	val := (*vals)[0]
	*vals = (*vals)[1:]
	if val == "nil" {
		return nil
	}
	v, _ := strconv.Atoi(val)
	node := &TreeNode{Val: v}
	node.Left = this.buildTree(vals)
	node.Right = this.buildTree(vals)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
