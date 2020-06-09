package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tail := util.NewList([]int{8, 4, 5})

	headA := util.NewList([]int{4, 1})
	headA.Next.Next = tail

	headB := util.NewList([]int{5, 0, 1})
	headB.Next.Next.Next = tail

	want := 8
	got := getIntersectionNode(headA, headB).Val
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	headA := util.NewList([]int{2, 6, 4})
	headB := util.NewList([]int{1, 5})

	assert.Nil(t, getIntersectionNode(headA, headB))
}
