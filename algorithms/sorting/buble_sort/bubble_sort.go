package buble_sort

func BubbleSort[T any](array []T, less func(a, b T) bool) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if less(array[j], array[i]) {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
