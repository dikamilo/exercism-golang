package reverse

func Reverse(text string) string {
	index := len(text)
	result := make([]rune, index)

	for _, character := range text {
		index--
		result[index] = character
	}

	return string(result[index:])
}
