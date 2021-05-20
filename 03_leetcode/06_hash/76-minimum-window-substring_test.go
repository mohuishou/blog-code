package hash

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "",
			args: args{
				s: "q",
				t: "q",
			},
			want: "q",
		},
		{
			name: "",
			args: args{
				s: "ADOBECODEBANC",
				t: "XCD",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "",
			args: args{
				s: "aaaaaxxbbxxaacb",
				t: "aab",
			},
			want: "aacb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
