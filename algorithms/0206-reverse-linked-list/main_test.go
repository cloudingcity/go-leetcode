package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	list := util.NewList([]int{1, 2, 3, 4, 5})
	want := []int{5, 4, 3, 2, 1}
	got := util.ListToNums(reverseList(list))

	assert.Equal(t, want, got)
}

func TestRecursive(t *testing.T) {
	list := util.NewList([]int{1, 2, 3, 4, 5})
	want := []int{5, 4, 3, 2, 1}
	got := util.ListToNums(reverseListWithRecursive(list))

	assert.Equal(t, want, got)
}
