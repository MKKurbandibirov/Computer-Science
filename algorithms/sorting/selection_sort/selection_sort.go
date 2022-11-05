package selection_sort

func min[T any](array []T, start int, less func(a, b T) bool) int {
	min := start
	for i := start; i < len(array); i++ {
		if less(array[i], array[min]) {
			min = i
		}
	}
	return min
}

func SelectionSort[T any](array []T, less func(a, b T) bool) {
	for i := 0; i < len(array); i++ {
		m := min(array, i, less)
		array[i], array[m] = array[m], array[i]
	}
}
