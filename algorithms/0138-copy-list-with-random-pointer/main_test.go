package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	head := newList([]int{7, 13, 11, 10, 1})

	want := []int{7, 13, 11, 10, 1}
	got := listToNums(copyRandomList(head))

	head.Val = 99999

	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	head := newList([]int{7, 13, 11, 10, 1})

	want := []int{7, 13, 11, 10, 1}
	got := listToNums(copyRandomList2(head))

	head.Val = 99999

	assert.Equal(t, want, got)
}

func newList(nums []int) *Node {
	dummy := &Node{}
	node := dummy
	for _, n := range nums {
		node.Next = &Node{Val: n}
		node = node.Next
	}
	return dummy.Next
}

func listToNums(head *Node) []int {
	var nums []int
	for node := head; node != nil; node = node.Next {
		nums = append(nums, node.Val)
	}
	return nums
}
