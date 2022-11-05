package tree_sort

import "algos-and-structures/structures/binary_tree"

func TreeSort[T any](array []T, less func(a, b T) bool) {
	tree := new(binary_tree.BinaryTree[T])
	for i := 0; i < len(array); i++ {
		tree.Insert(array[i], less)
	}
	i := 0
	for x := range tree.Iter() {
		array[i] = x
		i++
	}
}
