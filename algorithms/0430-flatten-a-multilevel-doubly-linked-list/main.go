package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	var stack []*Node

	dummy := &Node{}
	dummy.Next = root
	node := root

	for node != nil {
		if node.Child != nil {
			if node.Next != nil {
				stack = append(stack, node.Next)
			}
			node.Next = node.Child
			node.Next.Prev = node
			node.Child = nil
		} else if node.Next == nil && len(stack) != 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.Next = n
			node.Next.Prev = node
		}
		node = node.Next
	}

	return dummy.Next
}
