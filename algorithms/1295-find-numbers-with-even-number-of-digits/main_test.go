package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{12, 345, 2, 6, 7896}, want: 2},
		{nums: []int{555, 901, 482, 1771}, want: 1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, findNumbers(tt.nums))
	}
}
