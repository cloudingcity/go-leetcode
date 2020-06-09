package main

import "github.com/cloudingcity/go-leetcode/util"

type ListNode = util.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	var length int
	first := head
	for first != nil {
		length++
		first = first.Next
	}
	length -= n

	first = dummy
	for length > 0 {
		length--
		first = first.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}
