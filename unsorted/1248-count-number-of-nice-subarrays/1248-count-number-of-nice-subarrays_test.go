package unsorted

import "testing"

func Test_numberOfSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{[]int{1, 1, 2, 1, 1}, 3}, 2},
		{"a", args{[]int{2, 4, 6}, 1}, 0},
		{"a", args{[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numberOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
