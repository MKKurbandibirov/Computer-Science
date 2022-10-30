package stack

import (
	"reflect"
	"testing"
)

func TestStack_PushBack(t *testing.T) {
	type fields struct {
		arr []int
	}
	type args struct {
		elem int
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
			name: "Case 1: Add to empty stack",
			fields: fields{
				arr: []int{},
			},
			args: args{
				elem: 42,
			},
			want: want{
				elems: []int{42},
			},
		},
		{
			name: "Case 2: Add to stack",
			fields: fields{
				arr: []int{1, 2, 3},
			},
			args: args{
				elem: 42,
			},
			want: want{
				elems: []int{1, 2, 3, 42},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack[int]{
				arr: tt.fields.arr,
			}
			want := &Stack[int]{
				arr: tt.want.elems,
			}
			s.PushBack(tt.args.elem)
			if !reflect.DeepEqual(s.arr, want.arr) {
				t.Errorf("PushBack() got = %v, want %v", s.arr, want.arr)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		arr []string
	}
	type want struct {
		elems []string
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "Case 1: Delete last elem",
			fields: fields{
				arr: []string{"Hello"},
			},
			want: want{
				elems: []string{},
			},
		},
		{
			name: "Case 2: Delete elem",
			fields: fields{
				arr: []string{"Hello", "World!"},
			},
			want: want{
				elems: []string{"Hello"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack[string]{
				arr: tt.fields.arr,
			}
			s.Pop()
			want := &Stack[string]{
				arr: tt.want.elems,
			}
			if !reflect.DeepEqual(s.arr, want.arr) {
				t.Errorf("Pop() got = %v, want %v", s.arr, want.arr)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type user struct {
		Name   string
		Number int
	}
	type fields struct {
		arr []user
	}
	tests := []struct {
		name    string
		fields  fields
		want    user
		wantErr bool
	}{
		{
			name: "Case 1: Get last elem",
			fields: fields{
				arr: []user{
					{
						Name:   "John",
						Number: 989898,
					},
				},
			},
			want: user{
				Name:   "John",
				Number: 989898,
			},
			wantErr: false,
		},
		{
			name: "Case 2: Get elem",
			fields: fields{
				arr: []user{
					{
						Name:   "John",
						Number: 989898,
					},
					{
						Name:   "Alice",
						Number: 121212,
					},
				},
			},
			want: user{
				Name:   "Alice",
				Number: 121212,
			},
			wantErr: false,
		},
		{
			name: "Case 3: Empty stack",
			fields: fields{
				arr: []user{},
			},
			want:    user{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack[user]{
				arr: tt.fields.arr,
			}
			got, err := s.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}
