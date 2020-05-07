package luhn

import (
	"unicode"
)

func Valid(code string) bool {
	numDigits := 0
	sum := 0

	for index := len(code) - 1; index >= 0; index-- {
		character := rune(code[index])

		switch {
		case unicode.IsSpace(character):
			continue
		case unicode.IsDigit(character):
			value := int(character) - '0'

			if numDigits%2 != 0 {
				value *= 2

				if value > 9 {
					value -= 9
				}
			}

			sum += value
			numDigits++
		default:
			return false
		}
	}

	return numDigits > 1 && sum%10 == 0
}
