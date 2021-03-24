package mergeIntervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name string
		arr  [][]int
		want [][]int
	}{
		{
			name: "first",
			arr:  [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "second",
			arr:  [][]int{{1, 4}, {4, 5}},
			want: [][]int{{1, 5}},
		},
		{
			name: "third",
			arr:  [][]int{{1, 10}},
			want: [][]int{{1, 10}},
		},
		{
			name: "fourth",
			arr:  [][]int{},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
