package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var line string
	var proverbialRhyme []string
	size := len(rhyme)

	for index, verb := range rhyme {
		if index < size-1 && size > 1 {
			line = fmt.Sprintf("For want of a %s the %s was lost.", verb, rhyme[index+1])
		} else {
			line = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		}

		proverbialRhyme = append(proverbialRhyme, line)
	}

	return proverbialRhyme
}
