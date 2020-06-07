package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	assert.Equal(t, 2, l.Get(1))
	l.DeleteAtIndex(1)
	assert.Equal(t, 3, l.Get(1))
}

func Test2(t *testing.T) {
	l := Constructor()
	l.AddAtHead(7)
	l.AddAtHead(2)
	l.AddAtHead(1)
	l.AddAtIndex(3, 0)
	l.DeleteAtIndex(2)
	l.AddAtHead(6)
	l.AddAtTail(4)
	assert.Equal(t, 4, l.Get(4))
	l.AddAtHead(4)
	l.AddAtIndex(5, 0)
	l.AddAtHead(6)
}

func TestGet(t *testing.T) {
	l := Constructor()
	assert.Equal(t, -1, l.Get(0))
}

func TestAddAtHead(t *testing.T) {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtHead(2)
	l.AddAtHead(3)
	assert.Equal(t, 3, l.Get(0))
	assert.Equal(t, 2, l.Get(1))
	assert.Equal(t, 1, l.Get(2))
	assert.Equal(t, -1, l.Get(3))
}

func TestAddAtTail(t *testing.T) {
	l := Constructor()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	assert.Equal(t, 1, l.Get(0))
	assert.Equal(t, 2, l.Get(1))
	assert.Equal(t, 3, l.Get(2))
	assert.Equal(t, -1, l.Get(3))
}

func TestAddAtIndex(t *testing.T) {
	l := Constructor()
	l.AddAtIndex(0, 1)
	l.AddAtIndex(1, 2)
	l.AddAtIndex(2, 3)
	l.AddAtIndex(1, 9)
	assert.Equal(t, 1, l.Get(0))
	assert.Equal(t, 9, l.Get(1))
	assert.Equal(t, 2, l.Get(2))
	assert.Equal(t, 3, l.Get(3))
	assert.Equal(t, -1, l.Get(4))
}

func TestDeleteAtIndex(t *testing.T) {
	l := Constructor()
	l.AddAtIndex(0, 1)
	l.AddAtIndex(0, 2)
	l.AddAtIndex(0, 3)
	l.DeleteAtIndex(3)
	l.DeleteAtIndex(1)
	l.DeleteAtIndex(0)
	assert.Equal(t, 1, l.Get(0))
	assert.Equal(t, -1, l.Get(1))
	assert.Equal(t, -1, l.Get(2))
	assert.Equal(t, -1, l.Get(3))
}

func printList(list MyLinkedList) {
	for node := list.head; node != nil; node = node.next {
		fmt.Print(node.val, " ")
	}
	fmt.Println()
}
