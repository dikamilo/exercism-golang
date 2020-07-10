package diffsquares

func SquareOfSum(size int) int {
	var result int

	for ; size > 0; size-- {
		result += size
	}

	return result * result
}

func SumOfSquares(size int) (result int) {
	for ; size > 0; size-- {
		result += size * size
	}

	return
}

func Difference(size int) int {
	return SquareOfSum(size) - SumOfSquares(size)
}
