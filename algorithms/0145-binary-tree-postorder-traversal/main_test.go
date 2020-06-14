package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node1.Right = node2
	node2.Left = node3

	want := []int{3, 2, 1}
	got := postorderTraversal(node1)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node1.Right = node2
	node2.Left = node3

	want := []int{3, 2, 1}
	got := postorderTraversalRecursive(node1)
	assert.Equal(t, want, got)
}
