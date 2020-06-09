package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4})
	want := []int{1, 2, 3, 4}
	got := ListToNums(list)
	assert.Equal(t, want, got)
}

func TestCycleList(t *testing.T) {
	list := NewCycleList([]int{1, 2, 3})
	want := 1
	got := list.Next.Next.Next.Val

	assert.Equal(t, want, got)
}
