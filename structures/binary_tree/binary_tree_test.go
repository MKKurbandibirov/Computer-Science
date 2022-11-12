package binary_tree

import (
	"reflect"
	"testing"
)

type person struct {
	Id   int
	Name string
}

var testTreeInt BinaryTree[int]
var testTreeString BinaryTree[string]

func TestBinaryTree_Find(t *testing.T) {
	testTreeInt.Root = new(node[int])
	testTreeInt.Root.Value = 10
	testTreeInt.Root.Left = new(node[int])
	testTreeInt.Root.Left.Value = 7
	testTreeInt.Root.Right = new(node[int])
	testTreeInt.Root.Right.Value = 13
	testTreeInt.Root.Left.Left = new(node[int])
	testTreeInt.Root.Left.Left.Value = 4
	testTreeInt.Root.Left.Right = new(node[int])
	testTreeInt.Root.Left.Right.Value = 9

	type fields struct {
		Root  *node[int]
		equal func(a, b int) bool
		less  func(a, b int) bool
	}
	type args struct {
		elem int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node[int]
	}{
		{
			name: "Case 1: Find existing element",
			fields: fields{
				Root: testTreeInt.Root,
				equal: func(a, b int) bool {
					return a == b
				},
				less: func(a, b int) bool {
					return a < b
				},
			},
			args: args{
				elem: 9,
			},
			want: testTreeInt.Root.Left.Right,
		},
		{
			name: "Case 2: Find not existing element",
			fields: fields{
				Root: testTreeInt.Root,
				equal: func(a, b int) bool {
					return a == b
				},
				less: func(a, b int) bool {
					return a < b
				},
			},
			args: args{
				elem: 346,
			},
			want: nil,
		},
		{
			name: "Case 3: Find in nil root",
			fields: fields{
				Root: nil,
				equal: func(a, b int) bool {
					return a == b
				},
				less: func(a, b int) bool {
					return a < b
				},
			},
			args: args{
				elem: 346,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree[int]{
				Root:  tt.fields.Root,
				less:  tt.fields.less,
				equal: tt.fields.equal,
			}
			if got := b.Find(tt.args.elem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Insert(t *testing.T) {
	testTreeString.Root = new(node[string])
	testTreeString.Root.Value = "H"
	testTreeString.Root.Left = new(node[string])
	testTreeString.Root.Left.Value = "C"
	testTreeString.Root.Right = new(node[string])
	testTreeString.Root.Right.Value = "W"
	testTreeString.Root.Left.Left = new(node[string])
	testTreeString.Root.Left.Left.Value = "B"
	testTreeString.Root.Left.Right = new(node[string])
	testTreeString.Root.Left.Right.Value = "D"

	var tmp BinaryTree[string]

	type fields struct {
		Root  *node[string]
		equal func(a, b string) bool
		less  func(a, b string) bool
	}
	type args struct {
		elem string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Case 1: Insert in empty tree",
			fields: fields{
				Root: tmp.Root,
				equal: func(a, b string) bool {
					return a == b
				},
				less: func(a, b string) bool {
					return a < b
				},
			},
			args: args{
				elem: "S",
			},
			want: "S",
		},
		{
			name: "Case 2: Insert in full tree",
			fields: fields{
				Root: testTreeString.Root,
				equal: func(a, b string) bool {
					return a == b
				},
				less: func(a, b string) bool {
					return a < b
				},
			},
			args: args{
				elem: "E",
			},
			want: "E",
		},
		{
			name: "Case 3: Insert in full tree",
			fields: fields{
				Root: testTreeString.Root,
				equal: func(a, b string) bool {
					return a == b
				},
				less: func(a, b string) bool {
					return a < b
				},
			},
			args: args{
				elem: "A",
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree[string]{
				Root:  tt.fields.Root,
				equal: tt.fields.equal,
				less:  tt.fields.less,
			}
			b.Insert(tt.args.elem)
			got := b.Find(tt.args.elem).Value
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
