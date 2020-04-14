package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) (acronym string) {
	s = strings.ReplaceAll(s, "-", " ")

	for _, word := range strings.Fields(s) {
		acronym += string(word[strings.IndexFunc(word, unicode.IsLetter)])
	}

	acronym = strings.ToUpper(acronym)
	return
}
