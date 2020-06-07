package main

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	var count int

	for node := this.head; node != nil; node = node.next {
		if count == index {
			return node.val
		}
		count++
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Node{val, this.head}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.head = &Node{
			val: val,
		}
		return
	}
	node := this.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{val: val}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	var count int
	prev := this.head

	for node := this.head; node != nil; node = node.next {
		if count == index {
			prev.next = &Node{
				val:  val,
				next: prev.next,
			}
			return
		}
		prev = node
		count++
	}

	if count == index {
		this.AddAtTail(val)
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.head.next != nil {
			this.head.val = this.head.next.val
			this.head.next = this.head.next.next
		}
		return
	}
	var count int
	prev := this.head

	for node := this.head; node != nil; node = node.next {
		if count == index {
			prev.next = node.next
			break
		}
		prev = node
		count++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
