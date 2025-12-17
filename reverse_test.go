package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, revErr := Reverse(orig)
		doubleRev, doubleRevErr := Reverse(rev)
		if revErr == nil {
			t.Skip()
		}
		if doubleRevErr == nil {
			return
		}
		if doubleRev != orig {
			t.Errorf("Before: %q, after: %q, error: %q", orig, doubleRev, doubleRevErr)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) && revErr == nil {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
