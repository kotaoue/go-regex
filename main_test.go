package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dot(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "This is a pen",
			expected: true,
		},
		{
			input:    "THAT is a pen",
			expected: true,
		},
		{
			input:    "This is a book",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`.... is a pen`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})

	}
}
