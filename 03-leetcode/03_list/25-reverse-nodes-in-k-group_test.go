package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		list []int
		k    int
		want []int
	}{
		// {
		// 	list: []int{1, 2, 3, 4, 5},
		// 	k:    3,
		// 	want: []int{3, 2, 1, 4, 5},
		// },
		{
			list: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{2, 1, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewList(tt.list)
			assert.Equal(t, tt.want, reverseKGroup(head, tt.k).array())
		})
	}
}
