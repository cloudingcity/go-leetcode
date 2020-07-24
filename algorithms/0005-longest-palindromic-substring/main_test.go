package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "babad",
			want: "bab",
		},
		{
			s:    "cbbd",
			want: "bb",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, longestPalindrome(tt.s))
	}
}
