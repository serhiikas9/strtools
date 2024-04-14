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
