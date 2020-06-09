package main

import (
	"testing"

	"github.com/cloudingcity/go-leetcode/util"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 2, 1},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 2, 1},
			want: true,
		},
		{
			nums: []int{1},
			want: true,
		},
		{
			nums: []int{},
			want: true,
		},
		{
			nums: []int{1, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isPalindrome(util.NewList(tt.nums))
		assert.Equal(t, tt.want, got)
	}
}
