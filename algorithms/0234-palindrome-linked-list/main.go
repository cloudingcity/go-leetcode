package main

import "github.com/cloudingcity/go-leetcode/util"

type ListNode = util.ListNode

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	slow = reverse(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	for node != nil {
		node.Next, prev, node = prev, node, node.Next
	}
	return prev
}
