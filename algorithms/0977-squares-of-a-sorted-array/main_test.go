package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		A    []int
		want []int
	}{
		{
			A:    []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			A:    []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, sortedSquares(tt.A))
	}
}
