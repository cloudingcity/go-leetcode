package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	root := &Node{Val: 1}
	nodeA1 := &Node{Val: 2}
	nodeA2 := &Node{Val: 3}
	nodeB1 := &Node{Val: 4}
	nodeB2 := &Node{Val: 5}
	nodeB3 := &Node{Val: 6}
	nodeB4 := &Node{Val: 7}
	root.Left = nodeA1
	root.Right = nodeA2
	nodeA1.Left = nodeB1
	nodeA1.Right = nodeB2
	nodeA2.Left = nodeB3
	nodeA2.Right = nodeB4

	node := connect(root)
	assert.Nil(t, node.Next)
	assert.Equal(t, 3, node.Left.Next.Val)
	assert.Nil(t, node.Right.Next)
	assert.Equal(t, 5, node.Left.Left.Next.Val)
	assert.Equal(t, 6, node.Left.Left.Next.Next.Val)
	assert.Equal(t, 7, node.Left.Left.Next.Next.Next.Val)
	assert.Nil(t, node.Left.Left.Next.Next.Next.Next)
}
