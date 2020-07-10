package dna

import (
	"errors"
)

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	histogram := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, nucleotide := range d {
		switch nucleotide {
		case 'A', 'C', 'G', 'T':
			histogram[nucleotide]++
		default:
			return nil, errors.New("invalid nucleotide")
		}
	}

	return histogram, nil
}
