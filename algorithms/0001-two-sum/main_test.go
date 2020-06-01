package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{2, 7, 11, 15},
			target: 22,
			want:   []int{1, 3},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, twoSum(tt.nums, tt.target))
	}
}
