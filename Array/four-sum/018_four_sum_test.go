package array

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "a",
			args: args{nums: []int{1, 0, -1, 0, -2, 2}, target: 0},
			want: [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
