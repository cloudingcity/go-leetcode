package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			height: []int{1, 2},
			want:   1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, maxArea(tt.height))
	}
}
