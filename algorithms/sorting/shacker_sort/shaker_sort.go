package shacker_sort

func ShakerSort[T any](array []T, less, more func(a, b T) bool) {
	beg, end := 0, len(array)-1
	b := true
	for b {
		b = false
		for i := beg; i < end; i++ {
			if less(array[i+1], array[i]) {
				array[i+1], array[i] = array[i], array[i+1]
				b = true
			}
		}
		if !b {
			break
		}
		for i := end; i > beg; i-- {
			if more(array[i-1], array[i]) {
				array[i-1], array[i] = array[i], array[i-1]
				b = true
			}
		}
		beg++
		end--
	}
}
