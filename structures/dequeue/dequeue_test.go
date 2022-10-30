package dequeue

import (
	"reflect"
	"testing"
)

func TestDequeue_AddFront(t *testing.T) {
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
			name: "Case 1: Add front to empty dequeue",
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
			name: "Case 2: Add front to full dequeue",
			fields: fields{
				arr: []int{1, 2, 3, 4},
			},
			args: args{
				elem: 42,
			},
			want: want{
				elems: []int{42, 1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dequeue[int]{
				arr: tt.fields.arr,
			}
			d.AddFront(tt.args.elem)
			want := &Dequeue[int]{
				arr: tt.want.elems,
			}
			if !reflect.DeepEqual(d, want) {
				t.Errorf("AddFront() got = %v, want %v", d, want)
			}
		})
	}
}

func TestDequeue_AddBack(t *testing.T) {
	type fields struct {
		arr []float64
	}
	type args struct {
		elem float64
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
			name: "Case 1: Add back to empty dequeue",
			fields: fields{
				arr: []float64{},
			},
			args: args{
				elem: 42.15,
			},
			want: want{
				elems: []float64{42.15},
			},
		},
		{
			name: "Case 1: Add back to empty dequeue",
			fields: fields{
				arr: []float64{1.2, 3.4, 5.6},
			},
			args: args{
				elem: 42.15,
			},
			want: want{
				elems: []float64{1.2, 3.4, 5.6, 42.15},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dequeue[float64]{
				arr: tt.fields.arr,
			}
			d.AddBack(tt.args.elem)
			want := &Dequeue[float64]{
				arr: tt.want.elems,
			}
			if !reflect.DeepEqual(d, want) {
				t.Errorf("AddBack() got = %v, want %v", d, want)
			}
		})
	}
}

func TestDequeue_RemoveFront(t *testing.T) {
	type fields struct {
		arr []string
	}
	type want struct {
		elems []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    want
		wantErr bool
	}{
		{
			name: "Case 1: Remove first elem in empty dequeue",
			fields: fields{
				arr: []string{},
			},
			want: want{
				elems: []string{},
			},
			wantErr: true,
		},
		{
			name: "Case 2: Remove first elem",
			fields: fields{
				arr: []string{"Hello"},
			},
			want: want{
				elems: []string{},
			},
			wantErr: false,
		},
		{
			name: "Case 3: Remove first elem in full dequeue",
			fields: fields{
				arr: []string{"Hello ", "World!"},
			},
			want: want{
				elems: []string{"World!"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dequeue[string]{
				arr: tt.fields.arr,
			}
			if err := d.RemoveFront(); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFront() error = %v, wantErr %v", err, tt.wantErr)
			}
			want := &Dequeue[string]{
				arr: tt.want.elems,
			}
			if !reflect.DeepEqual(d, want) {
				t.Errorf("RemoveFront() got = %v, want %v", d, want)
			}
		})
	}
}

func TestDequeue_RemoveBack(t *testing.T) {
	type fields struct {
		arr []bool
	}
	type want struct {
		elems []bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    want
		wantErr bool
	}{
		{
			name: "Case 1: Remove last elem in empty dequeue",
			fields: fields{
				arr: []bool{},
			},
			want: want{
				elems: []bool{},
			},
			wantErr: true,
		},
		{
			name: "Case 2: Remove last elem",
			fields: fields{
				arr: []bool{true},
			},
			want: want{
				elems: []bool{},
			},
			wantErr: false,
		},
		{
			name: "Case 3: Remove last elem in full dequeue",
			fields: fields{
				arr: []bool{true, false},
			},
			want: want{
				elems: []bool{true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dequeue[bool]{
				arr: tt.fields.arr,
			}
			if err := d.RemoveBack(); (err != nil) != tt.wantErr {
				t.Errorf("RemoveBack() error = %v, wantErr %v", err, tt.wantErr)
			}
			want := &Dequeue[bool]{
				arr: tt.want.elems,
			}
			if !reflect.DeepEqual(d, want) {
				t.Errorf("RemoveBack() got = %v, want %v", d, want)
			}
		})
	}
}

func TestDequeue_PeakFront(t *testing.T) {
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
			name: "Case 1: Get first elem in empty dequeue",
			fields: fields{
				arr: []user{},
			},
			want:    user{},
			wantErr: true,
		},
		{
			name: "Case 2: Get first elem in dequeue",
			fields: fields{
				arr: []user{
					{
						Name:   "John",
						Number: 898989,
					},
					{
						Name:   "Alice",
						Number: 010101,
					},
				},
			},
			want:    user{"John", 898989},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dequeue[user]{
				arr: tt.fields.arr,
			}
			got, err := d.PeakFront()
			if (err != nil) != tt.wantErr {
				t.Errorf("PeakFront() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeakFront() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDequeue_PeakBack(t *testing.T) {
	type fields struct {
		arr []complex64
	}
	tests := []struct {
		name    string
		fields  fields
		want    complex64
		wantErr bool
	}{
		{
			name: "Case 1: Get first elem in empty dequeue",
			fields: fields{
				arr: []complex64{},
			},
			want:    complex64(0),
			wantErr: true,
		},
		{
			name: "Case 2: Get first elem in dequeue",
			fields: fields{
				arr: []complex64{0, 1, 3},
			},
			want:    complex64(3),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dequeue[complex64]{
				arr: tt.fields.arr,
			}
			got, err := d.PeakBack()
			if (err != nil) != tt.wantErr {
				t.Errorf("PeakBack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeakBack() got = %v, want %v", got, tt.want)
			}
		})
	}
}
