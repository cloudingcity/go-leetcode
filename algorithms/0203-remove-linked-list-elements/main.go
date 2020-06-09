package main

import "github.com/cloudingcity/go-leetcode/util"

type ListNode = util.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	node := dummy
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return dummy.Next
}
