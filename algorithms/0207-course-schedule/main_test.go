package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, canFinish(tt.numCourses, tt.prerequisites))
	}
}
