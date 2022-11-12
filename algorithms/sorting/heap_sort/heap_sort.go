package heap_sort

import (
	"algos-and-structures/structures/heap"
)

func HeapSort[T any](array []T, less, more func(a, b T) bool) {
	h := heap.NewHeap(less, more, heap.MIN)
	cpy := make([]T, len(array))
	copy(cpy, array)
	h.Build(cpy)
	i := 0
	for x := range h.Iter() {
		array[i] = x
		i++
	}
}
