package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tail1 := &ListNode{Val: 8}
	tail2 := &ListNode{Val: 4}
	tail3 := &ListNode{Val: 5}
	tail1.Next = tail2
	tail2.Next = tail3

	nodeA1 := &ListNode{Val: 4}
	nodeA2 := &ListNode{Val: 1}
	nodeA1.Next = nodeA2
	nodeA2.Next = tail1

	nodeB1 := &ListNode{Val: 5}
	nodeB2 := &ListNode{Val: 0}
	nodeB3 := &ListNode{Val: 1}
	nodeB1.Next = nodeB2
	nodeB2.Next = nodeB3
	nodeB3.Next = tail1

	want := tail1
	got := getIntersectionNode(nodeA1, nodeB1)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	nodeA1 := &ListNode{Val: 2}
	nodeA2 := &ListNode{Val: 6}
	nodeA3 := &ListNode{Val: 4}
	nodeA1.Next = nodeA2
	nodeA2.Next = nodeA3

	nodeB1 := &ListNode{Val: 1}
	nodeB2 := &ListNode{Val: 5}
	nodeB1.Next = nodeB2

	got := getIntersectionNode(nodeA1, nodeB1)
	assert.Nil(t, got)
}
