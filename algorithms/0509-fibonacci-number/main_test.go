package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{N: 0, want: 0},
		{N: 1, want: 1},
		{N: 2, want: 1},
		{N: 3, want: 2},
		{N: 4, want: 3},
		{N: 5, want: 5},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, fib(tt.N))
		assert.Equal(t, tt.want, fibDP(tt.N))
		assert.Equal(t, tt.want, fibIter(tt.N))
	}
}
