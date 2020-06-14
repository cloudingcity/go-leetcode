package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	head1 := newList([]int{1, 2, 3, 4, 5, 6})
	head2 := newList([]int{7, 8, 9, 10})
	head3 := newList([]int{11, 12})
	head1.Next.Next.Child = head2
	head2.Next.Child = head3

	head := flatten(head1)

	want := []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6}
	got := listToNums(head)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	head1 := newList([]int{1, 2})
	head2 := newList([]int{3})
	head1.Child = head2

	head := flatten(head1)

	want := []int{1, 3, 2}
	got := listToNums(head)
	assert.Equal(t, want, got)
}

func Test3(t *testing.T) {
	head := flatten(nil)

	want := []int(nil)
	got := listToNums(head)
	assert.Equal(t, want, got)
}

func Test4(t *testing.T) {
	head1 := newList([]int{1})
	head2 := newList([]int{3})
	head1.Child = head2

	head := flatten(head1)

	want := []int{1, 3}
	got := listToNums(head)
	assert.Equal(t, want, got)
}

func newList(nums []int) *Node {
	dummy := &Node{}
	node := dummy
	for _, n := range nums {
		node.Next = &Node{Val: n}
		node.Next.Prev = node
		node = node.Next
	}
	dummy.Next.Prev = nil
	return dummy.Next
}

func listToNums(head *Node) []int {
	var nums []int
	for node := head; node != nil; node = node.Next {
		nums = append(nums, node.Val)
	}
	return nums
}
