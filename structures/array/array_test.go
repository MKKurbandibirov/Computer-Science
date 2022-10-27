package array

import (
	"reflect"
	"testing"
)

func TestArray_AddElems_Integer(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		elems []int
	}
	type want struct {
		elems []int
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
				arr: []int{},
			},
			args: args{
				elems: []int{42},
			},
			want: want{
				elems: []int{42},
			},
		},
		{
			name: "Case 2: For many element",
			fields: fields{
				arr: []int{42},
			},
			args: args{
				elems: []int{1, 2, 3, 4, 5, 6},
			},
			want: want{
				elems: []int{42, 1, 2, 3, 4, 5, 6},
			},
		},
		{
			name: "Case 3: For none element",
			fields: fields{
				arr: []int{42},
			},
			args: args{
				elems: []int{},
			},
			want: want{
				elems: []int{42},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[int]{
				arr: tt.fields.arr,
			}
			want := &Array[int]{
				arr: tt.want.elems,
			}
			a.AddElems(tt.args.elems...)
			if !reflect.DeepEqual(a.arr, want.arr) {
				t.Errorf("At() got = %v, want %v", a.arr, want.arr)
			}
		})
	}
}

func TestArray_AddElems_String(t *testing.T) {
	type fields struct {
		arr []string
	}
	type args struct {
		elems []string
	}
	type want struct {
		elems []string
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
				arr: []string{},
			},
			args: args{
				elems: []string{"Hello World!"},
			},
			want: want{
				elems: []string{"Hello World!"},
			},
		},
		{
			name: "Case 2: For many element",
			fields: fields{
				arr: []string{"Hello "},
			},
			args: args{
				elems: []string{"W", "o", "r", "l", "d", "!"},
			},
			want: want{
				elems: []string{"Hello ", "W", "o", "r", "l", "d", "!"},
			},
		},
		{
			name: "Case 3: For none element",
			fields: fields{
				arr: []string{"Hello World!"},
			},
			args: args{
				elems: []string{},
			},
			want: want{
				elems: []string{"Hello World!"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[string]{
				arr: tt.fields.arr,
			}
			want := &Array[string]{
				arr: tt.want.elems,
			}
			a.AddElems(tt.args.elems...)
			if !reflect.DeepEqual(a.arr, want.arr) {
				t.Errorf("At() got = %v, want %v", a.arr, want.arr)
			}
		})
	}
}

func TestArray_AddElems_Custom(t *testing.T) {
	type user struct {
		Name   string
		Number int
	}
	type fields struct {
		arr []user
	}
	type args struct {
		elems []user
	}
	type want struct {
		elems []user
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
				arr: []user{},
			},
			args: args{
				elems: []user{
					{
						Name:   "John",
						Number: 939393,
					},
				},
			},
			want: want{
				elems: []user{
					{
						Name:   "John",
						Number: 939393,
					},
				},
			},
		},
		{
			name: "Case 2: For many element",
			fields: fields{
				arr: []user{
					{
						Name:   "John",
						Number: 939393,
					},
				},
			},
			args: args{
				elems: []user{
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
			want: want{
				elems: []user{
					{
						Name:   "John",
						Number: 939393,
					},
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
		},
		{
			name: "Case 3: For none element",
			fields: fields{
				arr: []user{
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
			args: args{
				elems: []user{},
			},
			want: want{
				elems: []user{
					{
						Name:   "Sam",
						Number: 767676,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[user]{
				arr: tt.fields.arr,
			}
			want := &Array[user]{
				arr: tt.want.elems,
			}
			a.AddElems(tt.args.elems...)
			if !reflect.DeepEqual(a.arr, want.arr) {
				t.Errorf("At() got = %v, want %v", a.arr, want.arr)
			}
		})
	}
}

func TestArray_At_Integer(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		pos int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Case 1: Get first",
			fields:  fields{arr: []int{1, 2, 3, 4}},
			args:    args{pos: 0},
			want:    1,
			wantErr: false,
		},
		{
			name:    "Case 2: Get last",
			fields:  fields{arr: []int{1, 2, 3, 4}},
			args:    args{pos: 3},
			want:    4,
			wantErr: false,
		},
		{
			name:    "Case 3: Get intermediate",
			fields:  fields{arr: []int{1, 2, 3, 4}},
			args:    args{pos: 2},
			want:    3,
			wantErr: false,
		},
		{
			name:    "Case 4: Get negative",
			fields:  fields{arr: []int{1, 2, 3, 4}},
			args:    args{pos: -1},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Case 5: Get invalid",
			fields:  fields{arr: []int{1, 2, 3, 4}},
			args:    args{pos: 99},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[int]{
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

func TestArray_At_String(t *testing.T) {
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

func TestArray_At_Custom(t *testing.T) {
	type user struct {
		Name   string
		Number int
	}
	type fields struct {
		arr []user
	}
	type args struct {
		pos int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    user
		wantErr bool
	}{
		{
			name: "Case 1: Get first",
			fields: fields{arr: []user{
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
			}},
			args:    args{pos: 0},
			want:    user{"Alice", 010101},
			wantErr: false,
		},
		{
			name: "Case 2: Get last",
			fields: fields{arr: []user{
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
			}},
			args:    args{pos: 2},
			want:    user{"Sam", 767676},
			wantErr: false,
		},
		{
			name: "Case 3: Get intermediate",
			fields: fields{arr: []user{
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
			}},
			args:    args{pos: 1},
			want:    user{"Bob", 242424},
			wantErr: false,
		},
		{
			name: "Case 4: Get negative",
			fields: fields{arr: []user{
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
			}},
			args:    args{pos: -1},
			want:    user{},
			wantErr: true,
		},
		{
			name: "Case 5: Get invalid",
			fields: fields{arr: []user{
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
			}},
			args:    args{pos: 99},
			want:    user{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[user]{
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

func TestArray_InsertElems_Integer(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		pos   int
		elems []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			// Add cases
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[int]{
				arr: tt.fields.arr,
			}
			if err := a.InsertElems(tt.args.pos, tt.args.elems...); (err != nil) != tt.wantErr {
				t.Errorf("InsertElems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArray_InsertElems_String(t *testing.T) {
	type fields struct {
		arr []string
	}
	type args struct {
		pos   int
		elems []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			// Add cases
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[string]{
				arr: tt.fields.arr,
			}
			if err := a.InsertElems(tt.args.pos, tt.args.elems...); (err != nil) != tt.wantErr {
				t.Errorf("InsertElems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArray_InsertElems_Custom(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			// Add cases
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
		})
	}
}

func TestArray_DeleteElems_Integer(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		pos   int
		count int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[int]{
				arr: tt.fields.arr,
			}
			if err := a.DeleteElems(tt.args.pos, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("DeleteElems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArray_DeleteElems_String(t *testing.T) {
	type fields struct {
		arr []string
	}
	type args struct {
		pos   int
		count int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[string]{
				arr: tt.fields.arr,
			}
			if err := a.DeleteElems(tt.args.pos, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("DeleteElems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArray_DeleteElems(t *testing.T) {
	type user struct {
		Name   string
		Number int
	}
	type fields struct {
		arr []user
	}
	type args struct {
		pos   int
		count int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array[user]{
				arr: tt.fields.arr,
			}
			if err := a.DeleteElems(tt.args.pos, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("DeleteElems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}