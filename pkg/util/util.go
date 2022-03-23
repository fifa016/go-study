package util

func Swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func FakeSwap(a, b int) {
	a, b = b, a
}

