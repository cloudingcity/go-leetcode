package main

import (
	"github.com/cloudingcity/go-leetcode/util"
)

type ListNode = util.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	node := head
	length := 1
	for node.Next != nil {
		length++
		node = node.Next
	}

	k %= length
	if k == 0 {
		return head
	}

	node.Next = head
	for i := length - k; i > 0; i-- {
		node = node.Next
	}

	res := node.Next
	node.Next = nil

	return res
}
