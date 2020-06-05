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
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			height: []int{1, 0, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, trap(tt.height))
	}
}
