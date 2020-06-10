package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{
			l1:   []int{1, 2, 4},
			l2:   []int{1, 3, 4},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1:   []int{1},
			l2:   []int{2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		l1 := util.NewList(tt.l1)
		l2 := util.NewList(tt.l2)
		l := mergeTwoLists(l1, l2)
		assert.Equal(t, tt.want, util.ListToNums(l))
	}
}
