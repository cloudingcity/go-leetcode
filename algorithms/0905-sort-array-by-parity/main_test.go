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
			A:    []int{3, 1, 2, 4},
			want: []int{2, 4, 3, 1},
		},
		{
			A:    []int{0, 1, 2},
			want: []int{0, 2, 1},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, sortArrayByParity(tt.A))
	}
}
