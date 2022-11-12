package heap

import (
	"reflect"
	"testing"
)

func TestHeap_Add(t *testing.T) {
	type fields struct {
		arr      []int
		less     func(x, y int) bool
		more     func(x, y int) bool
		heapType Type
		size     int
	}
	type args struct {
		elem int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "Case 1: Add in empty max-heap",
			fields: fields{
				arr: []int{},
				less: func(x, y int) bool {
					return x < y
				},
				more: func(x, y int) bool {
					return x > y
				},
				heapType: MAX,
				size:     0,
			},
			args: args{
				elem: 42,
			},
			want: []int{42},
		},
		{
			name: "Case 2: Add in full max-heap",
			fields: fields{
				arr: []int{20, 15, 11, 6, 9, 7, 8, 1, 3, 5},
				less: func(x, y int) bool {
					return x < y
				},
				more: func(x, y int) bool {
					return x > y
				},
				heapType: MAX,
				size:     10,
			},
			args: args{
				elem: 17,
			},
			want: []int{20, 17, 11, 6, 15, 7, 8, 1, 3, 5, 9},
		},
		{
			name: "Case 3: Add in empty min-heap",
			fields: fields{
				arr: []int{},
				less: func(x, y int) bool {
					return x < y
				},
				more: func(x, y int) bool {
					return x > y
				},
				heapType: MIN,
				size:     0,
			},
			args: args{
				elem: 42,
			},
			want: []int{42},
		},
		{
			name: "Case 4: Add in full min-heap",
			fields: fields{
				arr: []int{1, 3, 7, 6, 5, 15, 8, 20, 9, 11},
				less: func(x, y int) bool {
					return x < y
				},
				more: func(x, y int) bool {
					return x > y
				},
				heapType: MIN,
				size:     10,
			},
			args: args{
				elem: 4,
			},
			want: []int{1, 3, 7, 6, 4, 15, 8, 20, 9, 11, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap[int]{
				arr:      tt.fields.arr,
				less:     tt.fields.less,
				more:     tt.fields.more,
				heapType: tt.fields.heapType,
				size:     tt.fields.size,
			}
			h.Add(tt.args.elem)
			got := h.arr
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Build(t *testing.T) {
	type fields struct {
		arr      []int
		less     func(x, y int) bool
		more     func(x, y int) bool
		heapType Type
		size     int
	}
	type args struct {
		elems []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "Case 1: Build min-heap",
			fields: fields{
				arr: []int{},
				less: func(x, y int) bool {
					return x < y
				},
				more: func(x, y int) bool {
					return x > y
				},
				heapType: MIN,
				size:     0,
			},
			args: args{
				elems: []int{31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: []int{1, 9, 2, 13, 10, 5, 3, 15, 14, 11, 21, 7, 6, 4, 17, 16, 24, 28, 23, 12, 22, 27, 30, 8, 20, 26, 19, 31, 18, 25, 29},
		},
		{
			name: "Case 2: Build max-heap",
			fields: fields{
				arr: []int{},
				less: func(x, y int) bool {
					return x < y
				},
				more: func(x, y int) bool {
					return x > y
				},
				heapType: MAX,
				size:     0,
			},
			args: args{
				elems: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			},
			want: []int{31, 23, 30, 19, 22, 27, 29, 17, 18, 21, 11, 25, 26, 28, 15, 16, 8, 4, 9, 20, 10, 5, 2, 24, 12, 6, 13, 1, 14, 7, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap[int]{
				arr:      tt.fields.arr,
				less:     tt.fields.less,
				more:     tt.fields.more,
				heapType: tt.fields.heapType,
				size:     tt.fields.size,
			}
			h.Build(tt.args.elems)
			got := h.arr
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_GetRoot(t *testing.T) {
	type fields struct {
		arr      []string
		less     func(x, y string) bool
		more     func(x, y string) bool
		heapType Type
		size     int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Case 1: Get from min-heap",
			fields: fields{
				arr: []string{"1", "9", "2", "13", "10", "5", "3", "15", "14", "11"},
				less: func(x, y string) bool {
					return x < y
				},
				more: func(x, y string) bool {
					return x > y
				},
				heapType: MIN,
				size:     10,
			},
			want: "1",
		},
		{
			name: "Case 2: Get from max-heap",
			fields: fields{
				arr: []string{"31", "23", "30", "19", "22", "27", "29", "17", "18", "21"},
				less: func(x, y string) bool {
					return x < y
				},
				more: func(x, y string) bool {
					return x > y
				},
				heapType: MAX,
				size:     10,
			},
			want: "31",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap[string]{
				arr:      tt.fields.arr,
				less:     tt.fields.less,
				more:     tt.fields.more,
				heapType: tt.fields.heapType,
				size:     tt.fields.size,
			}
			if got := h.GetRoot(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
