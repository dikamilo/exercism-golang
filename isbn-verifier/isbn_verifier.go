package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	result, counter := 0, 10

	for _, character := range isbn {
		if unicode.IsDigit(character) {
			result += int(character-'0') * counter
			counter--
		}
	}

	if strings.HasSuffix(isbn, "X") {
		result += 10
		counter = 0
	}

	return result%11 == 0 && counter == 0
}
