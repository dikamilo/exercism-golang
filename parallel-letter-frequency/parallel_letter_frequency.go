package letter

type FreqMap map[rune]int

func Frequency(input string) FreqMap {
	frequency := FreqMap{}

	for _, character := range input {
		frequency[character]++
	}

	return frequency
}

func ConcurrentFrequency(inputs []string) (frequency FreqMap) {
	channel := make(chan FreqMap)
	defer close(channel)

	for _, input := range inputs {
		go func(input string) {
			channel <- Frequency(input)
		}(input)
	}

	frequency = <-channel

	for i := 0; i < len(inputs)-1; i++ {
		for character, count := range <-channel {
			frequency[character] += count
		}
	}

	return
}
