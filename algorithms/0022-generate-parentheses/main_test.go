package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	want := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	assert.Equal(t, want, generateParenthesis(3))
}
