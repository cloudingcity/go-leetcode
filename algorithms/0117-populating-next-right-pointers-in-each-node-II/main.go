package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	node := root
	for node != nil {
		dummyHead := &Node{}
		dummy := dummyHead
		for node != nil {
			if node.Left != nil {
				dummy.Next = node.Left
				dummy = dummy.Next
			}
			if node.Right != nil {
				dummy.Next = node.Right
				dummy = dummy.Next
			}
			node = node.Next
		}
		node = dummyHead.Next
	}
	return root
}
