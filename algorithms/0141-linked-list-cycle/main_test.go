package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	node0 := &ListNode{Val: 3}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: -4}
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1

	assert.True(t, hasCycle(node0))
}
