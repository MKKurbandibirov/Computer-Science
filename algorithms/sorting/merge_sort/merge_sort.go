package merge_sort

func merge[T any](a, b []T, less func(a, b T) bool) []T {
	var final []T
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if less(a[i], b[j]) {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func sort[T any](arr []T, less func(a, b T) bool) []T {
	if len(arr) < 2 {
		return arr
	}
	first := sort(arr[:len(arr)/2], less)
	second := sort(arr[len(arr)/2:], less)
	return merge(first, second, less)
}

func MergeSort[T any](array []T, less func(a, b T) bool) []T {
	return sort(array, less)
}
