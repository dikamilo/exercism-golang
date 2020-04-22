package hamming

import "errors"

func Distance(a, b string) (int, error) {
	distance := 0

	if len(a) != len(b) {
		return -1, errors.New("a and b straps have different length")
	}

	for index := range a {
		if a[index] != b[index] {
			distance += 1
		}
	}

	return distance, nil
}
