package main

import "github.com/cloudingcity/go-leetcode/util"

type ListNode = util.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

func reverseListWithRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListWithRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
