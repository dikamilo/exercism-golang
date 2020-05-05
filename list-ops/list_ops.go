package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

func (l IntList) Foldr(f binFunc, initial int) int {
	for i := l.Length() - 1; i >= 0; i-- {
		initial = f(l[i], initial)
	}

	return initial
}

func (l IntList) Foldl(f binFunc, initial int) int {
	for _, value := range l {
		initial = f(initial, value)
	}

	return initial
}

func (l IntList) Filter(f predFunc) IntList {
	result := IntList{}

	for _, value := range l {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

func (l IntList) Length() int {
	return len(l)
}

func (l IntList) Map(f unaryFunc) IntList {
	result := make(IntList, l.Length())

	for index, value := range l {
		result[index] = f(value)
	}

	return result
}

func (l IntList) Reverse() IntList {
	size := l.Length()
	result := make(IntList, size)

	for index, value := range l {
		result[size-index-1] = value
	}

	return result
}

func (l IntList) Append(values IntList) IntList {
	result := make(IntList, 0, l.Length()+values.Length())

	result = append(result, l...)
	result = append(result, values...)

	return result
}

func (l IntList) Concat(values []IntList) IntList {
	result := make(IntList, 0, l.Length()+len(values))
	result = append(result, l...)

	for _, value := range values {
		result = append(result, value...)
	}

	return result
}
