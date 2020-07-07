package robotname

import (
	"errors"
	"math/rand"
)

const (
	namesLimit = 26 * 26 * 10 * 10 * 10
	nameFormat = "AA000"
)

var usedNames = make(map[string]bool, namesLimit)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if len(r.name) == 0 {
		return r.name, r.generateName()
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func (r *Robot) generateName() error {
	nameCandidate := prepareName()

	if len(usedNames) == namesLimit {
		return errors.New("namespace is exhausted")
	}

	for usedNames[nameCandidate] {
		nameCandidate = prepareName()
	}

	usedNames[nameCandidate] = true
	r.name = nameCandidate

	return nil
}

func prepareName() string {
	name := make([]rune, len(nameFormat))

	for position, character := range nameFormat {
		switch character {
		case 'A':
			name[position] = 'A' + rune(rand.Intn(26))
		default:
			name[position] = '0' + rune(rand.Intn(9))
		}
	}

	return string(name)
}
