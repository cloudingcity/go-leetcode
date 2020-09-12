package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	assert.Equal(t, 1, c.Get(1))
	c.Put(3, 3)
	assert.Equal(t, -1, c.Get(2))
	c.Put(4, 4)
	assert.Equal(t, -1, c.Get(1))
	assert.Equal(t, 3, c.Get(3))
	assert.Equal(t, 4, c.Get(4))
}

func debug(t *testing.T, c LRUCache) {
	var vals []int
	for _, node := range c.sets {
		vals = append(vals, node.Val)
	}
	t.Log(vals)

	var ss []string
	for node := c.head; node != nil; node = node.Next {
		ss = append(ss, strconv.Itoa(node.Val))
	}
	t.Log(strings.Join(ss, " -> "))
	t.Fatal("end")
}
