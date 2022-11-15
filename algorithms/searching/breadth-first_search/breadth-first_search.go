package breadth_first_search

import (
	"algos-and-structures/structures/dequeue"
	g "algos-and-structures/structures/graph"
)

func contains[T any](arr []T, val T, equal func(a, b T) bool) bool {
	for i := 0; i < len(arr); i++ {
		if equal(arr[i], val) {
			return true
		}
	}
	return false
}

func BFS[T any](gr *g.Graph[T], val T, equal func(a, b T) bool) *g.GraphNode[T] {
	if gr == nil {
		return nil
	}
	visited := make([]T, 0)
	deq := new(dequeue.Dequeue[*g.GraphNode[T]])
	deq.AddBack(gr.Start)
	for !deq.IsEmpty() {
		i := 0
		l := deq.Size()
		for i < l {
			node, _ := deq.PeakFront()
			if equal(node.Value, val) {
				return node
			}
			visited = append(visited, node.Value)
			deq.RemoveFront()
			for _, x := range node.Edges {
				if !contains(visited, x.Value, equal) {
					deq.AddBack(x)
				}
			}
			i++
		}
	}
	return nil
}
