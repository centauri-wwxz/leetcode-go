package array

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "a", args: args{[]int{3, 4, -1, 1}}, want: 2},
		{name: "b", args: args{[]int{7, 8, 9, 11, 12}}, want: 1},
		{name: "c", args: args{[]int{-1, 4, 2, 1, 9, 10}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("FirstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
