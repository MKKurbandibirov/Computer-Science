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
var testTreeCustom BinaryTree[person]
var testTreeCustom2 BinaryTree[person]

func TestBinaryTree_Find(t *testing.T) {
	testTreeInt.Root = new(node[int])
	testTreeInt.Root.Value = int(10)
	testTreeInt.Root.Left = new(node[int])
	testTreeInt.Root.Left.Value = int(7)
	testTreeInt.Root.Right = new(node[int])
	testTreeInt.Root.Right.Value = int(13)
	testTreeInt.Root.Left.Left = new(node[int])
	testTreeInt.Root.Left.Left.Value = int(4)
	testTreeInt.Root.Left.Right = new(node[int])
	testTreeInt.Root.Left.Right.Value = int(9)

	type fields struct {
		Root *node[int]
	}
	type args struct {
		elem  int
		equal func(a, b int) bool
		less  func(a, b int) bool
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
			},
			args: args{
				elem: 9,
				equal: func(a, b int) bool {
					return a == b
				},
				less: func(a, b int) bool {
					return a < b
				},
			},
			want: testTreeInt.Root.Left.Right,
		},
		{
			name: "Case 2: Find not existing element",
			fields: fields{
				Root: testTreeInt.Root,
			},
			args: args{
				elem: 346,
				equal: func(a, b int) bool {
					return a == b
				},
				less: func(a, b int) bool {
					return a < b
				},
			},
			want: nil,
		},
		{
			name: "Case 3: Find in nil root",
			fields: fields{
				Root: nil,
			},
			args: args{
				elem: 346,
				equal: func(a, b int) bool {
					return a == b
				},
				less: func(a, b int) bool {
					return a < b
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree[int]{
				Root: tt.fields.Root,
			}
			if got := b.Find(tt.args.elem, tt.args.equal, tt.args.less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Insert(t *testing.T) {
	testTreeString.Root = new(node[string])
	testTreeString.Root.Value = string("H")
	testTreeString.Root.Left = new(node[string])
	testTreeString.Root.Left.Value = string("C")
	testTreeString.Root.Right = new(node[string])
	testTreeString.Root.Right.Value = string("W")
	testTreeString.Root.Left.Left = new(node[string])
	testTreeString.Root.Left.Left.Value = string("B")
	testTreeString.Root.Left.Right = new(node[string])
	testTreeString.Root.Left.Right.Value = string("D")

	var tmp BinaryTree[string]

	type fields struct {
		Root *node[string]
	}
	type args struct {
		elem string
		less func(a, b string) bool
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
			},
			args: args{
				elem: "S",
				less: func(a, b string) bool {
					return a < b
				},
			},
			want: "S",
		},
		{
			name: "Case 2: Insert in full tree",
			fields: fields{
				Root: testTreeString.Root,
			},
			args: args{
				elem: "E",
				less: func(a, b string) bool {
					return a < b
				},
			},
			want: "E",
		},
		{
			name: "Case 3: Insert in full tree",
			fields: fields{
				Root: testTreeString.Root,
			},
			args: args{
				elem: "A",
				less: func(a, b string) bool {
					return a < b
				},
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree[string]{
				Root: tt.fields.Root,
			}
			b.Insert(tt.args.elem, tt.args.less)
			got := b.Find(tt.args.elem, func(a, b string) bool {
				return a == b
			}, func(a, b string) bool {
				return a < b
			}).Value
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestBinaryTree_Delete(t *testing.T) {
//
//	testTreeCustom.Root = new(node[person])
//	testTreeCustom.Root.Value = person{6, "Jack"}
//	testTreeCustom.Root.Left = new(node[person])
//	testTreeCustom.Root.Left.Value = person{2, "Simon"}
//	testTreeCustom.Root.Right = new(node[person])
//	testTreeCustom.Root.Right.Value = person{15, "Alice"}
//	testTreeCustom.Root.Left.Left = new(node[person])
//	testTreeCustom.Root.Left.Left.Value = person{1, "John"}
//	testTreeCustom.Root.Left.Right = new(node[person])
//	testTreeCustom.Root.Left.Right.Value = person{4, "Patric"}
//
//	type fields struct {
//		Root *node[person]
//	}
//	type args struct {
//		node *node[person]
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *node[person]
//	}{
//		{
//			name: "Case 1: Delete not existing node",
//			fields: fields{
//				Root: testTreeCustom.Root,
//			},
//			args: args{
//				node: nil,
//			},
//			want: nil,
//		},
//		{
//			name: "Case 2: Delete existing node",
//			fields: fields{
//				Root: testTreeCustom.Root,
//			},
//			args: args{
//				node: testTreeCustom.Root.Left.Right,
//			},
//			want: nil,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := &BinaryTree[person]{
//				Root: tt.fields.Root,
//			}
//			b.Delete(tt.args.node)
//			got := tt.args.node
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Delete() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
