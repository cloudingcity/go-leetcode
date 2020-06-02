package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{0, 0, 0},
			m:     0,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{2, 5, 6},
		},
	}
	for _, tt := range tests {
		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		assert.Equal(t, tt.want, tt.nums1)
	}
}
