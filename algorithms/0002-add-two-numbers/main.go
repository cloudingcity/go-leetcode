package main

import (
	"github.com/cloudingcity/go-leetcode/util"
)

type ListNode = util.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	var carry, sum int
	for l1 != nil || l2 != nil {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		node.Next = &ListNode{Val: sum % 10}
		node = node.Next
	}

	if carry == 1 {
		node.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
