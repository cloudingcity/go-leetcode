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
			want: []int{5, 6},
		},
		{
			nums: []int{1, 2, 2},
			want: []int{3},
		},
		{
			nums: []int{2, 2, 3, 3, 5},
			want: []int{1, 4},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, findDisappearedNumbers(tt.nums))
	}
}
