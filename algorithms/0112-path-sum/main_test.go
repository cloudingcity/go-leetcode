package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.True(t, hasPathSum(tree(), 22))
}

func tree() *TreeNode {
	root := &TreeNode{Val: 5}
	nodeA1 := &TreeNode{Val: 4}
	nodeA2 := &TreeNode{Val: 8}
	nodeB1 := &TreeNode{Val: 11}
	nodeB2 := &TreeNode{Val: 13}
	nodeB3 := &TreeNode{Val: 4}
	nodeC1 := &TreeNode{Val: 7}
	nodeC2 := &TreeNode{Val: 2}
	nodeC3 := &TreeNode{Val: 1}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Left = nodeB1
	nodeA2.Left = nodeB2
	nodeA2.Right = nodeB3
	nodeB1.Left = nodeC1
	nodeB1.Right = nodeC2
	nodeB3.Right = nodeC3
	return root
}
