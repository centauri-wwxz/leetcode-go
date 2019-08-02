package array

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "a", args: args{[]int{1, 3, 5, 6}, 5}, want: 2},
		{name: "b", args: args{[]int{1, 3, 5, 6}, 5}, want: 2},
		{name: "c", args: args{[]int{1, 3, 5, 6}, 7}, want: 4},
		{name: "d", args: args{[]int{1, 3, 5, 6}, 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
