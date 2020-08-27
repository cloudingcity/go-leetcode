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
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, canJump(tt.nums))
		assert.Equal(t, tt.want, canJumpBacktrack(tt.nums))
	}
}
