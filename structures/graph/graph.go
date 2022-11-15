package graph

import "algos-and-structures/structures/dequeue"

type GraphNode[T any] struct {
	Value T
	Edges []*GraphNode[T]
}

func NewGraphNode[T any](val T) *GraphNode[T] {
	return &GraphNode[T]{
		Value: val,
		Edges: nil,
	}
}

type Graph[T any] struct {
	Start *GraphNode[T]
	equal func(a, b T) bool
}

func NewGraph[T any](equal func(a, b T) bool) *Graph[T] {
	return &Graph[T]{
		Start: nil,
		equal: equal,
	}
}

func (g *Graph[T]) contains(arr []T, val T) bool {
	for i := 0; i < len(arr); i++ {
		if g.equal(arr[i], val) {
			return true
		}
	}
	return false
}

func (g *Graph[T]) BFS(val T) *GraphNode[T] {
	visited := make([]T, 0)
	deq := new(dequeue.Dequeue[*GraphNode[T]])
	deq.AddBack(g.Start)
	for !deq.IsEmpty() {
		i := 0
		l := deq.Size()
		for i < l {
			node, _ := deq.PeakFront()
			if g.equal(node.Value, val) {
				return node
			}
			visited = append(visited, node.Value)
			deq.RemoveFront()
			for _, x := range node.Edges {
				if !g.contains(visited, x.Value) {
					deq.AddBack(x)
				}
			}
			i++
		}
	}
	return nil
}

func (g *Graph[T]) Add(dst, new *GraphNode[T]) {
	node := g.BFS(dst.Value)
	node.Edges = append(node.Edges, new)
}
