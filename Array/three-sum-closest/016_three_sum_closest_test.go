package array

import "testing"

func TestThreeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "a", args: args{nums: []int{-1, 2, 1, -4}, target: 1}, want: 2},
		{name: "b", args: args{nums: []int{0, 1, 2}, target: 3}, want: 3},
		{name: "c", args: args{nums: []int{0, 2, 1, -3}, target: 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("ThreeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
