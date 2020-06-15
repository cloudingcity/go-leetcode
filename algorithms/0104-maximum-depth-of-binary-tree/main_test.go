package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	nodeA1 := &TreeNode{Val: 3}
	nodeB1 := &TreeNode{Val: 9}
	nodeB2 := &TreeNode{Val: 20}
	nodeC1 := &TreeNode{Val: 15}
	nodeC2 := &TreeNode{Val: 7}
	nodeA1.Left = nodeB1
	nodeA1.Right = nodeB2
	nodeB2.Left = nodeC1
	nodeB2.Right = nodeC2

	want := 3
	got := maxDepth(nodeA1)
	assert.Equal(t, want, got)
}
