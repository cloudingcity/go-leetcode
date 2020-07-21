package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, maxProfit(tt.prices))
	}
}
