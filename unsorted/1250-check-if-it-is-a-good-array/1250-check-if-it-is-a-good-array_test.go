package unsorted

import "testing"

func Test_isGoodArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"a", args{[]int{12, 5, 7, 23}}, true},
		{"a", args{[]int{29, 6, 10}}, true},
		{"a", args{[]int{3, 6}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGoodArray(tt.args.nums); got != tt.want {
				t.Errorf("isGoodArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
