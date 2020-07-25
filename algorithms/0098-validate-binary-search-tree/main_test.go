package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := &TreeNode{Val: 2}
	nodeA1 := &TreeNode{Val: 1}
	nodeA2 := &TreeNode{Val: 3}
	root.Left = nodeA1
	root.Right = nodeA2

	assert.True(t, isValidBST(root))
}

func Test2(t *testing.T) {
	root := &TreeNode{Val: 5}
	nodeA1 := &TreeNode{Val: 1}
	nodeA2 := &TreeNode{Val: 4}
	nodeB1 := &TreeNode{Val: 3}
	nodeB2 := &TreeNode{Val: 6}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA2.Left = nodeB1
	nodeA2.Right = nodeB2

	assert.False(t, isValidBST(root))
}
