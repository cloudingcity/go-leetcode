package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want []int
	}{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: []int{2, 2},
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: []int{0, 1, 3, 0, 4},
		},
	}
	for _, tt := range tests {
		l := removeElement(tt.nums, tt.val)
		assert.Equal(t, tt.want, tt.nums[:l])
	}
}
