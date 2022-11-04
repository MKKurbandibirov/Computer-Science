package comb_sort

import "algorithms/sorting/buble_sort"

func CombSort[T any](array []T, less func(a, b T) bool) {
	k := 1.2473309
	step := len(array)
	for step > 1 {
		for i := 0; i+step < len(array); i++ {
			if less(array[i+step], array[i]) {
				array[i+step], array[i] = array[i], array[i+step]
			}
		}
		step = int(float64(step) / k)
	}
	buble_sort.BubbleSort(array, less)
}
