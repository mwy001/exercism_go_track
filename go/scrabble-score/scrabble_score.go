package scrabble

import "unicode"

var score1 []rune = []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}
var score2 []rune = []rune{'D', 'G'}
var score3 []rune = []rune{'B', 'C', 'M', 'P'}
var score4 []rune = []rune{'F', 'H', 'V', 'W', 'Y'}
var score5 []rune = []rune{'K'}
var score8 []rune = []rune{'J', 'X'}
var score10 []rune = []rune{'Q', 'Z'}

func contains(s []rune, c rune) bool {
    for _, a := range s {
        if a == c {
            return true
        }
    }
    return false
}

func Score(input string) int {
	var res int = 0

	for _, c := range input {

		c = unicode.ToUpper(c)

		switch {
		case contains(score1, c):
			res += 1
		case contains(score2, c):
			res += 2
		case contains(score3, c):
			res += 3
		case contains(score4, c):
			res += 4
		case contains(score5, c):
			res += 5
		case contains(score8, c):
			res += 8
		case contains(score10, c):
			res += 10			
		}
	}

	return res
}
