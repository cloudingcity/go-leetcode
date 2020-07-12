package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{
			x:    2,
			n:    10,
			want: 1024,
		},
		{
			x:    2.1,
			n:    3,
			want: 9.261000000000001,
		},
		{
			x:    2,
			n:    -2,
			want: 0.25,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, myPow(tt.x, tt.n))
	}
}
