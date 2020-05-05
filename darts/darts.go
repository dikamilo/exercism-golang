package darts

import "math"

func Score(x, y float64) int {
	switch score := math.Hypot(x, y); {
	case score <= 1.:
		return 10
	case score <= 5.:
		return 5
	case score <= 10.:
		return 1
	default:
		return 0
	}
}
