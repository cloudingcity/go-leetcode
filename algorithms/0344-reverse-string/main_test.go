package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{
			s:    []byte("hello"),
			want: []byte("olleh"),
		},
		{
			s:    []byte("ab"),
			want: []byte("ba"),
		},
	}
	for _, tt := range tests {
		reverseString(tt.s)
		assert.Equal(t, tt.want, tt.s)
	}
}
