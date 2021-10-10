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

func Test_LowerW(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "This is a pen",
			expected: true,
		},
		{
			input:    "This is a !!!",
			expected: false,
		},
		{
			input:    "This is a B_2",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`This is a \w\w\w`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})

	}
}

func Test_UpperW(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "This is a pen",
			expected: false,
		},
		{
			input:    "This is a !!!",
			expected: true,
		},
		{
			input:    "This is a B_2",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`This is a \W\W\W`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})

	}
}

func Test_LowerS(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "ABCD EFGH",
			expected: true,
		},
		{
			input:    "ABCDEFGH",
			expected: false,
		},
		{
			input:    "ABCD EF GH",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`ABCD\sEFGH`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})

	}
}

func Test_UpperS(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "ABCD EFGH",
			expected: false,
		},
		{
			input:    "ABCDEFGH",
			expected: false,
		},
		{
			input:    "ABCD EF GH",
			expected: false,
		},
		{
			input:    "ABCDxEFGH",
			expected: true,
		},
		{
			input:    "ABCD-EFGH",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`ABCD\SEFGH`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})

	}
}
