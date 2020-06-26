package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := &TreeNode{Val: 3}
	nodeA1 := &TreeNode{Val: 5}
	nodeA2 := &TreeNode{Val: 1}
	nodeB1 := &TreeNode{Val: 6}
	nodeB2 := &TreeNode{Val: 2}
	nodeB3 := &TreeNode{Val: 0}
	nodeB4 := &TreeNode{Val: 8}
	nodeC1 := &TreeNode{Val: 7}
	nodeC2 := &TreeNode{Val: 4}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Left = nodeB1
	nodeA1.Right = nodeB2
	nodeA2.Left = nodeB3
	nodeA2.Right = nodeB4
	nodeB2.Left = nodeC1
	nodeB2.Right = nodeC2
	p := nodeA1
	q := nodeA2

	node := lowestCommonAncestor(root, p, q)

	assert.Equal(t, 3, node.Val)
}
