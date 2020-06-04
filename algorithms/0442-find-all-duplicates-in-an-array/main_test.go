package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{2, 3},
		},
		{
			nums: []int{10, 2, 5, 10, 9, 1, 1, 4, 3, 7},
			want: []int{10, 1},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, findDuplicates(tt.nums))
	}
}
