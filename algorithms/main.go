package main

import (
	"algorithms/sorting/buble_sort"
	"fmt"
)

func main() {
	arr := []int{3, 7, -2, 72, 67, 112, -67, 0, -12, 21}
	buble_sort.BubbleSort(arr, func(a, b int) bool {
		return a < b
	})
	fmt.Println(arr)
}
