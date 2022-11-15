package breadth_first_search

import (
	g "algos-and-structures/structures/graph"
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	gr := g.NewGraph[int](1, func(a, b int) bool {
		return a == b
	})
	gr.Add(1, 2)
	gr.Add(1, 3)
	gr.Add(2, 4)
	gr.Add(2, 5)
	gr.Add(1, 4)
	gr.Add(3, 4)
	gr.Add(3, 2)

	type args struct {
		start *g.Graph[int]
		val   int
		equal func(a, b int) bool
	}
	tests := []struct {
		name string
		args args
		want *g.GraphNode[int]
	}{
		{
			name: "Case 1: Search in empty graph",
			args: args{
				start: nil,
				val:   42,
				equal: func(a, b int) bool {
					return a == b
				},
			},
			want: nil,
		},
		{
			name: "Case 2: Search in full graph",
			args: args{
				start: gr,
				val:   4,
				equal: func(a, b int) bool {
					return a == b
				},
			},
			want: gr.Start.Edges[2],
		},
		{
			name: "Case 2: Search not existing element",
			args: args{
				start: gr,
				val:   25,
				equal: func(a, b int) bool {
					return a == b
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.start, tt.args.val, tt.args.equal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
