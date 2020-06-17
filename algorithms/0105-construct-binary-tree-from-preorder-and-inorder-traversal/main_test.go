package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)

	assert.Equal(t, 3, tree.Val)
	assert.Equal(t, 9, tree.Left.Val)
	assert.Equal(t, 20, tree.Right.Val)
	assert.Equal(t, 15, tree.Right.Left.Val)
	assert.Equal(t, 7, tree.Right.Right.Val)
}
