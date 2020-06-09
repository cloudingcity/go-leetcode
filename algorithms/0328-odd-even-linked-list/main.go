package main

import "github.com/cloudingcity/go-leetcode/util"

type ListNode = util.ListNode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// 1->2->3->4->5->NULL
// 1->3->5->2->4->NULL
