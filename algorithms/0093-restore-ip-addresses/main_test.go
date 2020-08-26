package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s:    "25525511135",
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			s:    "0000",
			want: []string{"0.0.0.0"},
		},
		{
			s:    "1111",
			want: []string{"1.1.1.1"},
		},
		{
			s:    "010010",
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			s:    "101023",
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			s:    strings.Repeat("1", 3000),
			want: nil,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, restoreIpAddresses(tt.s))
	}
}
