package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		want   int
	}{
		{
			amount: 5,
			coins:  []int{1, 2, 5},
			want:   4,
		},
		{
			amount: 3,
			coins:  []int{2},
			want:   0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, change(tt.amount, tt.coins))
	}
}
