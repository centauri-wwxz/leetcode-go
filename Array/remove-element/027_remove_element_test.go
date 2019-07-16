package array

import "testing"

func Test_RemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "a", args: args{[]int{3, 2, 2, 3}, 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
