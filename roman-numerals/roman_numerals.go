package romannumerals

import "errors"

const (
	min = 1
	max = 3000
)

type mapping struct {
	value  int
	letter string
}

var alphabet = []mapping{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(number int) (string, error) {
	var roman string

	if number < min || number > max {
		return roman, errors.New("invalid value")
	}

	for number > 0 {
		for _, mapping := range alphabet {
			counter := number / mapping.value

			for ; counter > 0; counter-- {
				roman += mapping.letter
				number -= mapping.value
			}
		}
	}

	return roman, nil
}
