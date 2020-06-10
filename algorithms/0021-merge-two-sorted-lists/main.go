package main

import (
	"github.com/cloudingcity/go-leetcode/util"
)

type ListNode = util.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}
	return dummy.Next
}
