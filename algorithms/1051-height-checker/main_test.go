package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{
			heights: []int{1, 1, 4, 2, 1, 3},
			want:    3,
		},
		{
			heights: []int{1, 2, 3, 4, 5},
			want:    0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, heightChecker(tt.heights))
	}
}
