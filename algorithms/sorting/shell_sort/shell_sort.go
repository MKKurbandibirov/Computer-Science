package shell_sort

func ShellSort[T any](array []T, less func(a, b T) bool) {
	for s := len(array) / 2; s > 0; s /= 2 {
		for i := s; i < len(array); i++ {
			for j := i - s; j >= 0 && less(array[j+s], array[j]); j -= s {
				array[j], array[j+s] = array[j+s], array[j]
			}
		}
	}
}
