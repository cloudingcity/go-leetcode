package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          []int{0, 1},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          []int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, findOrder(tt.numCourses, tt.prerequisites))
	}
}
