package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) *ListNode {
	dummy := &ListNode{}
	node := dummy
	for _, n := range nums {
		node.Next = &ListNode{Val: n}
		node = node.Next
	}
	return dummy.Next
}

func ListToNums(head *ListNode) []int {
	var nums []int
	for node := head; node != nil; node = node.Next {
		nums = append(nums, node.Val)
	}
	return nums
}

func NewCycleList(nums []int) *ListNode {
	dummy := &ListNode{}
	node := dummy
	for _, n := range nums {
		node.Next = &ListNode{Val: n}
		node = node.Next
	}
	node.Next = dummy.Next
	return dummy.Next
}
