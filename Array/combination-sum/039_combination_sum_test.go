package array

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "a", args: args{[]int{2, 3, 6, 7}, 7}, want: [][]int{{2, 2, 3}, {7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
