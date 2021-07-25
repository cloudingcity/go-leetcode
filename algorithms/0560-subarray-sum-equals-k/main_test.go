package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 1, 1}, k: 2, want: 2},
		{nums: []int{1, 2, 3}, k: 3, want: 2},
		{nums: []int{1, -1, 0}, k: 0, want: 3},
		{nums: []int{3, 4, 7, 2, -3, 1, 4, 2}, k: 7, want: 4},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, subarraySum(tt.nums, tt.k))
	}
}
