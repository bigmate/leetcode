package _89_Rotate_array

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
		k    int
	}{
		{
			name: "first",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
			k:    0,
		},
		{
			name: "second",
			nums: []int{1, 2, 3, 4},
			want: []int{4, 1, 2, 3},
			k:    5,
		},
		{
			name: "third",
			nums: []int{14},
			want: []int{14},
			k:    4,
		},
		{
			name: "fourth",
			nums: []int{},
			want: []int{},
			k:    18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.nums) != len(tt.want) {
				t.Fatal("bad args")
			}
			rotate(tt.nums, tt.k)
			for i := 0; i < len(tt.nums); i++ {
				if tt.nums[i] != tt.want[i] {
					t.Errorf("expected: %v; got: %v", tt.want, tt.nums)
				}
			}
		})
	}
}
