package binary_search

func BinarySearch[T any](array []T, num T, less, equal func(a, b T) bool) int {
	l, r := 0, len(array)-1
	var m int
	for l <= r {
		m = (l + r) / 2
		if equal(array[m], num) {
			return m
		} else if less(array[m], num) {
			l = m
		} else {
			r = m
		}
	}
	return -1
}
