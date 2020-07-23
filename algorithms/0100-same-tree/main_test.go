package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := &TreeNode{Val: 1}
	nodeA1 := &TreeNode{Val: 2}
	nodeA2 := &TreeNode{Val: 3}
	root.Left = nodeA1
	root.Right = nodeA2

	p := root
	q := root

	assert.True(t, isSameTree(p, q))
}
