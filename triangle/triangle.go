package triangle

import (
	"math"
	"sort"
)

type Kind int

const (
	NaT = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)

	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}

	switch {
	case sides[0]+sides[1] < sides[2]:
		return NaT
	case sides[0] == sides[1] && sides[0] == sides[2]:
		return Equ
	case sides[0] == sides[1] || sides[1] == sides[2]:
		return Iso
	default:
		return Sca
	}
}
