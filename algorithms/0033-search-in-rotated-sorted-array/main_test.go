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
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{nums: []int{1}, target: 0, want: -1},
		{nums: []int{1, 3}, target: 3, want: 1},
		{nums: []int{3, 1}, target: 3, want: 0},
		{nums: []int{1, 3, 5}, target: 0, want: -1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, search(tt.nums, tt.target))
	}
}
