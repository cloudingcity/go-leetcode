package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{
			coins:  []int{1, 2, 5},
			amount: 0,
			want:   0,
		},
		{
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, coinChange(tt.coins, tt.amount))
		assert.Equal(t, tt.want, coinChange2(tt.coins, tt.amount))
	}
}
