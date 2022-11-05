package insertion_sort

func InsertionSort[T any](array []T, less func(a, b T) bool) {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 && less(array[j], array[j-1]) {
			array[j], array[j-1] = array[j-1], array[j]
			j--
		}
	}
}
