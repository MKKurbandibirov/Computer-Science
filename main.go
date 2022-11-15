package main

import (
	"algos-and-structures/structures/graph"
	"fmt"
)

func main() {
	//{
	//	hMax := heap.NewHeap[int](func(a, b int) bool {
	//		return a < b
	//	}, func(a, b int) bool {
	//		return a > b
	//	}, heap.MAX)
	//
	//	hMax.Add(23)
	//	hMax.Add(45)
	//	hMax.Add(67)
	//	hMax.Add(98)
	//	hMax.Add(2)
	//	hMax.Add(13)
	//
	//	fmt.Println(hMax)
	//
	//	hMin := heap.NewHeap[int](func(a, b int) bool {
	//		return a < b
	//	}, func(a, b int) bool {
	//		return a > b
	//	}, heap.MIN)
	//
	//	hMin.Add(23)
	//	hMin.Add(45)
	//	hMin.Add(67)
	//	hMin.Add(98)
	//	hMin.Add(2)
	//	hMin.Add(13)
	//
	//	fmt.Println(hMin)
	//
	//	for x := range hMin.Iter() {
	//		fmt.Println(x)
	//	}
	//
	//	fmt.Println(hMin)
	//
	//	for x := range hMin.Iter() {
	//		fmt.Println(x)
	//	}
	//
	//	fmt.Println(hMin)
	//}

	//{
	//	b := binary_tree.NewBinaryTree[int](func(a, b int) bool {
	//		return a < b
	//	}, func(a, b int) bool {
	//		return a == b
	//	})
	//
	//	b.Insert(24)
	//	b.Insert(12)
	//	b.Insert(45)
	//	b.Insert(15)
	//	b.Insert(6)
	//	b.Insert(35)
	//	fmt.Println(b)
	//}

	{
		g := graph.NewGraph[int](1, func(a, b int) bool {
			return a == b
		})

		g.Add(1, 2)
		g.Add(1, 3)
		g.Add(2, 4)
		g.Add(2, 5)
		g.Add(1, 4)
		g.Add(3, 4)
		g.Add(3, 2)

		fmt.Println(g.Find(1))
		fmt.Println(g.Find(2))

		g.Remove(4)

		fmt.Println(g.Find(3))
	}

}
