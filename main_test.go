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

func Test_MoreThanNAndLessThanM(t *testing.T) {
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
			expected: true,
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
			r := regexp.MustCompile(`Yea{2,3}h!`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_Question(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "color",
			expected: true,
		},
		{
			input:    "colour",
			expected: true,
		},
		{
			input:    "colur",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`colou?r`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_SquareBrackets(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Tatsuya",
			expected: true,
		},
		{
			input:    "tatsuya",
			expected: true,
		},
		{
			input:    "Katsuya",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`[Tt]atsuya`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_Pipe(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Jack",
			expected: true,
		},
		{
			input:    "James",
			expected: true,
		},
		{
			input:    "John",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`Jack|James`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_Caret(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "you are happy",
			expected: false,
		},
		{
			input:    "happy new year",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`^happy`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_Dollar(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "you are happy",
			expected: true,
		},
		{
			input:    "happy new year",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`happy$`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_UpperB(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "mac",
			expected: false,
		},
		{
			input:    "emacs",
			expected: true,
		},
		{
			input:    "macintosh",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`\Bmac\B`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

/*
Go don't have backreference.
func Test_Backreference(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "A promise is promise",
			expected: true,
		},
		{
			input:    "A bargain is bargain",
			expected: true,
		},
		{
			input:    "A liar is worse",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`A (\w+) is \1`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}
*/

func Test_Escape(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "gihyo.co.jp",
			expected: true,
		},
		{
			input:    "gihyo\\.co\\.jp",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`gihyo\.co\.jp`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_HTML(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "<br>",
			expected: true,
		},
		{
			input:    "<a href='https://example.com>example.com</a>'",
			expected: true,
		},
		{
			input:    "[example.com](https://example.com)",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`<(".*?"|'.*?'|[^'"])*?>`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_PostalCode(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "123-0001",
			expected: true,
		},
		{
			input:    "1230001'",
			expected: false,
		},
		{
			input:    "1-0001",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`\d{3}-\d{4}`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_TelephoneNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "03-9999-9999",
			expected: true,
		},
		{
			input:    "072-9999-9999'",
			expected: true,
		},
		{
			input:    "123-0001",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`0\d{1,4}-\d{1,4}-\d{4}`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}

func Test_URL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "http://example.com",
			expected: true,
		},
		{
			input:    "072-9999-9999'",
			expected: true,
		},
		{
			input:    "123-0001",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := regexp.MustCompile(`0\d{1,4}-\d{1,4}-\d{4}`)
			assert.Equal(t, tt.expected, r.MatchString(tt.input))
		})
	}
}
