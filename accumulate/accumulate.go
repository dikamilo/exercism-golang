package accumulate

type Collection []string
type Converter func(string) string

func Accumulate(items Collection, converter Converter) Collection {
	result := make(Collection, 0, len(items))

	for _, item := range items {
		result = append(result, converter(item))
	}

	return result
}
