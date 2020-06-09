package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	list := util.NewCycleList([]int{3, 2, 0, -4})

	assert.True(t, hasCycle(list))
}

func TestNoCycle(t *testing.T) {
	list := util.NewList([]int{1, 2, 3, 4})

	assert.False(t, hasCycle(list))
}
