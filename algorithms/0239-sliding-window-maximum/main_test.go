package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, maxSlidingWindow(tt.nums, tt.k))
		assert.Equal(t, tt.want, maxSlidingWindow2(tt.nums, tt.k))
	}
}
