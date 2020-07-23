package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 5},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, canPartition(tt.nums))
	}
}
