package list

import (
	"reflect"
	"testing"
)

func TestList_AddFront(t *testing.T) {
	node1 := new(node[int])
	node1.Value = int(1)
	node1.Prev = nil
	node2 := new(node[int])
	node2.Value = int(2)
	node2.Next = nil
	node1.Next = node2
	node2.Prev = node1
	tl := new(List[int])
	tl.First = node1
	tl.Last = node2
	tl.length = 2

	type args struct {
		elem int
	}
	type want struct {
		list string
	}
	tests := []struct {
		name   string
		fields *List[int]
		args   args
		want   want
	}{
		{
			name:   "Case 1: Add to empty list",
			fields: new(List[int]),
			args: args{
				42,
			},
			want: want{
				list: "| 42 |",
			},
		},
		{
			name:   "Case 2: Add to full list",
			fields: tl,
			args: args{
				42,
			},
			want: want{
				list: "| 42 <-> 1 <-> 2 |",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.fields
			l.AddFront(tt.args.elem)
			want := tt.want.list
			if !reflect.DeepEqual(l.String(), want) {
				t.Errorf("AddFront() got = %v, want %v", l.String(), want)
			}
		})
	}
}

func TestList_AddBack(t *testing.T) {
	node1 := new(node[string])
	node1.Value = string("Hell")
	node1.Prev = nil
	node2 := new(node[string])
	node2.Value = string("oo!")
	node2.Next = nil
	node1.Next = node2
	node2.Prev = node1
	tl := new(List[string])
	tl.First = node1
	tl.Last = node2
	tl.length = 2

	type args struct {
		elem string
	}
	type want struct {
		list string
	}
	tests := []struct {
		name   string
		fields *List[string]
		args   args
		want   want
	}{
		{
			name:   "Case 1: Add to empty list",
			fields: new(List[string]),
			args: args{
				"42",
			},
			want: want{
				list: "| 42 |",
			},
		},
		{
			name:   "Case 2: Add to full list",
			fields: tl,
			args: args{
				"42",
			},
			want: want{
				list: "| Hell <-> oo! <-> 42 |",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.fields
			l.AddBack(tt.args.elem)
			want := tt.want.list
			if !reflect.DeepEqual(l.String(), want) {
				t.Errorf("AddBack() got = %v, want %v", l.String(), want)
			}
		})
	}
}

func TestList_Insert(t *testing.T) {
	node1 := new(node[float64])
	node1.Value = float64(54.7)
	node1.Prev = nil
	node2 := new(node[float64])
	node2.Value = float64(32.5)
	node2.Next = nil
	node1.Next = node2
	node2.Prev = node1
	tl := new(List[float64])
	tl.First = node1
	tl.Last = node2
	tl.length = 2

	type args struct {
		ind  int
		elem float64
	}
	type want struct {
		list string
	}
	tests := []struct {
		name    string
		fields  *List[float64]
		args    args
		want    want
		wantErr bool
	}{
		{
			name:   "Case 1: Add to empty list",
			fields: new(List[float64]),
			args: args{
				0,
				42.4,
			},
			want: want{
				list: "| 42.4 |",
			},
			wantErr: false,
		},
		{
			name:   "Case 2: Add to full list",
			fields: tl,
			args: args{
				1,
				42.5,
			},
			want: want{
				list: "| 54.7 <-> 42.5 <-> 32.5 |",
			},
			wantErr: false,
		},
		{
			name:   "Case 3: Invalid ind",
			fields: new(List[float64]),
			args: args{
				5,
				42.4,
			},
			want: want{
				list: "",
			},
			wantErr: true,
		},
		{
			name:   "Case 4: Negative ind",
			fields: new(List[float64]),
			args: args{
				-5,
				42.4,
			},
			want: want{
				list: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.fields
			if err := l.Insert(tt.args.ind, tt.args.elem); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			want := tt.want.list
			if !reflect.DeepEqual(l.String(), want) {
				t.Errorf("Insert() got = %v, want %v", l.String(), want)
			}
		})
	}
}

func TestList_At(t *testing.T) {
	type user struct {
		Name   string
		Number int
	}

	node1 := new(node[user])
	node1.Value = user{Name: "John", Number: 121212}
	node1.Prev = nil
	node2 := new(node[user])
	node2.Value = user{Name: "Alice", Number: 343434}
	node2.Next = nil
	node1.Next = node2
	node2.Prev = node1
	tl := new(List[user])
	tl.First = node1
	tl.Last = node2
	tl.length = 2

	type args struct {
		ind int
	}
	tests := []struct {
		name    string
		fields  *List[user]
		args    args
		want    user
		wantErr bool
	}{
		{
			name:   "Case 1: Get from empty list",
			fields: new(List[user]),
			args: args{
				0,
			},
			want:    user{},
			wantErr: true,
		},
		{
			name:   "Case 2: Get from full list",
			fields: tl,
			args: args{
				1,
			},
			want:    user{Name: "Alice", Number: 343434},
			wantErr: false,
		},
		{
			name:   "Case 3: Invalid ind",
			fields: new(List[user]),
			args: args{
				5,
			},
			want:    user{},
			wantErr: true,
		},
		{
			name:   "Case 4: Negative ind",
			fields: new(List[user]),
			args: args{
				-5,
			},
			want:    user{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.fields
			got, err := l.At(tt.args.ind)
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

func TestList_DeleteElem(t *testing.T) {
	node11 := new(node[bool])
	node11.Value = bool(true)
	node11.Prev = nil
	node21 := new(node[bool])
	node21.Value = bool(true)
	node21.Next = nil
	node11.Next = node21
	node21.Prev = node11
	tl1 := new(List[bool])
	tl1.First = node11
	tl1.Last = node21
	tl1.length = 2

	node1 := new(node[bool])
	node1.Value = bool(true)
	node1.Prev = nil
	node2 := new(node[bool])
	node2.Value = bool(true)
	node2.Next = nil
	node1.Next = node2
	node2.Prev = node1
	tl2 := new(List[bool])
	tl2.First = node1
	tl2.Last = node2
	tl2.length = 2

	type args struct {
		ind int
	}
	type want struct {
		list string
	}
	tests := []struct {
		name    string
		fields  *List[bool]
		args    args
		want    want
		wantErr bool
	}{
		{
			name:   "Case 1: Delete from empty list",
			fields: new(List[bool]),
			args: args{
				0,
			},
			want: want{
				"",
			},
			wantErr: true,
		},
		{
			name:   "Case 2: Delete from full list-1",
			fields: tl1,
			args: args{
				1,
			},
			want: want{
				"| true |",
			},
			wantErr: false,
		},
		{
			name:   "Case 2: Delete from full list-2",
			fields: tl2,
			args: args{
				0,
			},
			want: want{
				"| true |",
			},
			wantErr: false,
		},
		{
			name:   "Case 3: Invalid ind",
			fields: new(List[bool]),
			args: args{
				5,
			},
			want: want{
				"",
			},
			wantErr: true,
		},
		{
			name:   "Case 4: Negative ind",
			fields: new(List[bool]),
			args: args{
				-5,
			},
			want: want{
				"",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.fields
			if err := l.DeleteElem(tt.args.ind); (err != nil) != tt.wantErr {
				t.Errorf("DeleteElem() error = %v, wantErr %v", err, tt.wantErr)
			}
			want := tt.want.list
			if !reflect.DeepEqual(l.String(), want) {
				t.Errorf("DeleteElem() got = %v, want %v", l.String(), want)
			}
		})
	}
}
