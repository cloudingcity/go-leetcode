package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "abcabcbb", want: 3},
		{s: "abbbba", want: 2},
		{s: "bbbbb", want: 1},
		{s: "pwwkew", want: 3},
		{s: "", want: 0},
		{s: "dvdf", want: 3},
		{s: "au", want: 2},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, lengthOfLongestSubstring(tt.s))
		assert.Equal(t, tt.want, lengthOfLongestSubstring2(tt.s))
	}
}
