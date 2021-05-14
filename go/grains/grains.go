package grains

import (
	"errors"
)

func Square(input int) (uint64, error) {
	var res uint64 = 1
	if input < 1 || input > 64 {
		return 0, errors.New("")
	}

	for i := 1; i < input; i++ {
		res *= 2
	}

	return res, nil
}

func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i < 65; i++ {
		sq, _ := Square(i)
		sum += sq
	}

	return sum
}
