package array

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "a", args: args{[]int{1, 2, 3}}},
		{name: "b", args: args{[]int{3, 2, 1}}},
		{name: "c", args: args{[]int{1, 5, 8, 4, 7, 6, 5, 3, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NextPermutation(tt.args.nums)
		})
	}
}
