package array

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "a", args: args{[]int{10, 1, 2, 7, 6, 1, 5}, 8}, want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
