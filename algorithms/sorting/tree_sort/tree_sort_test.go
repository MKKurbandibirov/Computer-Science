package tree_sort

import (
	"reflect"
	"testing"
)

func TestTreeSort(t *testing.T) {
	type args struct {
		array []int
		less  func(a, b int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1: Sort sorted array",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				less: func(a, b int) bool {
					return a < b
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "Case 2: Sort some usually array",
			args: args{
				array: []int{-3, 89, -67, 88, 4, 56, 1, 0, 14, 35},
				less: func(a, b int) bool {
					return a < b
				},
			},
			want: []int{-67, -3, 0, 1, 4, 14, 35, 56, 88, 89},
		},
		{
			name: "Case 3: Sort big array",
			args: args{
				array: []int{452, 117, 583, 259, 594, 749, 530, 736, -810, 271, 493, -664, -21, -117, 203, 264, -242,
					582, 244, -861, -294, -787, -798, -688, -59, -493, -843, 536, -214, -80, 465, -763, -964, 47, 497,
					631, -204, 26, 366, -14, -703, -141, 322, 277, 742, -476, -460, -500, -894, 784, -361, -188, 998,
					-159, 125, -61, 349, 282, -525, -866, -799, 940, -628, -762, -13, 869, 869, 783, -105, 235, -232,
					193, -907, -910, -531, 835, -386, 9, -665, -280, -207, -26, 533, -209, 816, -343, 731, 164, 940,
					-795, 299, -859, -855, 671, -621, 132, 540, -752, -85, -565},
				less: func(a, b int) bool {
					return a < b
				},
			},
			want: []int{-964, -910, -907, -894, -866, -861, -859, -855, -843, -810, -799, -798, -795, -787, -763, -762,
				-752, -703, -688, -665, -664, -628, -621, -565, -531, -525, -500, -493, -476, -460, -386, -361, -343,
				-294, -280, -242, -232, -214, -209, -207, -204, -188, -159, -141, -117, -105, -85, -80, -61, -59, -26,
				-21, -14, -13, 9, 26, 47, 117, 125, 132, 164, 193, 203, 235, 244, 259, 264, 271, 277, 282, 299, 322, 349,
				366, 452, 465, 493, 497, 530, 533, 536, 540, 582, 583, 594, 631, 671, 731, 736, 742, 749, 783, 784, 816,
				835, 869, 869, 940, 940, 998},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TreeSort(tt.args.array, tt.args.less)
		})
		got := tt.args.array
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TreeSort() got = %v, want %v", got, tt.want)
		}
	}
}
