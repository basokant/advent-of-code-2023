package util

func Values[K comparable, V any](m map[K]V) []V {
	values := []V{}
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
