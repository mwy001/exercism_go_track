package twelve

import (
	"fmt"
)

var dayMapping map[int]string = map[int]string{
	1:"first",
	2:"second",
	3:"third",
	4:"fourth",
	5:"fifth",
	6:"sixth",
	7:"seventh",
	8:"eighth",
	9:"ninth",
	10:"tenth",
	11:"eleventh",
	12:"twelfth",
}

var itemMapping map[int]string = map[int]string{
	1:"a Partridge in a Pear Tree.",
	2:"two Turtle Doves",
	3:"three French Hens",
	4:"four Calling Birds",
	5:"five Gold Rings",
	6:"six Geese-a-Laying",
	7:"seven Swans-a-Swimming",
	8:"eight Maids-a-Milking",
	9:"nine Ladies Dancing",
	10:"ten Lords-a-Leaping",
	11:"eleven Pipers Piping",
	12:"twelve Drummers Drumming",
}

func Verse(input int) string {

	var part2 string

	for i := input; i > 1; i-- {
		part2 += itemMapping[i] + ", "
	}

	if input > 1 {
		part2 += "and " + itemMapping[1]
	} else {
		part2 += itemMapping[1]
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", dayMapping[input], part2)
}

func Song() string {
	var res string
	for i := 1; i < 12; i++ {
		res += Verse(i) + "\n"
	}

	res += Verse(12)

	return res
}
