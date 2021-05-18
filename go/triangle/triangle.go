package triangle

import (
	"math"
)

// Kind defines the type for different triangles
type Kind int

// Defines the triangle types
const (
    NaT = iota
    Equ = iota
    Iso = iota
    Sca = iota
)

func isValidSideLen(a float64) bool {
	return a > 0 && !math.IsNaN(a) && a != math.Inf(1)
}

// KindFromSides checks the triangle type from the 3 input sides
func KindFromSides(a, b, c float64) Kind {

	if !isValidSideLen(a) || !isValidSideLen(b) || !isValidSideLen(c) {
		return NaT
	}

	if a + b < c || a + c < b || b + c < a {
		return NaT

	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
