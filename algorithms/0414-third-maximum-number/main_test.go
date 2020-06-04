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
		{
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			nums: []int{1, 2},
			want: 2,
		},
		{
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, thirdMax(tt.nums))
	}
}
