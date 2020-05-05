package raindrops

import "strconv"

func Convert(drops int) string {
	var result string

	if drops%3 == 0 {
		result += "Pling"
	}

	if drops%5 == 0 {
		result += "Plang"
	}

	if drops%7 == 0 {
		result += "Plong"
	}

	if len(result) > 0 {
		return result
	}

	return strconv.Itoa(drops)
}
