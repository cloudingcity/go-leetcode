package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := &TreeNode{Val: 4}
	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 7}
	nodeB1 := &TreeNode{Val: 1}
	nodeB2 := &TreeNode{Val: 3}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Left = nodeB1
	nodeA1.Right = nodeB2

	node := searchBST(root, 2)
	assert.Equal(t, 2, node.Val)
	assert.Equal(t, 1, node.Left.Val)
	assert.Equal(t, 3, node.Right.Val)
}

func TestIter(t *testing.T) {
	root := &TreeNode{Val: 4}
	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 7}
	nodeB1 := &TreeNode{Val: 1}
	nodeB2 := &TreeNode{Val: 3}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Left = nodeB1
	nodeA1.Right = nodeB2

	node := searchBSTIter(root, 2)
	assert.Equal(t, 2, node.Val)
	assert.Equal(t, 1, node.Left.Val)
	assert.Equal(t, 3, node.Right.Val)
}
