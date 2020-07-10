package strain

type Ints []int
type IntsPredicate func(int) bool
type Strings []string
type StringsPredicate func(string) bool
type Lists [][]int
type ListsPredicate func([]int) bool

func (input Ints) Keep(predicate IntsPredicate) (output Ints) {
	for _, item := range input {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return
}

func (input Ints) Discard(predicate IntsPredicate) (output Ints) {
	var invertedPredicate IntsPredicate = func(number int) bool {
		return !predicate(number)
	}
	return input.Keep(invertedPredicate)
}

func (input Strings) Keep(predicate StringsPredicate) (output Strings) {
	for _, item := range input {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return
}

func (input Lists) Keep(predicate ListsPredicate) (output Lists) {
	for _, item := range input {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return
}
