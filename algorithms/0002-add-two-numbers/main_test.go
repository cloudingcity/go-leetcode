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
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			l1:   []int{5},
			l2:   []int{5},
			want: []int{0, 1},
		},
		{
			l1:   []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			l2:   []int{5, 6, 4},
			want: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		l1 := util.NewList(tt.l1)
		l2 := util.NewList(tt.l2)
		l := addTwoNumbers(l1, l2)
		assert.Equal(t, tt.want, util.ListToNums(l))
	}
}
