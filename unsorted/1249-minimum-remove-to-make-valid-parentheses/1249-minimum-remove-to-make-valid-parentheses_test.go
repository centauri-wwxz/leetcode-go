package unsorted

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"a", args{"lee(t(c)o)de)"}, "lee(t(c)o)de"},
		{"a", args{"a)b(c)d"}, "ab(c)d"},
		{"a", args{"))(("}, ""},
		{"a", args{"(a(b(c)d)"}, "(a(bc)d)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
