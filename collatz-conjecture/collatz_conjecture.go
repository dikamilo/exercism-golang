package collatzconjecture

import "errors"

func CollatzConjecture(value int) (int, error) {
	if value <= 0 {
		return -1, errors.New("value must be greater than 0")
	}

	steps := 0

	for ; value != 1; steps++ {
		if test := value % 2; test == 0 {
			value = value / 2
		} else {
			value = (value * 3) + 1
		}
	}

	return steps, nil
}
