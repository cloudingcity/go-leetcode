package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		want []int
	}{
		{
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, foo())
	}
}
