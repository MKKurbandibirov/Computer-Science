package gnome_sort

func GnomeSort[T any](array []T, less func(a, b T) bool) {
	for i := 0; i < len(array); {
		if i == 0 || !less(array[i], array[i-1]) {
			i++
		} else {
			array[i-1], array[i] = array[i], array[i-1]
			i--
		}
	}
}
