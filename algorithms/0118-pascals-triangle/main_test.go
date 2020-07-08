package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{
			numRows: 1,
			want:    [][]int{{1}},
		},
		{
			numRows: 2,
			want:    [][]int{{1}, {1, 1}},
		},
		{
			numRows: 3,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			numRows: 5,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, generate(tt.numRows))
	}
}
