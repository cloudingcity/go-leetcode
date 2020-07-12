package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 2, want: 2},
		{n: 3, want: 3},
		{n: 4, want: 5},
		{n: 5, want: 8},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, climbStairs(tt.n))
		assert.Equal(t, tt.want, climbStairsIter(tt.n))
		assert.Equal(t, tt.want, climbStairsDP(tt.n))
	}
}
