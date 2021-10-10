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

func Test_LowerD(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "2021/01/15",
			expected: true,
		},
		{
			input:    "2021/1/15",
			expected: false,
		},
		{
			input:    "2021/xx/15",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`\d\d\d\d/\d\d/\d\d`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_UpperD(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "123 01 2021",
			expected: false,
		},
		{
			input:    "Jan 01 2021",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`\D\D\D \d\d \d\d\d\d`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_Asterisk(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Yeah!",
			expected: true,
		},
		{
			input:    "Yeaaaaaaah!",
			expected: true,
		},
		{
			input:    "Yeh!",
			expected: true,
		},
		{
			input:    "YEAH!",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`Yea*h!`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_Plus(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Yeah!",
			expected: true,
		},
		{
			input:    "Yeaaaaaaah!",
			expected: true,
		},
		{
			input:    "Yeh!",
			expected: false,
		},
		{
			input:    "YEAH!",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`Yea+h!`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_EqualN(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Yeah!",
			expected: false,
		},
		{
			input:    "Yeaah!",
			expected: false,
		},
		{
			input:    "Yeaaah!",
			expected: true,
		},
		{
			input:    "Yeaaaah!",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`Yea{3}h!`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_MoreThanN(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Yeah!",
			expected: false,
		},
		{
			input:    "Yeaah!",
			expected: false,
		},
		{
			input:    "Yeaaah!",
			expected: true,
		},
		{
			input:    "Yeaaaah!",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`Yea{3,}h!`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}
