package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortList(t *testing.T) {

	tests := []struct {
		name string
		list []int
		want []int
	}{
		{
			list: []int{4, 2, 1, 3},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sortList2(NewList(tt.list)).array())
		})
	}
}
