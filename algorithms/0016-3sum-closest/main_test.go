package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			nums:   []int{0, 1, 2},
			target: 3,
			want:   3,
		},
		{
			nums:   []int{1, 1, 1, 1},
			target: 0,
			want:   3,
		},
		{
			nums:   []int{1, 1, 1, 0},
			target: -100,
			want:   2,
		},
		{
			nums:   []int{0, 2, 1, -3},
			target: 1,
			want:   0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, threeSumClosest(tt.nums, tt.target))
	}
}
