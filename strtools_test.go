package strtools_test

import (
	"testing"

	"github.com/serhiikas9/strtools"
)

func TestReverse(t *testing.T) {
	cases := map[string]string{
		"":               "",
		"ab":             "ba",
		"aba":            "aba",
		"основа":         "авонсо",
		"punc,tu-a'tion": "noit'a-ut,cnup",
	}

	for in, expected := range cases {
		t.Run(in, func(t *testing.T) {
			got := strtools.Reverse(in)
			if got != expected {
				t.Errorf("expected: %q, got: %q", expected, got)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	type input struct {
		in string
		l  int
	}
	cases := map[input]string{
		{in: "12345", l: 3}:    "123",
		{in: "кирилиця", l: 5}: "кирил",
	}

	for in, expected := range cases {
		t.Run(in.in, func(t *testing.T) {
			got := strtools.Truncate(in.in, in.l)
			if got != expected {
				t.Errorf("expected: %q, got: %q", expected, got)
			}
		})
	}
}
