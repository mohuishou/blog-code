package dp

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{3, 1, 2, 4},
				k:    2,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargestHeapSort(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
