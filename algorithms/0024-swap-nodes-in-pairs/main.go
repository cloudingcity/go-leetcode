package main

import (
	"github.com/cloudingcity/go-leetcode/util"
)

type ListNode = util.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	node := head.Next
	head.Next = swapPairs(head.Next.Next)
	node.Next = head
	return node
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	node := dummy
	for node.Next != nil && node.Next.Next != nil {
		first := node.Next
		second := node.Next.Next
		first.Next = second.Next
		second.Next = first
		node.Next = second
		node = node.Next.Next
	}
	return dummy.Next
}
