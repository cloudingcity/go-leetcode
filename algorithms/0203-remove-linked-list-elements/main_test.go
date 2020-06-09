package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	list := util.NewList([]int{1, 2, 6, 3, 4, 5, 6})
	want := []int{1, 2, 3, 4, 5}
	got := util.ListToNums(removeElements(list, 6))

	assert.Equal(t, want, got)
}
