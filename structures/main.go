package main

import (
	"data-structures/binary_tree"
	"fmt"
)

func main() {
	tr := new(binary_tree.BinaryTree[int])
	tr.Insert(23, func(a, b int) bool {
		return a < b
	})
	tr.Insert(14, func(a, b int) bool {
		return a < b
	})
	tr.Insert(45, func(a, b int) bool {
		return a < b
	})

	for x := range tr.Iter() {
		fmt.Printf("%v ", x)
	}
}
