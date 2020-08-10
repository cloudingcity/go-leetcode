package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	want := [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}
	got := pathSum(tree(), 22)
	assert.Equal(t, want, got)
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
	nodeC3 := &TreeNode{Val: 5}
	nodeC4 := &TreeNode{Val: 1}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Left = nodeB1
	nodeA2.Left = nodeB2
	nodeA2.Right = nodeB3
	nodeB1.Left = nodeC1
	nodeB1.Right = nodeC2
	nodeB3.Left = nodeC3
	nodeB3.Right = nodeC4

	return root
}
