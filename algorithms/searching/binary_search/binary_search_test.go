package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		array []int
		num   int
		less  func(a, b int) bool
		equal func(a, b int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1: Find in empty array",
			args: args{
				array: []int{},
				num:   42,
				less: func(a, b int) bool {
					return a < b
				},
				equal: func(a, b int) bool {
					return a == b
				},
			},
			want: -1,
		},
		{
			name: "Case 2: Find in full array",
			args: args{
				array: []int{-4, 0, 6, 18, 34, 42, 57, 78, 99},
				num:   42,
				less: func(a, b int) bool {
					return a < b
				},
				equal: func(a, b int) bool {
					return a == b
				},
			},
			want: 5,
		},
		{
			name: "Case 2: Find in full array",
			args: args{
				array: []int{-964, -910, -907, -894, -866, -861, -859, -855, -843, -810, -799, -798, -795, -787, -763, -762,
					-752, -703, -688, -665, -664, -628, -621, -565, -531, -525, -500, -493, -476, -460, -386, -361, -343,
					-294, -280, -242, -232, -214, -209, -207, -204, -188, -159, -141, -117, -105, -85, -80, -61, -59, -26,
					-21, -14, -13, 9, 26, 47, 117, 125, 132, 164, 193, 203, 235, 244, 259, 264, 271, 277, 282, 299, 322, 349,
					366, 452, 465, 493, 497, 530, 533, 536, 540, 582, 583, 594, 631, 671, 731, 736, 742, 749, 783, 784, 816,
					835, 869, 869, 940, 940, 998},
				num: 26,
				less: func(a, b int) bool {
					return a < b
				},
				equal: func(a, b int) bool {
					return a == b
				},
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.array, tt.args.num, tt.args.less, tt.args.equal); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
