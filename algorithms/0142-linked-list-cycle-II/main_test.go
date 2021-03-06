package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	list := util.NewCycleList([]int{2, 0, -4})
	head := &ListNode{Val: 3}
	head.Next = list

	want := 2
	got := detectCycle(head).Val
	assert.Equal(t, want, got)
}
