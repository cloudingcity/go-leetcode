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
			nums: []int{1,1,2},
			want: []int{1,2},
		},
		{
			nums: []int{0,0,1,1,1,2,2,3,3,4},
			want: []int{0,1,2,3,4},
		},
	}
	for _, tt := range tests {
		l := removeDuplicates(tt.nums)
		assert.Equal(t, tt.want, tt.nums[:l])
	}
}
