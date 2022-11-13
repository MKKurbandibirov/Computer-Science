package quick_sort

func partition[T any](arr []T, low, high int, less func(a, b T) bool) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if less(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func sort[T any](arr []T, low, high int, less func(a, b T) bool) {
	if low < high {
		var p int
		p = partition(arr, low, high, less)
		sort[T](arr, low, p-1, less)
		sort[T](arr, p+1, high, less)
	}
}

func QuickSort[T any](array []T, less func(a, b T) bool) {
	sort[T](array, 0, len(array)-1, less)
}
