package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := &TreeNode{Val: 3}
	nodeA1 := &TreeNode{Val: 9}
	nodeA2 := &TreeNode{Val: 20}
	nodeB1 := &TreeNode{Val: 15}
	nodeB2 := &TreeNode{Val: 7}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA2.Left = nodeB1
	nodeA2.Right = nodeB2

	assert.Equal(t, 2, minDepth(root))
}
