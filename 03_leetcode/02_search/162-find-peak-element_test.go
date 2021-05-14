package search

import "testing"

func Test_findPeakElement2(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			args: []int{1, 2, 3, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement2(tt.args); got != tt.want {
				t.Errorf("findPeakElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}
