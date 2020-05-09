package scrabble

import "strings"

var scores = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(input string) int {
	var points int

	for _, letter := range strings.ToUpper(input) {
		for letters, score := range scores {
			if strings.ContainsRune(letters, letter) {
				points += score
			}
		}
	}

	return points
}
