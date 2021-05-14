// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"unicode"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var result string

	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '-'
	}

	words := strings.FieldsFunc(s, f)

	for _, w := range words {
		if len(w) > 0 {
			for _, c := range w {
				if unicode.IsLetter(c) {
					result += strings.ToUpper(string(c))
					break
				}
			}
		}
	}
	return result
}
