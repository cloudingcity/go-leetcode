package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	head := util.NewList([]int{1, 2, 3, 4, 5})
	list := removeNthFromEnd(head, 2)

	want := []int{1, 2, 3, 5}
	got := util.ListToNums(list)
	assert.Equal(t, want, got)
}

func TestNil(t *testing.T) {
	head := util.NewList([]int{1})
	list := removeNthFromEnd(head, 1)

	assert.Nil(t, list)
}
