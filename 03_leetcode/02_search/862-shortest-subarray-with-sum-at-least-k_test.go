package search

import "testing"

func Test_shortestSubarray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{-34, 37, 51, 3, -12, -50, 51, 100, -47, 99, 34, 14, -13, 89, 31, -14, -44, 23, -38, 6},
				k:    151,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{48, 99, 37, 4, -31},
				k:    140,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{-28, 81, -20, 28, -29},
				k:    89,
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{2, -1, 2},
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSubarray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("shortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
