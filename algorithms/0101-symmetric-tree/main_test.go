package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.True(t, isSymmetric(tree1()))
	assert.False(t, isSymmetric(tree2()))
}

func TestRecursive(t *testing.T) {
	assert.True(t, isSymmetricRecursive(tree1()))
	assert.False(t, isSymmetricRecursive(tree2()))
}

func tree1() *TreeNode {
	root := &TreeNode{Val: 1}
	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 2}
	nodeB1 := &TreeNode{Val: 3}
	nodeB2 := &TreeNode{Val: 4}
	nodeB3 := &TreeNode{Val: 4}
	nodeB4 := &TreeNode{Val: 3}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Left = nodeB1
	nodeA1.Right = nodeB2
	nodeA2.Left = nodeB3
	nodeA2.Right = nodeB4
	return root
}

func tree2() *TreeNode {
	root := &TreeNode{Val: 1}
	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 2}
	nodeB1 := &TreeNode{Val: 3}
	nodeB2 := &TreeNode{Val: 3}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Right = nodeB1
	nodeA1.Right = nodeB2
	return root
}
