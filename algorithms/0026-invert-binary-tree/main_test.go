package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tree := invertTree(tree())
	assert.Equal(t, 4, tree.Val)
	assert.Equal(t, 7, tree.Left.Val)
	assert.Equal(t, 2, tree.Right.Val)
	assert.Equal(t, 9, tree.Left.Left.Val)
	assert.Equal(t, 6, tree.Left.Right.Val)
	assert.Equal(t, 3, tree.Right.Left.Val)
	assert.Equal(t, 1, tree.Right.Right.Val)
}

func TestIter(t *testing.T) {
	tree := invertTreeIter(tree())
	assert.Equal(t, 4, tree.Val)
	assert.Equal(t, 7, tree.Left.Val)
	assert.Equal(t, 2, tree.Right.Val)
	assert.Equal(t, 9, tree.Left.Left.Val)
	assert.Equal(t, 6, tree.Left.Right.Val)
	assert.Equal(t, 3, tree.Right.Left.Val)
	assert.Equal(t, 1, tree.Right.Right.Val)
}

func tree() *TreeNode {
	root := &TreeNode{Val: 4}
	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 7}
	nodeB1 := &TreeNode{Val: 1}
	nodeB2 := &TreeNode{Val: 3}
	nodeB3 := &TreeNode{Val: 6}
	nodeB4 := &TreeNode{Val: 9}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Left = nodeB1
	nodeA1.Right = nodeB2
	nodeA2.Left = nodeB3
	nodeA2.Right = nodeB4

	return root
}
