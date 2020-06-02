package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			arr:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		duplicateZeros(tt.arr)
		assert.Equal(t, tt.want, tt.arr)
	}
}
