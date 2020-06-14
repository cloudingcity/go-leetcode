package main

import (
	"github.com/cloudingcity/go-leetcode/util"
)

type ListNode = util.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	var count = 1
	node := head
	for node.Next != nil {
		node = node.Next
		count++
	}
	node.Next = head
	k %= count

	node = head
	for i := 1; i < count-k; i++ {
		node = node.Next
	}

	dummy := &ListNode{}
	dummy.Next = node.Next
	node.Next = nil

	return dummy.Next
}
