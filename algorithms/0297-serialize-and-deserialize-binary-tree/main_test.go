package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := &TreeNode{Val: 1}
	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 3}
	nodeB1 := &TreeNode{Val: 4}
	nodeB2 := &TreeNode{Val: 5}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA2.Left = nodeB1
	nodeA2.Right = nodeB2

	obj := Constructor()
	s := obj.serialize(root)
	node := obj.deserialize(s)

	assert.Equal(t, "1,2,nil,nil,3,4,nil,nil,5,nil,nil,", s)

	assert.Equal(t, node.Val, root.Val)
	assert.Equal(t, node.Left.Val, root.Left.Val)
	assert.Equal(t, node.Right.Val, root.Right.Val)
}
