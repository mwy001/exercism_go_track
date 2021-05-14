package clock

import (
	"fmt"
)

type Clock struct {
	h int
	m int	
}

func regulate(h int, m int) (int, int) {
	var total = h * 60 + m

	for; total < 0; {
		total += 1440
	}

	for; total >= 1440; {
		total -= 1440
	}

	return total / 60, total % 60
}

func New(h int, m int) Clock {
	h, m = regulate(h, m)

	return Clock{h:h, m:m}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.h, clock.m)
}

func (clock Clock) Add(m int) Clock {
	return New(clock.h, clock.m + m)
}


func (clock Clock) Subtract(m int) Clock {
	return New(clock.h, clock.m - m)
}
