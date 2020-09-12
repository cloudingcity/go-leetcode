package main

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
	Prev *ListNode
}

type LRUCache struct {
	sets map[int]*ListNode
	cap  int
	head *ListNode
	tail *ListNode
}

func Constructor(capacity int) LRUCache {
	dummyHead := &ListNode{}
	dummyTail := &ListNode{}
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead

	return LRUCache{
		sets: make(map[int]*ListNode, capacity),
		cap:  capacity,
		head: dummyHead,
		tail: dummyTail,
	}
}

func (c *LRUCache) Get(key int) int {
	if n, ok := c.sets[key]; ok {
		c.removeCurrent(n)
		c.insertToHead(n)
		return n.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if n, ok := c.sets[key]; ok {
		n.Val = value
		c.removeCurrent(n)
		c.insertToHead(n)
		return
	}
	if len(c.sets) == c.cap {
		delete(c.sets, c.tail.Prev.Key)
		c.removeTail()
	}
	n := &ListNode{Key: key, Val: value}
	c.sets[key] = n
	c.insertToHead(n)
}

func (c *LRUCache) removeCurrent(n *ListNode) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (c *LRUCache) insertToHead(n *ListNode) {
	n.Next = c.head.Next
	n.Prev = c.head
	n.Next.Prev = n
	n.Prev.Next = n
}

func (c *LRUCache) removeTail() {
	c.tail.Prev = c.tail.Prev.Prev
	c.tail.Prev.Next = c.tail
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
