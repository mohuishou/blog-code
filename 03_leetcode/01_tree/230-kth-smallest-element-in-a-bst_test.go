package tree

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		data []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				data: []int{5, 3, 6, 2, 4, null, null, 1},
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(NewTreeNode(tt.args.data, 0), tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
