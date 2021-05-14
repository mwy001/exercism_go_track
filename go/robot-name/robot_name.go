package robotname

import (
	"fmt"
	"errors"
)

type Robot struct{
	MyName string
}

var counter1 int = 0
var counter2 int = 0

func newName() string {
	if counter1 == 1000 {
		fmt.Println(counter2)
		counter1 = 0
		counter2 ++
		if counter2 == 26*26 {
			return ""
		}
	}

	c0 := '0' + counter1 % 10
	c1 := '0' + (counter1 / 10) % 10
	c2 := '0' + (counter1 / (10 * 10)) % 10

	c3 := 'A' + counter2 % 26
	c4 := 'A' + (counter2 / 26) % 26

	counter1 ++

	return fmt.Sprintf("%c%c%c%c%c", c4, c3, c2, c1, c0)
}

func New() *Robot {
	return &Robot{MyName:newName()}
}

func (r *Robot) Name() (string, error) {
	if r.MyName != "" {
		return r.MyName, nil
	} else {
		return "", errors.New("namespace is exhausted")
	}
}

func (r *Robot) Reset() {
	r.MyName = newName()
}
