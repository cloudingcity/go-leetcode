package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		cost []int
		want int
	}{
		{
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, minCostClimbingStairs(tt.cost))
	}
}
