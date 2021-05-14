package luhn

import (
	"strings"
	"unicode"
	"strconv"
)

// What i have learnt?
// ==> srings.TrimSpace only trim leading and trailing spaces, but not spaces
//      in the middle of the string

func Valid(input string) bool {

	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}

	for _, c := range input {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	var sum int

	len := len(input)

	for i:= len - 1; i >= 0; i-- {
		if (len - i) % 2 == 0 {
			v, _ := strconv.Atoi(string(input[i]))
			v *= 2

			if v > 9 {
				v = v - 9
			}

			sum += v

		} else {
			v, _ := strconv.Atoi(string(input[i]))
			sum += v
		}
	}

	if sum % 10 == 0 {
		return true
	} else {
		return false
	}
}
