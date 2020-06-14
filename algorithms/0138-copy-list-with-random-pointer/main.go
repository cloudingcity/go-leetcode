package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// O(N) space
func copyRandomList(head *Node) *Node {
	dummy := &Node{}
	m := make(map[*Node]*Node)

	orig := head
	node := dummy
	for orig != nil {
		newNode := &Node{Val: orig.Val}
		m[orig] = newNode
		node.Next = newNode
		node = node.Next
		orig = orig.Next
	}

	orig = head
	node = dummy.Next
	for orig != nil {
		node.Random = m[orig.Random]
		node = node.Next
		orig = orig.Next
	}

	return dummy.Next
}

// O(1) space
func copyRandomList2(head *Node) *Node {
	orig := head
	for orig != nil {
		orig.Next = &Node{
			Val:    orig.Val,
			Next:   orig.Next,
			Random: orig.Random,
		}
		orig = orig.Next.Next
	}

	orig = head
	for orig != nil {
		orig = orig.Next
		if orig.Random != nil {
			orig.Random = orig.Random.Next
		}
		orig = orig.Next
	}

	dummy := &Node{}
	orig = head
	node := dummy
	for orig != nil {
		node.Next = orig.Next
		node = node.Next

		orig.Next = orig.Next.Next
		orig = orig.Next
	}
	return dummy.Next
}
