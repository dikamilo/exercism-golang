package letter

import (
	"sync"
)

type FreqMap map[rune]int

func Frequency(input string) FreqMap {
	frequency := FreqMap{}

	for _, character := range input {
		frequency[character]++
	}

	return frequency
}

func ConcurrentFrequency(inputs []string) FreqMap {
	frequency := FreqMap{}
	wg := sync.WaitGroup{}
	mutex := sync.RWMutex{}

	for _, input := range inputs {
		wg.Add(1)

		go func(input string) {
			defer wg.Done()

			for character, count := range Frequency(input) {
				mutex.Lock()
				frequency[character] += count
				mutex.Unlock()
			}
		}(input)
	}

	wg.Wait()
	return frequency
}
