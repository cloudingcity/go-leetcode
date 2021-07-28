package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{nums: []int{1, 2, 3, 4, 5}, want: true},
		{nums: []int{5, 4, 3, 2, 1}, want: false},
		{nums: []int{2, 1, 5, 0, 4, 6}, want: true},
		{nums: []int{1, 1, 1}, want: false},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, increasingTriplet(tt.nums))
	}
}
