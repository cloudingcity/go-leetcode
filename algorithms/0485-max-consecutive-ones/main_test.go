package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 1, 0, 1, 1, 1}, want: 3},
		{nums: []int{1, 0, 1, 1, 0, 1}, want: 2},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, findMaxConsecutiveOnes(tt.nums))
	}
}
