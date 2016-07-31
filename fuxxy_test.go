package fuxxy

import (
	"reflect"
	"testing"
)

func TestMatches(t *testing.T) {

	var tests = []struct {
		s      string
		input  []string
		output []string
	}{
		{"slc",
			[]string{"salt lake city", "slick", "sdd", "ddslckk", "ds", "dddd", "slc"},
			[]string{"slc", "ddslckk", "slick", "salt lake city"},
		},
	}

	for _, test := range tests {
		if matches := Matches(test.input, test.s); !reflect.DeepEqual(matches, test.output) {
			t.Errorf("Expected %q, got %q", test.output, matches)
		}
	}

}

func BenchmarkMatches(b *testing.B) {

	s, names := "slc", []string{"salt lake city", "slick", "sdd", "ddslckk", "ds", "dddd", "slc"}

	for i := 0; i < b.N; i++ {
		Matches(names, s)
	}

}
