package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		rowIndex int
		want     []int
	}{
		{
			rowIndex: 0,
			want:     []int{1},
		},
		{
			rowIndex: 1,
			want:     []int{1, 1},
		},
		{
			rowIndex: 2,
			want:     []int{1, 2, 1},
		},
		{
			rowIndex: 4,
			want:     []int{1, 4, 6, 4, 1},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, getRow(tt.rowIndex))
	}
}
