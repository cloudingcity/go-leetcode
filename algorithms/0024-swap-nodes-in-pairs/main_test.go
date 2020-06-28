package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		head []int
		want []int
	}{
		{
			head: []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
	}
	for _, tt := range tests {
		list := util.NewList(tt.head)
		got := swapPairs(list)
		assert.Equal(t, tt.want, util.ListToNums(got))
	}
}

func Test2(t *testing.T) {
	tests := []struct {
		head []int
		want []int
	}{
		{
			head: []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
	}
	for _, tt := range tests {
		list := util.NewList(tt.head)
		got := swapPairs2(list)
		assert.Equal(t, tt.want, util.ListToNums(got))
	}
}
