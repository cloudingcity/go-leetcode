package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		A    []int
		want bool
	}{
		{
			A:    []int{2, 1},
			want: false,
		},
		{
			A:    []int{3, 5, 5},
			want: false,
		},
		{
			A:    []int{0, 3, 2, 1},
			want: true,
		},
		{
			A:    []int{0, 3, 2, 1, 2},
			want: false,
		},
		{
			A:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.A), func(t *testing.T) {
			assert.Equal(t, tt.want, validMountainArray(tt.A))
		})
	}
}
