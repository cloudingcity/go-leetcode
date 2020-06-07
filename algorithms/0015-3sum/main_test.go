package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, threeSum(tt.nums))
	}
}
