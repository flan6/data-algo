package sorting

func swap[T any](a, b *T) {
	temp := *a
	*a = *b
	*b = temp
}
