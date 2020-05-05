package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	filteredWord := strings.ToLower(word)

	for _, character := range filteredWord {
		isLetter := unicode.IsLetter(character)
		hasDuplicates := strings.Count(filteredWord, string(character)) > 1

		if isLetter && hasDuplicates {
			return false
		}
	}

	return true
}
