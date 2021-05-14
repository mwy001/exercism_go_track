// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func IsUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func HasLetter(s string) bool {
    for _, r := range s {
        if unicode.IsLetter(r) {
            return true
        }
    }
    return false
}

func IsLower(s string) bool {
    for _, r := range s {
        if !unicode.IsLower(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case strings.HasSuffix(remark, "?"):
		if IsUpper(remark) && HasLetter(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	case IsUpper(remark) && HasLetter(remark):
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
