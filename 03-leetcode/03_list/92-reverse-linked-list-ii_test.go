package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		list  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				list:  []int{1, 2, 3},
				left:  3,
				right: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			args: args{
				list:  []int{1, 2, 3, 4, 5},
				left:  2,
				right: 4,
			},
			want: []int{1, 4, 3, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseBetween(NewList(tt.args.list), tt.args.left, tt.args.right).array())
		})
	}
}
