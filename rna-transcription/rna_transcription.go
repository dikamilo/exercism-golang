package strand

import "strings"

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	return strings.Map(func(nucleotide rune) rune {
		return complements[nucleotide]
	}, dna)
}
