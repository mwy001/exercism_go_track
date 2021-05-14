package isogram

import "unicode"

func IsIsogram(input string) bool {
	var letters map[rune]int = map[rune]int{}

	for _, c := range input {
		if unicode.IsLetter(c) {
			c = unicode.ToUpper(c)

			if _, ok := letters[c]; ok {
				return false
			}
			
			letters[c] = 1
		}
	}
	return true
}
