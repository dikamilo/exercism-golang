package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Number(number string) (string, error) {
	return prepare(extractDigits(number))
}

func AreaCode(number string) (string, error) {
	n, err := Number(number)

	if err != nil {
		return "", err
	}

	return n[0:3], nil
}

func Format(number string) (string, error) {
	n, err := Number(number)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", n[0:3], n[3:6], n[6:]), nil
}

func extractDigits(number string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, number)
}

func prepare(number string) (string, error) {
	size := len(number)

	switch {
	case size < 10 || size > 11:
		return "", errors.New("incorrect number of digits")
	case size == 11:
		if number[0] != '1' {
			return "", errors.New("11 digits must start with 1")
		}
		number = number[1:]
	}

	areaCodePrefix, exchangePrefix := number[0], number[3]

	if areaCodePrefix == '0' || areaCodePrefix == '1' {
		return "", errors.New("invalid area code")
	}

	if exchangePrefix == '0' || exchangePrefix == '1' {
		return "", errors.New("invalid exchange code")
	}
	return number, nil
}
