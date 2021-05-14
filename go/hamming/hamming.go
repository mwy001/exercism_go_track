package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Invalid parameters")
	}

	var result int

	length := len(a)

	for i:=0; i < length; i++ {
		if a[i] != b[i] {
			result += 1
		}
	}

	return result, nil
}
