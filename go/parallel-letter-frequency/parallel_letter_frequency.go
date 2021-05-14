package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}


func Frequency2(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}
	var rwm sync.RWMutex

	var wg sync.WaitGroup

	for _, s := range ss {
		wg.Add(1)

		go func(s string) {
			defer wg.Done()

			for _, r := range s {
				rwm.Lock()
				m[r]++
				rwm.Unlock()
			}
		}(s)
	}
    wg.Wait()

	return m
}
