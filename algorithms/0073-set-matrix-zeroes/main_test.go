package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
	}
	for _, tt := range tests {
		setZeroes(tt.matrix)
		assert.Equal(t, tt.want, tt.matrix)
	}
}
