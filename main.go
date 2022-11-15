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
		g := graph.NewGraph[int](func(a, b int) bool {
			return a == b
		})
		g.Start = graph.NewGraphNode(101)
		n1 := graph.NewGraphNode(102)
		n2 := graph.NewGraphNode(103)
		n3 := graph.NewGraphNode(104)
		n4 := graph.NewGraphNode(105)
		n5 := graph.NewGraphNode(106)
		n6 := graph.NewGraphNode(108)

		g.Start.Edges = append(g.Start.Edges, n1)
		g.Start.Edges = append(g.Start.Edges, n2)

		n1.Edges = append(n1.Edges, n4)
		n1.Edges = append(n1.Edges, n5)

		n2.Edges = append(n2.Edges, n1)
		n2.Edges = append(n2.Edges, g.Start)

		n5.Edges = append(n5.Edges, n3)

		node := g.BFS(101)
		fmt.Println(node)

		g.Add(node, n6)

		node = g.BFS(101)
		fmt.Println(node)
	}

}
