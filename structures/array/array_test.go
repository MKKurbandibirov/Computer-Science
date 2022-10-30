package array

import (
	"reflect"
	"testing"
)

func TestArray_AddElems(t *testing.T) {
	type fields struct {
		arr []float64
	}
	type args struct {
		elems []float64
	}
	type want struct {
		elems []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "Case 1: For one element",
			fields: fields{
				arr: []float64{},
			},
			args: args{
				elems: []float64{42},
			},
			want: want{
				elems: []float64{42},
			},
		},
		{
			name: "Case 2: For many element",
			fields: fields{
				arr: []float64{42},
			},
			args: args{
				elems: []float64{1, 2, 3, 4, 5, 6},
			},
			want: want{
				elems: []float64{42, 1, 2, 3, 4, 5, 6},
			},
		},
		{
			name: "Case 3: For none element",
			fields: fields{
				arr: []float64{42},
			},
			args: args{
				elems: []float64{},
			},
			want: want{
				elems: []float64{42},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[float64]{
				arr: tt.fields.arr,
			}
			want := &Array[float64]{
				arr: tt.want.elems,
			}
			a.AddElems(tt.args.elems...)
			if !reflect.DeepEqual(a.arr, want.arr) {
				t.Errorf("AddElems() got = %v, want %v", a.arr, want.arr)
			}
		})
	}
}

func TestArray_At(t *testing.T) {
	type fields struct {
		arr []string
	}
	type args struct {
		pos int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Case 1: Get first",
			fields:  fields{arr: []string{"H", "e", "l", "l", "o"}},
			args:    args{pos: 0},
			want:    "H",
			wantErr: false,
		},
		{
			name:    "Case 2: Get last",
			fields:  fields{arr: []string{"H", "e", "l", "l", "o"}},
			args:    args{pos: 4},
			want:    "o",
			wantErr: false,
		},
		{
			name:    "Case 3: Get intermediate",
			fields:  fields{arr: []string{"H", "e", "l", "l", "o"}},
			args:    args{pos: 2},
			want:    "l",
			wantErr: false,
		},
		{
			name:    "Case 4: Get negative",
			fields:  fields{arr: []string{"H", "e", "l", "l", "o"}},
			args:    args{pos: -1},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Case 5: Get invalid",
			fields:  fields{arr: []string{"H", "e", "l", "l", "o"}},
			args:    args{pos: 99},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[string]{
				arr: tt.fields.arr,
			}
			got, err := a.At(tt.args.pos)
			if (err != nil) != tt.wantErr {
				t.Errorf("At() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("At() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_InsertElems(t *testing.T) {
	type user struct {
		Name   string
		Number int
	}
	type fields struct {
		arr []user
	}
	type args struct {
		pos   int
		elems []user
	}
	type want struct {
		elems []user
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Case 1: For one element at start",
			fields: fields{
				arr: []user{},
			},
			args: args{
				pos: 0,
				elems: []user{
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			want: want{
				elems: []user{
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Case 2: For one element in the middle",
			fields: fields{
				arr: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "Bob",
						Number: 242424,
					},
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			args: args{
				pos: 1,
				elems: []user{
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			want: want{
				elems: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
					{
						Name:   "Bob",
						Number: 242424,
					},
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Case 3: For many elements at start",
			fields: fields{
				arr: []user{
					{
						Name:   "Bob",
						Number: 242424,
					},
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			args: args{
				pos: 0,
				elems: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			want: want{
				elems: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
					{
						Name:   "Bob",
						Number: 242424,
					},
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Case 4: For many elements in the middle",
			fields: fields{
				arr: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			args: args{
				pos: 1,
				elems: []user{
					{
						Name:   "John",
						Number: 989898,
					},
					{
						Name:   "Bob",
						Number: 242424,
					},
				},
			},
			want: want{
				elems: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
					{
						Name:   "Bob",
						Number: 242424,
					},
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Case 5: For none elements",
			fields: fields{
				arr: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
					{
						Name:   "Bob",
						Number: 242424,
					},
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			args: args{
				pos:   2,
				elems: []user{},
			},
			want: want{
				elems: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
					{
						Name:   "Bob",
						Number: 242424,
					},
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Case 6: For negative position",
			fields: fields{
				arr: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			args: args{
				pos: -4,
				elems: []user{
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			want: want{
				elems: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Case 6: For invalid position",
			fields: fields{
				arr: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			args: args{
				pos: 78,
				elems: []user{
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			want: want{
				elems: []user{
					{
						Name:   "Alice",
						Number: 010101,
					},
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[user]{
				arr: tt.fields.arr,
			}
			if err := a.InsertElems(tt.args.pos, tt.args.elems...); (err != nil) != tt.wantErr {
				t.Errorf("InsertElems() error = %v, wantErr %v", err, tt.wantErr)
			}
			want := &Array[user]{
				arr: tt.want.elems,
			}
			if !reflect.DeepEqual(a, want) {
				t.Errorf("InsertElems() got = %v, want %v", a, want)
			}
		})
	}
}

func TestArray_DeleteElems(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		pos   int
		count int
	}
	type want struct {
		elems []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Case 1: For one element at start",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   0,
				count: 1,
			},
			want: want{
				elems: []int{2, 3, 4},
			},
			wantErr: false,
		},
		{
			name: "Case 2: For one element in the middle",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   2,
				count: 1,
			},
			want: want{
				elems: []int{1, 2, 4},
			},
			wantErr: false,
		},
		{
			name: "Case 3: For one element in end",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   3,
				count: 1,
			},
			want: want{
				elems: []int{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "Case 4: For many elements at start",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   0,
				count: 3,
			},
			want: want{
				elems: []int{4},
			},
			wantErr: false,
		},
		{
			name: "Case 5: For many elements in the middle",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   1,
				count: 2,
			},
			want: want{
				elems: []int{1, 4},
			},
			wantErr: false,
		},
		{
			name: "Case 6: For many elements in end",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   2,
				count: 2,
			},
			want: want{
				elems: []int{1, 2},
			},
			wantErr: false,
		},
		{
			name: "Case 7: For negative position",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   -1,
				count: 3,
			},
			want: want{
				elems: []int{1, 2, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "Case 8: For invalid position",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   89,
				count: 3,
			},
			want: want{
				elems: []int{1, 2, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "Case 9: For negative count",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   1,
				count: -3,
			},
			want: want{
				elems: []int{1, 2, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "Case 10: For invalid count",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				pos:   1,
				count: 53,
			},
			want: want{
				elems: []int{1, 2, 3, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[int]{
				arr: tt.fields.arr,
			}
			if err := a.DeleteElems(tt.args.pos, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("DeleteElems() error = %v, wantErr %v", err, tt.wantErr)
			}
			want := &Array[int]{
				arr: tt.want.elems,
			}
			if !reflect.DeepEqual(a, want) {
				t.Errorf("DeleteElems() got = %v, want %v", a, want)
			}
		})
	}
}
