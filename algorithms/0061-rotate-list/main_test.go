package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		head []int
		k    int
		want []int
	}{
		{
			head: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{4, 5, 1, 2, 3},
		},
		{
			head: []int{0, 1, 2},
			k:    4,
			want: []int{2, 0, 1},
		},
		{
			head: []int{},
			k:    0,
			want: []int(nil),
		},
		{
			head: []int{1},
			k:    0,
			want: []int{1},
		},
		{
			head: []int{1, 2},
			k:    1,
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		head := util.NewList(tt.head)
		list := rotateRight(head, tt.k)
		assert.Equal(t, tt.want, util.ListToNums(list))
	}
}
