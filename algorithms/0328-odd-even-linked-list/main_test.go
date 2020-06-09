package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 3, 5, 2, 4},
		},
		{
			nums: []int{1, 2},
			want: []int{1, 2},
		},
		{
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		head := oddEvenList(util.NewList(tt.nums))
		got := util.ListToNums(head)
		assert.Equal(t, tt.want, got)
	}
}
